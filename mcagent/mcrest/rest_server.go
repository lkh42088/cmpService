package mcrest

import (
	config2 "cmpService/mcagent/config"
	"github.com/gin-gonic/gin"
	"sync"
)

const (
	apiPathPrefix = "/api/v1"
)

var Router *gin.Engine

func Start(parentwg *sync.WaitGroup) {

	config := config2.GetGlobalConfig()
	address := config.McagentIp + ":" + config.McagentPort

	Router = gin.Default()

	rg := Router.Group(apiPathPrefix)

	rg.POST(apiPathPrefix+"/vms/create", addVmHandler)
	rg.GET(apiPathPrefix+"/vms/:id", getVmByIdHandler)

	Router.Run(address)
	if parentwg != nil {
		parentwg.Done()
	}
}

