package agent

import (
	"cmpService/winagent/common"
	"cmpService/winagent/config"
	"cmpService/winagent/winrest"
	"fmt"
	"io/ioutil"
	"sync"
)

func Start (conf string) {
	var wg sync.WaitGroup

	if !config.ApplyGlobalConfig(conf) {
		ioutil.WriteFile("C:/temp/winagent_log.txt", []byte("ApplyGlobalConfig failed.\n"), 0)
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
	err := common.AddFireWallRule("WindowAgentRule", "in", "allow", "tcp", "8083")
	fmt.Println("To add rule failed: ", err)

	return true
}



