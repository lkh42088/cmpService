package agent

import (
	config2 "cmpService/mcagent/config"
	"cmpService/mcagent/kvm"
	"cmpService/mcagent/mcinflux"
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

	wg.Add(4)

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

	//if kvm.LibvirtS != nil {
	//	go kvm.LibvirtS.Start(&wg)
	//} else {
	//	wg.Done()
	//}

	//Baremetal system info
	SendSysInfo()

	wg.Wait()
}

func configure() bool {
	// ConfigureMonitoring Mongo DB
	//if ! mcmongo.Configure() {
	//	fmt.Println("Failed to configure mongodb!")
	//	return false
	//}

	if !mcinflux.ConfigureInfluxDB() {
		fmt.Println("Failed to configure influxdb!")
		return false
	}

	// ConfigureMonitoring MonitorRoutine
	//if ! ConfigureMonitoring() {
	//	fmt.Println("Failed to configure agent!")
	//	return false
	//}

	//ConfigureVmList()

	kvm.ConfigureKvmRoutine()

	kvm.ConfigureLibvirtResource()

	kvm.ConfigureLibvirtStatstics()

	return true
}
