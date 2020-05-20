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

	// Search username query
	// Need to code

	comment := models.DeviceComment{
		DeviceCode: c.Param("devicecode"),
		Contents: c.Param("comment"),
		RegisterId: c.Param("userid"),
		//RegisterName:,
	}
	err := h.db.AddComment(comment)
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
	// test code by lkh
	fmt.Printf("[TEST BODY] %v\n", c.Request.Body)
	fmt.Printf("[TEST FULL] %v\n", c.Request)

	idx, err := strconv.Atoi(c.Param("commentidx"))
	if err != nil {
		c.JSON(http.StatusNoContent, gin.H{"Error":err.Error()})
		return
	}
	comment := models.DeviceComment{
		Idx: uint(idx),
		Contents: c.Param("comment"),
		RegisterId: c.Param("userid"),
	}

	// User-Id check
	content, err1 := h.db.GetCommentByIdx(idx)
	if err1 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if content.RegisterId != comment.RegisterId {
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
	err = h.db.DeleteComments(idx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, err)
}
