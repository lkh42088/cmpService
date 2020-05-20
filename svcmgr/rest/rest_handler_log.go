package rest

import (
	"cmpService/common/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) GetLogsByCode(c *gin.Context) {
	if h.db == nil {
		return
	}
	deviceCode := c.Param("devicecode")
	logs, err := h.db.GetLogs(deviceCode)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	//fmt.Println("[###] %v", comments)
	c.JSON(http.StatusOK, logs)
}

func (h *Handler) AddLog(c *gin.Context) {
	if h.db == nil {
		return
	}

	// Search username query
	// Need to code

	log := models.DeviceLog{
		DeviceCode: c.Param("devicecode"),
		//WorkCode: c.Param(""),
		//Field: c.Param(""),
		//OldStatus: c.Param(""),
		//NewStatus: c.Param(""),
		RegisterId: c.Param("userid"),
		//RegisterName:,
	}
	err := h.db.AddLog(log)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "OK")
}

func (h *Handler) UpdateLog(c *gin.Context) {
	if h.db == nil {
		return
	}
	idx, err := strconv.Atoi(c.Param("logidx"))
	if err != nil {
		c.JSON(http.StatusNoContent, gin.H{"Error":err.Error()})
		return
	}
	code, tmpErr := strconv.Atoi(c.Param("workcode"))
	if tmpErr != nil {
		c.JSON(http.StatusNoContent, gin.H{"Error":err.Error()})
		return
	}
	log := models.DeviceLog{
		Idx: uint(idx),
		WorkCode: code,
		Field: c.Param("field"),
		RegisterId: c.Param("userid"),
	}

	// User-Id check
	content, err1 := h.db.GetLogByIdx(idx)
	if err1 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if content.RegisterId != log.RegisterId {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Can modify data by create user."})
		return
	}

	err = h.db.UpdateLog(log.Field, c.Param("change"), log)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "OK")
}

func (h *Handler) DeleteLogByIdx(c *gin.Context) {
	if h.db == nil {
		return
	}
	idx, err := strconv.Atoi(c.Param("logidx"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	err = h.db.DeleteLog(idx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, err)
}
