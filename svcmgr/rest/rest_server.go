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
	UpdateComment(c *gin.Context)
	DeleteCommentByIdx(c *gin.Context)
	// Log
	GetLogsByCode(c *gin.Context)
	AddLog(c *gin.Context)
	UpdateLog(c *gin.Context)
	DeleteLogByIdx(c *gin.Context)
	// Page
	GetDeviceForPage(c *gin.Context)
	GetDevicesForPageSearch(c *gin.Context)
	// Monitoring
	GetDevicesMonitoring(c *gin.Context)
	AddDevicesMonitoring(c *gin.Context)
	DeleteDevicesMonitoring(c *gin.Context)
	// Login
	GetSession(c *gin.Context)
	LoginUserById(c *gin.Context)
	LoginSendEmail(c *gin.Context)
	Logout(c *gin.Context)
	LoginUserByEmail(c *gin.Context)
	EmailConfirm(c *gin.Context)
	// User
	CheckDuplicatedUser(c *gin.Context)
	GetUsersPage(c *gin.Context)
	GetUsersWithSearchParamPage(c *gin.Context)
	RegisterUser(c *gin.Context)
	UnRegisterUser(c *gin.Context)
	// Companies
	CheckDuplicatedCompany(c *gin.Context)
	GetCompaniesPage(c *gin.Context)
	GetCompaniesByName(c *gin.Context)
	GetUserDetailsByCpIdx(c *gin.Context)
	GetCompaniesWithUserByLikeCpName(c *gin.Context)
	GetCompanies(c *gin.Context)
	AddCompany(c *gin.Context)
	DeleteCompany(c *gin.Context)
	// Subnet
	GetSubnet(c *gin.Context)
	AddSubnet(c *gin.Context)
}

type Handler struct {
	db mariadblayer.MariaDBLayer
}

func NewHandler(db *mariadblayer.DBORM) (*Handler, error) {
	h := new(Handler)
	h.db = db
	return h, nil
}

const (
	ApiPrefix  = "/v1"
	ApiCode    = ApiPrefix + "/codes"
	ApiSubCode = ApiPrefix + "/subcodes"
	ApiDevice  = ApiPrefix + "/device"
	ApiDevices = ApiPrefix + "/devices"
	ApiLogin   = ApiPrefix + "/auth"
	ApiCompany = ApiPrefix + "/companies"
	ApiUser    = ApiPrefix + "/users"
)

func RunAPI(address string, db *mariadblayer.DBORM) error {
	router := gin.Default()

	// Middlewares
	router.Use(ErrorHandler)
	router.Use(CORSMiddleware())
	h, _ := NewHandler(db)

	// Code
	router.GET(ApiCode, h.GetCodes)
	router.GET(ApiCode+"/:code/:subcode", h.GetCodeList)
	router.POST(ApiCode+"/create", h.AddCode)
	router.DELETE(ApiCode+"/delete/:id", h.DeleteCode)
	router.DELETE(ApiCode+"/delete", h.DeleteCodes)

	// SubCode
	router.GET(ApiSubCode, h.GetSubCodes)
	router.GET(ApiSubCode+"/:c_idx", h.GetSubCodeList)
	router.POST(ApiSubCode+"/create", h.AddSubCode)
	router.DELETE(ApiSubCode+"/delete/:id", h.DeleteSubCode)
	router.DELETE(ApiSubCode+"/delete", h.DeleteSubCodes)

	// Devices
	router.GET(ApiDevice+"/:type/:value/:field", h.GetDevicesByCode)
	router.GET(ApiDevice+"/:type/:value", h.GetDevicesByCode)
	router.POST(ApiDevice+"/create/:type", h.AddDevice)
	router.PUT(ApiDevice+"/update/:type/:deviceCode", h.UpdateDevice)
	router.PUT(ApiDevices+"/update/:type", h.UpdateOutFlag)
	router.GET(ApiDevices+"/:type/:outFlag/list", h.GetDevicesByList)
	router.GET("/v1/raw/device/:type/:value", h.GetDeviceWithoutJoin)

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
	// API_ROUTE/page/ server / 0 / 1000 / 110 / deviceCode / 1
	//router.GET("/v1/page/:type/:outFlag/:size/:checkcnt/:order/:dir", h.GetDevicesForPage)
	//router.GET("/v1/page/:type/:outFlag/:size/:checkcnt", h.GetDevicesForPage2)
	//router.GET("/v1/page/:type/:outFlag/:row/:page/:order/:dir/:offsetPage", h.GetDevicesForPage)
	//todo outFlag 삭제 필요
	router.POST("/v1/search/devices/:type/:row/:page/:order/:dir/:offsetPage",
		h.GetDevicesForPageSearch)

	// LOG
	router.GET("/v1/log/device/:value", h.GetDevicesByLog)

	// Monitoring
	//router.GET("/v1/devices/monitoring", h.GetDevicesMonitoring)
	//router.POST("/v1/devices/monitoring", h.AddDevicesMonitoring)

	// Login
	router.POST(ApiLogin+"/login", h.LoginUserById)
	router.POST(ApiLogin+"/login-send-email", h.LoginSendEmail)
	router.POST(ApiLogin+"/grouplogin", h.LoginGroupEmail)
	router.POST(ApiLogin+"/input_email", h.LoginUserById)
	router.POST(ApiLogin+"/confirm", h.LoginFrontConfirm)
	router.POST(ApiLogin+"/email_confirm", h.EmailConfirm)
	router.GET(ApiLogin+"/check", h.GetSession)
	router.POST(ApiLogin+"/logout", h.Logout)

	pagingParam := "/:rows/:offset/:orderby/:order"

	// User
	router.GET(ApiUser+pagingParam, h.GetUsersPage)
	router.POST(ApiUser+"/page-with-search-param", h.GetUsersWithSearchParamPage)
	router.POST(ApiUser+"/register", h.RegisterUser)
	router.POST(ApiUser+"/unregister", h.UnRegisterUser)
	router.POST(ApiUser+"/check-user", h.CheckDuplicatedUser)

	// Companies
	router.GET("/v1/companies-with-user-like-cpname/:cpName", h.GetCompaniesWithUserByLikeCpName)
	router.GET("/v1/companies/:cpName", h.GetCompaniesByCpName)
	router.GET("/v1/users-about-companies/:cpIdx", h.GetUserDetailsByCpIdx)
	router.GET("/v1/companies", h.GetCompanies)
	router.GET("/v1/customers/companies"+pagingParam, h.GetCompaniesPage)
	router.POST("/v1/customers/register", h.AddCompany)
	router.POST("/v1/customers/unregister", h.DeleteCompany)
	router.POST("/v1/customers/check-company", h.CheckDuplicatedCompany)

	// Subnet
	router.POST("/v1/subnet/create", h.AddSubnet)
	router.GET("/v1/subnet"+pagingParam, h.GetSubnet)

	return router.Run(address)
}
