package rest

import (
	"cmpService/common/mcmodel"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func (h *Handler) GetSystemInfoByMac(c *gin.Context) {
	mac := c.Param("mac")
	info, err := h.db.GetSystemInfoByMac(mac)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
	c.JSON(http.StatusOK, info)
}

func (h *Handler) AddSystemInfo(c *gin.Context) {
	var msg mcmodel.SysInfo
	c.Bind(&msg)

	msg, err := h.db.AddSystemInfo(msg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, msg)
}

func (h *Handler) CheckNStoreSystemInfo(c *gin.Context) {
	var msg mcmodel.SysInfo
	c.Bind(&msg)

	if msg.IfMac == "" {
		c.JSON(http.StatusBadRequest, "Interface MAC isn't valid.")
		return
	}

	// Convert mac-address : because libvert mac error (fe -> 52)
	if strings.HasPrefix(msg.IfMac, strings.ToLower("fe")) {
		strings.Replace(msg.IfMac, "fe", "52", 1)
	}

	info, err := h.db.GetSystemInfoByMac(msg.IfMac)
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, err)
	//	return
	//}

	if info.IfMac == "" {
		info, err = h.db.AddSystemInfo(msg)
		fmt.Println("sysinfo add success", info)
	} else {
		info, err = h.db.UpdateSystemInfo(msg)
		fmt.Println("sysinfo update success", info)
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, info)
}
