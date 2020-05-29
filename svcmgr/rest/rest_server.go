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
	GetSubCodeList(c *gin.Context)
	AddSubCode(c *gin.Context)
	DeleteSubCode(c *gin.Context)
	DeleteSubCodes(c *gin.Context)
	// Device
	GetDevicesByList(c *gin.Context)
	GetDevicesByCode(c *gin.Context)
	GetDevicesForSearch(c *gin.Context)
	GetDeviceWithoutJoin(c *gin.Context)
	AddDevice(c *gin.Context)
	UpdateDevice(c *gin.Context)
	UpdateOutFlag(c *gin.Context)
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
	// Company
	GetCompaniesByName(c *gin.Context)
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
	router.POST("/v1/code/create", h.AddCode)
	router.DELETE("/v1/code/delete/:id", h.DeleteCode)
	router.DELETE("/v1/codes/delete", h.DeleteCodes)

	// SubCode
	router.GET("/v1/subcodes", h.GetSubCodes)
	router.GET("/v1/subcodes/:c_idx", h.GetSubCodeList)
	router.POST("/v1/subcode/create", h.AddSubCode)
	router.DELETE("/v1/subcode/delete/:id", h.DeleteSubCode)
	router.DELETE("/v1/subcodes/delete", h.DeleteSubCodes)

	// Devices
	router.GET("/v1/devices/:type/:outFlag/list", h.GetDevicesByList)
	router.GET("/v1/device/:type/:value/:field", h.GetDevicesByCode)
	router.GET("/v1/device/:type/:value", h.GetDevicesByCode)
	router.GET("/v1/devices/:type", h.GetDevicesForSearch)
	router.GET("/v1/raw/device/:type/:value", h.GetDeviceWithoutJoin)
	router.POST("/v1/device/create/:type", h.AddDevice)
	router.PUT("/v1/device/update/:type/:idx", h.UpdateDevice)
	router.PUT("/v1/devices/update/:type/:outFlag", h.UpdateOutFlag)

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

	// Company
	router.GET("/v1/companies/:name", h.GetCompaniesByName)

	// Monitoring
	//router.GET("/v1/devices/monitoring", h.GetDevicesMonitoring)
	//router.POST("/v1/devices/monitoring", h.AddDevicesMonitoring)

	// Login
	router.POST("/api/auth/login", h.LoginUserById)
	router.POST("/api/auth/grouplogin", h.LoginGroupEmail)
	router.POST("/api/auth/input_email", h.LoginUserById)
	router.POST("/api/auth/confirm", h.LoginFrontConfirm)
	router.POST("/api/auth/email_confirm", h.EmailConfirm)
	router.GET("/api/auth/check", h.GetSession)
	router.POST("/api/auth/logout", h.Logout)
	router.POST("/api/auth/register", h.RegisterUser)
	router.POST("/api/auth/unregister", h.UnRegisterUser)

	return router.Run(address)
}

