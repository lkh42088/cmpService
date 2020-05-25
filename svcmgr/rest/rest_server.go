package rest

import (
	"cmpService/common/mariadblayer"
	"github.com/gin-gonic/gin"
)

type HandlerInterface interface {
	// Code
	GetCodes(c *gin.Context)
	GetCodeList(c *gin.Context)
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
	GetDevicesByCode(c *gin.Context)
	AddDevice(c *gin.Context)
	// Comment
	GetCommentsByCode(c *gin.Context)
	AddComment(c *gin.Context)
	UpdateComment( c *gin.Context)
	DeleteCommentByIdx(c *gin.Context)
	// Log
	GetLogsByCode(c *gin.Context)
	AddLog(c *gin.Context)
	UpdateLog( c *gin.Context)
	DeleteLogByIdx(c *gin.Context)
	// Page
	GetDeviceForPage(c *gin.Context)
	// Monitoring
	GetDevicesMonitoring(c *gin.Context)
	AddDevicesMonitoring(c *gin.Context)
	DeleteDevicesMonitoring(c *gin.Context)
	// Login
	RegisterUser(c *gin.Context)
	LoginUserById(c *gin.Context)
	Logout(c *gin.Context)
	LoginUserByEmail(c *gin.Context)
	GetSession(c *gin.Context)
	EmailConfirm(c *gin.Context)
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
	router.Use(CORSMiddleware())
	h, _ := NewHandler(db)

	// Code
	router.GET("/v1/codes", h.GetCodes)
	router.GET("/v1/codes/:code/:subcode", h.GetCodeList)
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
	router.GET("/v1/device/:type/:value/:field", h.GetDevicesByCode)
	router.GET("/v1/device/:type/:value", h.GetDevicesByCode)
	router.POST("/v1/device/:type", h.AddDevice)
	router.PUT("/v1/device/:type/:outFlag", h.UpdateOutFlag)

	// Comment
	router.GET("/v1/comments/:devicecode", h.GetCommentsByCode)
	router.POST("/v1/comment/create", h.AddComment)
	router.PUT("/v1/comment/update", h.UpdateComment)
	router.DELETE("/v1/comment/delete/:userid/:commentidx", h.DeleteCommentByIdx)

	// log
	router.GET("/v1/logs/:devicecode", h.GetLogsByCode)
	//router.POST("/v1/log/create/:devicecode/:comment/:userid", h.AddLog)
	//router.PUT("/v1/log/update/:workcode/:field/:change/:userid/:logidx", h.UpdateLog)
	router.DELETE("/v1/logs/delete/:logidx", h.DeleteLogByIdx)


	// Page
	router.GET("/v1/page/:type/:outFlag/:size/:checkcnt/:order/:dir", h.GetDevicesForPage)
	router.GET("/v1/page/:type/:outFlag/:size/:checkcnt", h.GetDevicesForPage)

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
	router.POST("/api/auth/logout", h.Logout)
	router.POST("/api/auth/register", h.RegisterUser)
	router.GET("/api/auth/check", h.GetSession)
	router.GET("/api/auth/emailconfirm/:secret", h.EmailConfirm)
	return router.Run(address)
}

