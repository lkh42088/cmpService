package mcrest

import (
	"cmpService/common/mcmodel"
	"cmpService/common/utils"
	"cmpService/mcagent/config"
	"cmpService/mcagent/kvm"
	"cmpService/mcagent/mcmongo"
	"cmpService/mcagent/svcmgrapi"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func checkValidation(msg mcmodel.MgoVm) bool {
	if msg.Idx == 0 {
		fmt.Printf("error: idx is zero!\n")
		return false
	}
	if msg.McServerIdx == 0 {
		fmt.Printf("error: serverIdx is zero!\n")
		return false
	}
	if msg.CompanyIdx == 0 {
		fmt.Printf("error: cpIdx is zero!\n")
		return false
	}
	if msg.Name == "" {
		fmt.Printf("error: name is nil!\n")
		return false
	}
	if msg.OS == "" {
		fmt.Printf("error: os is nil!\n")
		return false
	}
	if msg.Image == "" {
		fmt.Printf("error: image is nil!\n")
		return false
	}
	return true
}

func GetMgoServer() (mcmodel.MgoServer, error) {
	var server mcmodel.MgoServer
	networks, err := kvm.GetNetworksFromXml()
	if err != nil {
		return server, err
	}
	images := kvm.GetImages()
	cfg := config.GetGlobalConfig()
	server.Mac = cfg.ServerMac
	server.Port = cfg.ServerPort
	server.Ip = cfg.ServerIp
	if len(networks) > 0 {
		server.Networks = &networks
	}
	if len(images) > 0 {
		server.Images = &images
	}
	fmt.Println("server:", server)
	fmt.Println("networks:", server.Networks)
	fmt.Println("images:", server.Images)
	return server, err
}

func registerServerHandler(c *gin.Context) {
	var msg mcmodel.McServerDetail
	err := c.ShouldBindJSON(&msg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	fmt.Printf("registerServerHandler: %v\n", msg)
	config.WriteServerStatus(msg.SerialNumber, msg.CompanyName, msg.CompanyIdx)

	server, _ := GetMgoServer()
	c.JSON(http.StatusOK, server)
}

func unRegisterServerHandler(c *gin.Context) {
	var msg mcmodel.McServerDetail
	err := c.ShouldBindJSON(&msg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	fmt.Printf("unRegisterServerHandler: %v\n", msg)
	config.DeleteServerStatus()
	c.JSON(http.StatusOK, msg)
}

func addVmHandler(c *gin.Context) {
	var msg mcmodel.MgoVm
	err := c.ShouldBindJSON(&msg)
	fmt.Printf("addVmHandler: %v\n", msg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if !checkValidation(msg) {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid message"})
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
	_, err = mcmongo.McMongo.UpdateVmByInternal(&msg)

	// 1. Copy image
	fmt.Printf("addVmHandler: before copy - %v\n", msg)
	filepath := cfg.VmInstanceDir+"/"+msg.Filename+".qcow2"
	fmt.Printf("addVmHandler: %s\n", filepath)
	if ! utils.IsExistFile(filepath) {
		kvm.CopyVmInstance(&msg)
	}
	fmt.Printf("addVmHandler: after copy - %v\n", msg)

	// 2. Create vm
	kvm.CreateVmInstance(msg)

	msg.CurrentStatus = "finished to create vm"

	// Update vm
	_, err = mcmongo.McMongo.UpdateVmByInternal(&msg)

	fmt.Printf("addVmHandler: finished - %v\n", msg)
	// Send rest api : status update - finished creating vm
	svcmgrapi.SendUpdateVm2Svcmgr(msg, svcmgrRestAddr)

	// Update Vm
	_, err = mcmongo.McMongo.UpdateVmByInternal(&msg)
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
