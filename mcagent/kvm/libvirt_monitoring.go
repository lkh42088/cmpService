package kvm

import (
	"cmpService/common/mcmodel"
	"cmpService/mcagent/config"
	"cmpService/mcagent/svcmgrapi"
	"fmt"
	"sync"
	"time"
)

type LibvirtResource struct {
	Interval int
	Old *mcmodel.MgoServer
}

var LibvirtR *LibvirtResource

func (l *LibvirtResource) GetVmByName(name string) *mcmodel.MgoVm {
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

func ConfigureLibvirtResource() {
	l := NewLibvirtResource(10)
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
}

func (l *LibvirtResource) Run() {
	server := GetMcServerInfo()
	server.DumpSummary()

	isChanged := false
	if l.Old != nil {
		isChanged = l.Old.Compare(server)
	} else {
		l.Old = &server
		isChanged = true
	}
	if isChanged {
		cfg := config.GetGlobalConfig()
		svcmgrRestAddr := fmt.Sprintf("%s:%s", cfg.SvcmgrIp, cfg.SvcmgrPort)
		// Notify ...
		fmt.Printf("Changed! --> Notify...\n")
		svcmgrapi.SendUpdateServer2Svcmgr(*l.Old, svcmgrRestAddr)
	}
}