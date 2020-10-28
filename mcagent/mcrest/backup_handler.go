package mcrest

import (
	"cmpService/common/messages"
	"cmpService/mcagent/config"
	"cmpService/mcagent/ktrest"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
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
	if backupInfo.NasBackupName != "" {
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
		// todo: Delete at NAS
	}
	// DB update
	_, err = config.GetMcGlobalConfig().DbOrm.DeleteMcVmBackup(backupInfo)
	if err != nil {
		fmt.Printf("! Error: MC-server DB updating is failed.(%v)\n", err)
		return err
	}
	return nil
}

