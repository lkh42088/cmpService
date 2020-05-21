package rest

import (
	"cmpService/common/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) GetCommentsByCode(c *gin.Context) {
	if h.db == nil {
		return
	}
	deviceCode := c.Param("devicecode")
	comments, err := h.db.GetComments(deviceCode)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	//fmt.Println("[###] %v", comments)
	c.JSON(http.StatusOK, comments)
}

func (h *Handler) AddComment(c *gin.Context) {
	if h.db == nil {
		return
	}

	// data parsing
	m, err := JsonUnmarshal(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":"Message parameter abnormal."})
		return
	}
	comment := models.DeviceComment{
		DeviceCode: m["deviceCode"].(string),
		Contents: m["comment"].(string),
		RegisterId: m["registerId"].(string),
	}

	err = h.db.AddComment(comment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "OK")
}

func (h *Handler) UpdateComment(c *gin.Context) {
	if h.db == nil {
		return
	}

	// data parsing
	m, err := JsonUnmarshal(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":"Message parameter abnormal."})
		return
	}

	value, _ := m["idx"].(float64)
	comment := models.DeviceComment{
		Idx: uint(int(value)),
		Contents: m["comment"].(string),
		RegisterId: m["registerId"].(string),
	}

	// User-Id check
	content, err1 := h.db.GetCommentByIdx(int(comment.Idx))
	if err1 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	} else if content.RegisterId != comment.RegisterId {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Can modify data by create user."})
		return
	}

	err = h.db.UpdateComment(comment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "OK")
}

func (h *Handler) DeleteCommentByIdx(c *gin.Context) {
	if h.db == nil {
		return
	}
	idx, err := strconv.Atoi(c.Param("commentidx"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// User-Id check
	userId := c.Param("userid")
	fmt.Println(userId)
	content, err1 := h.db.GetCommentByIdx(idx)
	if err1 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	} else if content.RegisterId != userId {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Can modify data by create user."})
		return
	}

	err = h.db.DeleteComments(idx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, err)
}
