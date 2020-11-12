package winrest

import (
	"cmpService/common/lib"
	"cmpService/winagent/config"
	"cmpService/winagent/winapi"
	"fmt"
	"github.com/gin-gonic/gin"
	"sync"
)

var Router *gin.Engine

func Start(parentwg *sync.WaitGroup) {

	conf := config.GetGlobalConfig()
	address := conf.WinAgentIp + ":" + conf.WinAgentPort

	fmt.Printf("CONFIG: %v\n", conf)
	fmt.Printf("REST API Server: address %s\n", address)

	// Health Check
	if !winapi.SendMsgToMcAgent("CHECK", lib.ToMcUrlHealth) {
		fmt.Println("\n###### Can't connect to MC-AGENT. Please check this. ######\n")
	}

	Router = gin.Default()
	rg := Router.Group(lib.WinUrlPrefix)

	// Health Check
	rg.GET(lib.WinUrlHealth, HealthCheck)
	rg.POST(lib.WinUrlModifyConf, ModifyConfVariable)
	rg.POST(lib.WinUrlAgentRestart, ReConfiguration)

	Router.Run(address)
	if parentwg != nil {
		parentwg.Done()
	}
}
