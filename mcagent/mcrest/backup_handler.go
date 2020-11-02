package mcrest

import (
	"cmpService/common/mcmodel"
	"cmpService/common/messages"
	"cmpService/mcagent/config"
	"cmpService/mcagent/ktrest"
	"cmpService/mcagent/kvm"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
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
		for i := 1; i <= (size / (ktrest.FILE_BLOCK_500M)); i++ {
			fileList = append (fileList, backupInfo.Name + "/" + backupInfo.Name + fmt.Sprintf(".z%.2d", i))
		}
		fileList = append (fileList, backupInfo.Name + "/" + backupInfo.Name + fmt.Sprintf(".zip"))
		fileList = append (fileList, backupInfo.Name)
		fmt.Println("# FILE LIST: ", fileList)

		// Delete at KT Storage
		ktrest.PostAuthTokens()
		for _, fileName := range fileList {
			err = ktrest.DeleteStorageObject(backupInfo.KtContainerName, fileName)
			if err != nil {
				fmt.Printf("! Error: kt storage file deleting is failed.(%v)\n", err)
			}
		}
	} else {
		filePath := os.Getenv("HOME") + "/nas/" + backupInfo.NasBackupName
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
		_ = ktrest.PostAuthTokens()
		go ktrest.GetStorageObjectByDLO(data.KtContainerName, data.Name, ch)

		for {
			v := <-ch
			if v == 5 {
				//Unzip file
				currentPath, _ := os.Getwd()
				// File unzip
				fmt.Println("# Backup File Unzip......\n")
				ktrest.UnZipVmBackupFile(currentPath+"/"+data.Name, "./.")

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
		if _, err := os.Stat(os.Getenv("HOME") + "/" + "nas"); err != nil {
			fmt.Println("!! NAS directory is not mounted.\n")
		}
		src := os.Getenv("HOME") + "/nas/" + data.NasBackupName
		dst := dstPath + "/" + vm.FullPath
		err := kvm.CopyFile(src, dst)
		if err != nil {
			c.JSON(http.StatusRequestTimeout, err)
		}
		kvm.RebootingByBackupFile(src, dst, data, vm)
	}
	c.JSON(http.StatusOK, "OK")
}

