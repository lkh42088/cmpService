package mcrest

import (
	"cmpService/common/mcmodel"
	"github.com/gin-gonic/gin"
	"net/http"
)

func getMcFilterAddress(c *gin.Context) {
	var msg []mcmodel.McFilterRule
	// todo
	c.JSON(http.StatusOK, msg)
}

func addMcFilterAddress(c *gin.Context) {
	var msg mcmodel.McFilterRule
	c.ShouldBindJSON(&msg)
	// todo
	c.JSON(http.StatusOK, msg)
}

func deleteMcFilterAddress(c *gin.Context) {
	var msg mcmodel.McFilterRule
	c.ShouldBindJSON(&msg)
	// todo
	c.JSON(http.StatusOK, msg)
}
