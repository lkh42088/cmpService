package mcrest

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CheckHealth (c *gin.Context) {
	c.JSON(http.StatusOK, "Health Check OK!")
}