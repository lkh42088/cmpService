package agent

import (
	"cmpService/winagent/config"
	"cmpService/winagent/winrest"
	"fmt"
	"sync"
)

func Start (conf string) {
	var wg sync.WaitGroup

	if !config.ApplyGlobalConfig(conf) {
		return
	}

	if ! configure() {
		fmt.Println("Fatal: Failed configuration!")
		return
	}

	wg.Add(1)

	// Rest Api Server
	go winrest.Start(&wg)

	// MonitorRoutine VMs
	//if MonitorR != nil {
	//	go MonitorR.StartByVirsh(&wg)
	//} else {
	//	wg.Done()
	//}

	wg.Wait()
}

func configure() bool {

	return true
}
