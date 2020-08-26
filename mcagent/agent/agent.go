package agent

import (
	config2 "cmpService/mcagent/config"
	"cmpService/mcagent/kvm"
	"cmpService/mcagent/mcinflux"
	"cmpService/mcagent/mcmongo"
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
	go MonitorR.StartByVirsh(&wg)

	go kvm.KvmR.Start(&wg)

	wg.Wait()
}

func configure() bool {
	// ConfigureMonitoring Mongo DB
	if ! mcmongo.Configure() {
		fmt.Println("Failed to configure mongodb!")
		return false
	}

	if !mcinflux.ConfigureInfluxDB() {
		fmt.Println("Failed to configure influxdb!")
		return false
	}

	// ConfigureMonitoring MonitorRoutine
	if ! ConfigureMonitoring() {
		fmt.Println("Failed to configure agent!")
		return false
	}

	ConfigureVmList()
	kvm.ConfigureKvmRoutine()

	return true
}
