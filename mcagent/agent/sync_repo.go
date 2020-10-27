package agent

import (
	"cmpService/common/mcmodel"
	"cmpService/mcagent/config"
	"cmpService/mcagent/kvm"
	"cmpService/mcagent/mciptables"
	"cmpService/mcagent/repo"
	"cmpService/mcagent/svcmgrapi"
	"fmt"
)

func SyncRepoWithCurrentInfo() *mcmodel.McServerMsg {
	server := kvm.GetMcServerInfo()
	if server.Vms == nil {
		return nil
	}

	vmCache := repo.GetVmCache()
	for _, vm := range *server.Vms {
		obj := repo.GetVmFromRepoByName(vm.Name)
		if obj == nil {
			repo.AddVm2RepoForSync(&vm)
		}
		/************************************
		 * Apply DNAT
		 ************************************/
		kvm.AddDnatRuleByVm(&vm)
	}

	/************************************
	 * Delete Reject Filter Rule
	 ************************************/
	if server.Networks != nil {
		mciptables.DeleteFilterForwardRejectAllRule()
	}

	/************************************
	 * Delete Reject Filter Rule
	 ************************************/
	cfg := config.GetMcGlobalConfig()
	svcmgrRestAddr := fmt.Sprintf("%s:%s", cfg.SvcmgrIp, cfg.SvcmgrPort)
	svcmgrapi.SendUpdateServer2Svcmgr(server, svcmgrRestAddr)

	for _, vm := range *vmCache {
		res, _, _:= vm.LookupList(server.Vms)
		if res == false {
			/**************
			 * Delete Vm
			 **************/
			repo.DeleteVmFromRepo(vm)
			fmt.Println(">>>>> SyncRepoWithCurrentInfo: Delete Vm to repo -", vm.Name, " !!!!!")
		}
	}
	return &server
}
