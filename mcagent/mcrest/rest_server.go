package mcrest

import (
	"cmpService/common/lib"
	config2 "cmpService/mcagent/config"
	"fmt"
	"github.com/gin-gonic/gin"
	"sync"
)


var Router *gin.Engine

func Start(parentwg *sync.WaitGroup) {

	config := config2.GetGlobalConfig()
	address := config.McagentIp + ":" + config.McagentPort

	fmt.Printf("global config: %v\n", config)
	fmt.Printf("REST API Server: ip %s\n", config.McagentIp)
	fmt.Printf("REST API Server: port %s\n", config.McagentPort)
	fmt.Printf("REST API Server: address %s\n", address)

	Router = gin.Default()

	rg := Router.Group(lib.McUrlPrefix)

	// Registration
	rg.POST(lib.McUrlRegisterServer, registerServerHandler)
	rg.POST(lib.McUrlUnRegisterServer, unRegisterServerHandler)

	// VM
	rg.POST(lib.McUrlCreateVm, addVmHandler)
	rg.POST(lib.McUrlDeleteVm, deleteVmHandler)
	rg.GET(lib.McUrlGetVmById, getVmByIdHandler)
	rg.GET(lib.McUrlVm, getVmAllHandler)

	// Network
	rg.POST(lib.McUrlNetworkAdd, addNetworkHandler)
	rg.POST(lib.McUrlNetworkDelete, deleteNetworkHandler)

	// Get info
	rg.GET(lib.McUrlMonServer, getServerHandler)

	// Search Client Ip
	rg.GET(lib.McUrlPublicIp, GetClientIp)

	// Get Vm interface traffic
	rg.GET(lib.McUrlVmInterfaceTraffic, GetVmInterfaceTrafficByMac)

	Router.Run(address)
	if parentwg != nil {
		parentwg.Done()
	}
}
