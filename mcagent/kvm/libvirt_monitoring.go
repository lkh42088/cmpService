package kvm

import (
	"cmpService/common/mcmodel"
	"cmpService/mcagent/config"
	"cmpService/mcagent/mciptables"
	"cmpService/mcagent/repo"
	"cmpService/mcagent/svcmgrapi"
	"fmt"
	"sync"
	"time"
)

type LibvirtResource struct {
	Interval int
	Old *mcmodel.McServerMsg
}

var LibvirtR *LibvirtResource

func (l *LibvirtResource) GetVmByName(name string) *mcmodel.McVm {
	server := l.Old
	if server == nil && server.Vms == nil {
		return nil
	}
	for _, vm := range *server.Vms {
		if vm.Name == name {
			return &vm
		}
	}
	return nil
}

func NewLibvirtResource(interval int) *LibvirtResource {
	return &LibvirtResource{
		Interval: interval,
	}
}

func SetLibvirtResource(l *LibvirtResource) {
	LibvirtR = l
}

func ConfigureLibvirtResource(server *mcmodel.McServerMsg) {
	l := NewLibvirtResource(10)
	l.Old = server
	SetLibvirtResource(l)
}

func (l *LibvirtResource) Start(parentwg *sync.WaitGroup) {
	loop := 1
	for {
		l.Run()
		time.Sleep(time.Duration(l.Interval * int(time.Second)))
		fmt.Printf("LibvirtResource Loop %d -----------------\n", loop)
		loop += 1
	}
	parentwg.Done()
}

func (l *LibvirtResource) Run() {
	server := GetMcServerInfo()
	server.DumpSummary()

	isChanged := false
	if l.Old != nil {
		isChanged = l.Old.Compare(&server)
	} else {
		// First after starting
		isChanged = true
	}
	if isChanged {
		l.Old = &server
		ApplyChangeFactor(l.Old)
		repo.UpdateVmList(server.Vms)
		cfg := config.GetMcGlobalConfig()
		svcmgrRestAddr := fmt.Sprintf("%s:%s", cfg.SvcmgrIp, cfg.SvcmgrPort)
		// Notify ...
		fmt.Printf("Changed! --> Notify...\n")
		svcmgrapi.SendUpdateServer2Svcmgr(*l.Old, svcmgrRestAddr)
	}
}

func ApplyChangeFactor(server *mcmodel.McServerMsg) {
	if server == nil { return }
	if server.Vms != nil {
		for _, vm := range *server.Vms {
			//vm.IsChangeIpAddr = false
			// apply DNAT
			AddDnatRuleByVm(&vm)
		}
	}
	if server.Networks != nil {
		mciptables.DeleteFilterForwardRejectAllRule()
	}
}

func GetDnatRuleConfigByVm(vm *mcmodel.McVm) *mciptables.DnatRule {
	cfg := config.GetMcGlobalConfig()
	// apply DNAT
	return &mciptables.DnatRule{
		vm.IpAddr,
		"3389",
		cfg.ServerIp,
		GetDnatPort(vm.VmIndex),
	}
}

func AddDnatRuleByVm(vm *mcmodel.McVm) {
	if vm.IpAddr == "" {
		return
	}
	rule := GetDnatRuleConfigByVm(vm)
	// Get Dnat Rules
	dnatList := mciptables.GetDnatList()
	isExist := false
	for _, nat := range *dnatList {
		if nat.Compare(rule) {
			isExist = true
			break
		}
	}
	if isExist == false {
		fmt.Println("AddDnatRuleByVm: do it")
		mciptables.AddDNATRule(rule)
	}
}

func DeleteDnatRulByVm(vm *mcmodel.McVm) {
	rule := GetDnatRuleConfigByVm(vm)
	mciptables.DeleteDNATRule(rule)
}

func GetDnatPort(vmIndex int) string {
	cfg := config.GetMcGlobalConfig()
	return fmt.Sprintf("%d", cfg.DnatBasePortNum + vmIndex)
}