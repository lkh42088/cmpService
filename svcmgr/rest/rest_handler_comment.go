package rest

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"nubes/common/models"
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

func (h *Handler) DeleteCommentsByCode(c *gin.Context) {
	if h.db == nil {
		return
	}
	deviceCode := c.Param("devicecode")
	err := h.db.DeleteComments(deviceCode)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, err)
}
