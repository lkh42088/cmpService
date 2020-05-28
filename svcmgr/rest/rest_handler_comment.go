package rest

import (
	"cmpService/common/lib"
	"cmpService/common/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"fmt"
)

func (h *Handler) GetCommentsByCode(c *gin.Context) {
	if h.db == nil {
		return
	}
	deviceCode := c.Param("devicecode")
	comments, err := h.db.GetComments(deviceCode)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
		return
	}
	fmt.Println("[###] %v", comments)
	c.JSON(http.StatusOK, comments)
}

func (h *Handler) AddComment(c *gin.Context) {
	if h.db == nil {
		return
	}

	// data parsing
	m, err := JsonUnmarshal(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":lib.RestAbnormalParam})
		return
	}

	// Find user name
	userId := m["registerId"].(string)
	user, err := h.db.GetUserByUserId(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}

	comment := models.DeviceComment{
		DeviceCode: m["deviceCode"].(string),
		Contents: m["comment"].(string),
		RegisterId: userId,
		RegisterName: user.Name,
	}

	err = h.db.AddComment(comment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
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
		c.JSON(http.StatusBadRequest, gin.H{"error":lib.RestAbnormalParam})
		return
	}

	value, _ := m["idx"].(float64)
	comment := models.DeviceComment{
		Idx: uint(int(value)),
		Contents: m["comment"].(string),
	}

	// User-Id check
	content, err1 := h.db.GetCommentByIdx(int(comment.Idx))
	if err1 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":err1.Error()})
		return
	} else if content.RegisterId != comment.RegisterId {
		c.JSON(http.StatusInternalServerError, gin.H{"error":lib.RestDoNotCreateUser})
		return
	}

	err = h.db.UpdateComment(comment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to updata data."})
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
		c.JSON(http.StatusInternalServerError, gin.H{"error":lib.RestAbnormalParam})
		return
	}

	// User-Id check
	userId := c.Param("userid")
	content, err1 := h.db.GetCommentByIdx(idx)
	if err1 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":err1.Error()})
		return
	} else if content.RegisterId != userId {
		c.JSON(http.StatusInternalServerError, gin.H{"error":lib.RestDoNotCreateUser})
		return
	}

	err = h.db.DeleteComments(idx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
		return
	}
	c.JSON(http.StatusOK, err)
}
