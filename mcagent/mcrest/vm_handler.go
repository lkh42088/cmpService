package mcrest

import (
	"cmpService/common/mcmodel"
	"cmpService/common/utils"
	"cmpService/mcagent/config"
	"cmpService/mcagent/kvm"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func checkValidation(msg mcmodel.McVm) bool {
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

func addVmHandler(c *gin.Context) {
	var msg mcmodel.McVm
	err := c.ShouldBindJSON(&msg)
	fmt.Printf("addVmHandler: %s\n", msg.Dump())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if !checkValidation(msg) {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid message"})
		return
	}

	// Insert VM to Mongodb
	//_, err = mcmongo.McMongo.AddVm(&msg)
	//if err != nil {
	//	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	//	return
	//}

	msg.CurrentStatus = "Ready"

	fmt.Printf("addVmHandler: success - %v\n", msg)
	c.JSON(http.StatusOK, msg)

	// Update Vm
	msg.Filename = kvm.MakeFilename(&msg)
	msg.IsCreated = false
	msg.IsProcess = true

	cfg := config.GetGlobalConfig()
	//_, err = mcmongo.McMongo.UpdateVmByInternal(&msg)

	filepath := cfg.VmInstanceDir+"/"+msg.Filename+".qcow2"
	if ! utils.IsExistFile(filepath) {
		//kvm.KvmR.Vms = append(kvm.KvmR.Vms, msg)
		kvm.KvmR.Vms[msg.Idx] = msg
	}
}

func deleteVmHandler(c *gin.Context) {
	var msg mcmodel.McVm
	err := c.ShouldBindJSON(&msg)
	fmt.Printf("deleteVmHandler: %v\n", msg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//vm, err := mcmongo.McMongo.GetVmById(int(msg.Idx))
	vm := kvm.LibvirtR.GetVmByName(msg.Name)
	if vm == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "The vm does not exist!"})
		return
	}

	//err = mcmongo.McMongo.DeleteVm(int(msg.Idx))
	//if err != nil {
	//	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	//	return
	//}

	config.SetGlobalConfigByVmNumber(uint(vm.VmIndex), 0)
	// 1. Delete Dnat Rule
	kvm.DeleteDnatRulByVm(vm)

	// 2. Delete Vm instance
	kvm.DeleteVm(*vm)

	// 3. Delete Vm image
	kvm.DeleteVmInstance(*vm)

	fmt.Printf("deleteVmHandler: success\n")
	c.JSON(http.StatusOK, msg)
}

