package kvm

import (
	"cmpService/common/mcmodel"
	"cmpService/common/utils"
	"cmpService/mcagent/config"
	"cmpService/mcagent/svcmgrapi"
	"fmt"
	"strings"
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
		cfg := config.GetGlobalConfig()
		svcmgrRestAddr := fmt.Sprintf("%s:%s", cfg.SvcmgrIp, cfg.SvcmgrPort)
		// Notify ...
		fmt.Printf("Changed! --> Notify...\n")
		svcmgrapi.SendUpdateServer2Svcmgr(*l.Old, svcmgrRestAddr)
	}
}

func ApplyChangeFactor(server *mcmodel.MgoServer) {
	if server == nil { return }
	if server.Vms != nil {
		for _, vm := range *server.Vms {
			//vm.IsChangeIpAddr = false
			// apply DNAT
			AddDnatRuleByVm(&vm)
		}
	}
	if server.Networks != nil {
		utils.DeleteFilterReject()
	}
}

func GetDnatRuleConfigByVm(vm *mcmodel.MgoVm) *utils.DnatRule{
	cfg := config.GetGlobalConfig()
	// apply DNAT
	return &utils.DnatRule{
		vm.IpAddr,
		"3389",
		cfg.ServerIp,
		GetDnatPort(vm.VmIndex),
	}
}

const (
	DNAT_NEXT_DST_IP = 1
	DNAT_NEXT_DPORT = 2
	DNAT_NEXT_TO_DEST = 3
)

func GetDnatRuleConfigByRule(rule string) *utils.DnatRule {
	var dnat utils.DnatRule
	if strings.Contains(rule, "DNAT") == false {
		return nil
	}
	arr := strings.Fields(rule)
	var next int
	for _, obj := range arr {
		//fmt.Println("GetDnatRuleConfigByRule:", obj)
		if next > 0 {
			switch next {
			case DNAT_NEXT_DST_IP:
				if strings.Contains(obj, "/") {
					tmp := strings.Split(obj, "/")
					dnat.WantAddr = tmp[0]
				} else {
					dnat.WantAddr = obj
				}
			case DNAT_NEXT_DPORT:
				dnat.WantPort = obj
			case DNAT_NEXT_TO_DEST:
				tmp := strings.Split(obj, ":")
				dnat.ToAddr = tmp[0]
				dnat.ToPort = tmp[1]
			default:
			}
			next = 0
			continue
		}
		if obj == "-d" {
			next = DNAT_NEXT_DST_IP
		} else if obj == "--dport" {
			next = DNAT_NEXT_DPORT
		} else if obj == "--to-destination" {
			next = DNAT_NEXT_TO_DEST
		} else {
			next = 0
		}
	}
	return &dnat
}

func GetDnatList() *[]utils.DnatRule {
	var DNATList []utils.DnatRule
	natList := utils.GetNATRule()
	for _, rule := range natList {
		dnat := GetDnatRuleConfigByRule(rule)
		if dnat != nil {
			DNATList = append(DNATList, *dnat)
		}
	}
	return &DNATList
}

func AddDnatRuleByVm(vm *mcmodel.MgoVm) {
	if vm.IpAddr == "" {
		return
	}
	rule := GetDnatRuleConfigByVm(vm)
	// Get Dnat Rules
	dnatList := GetDnatList()
	isExist := false
	for _, nat := range *dnatList {
		if nat.Compare(rule) {
			isExist = true
			break
		}
	}
	if isExist == false {
		fmt.Println("AddDnatRuleByVm: do it")
		utils.AddDNATRule(rule)
	}
}

func DeleteDnatRulByVm(vm *mcmodel.MgoVm) {
	rule := GetDnatRuleConfigByVm(vm)
	utils.DeleteDNATRule(rule)
}

func GetDnatPort(vmIndex int) string {
	cfg := config.GetGlobalConfig()
	return fmt.Sprintf("%d", cfg.DnatBasePortNum + vmIndex)
}