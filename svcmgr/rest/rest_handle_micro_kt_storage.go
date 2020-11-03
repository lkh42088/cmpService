package rest

import (
	"cmpService/common/lib"
	"cmpService/common/mcmodel"
	"cmpService/common/messages"
	"cmpService/common/models"
	"cmpService/svcmgr/config"
	"cmpService/svcmgr/mcapi"
	"errors"
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
	var sendMsg messages.BackupEntryMsg
	var entryList []messages.BackupEntry
	var serverIdx int

	fmt.Println("DeleteVmBackup:", msg)
	for _, idx := range msg.IdxList {
		var entry messages.BackupEntry
		backup, err := config.SvcmgrGlobalConfig.Mariadb.GetMcVmBackupByIdx(uint(idx))
		if err == nil {
			entry.VmName = backup.VmName
			if backup.Name != "" {
				entry.BackupName = backup.Name
			} else if (backup.NasBackupName != "") {
				entry.BackupName = backup.NasBackupName
			}
			entryList = append(entryList, entry)
		}
		if serverIdx == 0 {
			serverIdx = backup.McServerIdx
		}
	}
	sendMsg.Entry = &entryList
	//fmt.Println("# ", sendMsg.Entry, serverIdx)
	if serverIdx != 0 {
		server, err := h.db.GetMcServerByServerIdx(uint(serverIdx))
		if err != nil {
			return
		}
		//fmt.Println("# ", server)

		if mcapi.SendDeleteVmBackupList(sendMsg, server) {
			// svcmgr db update
			for _, entry := range entryList {
				backup, err := config.SvcmgrGlobalConfig.Mariadb.DeleteMcVmBackupByVmName(entry.VmName)
				if err != nil {
					fmt.Printf("# %s Backup file deleting is failed.(%s)\n", backup.VmName, backup.Name)
					c.JSON(http.StatusInternalServerError, err)
					return
				} else {
					fmt.Printf("# %s Backup file deleted.(%s)\n", backup.VmName, backup.Name)
				}
			}
		}
	}
	c.JSON(http.StatusOK, "OK")
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

// ADD BACKUP INFO IN DB
func AddVmBackupFromMc(backup mcmodel.McVmBackup) error {
	fmt.Println("# Add: ", backup)
	_, err := config.SvcmgrGlobalConfig.Mariadb.AddMcVmBackup(backup)
	if err != nil {
		return err
	}
	return nil
}

func UpdateVmBackupFromMc(backup mcmodel.McVmBackup) error {
	fmt.Println("# Update: ", backup)
	_, err := config.SvcmgrGlobalConfig.Mariadb.UpdateMcVmBackup(backup)
	if err != nil {
		return err
	}
	return nil
}

func (h *Handler) StoreVmBackupFromMc(c *gin.Context) {
	var backup mcmodel.McVmBackup
	c.Bind(&backup)
	fmt.Println("# Store: ", backup)
	_, err := config.SvcmgrGlobalConfig.Mariadb.GetMcVmBackupByVmName(backup.VmName)
	if err != nil {
		err = AddVmBackupFromMc(backup)
	} else {
		err = UpdateVmBackupFromMc(backup)
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, "OK")
}

// BACKUP FILE RESTORE
func (h *Handler) RestoreBackupStart(c *gin.Context) {
	var data mcmodel.McVmBackup
	c.Bind(&data)
	fmt.Print("# Restore : ", data)

	// send backup restore action
	v, err := config.SvcmgrGlobalConfig.Mariadb.GetMcVmBackupByVmName(data.VmName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	server, err := h.db.GetMcServerByServerIdx(uint(data.McServerIdx))
	if !mcapi.SendRestoreBackup2Mc(v, server) {
		c.JSON(http.StatusInternalServerError, errors.New("Send action message to MC server.\n"))
		return
	}

	c.JSON(http.StatusOK, "OK")
}

func (h *Handler) UpdateMcVmList(c *gin.Context) {
	var recvMsg []mcmodel.McVm
	c.Bind(&recvMsg)
	fmt.Print("# VM List : ", recvMsg)

	s, _ := config.SvcmgrGlobalConfig.Mariadb.GetMcServerByIp(c.ClientIP())

	if recvMsg != nil {
		vmList, _ := config.SvcmgrGlobalConfig.Mariadb.GetMcVmsByServerIdx(int(s.Idx))
		for _, vm := range recvMsg {
			old := mcmodel.LookupVm(&vmList, vm)
			if old != nil {
				vm.Idx = old.Idx
				vm.CompanyIdx = old.CompanyIdx
				vm.McServerIdx = old.McServerIdx
				obj, _ := config.SvcmgrGlobalConfig.Mariadb.UpdateMcVm(vm)
				fmt.Println("update vm: ", obj)
			} else {
				vm.Idx = 0
				vm.McServerIdx = int(s.Idx)
				vm.CompanyIdx = s.CompanyIdx
				obj, _ := config.SvcmgrGlobalConfig.Mariadb.AddMcVm(vm)
				fmt.Println("insert vm: ", obj)
			}
		}
		for _, vm := range vmList {
			obj := mcmodel.LookupVm(&vmList, vm)
			if obj == nil {
				config.SvcmgrGlobalConfig.Mariadb.DeleteMcVm(vm)
			}
		}
	} else {
		vmList, _ := config.SvcmgrGlobalConfig.Mariadb.GetMcVmsByServerIdx(int(s.Idx))
		for _, vm := range vmList {
			config.SvcmgrGlobalConfig.Mariadb.DeleteMcVm(vm)
		}
	}
	c.JSON(http.StatusOK, "OK")
}
