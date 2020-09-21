package rest

import (
	"cmpService/common/lib"
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
	// Code Setting
	GetCodesMenu(c *gin.Context)
	GetCodesMainByType(c *gin.Context)
	GetCodesSubByIdx(c *gin.Context)
	// Device
	GetDevicesByList(c *gin.Context)
	GetDevicesByCode(c *gin.Context)
	GetDevicesForSearch(c *gin.Context)
	GetDeviceWithoutJoin(c *gin.Context)
	// Device count
	GetDevicesTypeCount(c *gin.Context)
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
	CheckPassword(c *gin.Context)
	// User
	CheckDuplicatedUser(c *gin.Context)
	GetUsersPage(c *gin.Context)
	GetUsersWithSearchParamPage(c *gin.Context)
	RegisterUser(c *gin.Context)
	ModifyUser(c *gin.Context)
	UnRegisterUser(c *gin.Context)
	// Capcha
	GetCaptcha(c *gin.Context)
	// Companies
	CheckDuplicatedCompany(c *gin.Context)
	GetCompaniesPage(c *gin.Context)
	GetCompaniesPageWithSearchParam(c *gin.Context)
	GetCompaniesByName(c *gin.Context)
	GetUserDetailsByCpIdx(c *gin.Context)
	GetCompaniesWithUserByLikeCpName(c *gin.Context)
	GetCompanies(c *gin.Context)
	AddCompany(c *gin.Context)
	DeleteCompany(c *gin.Context)
	ModifyCompany(c *gin.Context)
	// Subnet
	GetSubnets(c *gin.Context)
	AddSubnet(c *gin.Context)
	UpdateSubnet(c *gin.Context)
	DeleteSubnets(c *gin.Context)

	//Micro Cloud
	GetMcServers(c *gin.Context)
	AddMcServer(c *gin.Context)
	DeleteMcServer(c *gin.Context)
	GetMcServersByCpIdx(c *gin.Context)
	UpdateMcServerResource(c *gin.Context)
	CheckNStoreSystemInfo(c *gin.Context)

	GetVmSnapshotConfig(c *gin.Context)
	AddVmSnapshot(c *gin.Context)
	DeleteVmSnapshot(c *gin.Context)
	UpdateVmSnapshot(c *gin.Context)

	GetMcVms(c *gin.Context)
	AddMcVm(c *gin.Context)
	DeleteMcVm(c *gin.Context)
	GetMcVmVnc(c *gin.Context)

	GetMcImages(c *gin.Context)
	GetMcImagesByServerIdx(c *gin.Context)

	AddMcNetwork(c *gin.Context)
	DeleteMcNetwork(c *gin.Context)
	GetMcNetworks(c *gin.Context)
	GetMcNetworksByServerIdx(c *gin.Context)

	UpdateMcVm(c *gin.Context)

	GetVmInterfaceTrafficByMac(c *gin.Context)
	GetVmInterfaceCpu(c *gin.Context)
	GetVmInterfaceMem(c *gin.Context)
	GetVmInterfaceDisk(c *gin.Context)

	GetVmWinInterface(c *gin.Context)
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

	router.Static("/image", "./svcmgr/files/img/")

	// Code
	router.GET(lib.SvcmgrApiCode, h.GetCodes)
	router.GET(lib.SvcmgrApiCode+"/:code/:subcode", h.GetCodeList)
	router.POST(lib.SvcmgrApiCode+"/create", h.AddCode)
	router.DELETE(lib.SvcmgrApiCode+"/delete/:id", h.DeleteCode)
	router.DELETE(lib.SvcmgrApiCode+"/delete", h.DeleteCodes)

	// SubCode
	router.GET(lib.SvcmgrApiSubCode, h.GetSubCodes)
	router.GET(lib.SvcmgrApiSubCode+"/:c_idx", h.GetSubCodeList)
	router.POST(lib.SvcmgrApiSubCode+"/create", h.AddSubCode)
	router.DELETE(lib.SvcmgrApiSubCode+"/delete/:id", h.DeleteSubCode)
	router.DELETE(lib.SvcmgrApiSubCode+"/delete", h.DeleteSubCodes)

	// Code Setting
	router.GET(lib.SvcmgrApiCodeSetting+"/menu/tag", h.GetCodesMenu)
	//client.get(`/v1/codesetting/code/${type}/${subType}`);
	router.GET(lib.SvcmgrApiCodeSetting+"/code/:type/:subType", h.GetCodesMainByType)
	router.GET(lib.SvcmgrApiCodeSetting+"/code/:type/:subType/:idx", h.GetCodesSubByIdx)

	// Devices
	router.GET(lib.SvcmgrApiDevice+"/:type/:value/:field", h.GetDevicesByCode)
	router.GET(lib.SvcmgrApiDevice+"/:type/:value", h.GetDevicesByCode)
	router.POST(lib.SvcmgrApiDevice+"/create/:type", h.AddDevice)
	router.PUT(lib.SvcmgrApiDevice+"/update/:type/:deviceCode", h.UpdateDevice)
	router.PUT(lib.SvcmgrApiDevices+"/update/:type", h.UpdateOutFlag)
	router.GET(lib.SvcmgrApiDevices+"/:type/:outFlag/list", h.GetDevicesByList)
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

	// Device Count
	// http://127.0.0.1:8081/v1/search/devices/count/server
	router.POST("/v1/search/count/devices/:type", h.GetDevicesTypeCount)

	// Monitoring
	//router.GET("/v1/devices/monitoring", h.GetDevicesMonitoring)
	//router.POST("/v1/devices/monitoring", h.AddDevicesMonitoring)

	// Login
	router.POST(lib.SvcmgrApiLogin+"/login", h.LoginUserById)
	router.POST(lib.SvcmgrApiLogin+"/login-send-email", h.LoginSendEmail)
	router.POST(lib.SvcmgrApiLogin+"/grouplogin", h.LoginGroupEmail)
	router.POST(lib.SvcmgrApiLogin+"/input_email", h.LoginUserById)
	router.POST(lib.SvcmgrApiLogin+"/confirm", h.LoginFrontConfirm)
	router.POST(lib.SvcmgrApiLogin+"/email_confirm", h.EmailConfirm)
	router.GET(lib.SvcmgrApiLogin+"/check", h.GetSession)
	router.POST(lib.SvcmgrApiLogin+"/logout", h.Logout)
	router.POST(lib.SvcmgrApiLogin+"/check-password", h.CheckPassword)

	pagingParam := "/:rows/:offset/:orderby/:order"

	// User
	router.GET(lib.SvcmgrApiUser+pagingParam, h.GetUsersPage)
	router.POST(lib.SvcmgrApiUser+"/get-user/:value", h.GetUserById)
	router.POST(lib.SvcmgrApiUser+"/page-with-search-param", h.GetUsersPageWithSearchParam)
	router.POST(lib.SvcmgrApiUser+"/register", h.RegisterUser)
	router.POST(lib.SvcmgrApiUser+"/modify", h.ModifyUser)
	router.POST(lib.SvcmgrApiUser+"/unregister", h.UnRegisterUser)
	router.POST(lib.SvcmgrApiUser+"/check-user", h.CheckDuplicatedUser)
	router.POST(lib.SvcmgrApiUser+"/fileUpload", h.UploadFileUser)

	//http.HandleFunc("/v1/users/fileUpload", uploadFile)
	//http.ListenAndServe(":4000", nil)

	// Auth
	router.GET(lib.SvcmgrApiPrefix+"/auth", h.GetAuth)
	// ReCAPTCHA
	router.POST(lib.SvcmgrApiPrefix+"/captcha", h.GetCaptcha)

	// Companies
	router.GET("/v1/companies-with-user-like-cpname/:cpName", h.GetCompaniesWithUserByLikeCpName)
	router.GET("/v1/companies/:cpName", h.GetCompaniesByCpName)
	router.GET("/v1/users-about-companies/:cpIdx", h.GetUserDetailsByCpIdx)
	router.GET("/v1/companies", h.GetCompanies)
	router.GET("/v1/customers/companies"+pagingParam, h.GetCompaniesPage)
	router.POST("/v1/company/get-company/:value", h.GetCompanyByName)

	router.POST("/v1/customers/register", h.AddCompany)
	router.POST("/v1/customers/unregister", h.DeleteCompany)
	router.POST("/v1/customers/check-company", h.CheckDuplicatedCompany)
	router.POST("/v1/customers/modify-company", h.ModifyCompany)
	router.POST("/v1/customers/companies/page-with-search-param", h.GetCompaniesPageWithSearchParam)

	// Subnet
	router.POST("/v1/subnet/create", h.AddSubnet)
	router.POST("/v1/subnet", h.GetSubnets)
	router.POST("/v1/subnet/update", h.UpdateSubnet)
	router.DELETE("/v1/subnet/:idx", h.DeleteSubnets)

	// Micro Cloud
	router.POST(lib.SvcmgrApiMicroServerRegister, h.AddMcServer)
	router.POST(lib.SvcmgrApiMicroServerUnRegister, h.DeleteMcServer)
	router.GET(lib.SvcmgrApiMicroServerSearchCompany+"/:cpIdx", h.GetMcServersByCpIdx)
	router.GET(lib.SvcmgrApiMicroServerPaging+pagingParam+"/:cpName", h.GetMcServers)

	router.POST(lib.SvcmgrApiMicroVmRegister, h.AddMcVm)
	router.POST(lib.SvcmgrApiMicroVmUnRegister, h.DeleteMcVm)
	router.GET(lib.SvcmgrApiMicroVmPaging+pagingParam+"/:cpName", h.GetMcVms)
	router.POST(lib.SvcmgrApiMicroVmUpdateFromMc, h.UpdateMcVmFromMc)
	router.GET(lib.SvcmgrApiMicroVmVnc+"/:target/:port", h.GetMcVmVnc)

	router.GET(lib.SvcmgrApiMicroImagePaging+pagingParam, h.GetMcImages)
	router.GET(lib.SvcmgrApiMicroImage+"/:serverIdx", h.GetMcImagesByServerIdx)

	router.POST(lib.SvcmgrApiMicroNetworkRegister, h.AddMcNetwork)
	router.POST(lib.SvcmgrApiMicroNetworkUnRegister, h.DeleteMcNetwork)
	router.GET(lib.SvcmgrApiMicroNetworkPaging+pagingParam, h.GetMcNetworks)
	router.GET(lib.SvcmgrApiMicroNetwork+"/:serverIdx", h.GetMcNetworksByServerIdx)

	router.GET(lib.SvcmgrApiMicroVmStats+"/:mac", GetVmInterfaceTrafficByMac)
	router.POST(lib.SvcmgrApiMicroServerResource, h.UpdateMcServerResource)

	router.POST(lib.SvcmgrApiMicroSystemInfo, h.CheckNStoreSystemInfo)
	// Micro Cloud GRAPH
	router.GET(lib.SvcmgrApiMicroVmSnapshotConfig+"/:serverIdx", h.GetVmSnapshotConfig)
	router.POST(lib.SvcmgrApiMicroVmAddSnapshot, h.AddVmSnapshot)
	router.POST(lib.SvcmgrApiMicroVmDeleteSnapshot, h.DeleteVmSnapshot)
	router.POST(lib.SvcmgrApiMicroVmUpdateSnapshot, h.UpdateVmSnapshot)
	router.POST(lib.SvcmgrApiMicroVmStatus, h.UpdateVmStatus)

	// Micro Cloud CPU
	router.GET(lib.SvcmgrApiMicroVmMonitor+"/cpu", GetVmInterfaceCpu)
	router.GET(lib.SvcmgrApiMicroVmMonitor+"/mem", GetVmInterfaceMem)
	router.GET(lib.SvcmgrApiMicroVmMonitor+"/disk", GetVmInterfaceDisk)

	// Micro Cloud Graph
	router.GET(lib.SvcmgrApiMicroVmGraph+"/:mac", GetVmWinInterface)
	router.GET(lib.SvcmgrApiMicroVmGraph+"/:mac/:currentStatus", GetVmWinInterface)

	return router.Run(address)
}
