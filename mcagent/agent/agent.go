package agent

import (
	config2 "cmpService/mcagent/config"
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

	wg.Add(2)

	// Rest Api Server
	go mcrest.Start(&wg)

	// Monitoring VMs
	go Mon.Start(&wg)

	wg.Wait()
}

func configure() bool {
	// ConfigureMonitoring Mongo DB
	if ! mcmongo.Configure() {
		fmt.Println("Failed to configure mongodb!")
		return false
	}
	// ConfigureMonitoring Monitoring
	if ! ConfigureMonitoring() {
		fmt.Println("Failed to configure agent!")
		return false
	}

	ConfigureVmList()
	InitVmList()

	return true
}
