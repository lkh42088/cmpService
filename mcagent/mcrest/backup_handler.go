package mcrest

import (
	"cmpService/common/ktapi"
	"cmpService/common/mcmodel"
	"cmpService/common/messages"
	"cmpService/mcagent/config"
	"cmpService/mcagent/ktrest"
	"cmpService/mcagent/kvm"
	"cmpService/mcagent/repo"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strconv"
	time "time"
)

func DeleteVmBackupEntryList(c *gin.Context) {
	var msg messages.BackupEntryMsg
	c.ShouldBindJSON(&msg)
	fmt.Println("deleteVmBackupEntryList start.")
	for _, entry := range *msg.Entry {
		// Delete Backup
		fmt.Println("[INFO] Entry to delete backup : ", entry)
		err := DeleteVmBackup(entry.VmName)
		if err != nil {
			c.JSON(http.StatusInternalServerError, errors.New("Backup file deleting is failed.\n"))
		}
	}
	c.JSON(http.StatusOK, "OK")
}

func DeleteVmBackup(vmName string) error {
	var fileList []string
	// Get Vm backup info
	backupInfo, err := config.GetMcGlobalConfig().DbOrm.GetMcVmBackupByVmName(vmName)
	if err != nil {
		fmt.Printf("! Error: No found to vm backup info.(%v)\n", err)
		return err
	}
	if backupInfo.NasBackupName == "" {
		size := backupInfo.BackupSize
		// Make backup file name
		// 1. manifest file : NAME-cronsh.qcow2.decrease
		// 2. partial zip file : NAME-cronsh.qcow2.decrease/NAME-cronsh.qcow2.decrease.z01
		for i := 1; i <= (size / (ktapi.FILE_BLOCK_500M)); i++ {
			fileList = append (fileList, backupInfo.Name + "/" + backupInfo.Name + fmt.Sprintf(".z%.2d", i))
		}
		fileList = append (fileList, backupInfo.Name + "/" + backupInfo.Name + fmt.Sprintf(".zip"))
		fileList = append (fileList, backupInfo.Name)
		fmt.Println("# FILE LIST: ", fileList)

		// Delete at KT Storage
		token, _ := ktapi.PostAuthTokens()
		ktapi.GlobalToken = token
		for _, fileName := range fileList {
			err = ktrest.DeleteStorageObject(backupInfo.KtContainerName, fileName)
			if err != nil {
				fmt.Printf("! Error: kt storage file deleting is failed.(%v)\n", err)
			}
		}
	} else {
		filePath := os.Getenv("HOME") + "/nas/backup/" + backupInfo.NasBackupName
		kvm.DeleteFile(filePath)
	}
	// DB update
	_, err = config.GetMcGlobalConfig().DbOrm.DeleteMcVmBackup(backupInfo)
	if err != nil {
		fmt.Printf("! Error: MC-server DB updating is failed.(%v)\n", err)
		return err
	}
	return nil
}

// Restore Backup
func RestoreVmBackup(c *gin.Context) {
	var data mcmodel.McVmBackup
	c.Bind(&data)
	fmt.Println("# RestoreVmBackup : ", data)

	dstPath := config.GetMcGlobalConfig().VmBackupDir
	vm, _ := config.GetMcGlobalConfig().DbOrm.GetMcVmByName(data.VmName)

	// Backup Type check
	if data.NasBackupName == "" {
		// KT Backup file download
		ch := make(chan int)
		token, _ := ktapi.PostAuthTokens()
		ktapi.GlobalToken = token
		go ktrest.GetStorageObjectByDLO(data.KtContainerName, data.Name, ch)

		for {
			v := <-ch
			if v == 5 {
				//Unzip file
				currentPath, _ := os.Getwd()
				// File unzip
				fmt.Println("# Backup File Unzip......")
				err := ktrest.UnZipVmBackupFile(currentPath+"/"+data.Name, "./.")
				if err != nil {
					c.JSON(http.StatusInternalServerError, err)
					return
				}

				// Move file and Operating
				src := currentPath + dstPath + "/" + data.Name
				dst := vm.FullPath
				fmt.Println("# dst : ", dst)

				// Move File & Delete Unnecessary File & Reboot VM
				kvm.RebootingByBackupFile(src, dst, data, vm)
				break
			}
			time.Sleep(100 * time.Millisecond)
		}
	} else {
		// NAS type
		err := RestoreFromNas(data, vm)
		if err != nil {
			c.JSON(http.StatusRequestTimeout, err)
		}
	}
	c.JSON(http.StatusOK, "OK")
}

func RestoreFromNas(data mcmodel.McVmBackup, vm mcmodel.McVm) error {
	// Check Directory
	if _, err := os.Stat(os.Getenv("HOME") + "/nas/backup"); err != nil {
		fmt.Println("!! NAS directory is not mounted.")
		return err
	}
	// Get Image from NAS
	src := os.Getenv("HOME") + "/nas/backup/" + data.NasBackupName
	var dst string
	if vm.FullPath == "" {
		return fmt.Errorf("RestoreFromNas Failed. (Not exist vm instance)")
	} else {
		dst = vm.FullPath
	}
	fmt.Println("NAS (RestoreVmBackup) : ", src, dst)
	kvm.RebootingByBackupFileWithNas(src, dst, data, vm)
	return nil
}


func UpdateVmBackup(c *gin.Context) {
	var msg messages.BackupConfigMsg
	c.ShouldBindJSON(&msg)
	kvm.UpdateVmBackupByConfig(&msg)

	vm := mcmodel.McVm{}
	vm.Name = msg.VmName
	vm.BackupType, _ = strconv.ParseBool(msg.Type)
	vm.BackupDays, _ = strconv.Atoi(msg.Days)
	vm.BackupHours, _ = strconv.Atoi(msg.Hours)
	vm.BackupMinutes, _ = strconv.Atoi(msg.Minutes)
	repo.UpdateVm2DbForBackup(vm)

	c.JSON(http.StatusOK, msg)
}