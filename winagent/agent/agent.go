package agent

import (
	"cmpService/common/mcmodel"
	"cmpService/winagent/common"
	"cmpService/winagent/config"
	"cmpService/winagent/winrest"
	"fmt"
	"sync"
)

var GlobalSysInfo mcmodel.SysInfo

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
	winrest.Start(nil)

	wg.Wait()
}

func configure() bool {

	CheckMySystem()
	InsertMacInTelegrafConf(GlobalSysInfo.IfMac)
	common.RestartTelegraf()

	return true
}



