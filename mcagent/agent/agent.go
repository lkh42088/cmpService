package agent

import (
	config2 "cmpService/mcagent/config"
	"cmpService/mcagent/mcmongo"
	"cmpService/mcagent/mcrest"
	"sync"
)

func Start (config string) {
	var wg sync.WaitGroup

	if !config2.ApplyGlobalConfig(config) {
		return
	}

	configure()

	wg.Add(1)

	go mcrest.Start(&wg)

	wg.Wait()
}

func configure () {
	// Configure Mongo DB
	mcmongo.Configure()
}
