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

	wg.Add(1)

	go mcrest.Start(&wg)

	wg.Wait()
}

func configure() bool {
	// Configure Mongo DB
	if ! mcmongo.Configure() {
		fmt.Println("Failed to configure mongodb!")
		return false
	}
	return true
}
