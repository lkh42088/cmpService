package rest

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		webServerUrl := fmt.Sprintf("http://%s", WebServerAddress)
		c.Writer.Header().Set("Access-Control-Allow-Origin", webServerUrl)
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			fmt.Println("CORSMiddleware: 204 error!", WebServerAddress, webServerUrl)
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

func ErrorHandler(c *gin.Context) {
	c.Next()
	if len(c.Errors) > 0 {
		fmt.Println("ErrorHandler:", c.Errors)
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": c.Errors,
		})
	}
}
