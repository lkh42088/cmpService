package mcrest

import (
	"cmpService/common/messages"
	"cmpService/mcagent/kvm"
	"fmt"
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

func deleteVmSnapshotEntryList(c *gin.Context) {
	var msg messages.SnapshotEntryMsg
	c.ShouldBindJSON(&msg)
	fmt.Println("^^^||||| deleteVmSnapshotEntryList")
	for _, entry := range *msg.Entry {
		dom, err := kvm.GetDomainByName(entry.VmName)
		if err != nil {
			continue
		}
		snap, err := dom.SnapshotLookupByName(entry.SnapName, 0)
		if err != nil {
			continue
		}
		// Delete snapshot
		kvm.DeleteSnap(entry.VmName, snap)
	}
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

func recoveryVmSnapshot(c *gin.Context) {
	var msg messages.SnapshotEntry
	c.ShouldBindJSON(&msg)
	kvm.Revert2Snapshot(msg.VmName, msg.SnapName)
	c.JSON(http.StatusOK, msg)
}
