package agent

import (
	"cmpService/common/mcmodel"
	config2 "cmpService/mcagent/config"
	"cmpService/mcagent/kvm"
	"cmpService/mcagent/mcmongo"
	"fmt"
	"sync"
	"time"
)

type Monitoring struct {
	Interval int
}

var Mon *Monitoring

func NewMonitoring(interval int) *Monitoring{
	return &Monitoring{
		interval,
	}
}

func SetMonitoring(m *Monitoring) {
	Mon = m
}

func ConfigureMonitoring() bool {
	cfg := config2.GetGlobalConfig()

	// ConfigureMonitoring Monitoring
	monitoring := NewMonitoring(cfg.MonitoringInterval)
	if monitoring != nil {
		SetMonitoring(monitoring)
		return true
	}
	return false
}

func (m *Monitoring)Start(parentwg *sync.WaitGroup) {
	loop := 1
	for {
		m.Run()
		time.Sleep(time.Duration(m.Interval * int(time.Second)))
		fmt.Printf("%d. monitoring check...\n", loop)
		loop += 1
	}
	if parentwg != nil {
		parentwg.Done()
	}
}

func (m *Monitoring)Run() {

	var wg sync.WaitGroup
	wg.Add(len(McVms.List))

	for _, vm := range McVms.List {
		go func(vm *mcmodel.MgoVm) {
			defer wg.Done()
			updated := false

			fmt.Printf("check vm: %s, %v\n", vm.Name, *vm)
			// check if copy vm instance, skip

			// check status
			if UpdateVmStatus(vm) {
				fmt.Println("Changed status!");
				updated = true
			}

			// check mac/ip address
			if UpdateVmAddress(vm) {
				fmt.Println("Changed Address!");
				updated = true
			}

			// update mongodb
			if updated {
				fmt.Println("Update vm: ", *vm)
				mcmongo.McMongo.UpdateVmByInternal(vm)
				// notify svcmgr
			}
		}(&vm)
	}

	wg.Wait()
}

func UpdateVmAddress (vm *mcmodel.MgoVm) bool {
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

func UpdateVmStatus (vm *mcmodel.MgoVm) bool {
	updated := false
	status := kvm.StatusVm(*vm)
	if vm.CurrentStatus != status {
		fmt.Println("change status: ", vm.CurrentStatus, "-->", status)
		vm.CurrentStatus = status
		updated = true
	}
	return updated
}

