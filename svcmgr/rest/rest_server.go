package rest

import (
	"github.com/gin-gonic/gin"
	"nubes/common/mariadblayer"
)

type HandlerInterface interface {
	// Code
	GetCodes(c *gin.Context)
	AddCode(c *gin.Context)
	DeleteCode(c *gin.Context)
	DeleteCodes(c *gin.Context)
	// SubCode
	GetSubCodes(c *gin.Context)
	AddSubCode(c *gin.Context)
	DeleteSubCode(c *gin.Context)
	DeleteSubCodes(c *gin.Context)
	// Device
	GetDevicesByList(c *gin.Context)
}

type Handler struct {
	db mariadblayer.MariaDBLayer
}

func NewHandler(db *mariadblayer.DBORM) (*Handler, error) {
	h := new(Handler)
	h.db = db
	return h, nil
}

func RunAPI(address string, db *mariadblayer.DBORM) error {
	r := gin.Default()
	r.Use(CORSMiddlewre())
	h, _ := NewHandler(db)

	// Code
	r.GET("/v1/codes", h.GetCodes)
	r.POST("/v1/code", h.AddCode)
	r.DELETE("/v1/code/:id", h.DeleteCode)
	r.DELETE("/v1/codes", h.DeleteCodes)

	// SubCode
	r.GET("/v1/subcodes", h.GetSubCodes)
	r.POST("/v1/subcode", h.AddSubCode)
	r.DELETE("/v1/subcode/:id", h.DeleteSubCode)
	r.DELETE("/v1/subcodes", h.DeleteSubCodes)

	// Devices
	r.GET("/v1/devices/:type/:outFlag/list", h.GetDevicesByList)

	return r.Run(address)
}

func CORSMiddlewre() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Headers", "Contest-Type, Authorization, Origin")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}
