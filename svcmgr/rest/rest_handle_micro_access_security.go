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

func (h *Handler) GetMcAccessSecurity(c *gin.Context) {
	rowsPerPage, err := strconv.Atoi(c.Param("rows"))
	if err != nil {
		fmt.Println("GetMcAccessSecurity error 1")
		c.JSON(http.StatusBadRequest, gin.H{"error": lib.RestAbnormalParam})
		return
	}
	offset, err := strconv.Atoi(c.Param("offset"))
	if err != nil {
		fmt.Println("GetMcAccessSecurity error 2")
		c.JSON(http.StatusBadRequest, gin.H{"error": lib.RestAbnormalParam})
		return
	}
	orderBy := c.Param("orderby")
	if err != nil {
		fmt.Println("GetMcAccessSecurity error 3")
		c.JSON(http.StatusBadRequest, gin.H{"error": lib.RestAbnormalParam})
		return
	}
	order := c.Param("order")
	if err != nil {
		fmt.Println("GetMcAccessSecurity error 4")
		c.JSON(http.StatusBadRequest, gin.H{"error": lib.RestAbnormalParam})
		return
	}
	cpName := c.Param("cpName")
	if err != nil {
		fmt.Println("GetMcAccessSecurity error 5")
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
	rules, err := config.SvcmgrGlobalConfig.Mariadb.GetMcFilterRulePage(page, cpName)
	if err != nil {
		fmt.Println("GetMcAccessSecurity error 6", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, rules)
}

func (h *Handler) AddMcAccessSecurity(c *gin.Context) {
	var msg mcmodel.McFilterRule
	c.Bind(&msg)
	fmt.Println("AddMcAccessSecurity:")
	msg.Dump()

	// Check
	_, err := h.db.GetMcFilterRuleBySerialNumberAndAddr(msg.SerialNumber, msg.IpAddr)
	if err == nil {
		fmt.Println("AddMcNetwork: Already exist")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Already exist"})
		return
	}

	// Add
	rule, err := h.db.AddMcFilterRule(msg)
	if err != nil {
		fmt.Println("AddMcNetwork: failed to add - ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	// Get rule detail
	sendMsg, err := h.db.GetMcFilterRuleByIdx(rule.Idx)
	if err != nil {
		fmt.Println("AddMcNetwork: failed to get detail rule - ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	// Send to mcagent
	fmt.Println("Send to mcagent", sendMsg)
	server, err := h.db.GetMcServerBySerialNumber(msg.SerialNumber)
	if err != nil {
		fmt.Println("AddMcNetwork: failed to get server - ", err)
		return
	}
	mcapi.SendAddAccessSecurity(sendMsg, server)
	c.JSON(http.StatusOK, sendMsg)
}

func (h *Handler) DeleteMcAccessSecurity(c *gin.Context) {
	var msg mcmodel.McFilterRule
	c.Bind(&msg)
	fmt.Println("DeleteMcAccessSecurity:")
	msg.Dump()

	// Check
	sendMsg, err := h.db.GetMcFilterRuleByIdx(msg.Idx)
	if err != nil {
		fmt.Println("DeleteMcAccessSecurity: failed to get rule - ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Delete
	_, err = h.db.DeleteMcFilterRule(msg)
	if err != nil {
		fmt.Println("DeleteMcAccessSecurity: failed to delete rule - ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Send to mcagent
	server, err := h.db.GetMcServerBySerialNumber(msg.SerialNumber)
	if err != nil {
		fmt.Println("DeleteMcAccessSecurity: failed to get server - ", err)
		return
	}
	fmt.Println("DeleteMcAccessSecurity: success")
	sendMsg.Dump()
	mcapi.SendDeleteAccessSecurity(sendMsg, server)
	c.JSON(http.StatusOK, "success")
}

func (h *Handler) DeleteMcAccessSecurityList(c *gin.Context) {
	var msg messages.DeleteDataMessage
	c.Bind(&msg)
	fmt.Println("DeleteMcAccessSecurityList:", msg)

	// Check
	for _, idx := range msg.IdxList {
		var rule mcmodel.McFilterRuleDetail
		rule, err := h.db.GetMcFilterRuleByIdx(uint(idx))
		if err != nil {
			fmt.Println("DeleteMcAccessSecurityList: failed to get rule - ", err)
			continue
		}
		fmt.Println("DeleteMcAccessSecurityList: delete", idx)
		_, err = h.db.DeleteMcFilterRule(rule.McFilterRule)
		if err != nil {
			fmt.Println("DeleteMcAccessSecurityList: failed to delete rule - ", err)
			continue
		}
		// Send to mcagent
		server, err := h.db.GetMcServerBySerialNumber(rule.SerialNumber)
		if err != nil {
			fmt.Println("DeleteMcAccessSecurityList: failed to get server - ", err)
			continue
		}
		fmt.Println("DeleteMcAccessSecurityList: send to mcagent")
		rule.Dump()
		mcapi.SendDeleteAccessSecurity(rule, server)
	}

	fmt.Println("DeleteMcAccessSecurityList: success")
	c.JSON(http.StatusOK, "success")
}
