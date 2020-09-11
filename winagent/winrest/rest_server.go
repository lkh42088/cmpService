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
	address := conf.VmAgentIp + ":" + conf.VmAgentPort

	fmt.Printf("global config: %v\n", conf)
	fmt.Printf("REST API Server: ip %s\n", conf.VmAgentIp)
	fmt.Printf("REST API Server: port %s\n", conf.VmAgentPort)
	fmt.Printf("REST API Server: address %s\n", address)

	// Health Check
	winapi.SendMsgToMcAgent("CHECK", lib.ToMcUrlHealth)

	Router = gin.Default()

	//rg := Router.Group(lib.McUrlPrefix)
	rg := Router.Group(lib.WinUrlPrefix)

	// Health Check
	rg.GET(lib.WinUrlHealth, HealthCheck)

	// Registration
	//rg.POST(lib.McUrlRegisterServer, registerServerHandler)
	//rg.POST(lib.McUrlUnRegisterServer, unRegisterServerHandler)
	//
	//// VM
	//rg.POST(lib.McUrlCreateVm, addVmHandler)
	//rg.POST(lib.McUrlDeleteVm, deleteVmHandler)
	//rg.GET(lib.McUrlGetVmById, getVmByIdHandler)
	//rg.GET(lib.McUrlVm, getVmAllHandler)
	//
	//// Network
	//rg.POST(lib.McUrlNetworkAdd, addNetworkHandler)
	//rg.POST(lib.McUrlNetworkDelete, deleteNetworkHandler)
	//
	//// Get info
	//rg.GET(lib.McUrlMonServer, getServerHandler)
	//
	//// Search Client Ip
	//rg.GET(lib.McUrlPublicIp, GetClientIp)

	Router.Run(address)
	if parentwg != nil {
		parentwg.Done()
	}
}
