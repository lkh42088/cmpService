package kvm

import (
	"cmpService/common/mcmodel"
	"cmpService/mcagent/config"
	"cmpService/mcagent/mcmongo"
	"cmpService/mcagent/svcmgrapi"
	"fmt"
	"sync"
	"time"
)

type KvmRoutine struct {
	Interval int
	Vms map[uint]mcmodel.MgoVm
}

var KvmR *KvmRoutine

func InitKvmR() {
	for _, vm := range KvmR.Vms {
		delete(KvmR.Vms, vm.Idx)
	}
	vms, err := mcmongo.McMongo.GetVmAll()
	if err != nil {
		return
	}
	for _, vm := range vms {
		if vm.IsProcess {
			KvmR.Vms[vm.Idx] = vm
		}
	}
}

func NewKvmRoutine(interval int) *KvmRoutine {
	return &KvmRoutine{
		Interval: interval,
		Vms: map[uint]mcmodel.MgoVm{},
	}
}

func SetKvmRoutine(k *KvmRoutine) {
	KvmR = k
}

func ConfigureKvmRoutine() {
	k := NewKvmRoutine(5)
	SetKvmRoutine(k)
}

func (k *KvmRoutine) Start(parentwg *sync.WaitGroup) {
	loop := 1
	for {
		k.Run()
		time.Sleep(time.Duration(k.Interval * int(time.Second)))
		fmt.Printf("%d. KvmRoutine(%ds)\n", loop, k.Interval)
		loop += 1
	}
}

func (k *KvmRoutine) Run() {

	//InitKvmR()

	if len(k.Vms) == 0 {
		return
	}

	list := k.Vms

	var wg sync.WaitGroup
	wg.Add(len(list))

	cfg := config.GetGlobalConfig()
	svcmgrRestAddr := fmt.Sprintf("%s:%s", cfg.SvcmgrIp, cfg.SvcmgrPort)

	for _, vm := range list {
		go func(vm *mcmodel.MgoVm) {
			defer wg.Done()

			delete(k.Vms, vm.Idx)

			// 1. copy image
			vm.CurrentStatus = "coping image"
			vm.IsProcess = false

			mcmongo.McMongo.UpdateVmByInternal(vm)
			fmt.Println("KvmRoutine: copy image - ", vm)
			svcmgrapi.SendUpdateVm2Svcmgr(*vm, svcmgrRestAddr)
			CopyVmInstance(vm)
			// 2. Create vm
			CreateVmInstance(*vm)
			vm.CurrentStatus = "created vm"
			vm.IsCreated = true
			fmt.Println("KvmRoutine: create vm- ", vm)
			mcmongo.McMongo.UpdateVmByInternal(vm)
			// 3. Notify svcmgr
			svcmgrapi.SendUpdateVm2Svcmgr(*vm, svcmgrRestAddr)

		}(&vm)
	}
	wg.Wait()
}

func (k *KvmRoutine) RunGetLibvirtInfo() {
	vmList, netList, imgList := GetMcVirtInfo()
	mcmodel.DumpVmList(vmList)
	mcmodel.DumpNetworkList(netList)
	mcmodel.DumpImageList(imgList)
}

