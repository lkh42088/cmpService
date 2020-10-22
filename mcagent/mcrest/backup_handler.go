package mcrest

import (
	"cmpService/common/messages"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DeleteVmBackupEntryList(c *gin.Context) {
	var msg messages.BackupEntryMsg
	c.ShouldBindJSON(&msg)
	fmt.Println("deleteVmBackupEntryList start.")
	for _, entry := range *msg.Entry {
		//// Delete Backup
		fmt.Println(entry)
		//todo : Bakcup file delete function
		//DeleteBackup(entry.VmName, snap)
		//todo : mc backup db update
	}
	c.JSON(http.StatusOK, msg)
}

