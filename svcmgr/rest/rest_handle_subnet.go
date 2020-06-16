package rest

import (
	"cmpService/common/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) AddSubnet(c *gin.Context) {
	if h.db == nil {
		return
	}

	var subnet models.SubnetMgmt
	c.ShouldBindJSON(&subnet)
	//fmt.Printf("%+v\n", subnet)

	err := h.db.AddSubnet(subnet)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":err.Error()})
		return
	}
	c.JSON(http.StatusOK, "OK")
}
