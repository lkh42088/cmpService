package agent

import (
	"cmpService/winagent/common"
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

	if ! Configure() {
		fmt.Println("Fatal: Failed configuration!")
		return
	}

	wg.Add(1)

	// Rest Api Server
	winrest.Start(nil)

	wg.Wait()
}

func Configure() bool {

	common.CheckMySystem()
	common.InsertMacInTelegrafConf(common.GlobalSysInfo.IfMac)
	common.RestartTelegraf()

	return true
}



