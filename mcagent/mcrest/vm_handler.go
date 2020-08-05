package mcrest

import (
	"cmpService/common/mcmodel"
	"cmpService/mcagent/config"
	"cmpService/mcagent/kvm"
	"cmpService/mcagent/mcmongo"
	"cmpService/mcagent/svcmgrapi"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func addVmHandler(c *gin.Context) {
	var msg mcmodel.MgoVm
	err := c.ShouldBindJSON(&msg)
	fmt.Printf("addVmHandler: %v\n", msg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Insert VM to Mongodb
	_, err = mcmongo.McMongo.AddVm(&msg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	fmt.Printf("addVmHandler: success - %v\n", msg)
	c.JSON(http.StatusOK, msg)

	// Update Vm
	msg.Filename = kvm.MakeFilename(msg)
	msg.CurrentStatus = "Be copy instance"

	cfg := config.GetGlobalConfig()
	svcmgrRestAddr := fmt.Sprintf("%s:%s", cfg.SvcmgrIp, cfg.SvcmgrPort)

	// Send rest api : status update - image copy
	svcmgrapi.SendUpdateVm2Svcmgr(msg, svcmgrRestAddr)

	// Update vm
	_, err = mcmongo.McMongo.UpdateVm(&msg)

	// 1. Copy image
	fmt.Printf("addVmHandler: before copy - %v\n", msg)
	kvm.CopyVmInstance(&msg)
	fmt.Printf("addVmHandler: after copy - %v\n", msg)

	// 2. Create vm
	kvm.CreateVmInstance(msg)

	// Update vm
	_, err = mcmongo.McMongo.UpdateVm(&msg)

	fmt.Printf("addVmHandler: finished - %v\n", msg)
	// Send rest api : status update - finished creating vm
	svcmgrapi.SendUpdateVm2Svcmgr(msg, svcmgrRestAddr)
	// Update Vm
	_, err = mcmongo.McMongo.UpdateVm(&msg)
}

func deleteVmHandler(c *gin.Context) {
	var msg mcmodel.MgoVm
	err := c.ShouldBindJSON(&msg)
	fmt.Printf("deleteVmHandler: %v\n", msg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	vm, err := mcmongo.McMongo.GetVmById(int(msg.Idx))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = mcmongo.McMongo.DeleteVm(int(msg.Idx))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	fmt.Printf("deleteVmHandler: success\n")
	c.JSON(http.StatusOK, msg)

	// 1. Delete Vm instance
	kvm.DeleteVm(vm)
	// 2. Delete Vm image
	kvm.DeleteVmInstance(vm)
}

func getVmByIdHandler(c *gin.Context) {
	idStr := c.Param("id")

	// Get VMs from Mongodb
	id, _ := strconv.Atoi(idStr)
	vm, err := mcmongo.McMongo.GetVmById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, vm)
}

func getVmAllHandler(c *gin.Context) {
	// Get VMs from Mongodb
	vm, err := mcmongo.McMongo.GetVmAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, vm)
}
