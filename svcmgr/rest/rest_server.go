package rest

import (
	"nubes/common/mariadblayer"

	"github.com/gin-gonic/gin"
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
	// Monitoring
	GetDevicesMonitoring(c *gin.Context)
	AddDevicesMonitoring(c *gin.Context)
	DeleteDevicesMonitoring(c *gin.Context)
	// Login
	RegisterUser(c *gin.Context)
	LoginUserById(c *gin.Context)
	LoginUserByEmail(c *gin.Context)
	GetSession(c *gin.Context)
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
	router := gin.Default()

	// Middlewares
	router.Use(ErrorHandler)
	router.Use(CORSMiddlewre())
	h, _ := NewHandler(db)

	// Code
	router.GET("/v1/codes", h.GetCodes)
	router.POST("/v1/code", h.AddCode)
	router.DELETE("/v1/code/:id", h.DeleteCode)
	router.DELETE("/v1/codes", h.DeleteCodes)

	// SubCode
	router.GET("/v1/subcodes", h.GetSubCodes)
	router.POST("/v1/subcode", h.AddSubCode)
	router.DELETE("/v1/subcode/:id", h.DeleteSubCode)
	router.DELETE("/v1/subcodes", h.DeleteSubCodes)

	// Devices
	router.GET("/v1/devices/:type/:outFlag/list", h.GetDevicesByList)

	// Monitoring
	//router.GET("/v1/devices/monitoring", h.GetDevicesMonitoring)
	//router.POST("/v1/devices/monitoring", h.AddDevicesMonitoring)

	// Login
	router.POST("/register", h.RegisterUser)
	router.POST("/loginemail", h.LoginUserByEmail)
	router.POST("/login", h.LoginUserById)
	router.GET("/session", h.GetSession)

	// Test
	router.POST("/api/auth/login", h.LoginUserById)
	router.GET("/api/auth/check", h.GetSession)
	return router.Run(address)
}

