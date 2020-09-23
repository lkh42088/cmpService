package rest

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type DeviceCount struct {
	Total			int		`json:"total"`
	Operate			int		`json:"operate"`
}

// FOR MARIA-DB
func (h *Handler) GetServerTotalCount(c *gin.Context) {
	var deviceCount	DeviceCount
	count, operCount, err := h.db.GetServerTotalCount()
	deviceCount.Total = count
	deviceCount.Operate = operCount
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}
	c.JSON(http.StatusOK, deviceCount)
}

func (h *Handler) GetVmTotalCount(c *gin.Context) {
	var deviceCount DeviceCount
	count, operCount, err := h.db.GetVmTotalCount()
	deviceCount.Total = count
	deviceCount.Operate = operCount
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}
	c.JSON(http.StatusOK, deviceCount)
}

// FOR INFLUX-DB