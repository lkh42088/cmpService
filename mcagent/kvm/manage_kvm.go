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
	Vms []mcmodel.MgoVm
}


var KvmR *KvmRoutine

func NewKvmRoutine(interval int) *KvmRoutine {
	return &KvmRoutine{
		Interval: interval,
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
		if loop%10 == 0 {
			fmt.Printf("%d. KvmRoutine(%ds)\n", loop, k.Interval)
		}
		loop += 1
	}
}

func (k *KvmRoutine) Run() {

	if len(k.Vms) == 0 {
		return
	}

	var wg sync.WaitGroup
	wg.Add(len(k.Vms))

	cfg := config.GetGlobalConfig()
	svcmgrRestAddr := fmt.Sprintf("%s:%s", cfg.SvcmgrIp, cfg.SvcmgrPort)

	for _, vm := range KvmR.Vms {
		go func(vm *mcmodel.MgoVm) {
			defer wg.Done()
			// 1. copy image
			vm.CurrentStatus = "coping image"
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
}