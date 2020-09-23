package kvm

import (
	"cmpService/common/mcmodel"
	"cmpService/mcagent/config"
	"cmpService/mcagent/repo"
	"cmpService/mcagent/svcmgrapi"
	"fmt"
	"sync"
	"time"
)

type CreateVmFSM struct {
	Interval int
	Vms map[uint]mcmodel.McVm
}

var CreateVmFsm *CreateVmFSM

func NewCreateVmFsm(interval int) *CreateVmFSM {
	return &CreateVmFSM{
		Interval: interval,
		Vms: map[uint]mcmodel.McVm{},
	}
}

func SetCreateVmFsm(k *CreateVmFSM) {
	CreateVmFsm = k
}

func ConfigCreateVmFsm() {
	k := NewCreateVmFsm(5)
	SetCreateVmFsm(k)
}

func (k *CreateVmFSM) Start(parentwg *sync.WaitGroup) {
	loop := 1
	for {
		k.Run()
		time.Sleep(time.Duration(k.Interval * int(time.Second)))
		//fmt.Printf("%d. CreateVmFSM(%ds)\n", loop, k.Interval)
		loop += 1
	}
	parentwg.Done()
}

func (k *CreateVmFSM) Run() {

	if len(k.Vms) == 0 {
		return
	}

	list := k.Vms

	var wg sync.WaitGroup
	wg.Add(len(list))

	cfg := config.GetMcGlobalConfig()
	svcmgrRestAddr := fmt.Sprintf("%s:%s", cfg.SvcmgrIp, cfg.SvcmgrPort)

	for _, vm := range list {
		go func(vm *mcmodel.McVm) {
			defer wg.Done()

			delete(k.Vms, vm.Idx)

			/*****************************************************************
			 * 1. copy image
			 *****************************************************************/
			vm.CurrentStatus = "coping image"
			vm.IsProcess = false

			fmt.Println("CreateVmFSM: copy image - ", vm)
			svcmgrapi.SendUpdateVm2Svcmgr(*vm, svcmgrRestAddr)
			CopyVmInstance(vm)

			/*****************************************************************
			 * 2. Create vm
			 *****************************************************************/
			CreateVmInstance(*vm)
			vm.CurrentStatus = "created vm"
			vm.IsCreated = true
			fmt.Println("CreateVmFSM: create vm- ", vm)
			//mcmongo.McMongo.UpdateVmByInternal(vm)

			repo.UpdateVm2Repo(vm)

			/*****************************************************************
			 * 3. Notify svcmgr
			 *****************************************************************/
			svcmgrapi.SendUpdateVm2Svcmgr(*vm, svcmgrRestAddr)
		}(&vm)
	}
	wg.Wait()
}

func (k *CreateVmFSM) RunGetLibvirtInfo() {
	vmList, netList, imgList := GetMcVirtInfo()
	mcmodel.DumpVmList(vmList)
	mcmodel.DumpNetworkList(netList)
	mcmodel.DumpImageList(imgList)
}
