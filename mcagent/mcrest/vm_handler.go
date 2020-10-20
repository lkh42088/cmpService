package mcrest

import (
	"cmpService/common/mcmodel"
	"cmpService/common/messages"
	"cmpService/common/utils"
	"cmpService/mcagent/config"
	"cmpService/mcagent/kvm"
	"cmpService/mcagent/repo"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func checkValidation(msg mcmodel.McVm) bool {
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
	fmt.Println("vm-handler addVmHandler start----------------------------------------------------------")
	fmt.Println("msg : ", msg)
	fmt.Println("msg.UserId : ", msg.UserId)
	fmt.Println("err : ", err)
	fmt.Printf("addVmHandler: %s\n", msg.Dump())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	msg.Idx = 0

	if !checkValidation(msg) {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid message"})
		return
	}

	msg.CurrentStatus = "Ready"

	repo.AddVm2Repo(&msg)

	fmt.Printf("addVmHandler: success - %v\n", msg)
	c.JSON(http.StatusOK, msg)

	// Update Vm
	msg.Filename = kvm.MakeFilename(&msg)
	msg.IsCreated = false
	msg.IsProcess = true

	cfg := config.GetMcGlobalConfig()

	filepath := cfg.VmInstanceDir+"/"+msg.Filename+".qcow2"
	if ! utils.IsExistFile(filepath) {
		fmt.Println("addVmHandler: not exist file!")
		kvm.CreateVmFsm.Vms[msg.Idx] = msg
	}
	repo.UpdateVm2Repo(&msg)
}

func deleteVm(vmName string) bool {

	// Get VM Domain from Libvirt
	vm := kvm.LibvirtR.GetVmByName(vmName)
	if vm == nil {
		return false
	}

	// 1. Delete Dnat Rule
	kvm.DeleteDnatRulByVm(vm)

	// 2. Delete Cron Rule
	inVm := repo.GetVmFromRepoByName(vm.Name)
	if inVm != nil && inVm.SnapType == true {
		kvm.CronSnap.DeleteVm(inVm.Name)
	}

	// 3. Delete Vm snapshot
	kvm.DeleteAllSnapshot(vm.Name)

	// 4. Delete Vm instance
	kvm.DeleteVm(*vm)

	// 5. Delete Vm image
	kvm.DeleteVmInstance(*vm)

	repo.DeleteVmFromRepo(*vm)

	return true
}

func deleteVmHandler(c *gin.Context) {
	var msg mcmodel.McVm
	err := c.ShouldBindJSON(&msg)
	fmt.Printf("deleteVmHandler: %v\n", msg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	res := deleteVm(msg.Name)
	if res == false {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "The vm does not exist!"})
		return
	}

	fmt.Printf("deleteVmHandler: success\n")
	c.JSON(http.StatusOK, msg)
}

func applyVmActionHandler(c *gin.Context) {
	var msg messages.McVmActionMsg
	err := c.ShouldBindJSON(&msg)
	fmt.Printf("applyVmACtionHandler: %v\n", msg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	switch(msg.VmAction) {
	case 1:
		// shutdown
		kvm.LibvirtDestroyVm(msg.VmName)
	case 2:
		// start
		kvm.LibvirtStartVm(msg.VmName)
	case 3:
		// restart
		kvm.LibvirtResetVm(msg.VmName)
	case 4:
		// suspend
		kvm.LibvirtSuspendVm(msg.VmName)
	case 5:
		// resume
		kvm.LibvirtResumeVm(msg.VmName)
	case 6:
		// snapshot
		kvm.SafeSnapshot(msg.VmName, kvm.GetTimeWord(), "By action command")
	default:
	}
	c.JSON(http.StatusOK, msg)
}
