package agent

import (
	"cmpService/common/mcmodel"
	config2 "cmpService/mcagent/config"
	"cmpService/mcagent/kvm"
	"cmpService/mcagent/svcmgrapi"
	"fmt"
	"sync"
	"time"
)

type MonitorRoutine struct {
	Interval int
}

var MonitorR *MonitorRoutine

func NewMonitorRoutine(interval int) *MonitorRoutine {
	return &MonitorRoutine{
		interval,
	}
}

func SetMonitoring(m *MonitorRoutine) {
	MonitorR = m
}

func ConfigureMonitoring() bool {
	cfg := config2.GetMcGlobalConfig()

	// ConfigureMonitoring MonitorRoutine
	monitoring := NewMonitorRoutine(cfg.MonitoringInterval)
	if monitoring != nil {
		SetMonitoring(monitoring)
		return true
	}
	return false
}

func (m *MonitorRoutine) StartByVirsh(parentwg *sync.WaitGroup) {
	loop := 1
	for {
		//GetVmListByDb()
		m.RunByVirsh()
		time.Sleep(time.Duration(m.Interval * int(time.Second)))
		fmt.Printf("%d. monitoring check(%ds)\n", loop, m.Interval)
		loop += 1
	}
	if parentwg != nil {
		parentwg.Done()
	}
}

func (m *MonitorRoutine) RunByVirsh() {

	var wg sync.WaitGroup
	wg.Add(len(McVms.List))

	for _, vm := range McVms.List {
		go func(vm *mcmodel.McVm) {
			defer wg.Done()
			updated := false

			if !vm.IsCreated {
				return
			}

			fmt.Printf("check vm: %s, %v\n", vm.Name, *vm)
			// check if copy vm instance, skip

			// check status: virsh
			if UpdateVmStatus(vm) {
				fmt.Println("Changed status!")
				updated = true
			}

			// check mac/ip address: virsh
			if UpdateVmAddress(vm) {
				cfg := config2.GetMcGlobalConfig()
				fmt.Println("Changed Address!")
				updated = true
				// NAT setup
				kvm.ConfigDNAT(vm)
				vm.RemoteAddr = fmt.Sprintf("%s:%d",
					cfg.ServerIp,
					cfg.DnatBasePortNum + vm.VmIndex)
			}

			// update mongodb
			if updated {
				fmt.Println("Update vm: ", *vm)
				//mcmongo.McMongo.UpdateVmByInternal(vm)
				// notify svcmgr
				svcmgrapi.SendUpdateVm2Svcmgr(*vm,"192.168.0.72:8081")
			}
		}(&vm)
	}

	wg.Wait()
}

func UpdateVmAddress (vm *mcmodel.McVm) bool {
	updated := false
	ip, mac, res := kvm.GetIpAddressOfVm(*vm)
	if res < 0 {
		return false
	}
	if vm.IpAddr != ip {
		fmt.Println("change ip: ", vm.IpAddr, "-->", ip)
		vm.IpAddr = ip
		updated = true
	}
	if vm.Mac != mac {
		fmt.Println("change mac: ", vm.Mac, "-->", mac)
		vm.Mac = mac
		updated = true
	}
	return updated
}

func UpdateVmStatus (vm *mcmodel.McVm) bool {
	updated := false
	status := kvm.StatusVm(*vm)
	if vm.CurrentStatus != status {
		fmt.Println("change status: ", vm.CurrentStatus, "-->", status)
		vm.CurrentStatus = status
		updated = true
	}
	return updated
}

