package agent

import (
	config2 "cmpService/mcagent/config"
	"cmpService/mcagent/kvm"
	"cmpService/mcagent/mcrest"
	"fmt"
	"sync"
)

func Start (config string) {
	var wg sync.WaitGroup

	if !config2.ApplyGlobalConfig(config) {
		return
	}

	if ! configure() {
		fmt.Println("Fatal: Failed configuration!")
		return
	}

	wg.Add(3)

	// Rest Api Server
	go mcrest.Start(&wg)

	// MonitorRoutine VMs
	if MonitorR != nil {
		go MonitorR.StartByVirsh(&wg)
	} else {
		wg.Done()
	}

	if kvm.KvmR != nil {
		go kvm.KvmR.Start(&wg)
	} else {
		wg.Done()
	}

	if kvm.LibvirtR != nil {
		go kvm.LibvirtR.Start(&wg)
	} else {
		wg.Done()
	}

	wg.Wait()
}

func configure() bool {
	// ConfigureMonitoring Mongo DB
	//if ! mcmongo.Configure() {
	//	fmt.Println("Failed to configure mongodb!")
	//	return false
	//}

	// ConfigureMonitoring MonitorRoutine
	//if ! ConfigureMonitoring() {
	//	fmt.Println("Failed to configure agent!")
	//	return false
	//}

	//ConfigureVmList()
	//kvm.ConfigureKvmRoutine()

	kvm.ConfigureLibvirtResource()

	return true
}
