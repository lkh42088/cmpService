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

	config := config2.GetMcGlobalConfig()
	address := config.McagentIp + ":" + config.McagentPort

	fmt.Printf("global config: %v\n", config)
	fmt.Printf("REST API Server: ip %s\n", config.McagentIp)
	fmt.Printf("REST API Server: port %s\n", config.McagentPort)
	fmt.Printf("REST API Server: address %s\n", address)

	Router = gin.Default()

	rg := Router.Group(lib.McUrlPrefix)

	// Internal Resource
	rg.GET(lib.McUrlResource, getResourceHandler)
	rg.POST(lib.McUrlResource, resourceControlHandler)

	// Registration
	rg.POST(lib.McUrlRegisterServer, registerServerHandler)
	rg.POST(lib.McUrlUnRegisterServer, unRegisterServerHandler)

	// VM
	rg.POST(lib.McUrlCreateVm, addVmHandler)
	rg.POST(lib.McUrlDeleteVm, deleteVmHandler)
	rg.POST(lib.McUrlApplyVmAction, applyVmActionHandler)
	//rg.GET(lib.McUrlGetVmById, getVmByIdHandler)
	//rg.GET(lib.McUrlVm, getVmAllHandler)
	//rg.GET(lib.McUrlUpdateVmStatus, updateVmStatus)

	// Snapshot
	rg.GET(lib.McUrlGetVmSnapshot, getVmSnapshot)
	rg.POST(lib.McUrlAddVmSnapshot, addVmSnapshot)
	rg.POST(lib.McUrlDeleteVmSnapshot, deleteVmSnapshot)
	rg.POST(lib.McUrlDeleteVmSnapshotList, deleteVmSnapshotEntryList)
	rg.POST(lib.McUrlUpdateVmSnapshot, updateVmSnapshot)
	rg.POST(lib.McUrlRecoveryVmSnapshot, recoveryVmSnapshot)

	// Backup

	// Network
	rg.POST(lib.McUrlNetworkAdd, addNetworkHandler)
	rg.POST(lib.McUrlNetworkDelete, deleteNetworkHandler)

	// Get info
	rg.GET(lib.McUrlMonServer, getServerHandler)

	// Search Client Ip
	rg.GET(lib.McUrlPublicIp, GetClientIp)

	// Windows System API
	rg.POST(lib.McUrlHealthCheckFromWin, CheckHealth)

	// Backup
	rg.POST(lib.McUrlDeleteVmBackupList, DeleteVmBackupEntryList)

	Router.Run(address)
	if parentwg != nil {
		parentwg.Done()
	}
}
