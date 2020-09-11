package mcrest

import (
	"cmpService/common/messages"
	"cmpService/mcagent/kvm"
	"github.com/gin-gonic/gin"
	"net/http"
)

func getVmSnapshot(c *gin.Context) {
	var msg messages.SnapshotEntryMsg
	msg.Entry = kvm.GetVmSnapshotAll()
	c.JSON(http.StatusOK, msg)
}

func addVmSnapshot(c *gin.Context) {
	var msg messages.SnapshotConfigMsg
	c.ShouldBindJSON(&msg)
	kvm.AddVmSnapshotByConfig(&msg)
	c.JSON(http.StatusOK, msg)
}

func deleteVmSnapshot(c *gin.Context) {
	var msg messages.SnapshotConfigMsg
	c.ShouldBindJSON(&msg)
	kvm.DeleteVmSnapshotByConfig(&msg)
	c.JSON(http.StatusOK, msg)
}

func updateVmSnapshot(c *gin.Context) {
	var msg messages.SnapshotConfigMsg
	c.ShouldBindJSON(&msg)
	kvm.UpdateVmSnapshotByConfig(&msg)
	c.JSON(http.StatusOK, msg)
}

func updateVmStatus(c *gin.Context) {
	var msg messages.VmStatusActionMsg
	c.ShouldBindJSON(&msg)
	kvm.UpdateVmStatus(&msg)
	c.JSON(http.StatusOK, msg)
}
