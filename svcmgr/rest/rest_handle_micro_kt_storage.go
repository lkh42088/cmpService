package rest

import (
	"cmpService/common/lib"
	"cmpService/common/mcmodel"
	"cmpService/common/messages"
	"cmpService/common/models"
	"cmpService/svcmgr/config"
	"cmpService/svcmgr/mcapi"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) UpdateKtAuthUrl(c *gin.Context) {
	var msg messages.KtAuthUrl
	c.Bind(&msg)
	fmt.Println("UpdateAuthUrl:", msg)
	server, err := h.db.UpdateKtAuthUrl(msg.Ip, msg.AuthUrl)
	if err != nil {
		return
	}
	fmt.Printf("UpdateAuthUrl: %+v\n", server)

	c.JSON(http.StatusOK, server)
}

func (h *Handler) GetVmBackupConfig(c *gin.Context) {
	serverIdx, _ := strconv.Atoi(c.Param("serverIdx"))
	fmt.Println("GetVmBackupConfig")
	server, err := h.db.GetMcServerByServerIdx(uint(serverIdx))
	if err != nil {
		return
	}
	vms, err := config.SvcmgrGlobalConfig.Mariadb.GetMcVmsByServerIdx(int(server.Idx))
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, vms)
}

func (h *Handler) NotifyMcAgentVmBackup(c *gin.Context) {
	var msg mcmodel.McVmBackup
	c.Bind(&msg)
	fmt.Println("NotifyMcAgentVmBackup:", msg.Command)
	server, err := h.db.GetMcServerBySerialNumber(msg.McServerSn)
	if err != nil {
		return
	}
	msg.McServerIdx = int(server.Idx)
	if msg.Command == "add" {
		// Add Backup
		config.SvcmgrGlobalConfig.Mariadb.AddMcVmBackup(msg)
	} else {
		// Delete Backup
		backup, err := config.SvcmgrGlobalConfig.Mariadb.GetMcVmBackupByName(msg.Name)
		if err == nil {
			config.SvcmgrGlobalConfig.Mariadb.DeleteMcVmBackup(backup)
		}
	}
	c.JSON(http.StatusOK, msg)
}

func (h *Handler) GetMcVmBackup(c *gin.Context) {
	rowsPerPage, err := strconv.Atoi(c.Param("rows"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": lib.RestAbnormalParam})
		return
	}
	offset, err := strconv.Atoi(c.Param("offset"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": lib.RestAbnormalParam})
		return
	}
	orderBy := c.Param("orderby")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": lib.RestAbnormalParam})
		return
	}
	order := c.Param("order")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": lib.RestAbnormalParam})
		return
	}
	cpName := c.Param("cpName")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": lib.RestAbnormalParam})
		return
	}

	page := models.Pagination{
		TotalCount:  0,
		RowsPerPage: rowsPerPage,
		Offset:      offset,
		OrderBy:     orderBy,
		Order:       order,
	}
	fmt.Println("1. page:")
	page.String()
	vms, err := config.SvcmgrGlobalConfig.Mariadb.GetMcBackupPage(page, cpName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, vms)
}

func (h *Handler) AddVmBackup(c *gin.Context) {
	var msg messages.BackupConfigMsg
	c.Bind(&msg)
	fmt.Println("AddVmBackup:", msg)
	server, err := h.db.GetMcServerByServerIdx(msg.ServerIdx)
	if err != nil {
		return
	}
	mcapi.SendAddVmBackup(msg, server)
}

func (h *Handler) DeleteVmBackup(c *gin.Context) {
	var msg messages.BackupConfigMsg
	c.Bind(&msg)
	fmt.Println("DeleteVmBackup:", msg)
	server, err := h.db.GetMcServerByServerIdx(msg.ServerIdx)
	if err != nil {
		return
	}
	mcapi.SendDeleteVmBackup(msg, server)
}

func (h *Handler) DeleteVmBackupEntryList(c *gin.Context) {
	var msg messages.DeleteDataMessage
	c.Bind(&msg)
	fmt.Println("DeleteVmBackup:", msg)
	var sendMsg messages.BackupEntryMsg
	var entryList []messages.BackupEntry
	var serverIdx int
	for _, idx := range msg.IdxList {
		var entry messages.BackupEntry
		backup, err := config.SvcmgrGlobalConfig.Mariadb.GetMcVmBackupByIdx(uint(idx))
		if err == nil {
			entry.VmName = backup.VmName
			entry.BackupName = backup.Name
			entryList = append(entryList, entry)
		}
		if serverIdx == 0 {
			serverIdx = backup.McServerIdx
		}
	}
	sendMsg.Entry = &entryList
	if serverIdx != 0 {
		server, err := h.db.GetMcServerByServerIdx(uint(serverIdx))
		if err != nil {
			return
		}
		mcapi.SendDeleteVmBackupList(sendMsg, server)
	}
}

func (h *Handler) UpdateVmBackup(c *gin.Context) {
	var msg messages.BackupConfigMsg
	c.Bind(&msg)
	fmt.Println("UpdateVmBackup:", msg)
	server, err := h.db.GetMcServerByServerIdx(msg.ServerIdx)
	if err != nil {
		return
	}
	mcapi.SendUpdateVmBackup(msg, server)
}
