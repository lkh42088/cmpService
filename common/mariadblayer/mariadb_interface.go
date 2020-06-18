package mariadblayer

import (
	"cmpService/common/models"
)

type MariaDBLayer interface {
	// Code
	GetAllCodes() ([]models.Code, error)
	GetCodeList(code string, subCode string) ([]models.Code, error)
	AddCode(code models.Code) (models.Code, error)
	DeleteCode(code models.Code) (models.Code, error)
	DeleteCodes() error

	// SubCode
	GetAllSubCodes() ([]models.SubCode, error)
	GetSubCodeList(cIdx []string) ([]models.SubCodeResponse, error)
	AddSubCode(subCode models.SubCode) (models.SubCode, error)
	DeleteSubCode(subCode models.SubCode) (models.SubCode, error)
	DeleteSubCodes() error

	// Devices
	GetAllDevicesServer(deviceType string, outFlag int) ([]models.DeviceServer, error)
	GetAllDevicesNetwork(deviceType string, outFlag int) ([]models.DeviceNetwork, error)
	GetAllDevicesPart(deviceType string, outFlag int) ([]models.DevicePart, error)
	GetDeviceServer(code string) (models.DeviceServer, error)
	GetDeviceNetwork(code string) (models.DeviceNetwork, error)
	GetDevicePart(code string) (models.DevicePart, error)
	GetDevicesServerForSearch(dc models.DeviceServer) ([]models.DeviceServerResponse, error)
	GetDevicesNetworkForSearch(dc models.DeviceNetwork) ([]models.DeviceNetworkResponse, error)
	GetDevicesPartForSearch(dc models.DevicePart) ([]models.DevicePartResponse, error)
	GetDeviceWithJoin(device string, field string, condition string) (interface{}, error)
	GetDeviceWithoutJoin(device string, code string) (interface{}, error)
	GetDeviceWithSplaJoin(spla []string) ([]models.Code, error)
	GetLastDeviceCodeInServer() (models.DeviceServer, error)
	GetLastDeviceCodeInNetwork() (models.DeviceNetwork, error)
	GetLastDeviceCodeInPart() (models.DevicePart, error)
	AddDeviceServer(server models.DeviceServer) (models.DeviceServer, error)
	AddDeviceNetwork(network models.DeviceNetwork) (models.DeviceNetwork, error)
	AddDevicePart(part models.DevicePart) (models.DevicePart, error)
	AddDevice(data interface{}, tableName string) error
	UpdateDevice(device interface{}, tableName string, deviceCode string) (interface{}, error)
	UpdateOutFlag(codes []string, tableName string, flag int) error
	DeleteAllDevicesServer() error
	DeleteAllDevicesNetwork() error
	DeleteAllDevicesPart() error
	DeleteDeviceServer(server models.DeviceServer) (models.DeviceServer, error)
	DeleteDeviceNetwork(network models.DeviceNetwork) (models.DeviceNetwork, error)
	DeleteDevicePart(part models.DevicePart) (models.DevicePart, error)

	// Page
	GetDevicesServerForPage(creteria models.PageCreteria) (models.DeviceServerPage, error)
	GetDevicesNetworkForPage(creteria models.PageCreteria) (models.DeviceNetworkPage, error)
	GetDevicesPartForPage(creteria models.PageCreteria) (models.DevicePartPage, error)

	GetDevicesServerWithJoin(creteria models.PageCreteria) (models.DeviceServerPage, error)
	GetDevicesNetworkWithJoin(creteria models.PageCreteria) (models.DeviceNetworkPage, error)
	GetDevicesPartWithJoin(creteria models.PageCreteria) (models.DevicePartPage, error)

	GetDevicesServerSearchWithJoin(creteria models.PageCreteria,
		dc models.DeviceServer) (models.DeviceServerPage, error)
	GetDevicesNetworkSearchWithJoin(creteria models.PageCreteria,
		dc models.DeviceNetwork) (models.DeviceNetworkPage, error)
	GetDevicesPartSearchWithJoin(creteria models.PageCreteria,
		dc models.DevicePart) (models.DevicePartPage, error)

	// Comment
	GetAllComments() ([]models.DeviceComment, error)
	GetComments(code string) ([]models.DeviceComment, error)
	GetCommentByIdx(idx int) (models.DeviceComment, error)
	UpdateComment(comment models.DeviceComment) error
	AddComment(comment models.DeviceComment) error
	DeleteAllComments() error
	DeleteComments(idx int) error

	// Log
	GetAllLogs() ([]models.DeviceLog, error)
	GetLogs(code string) ([]models.DeviceLog, error)
	GetLogByIdx(idx int) (models.DeviceLog, error)
	UpdateLog(field string, change string, comment models.DeviceLog) error
	AddLog(comment models.DeviceLog) error
	DeleteAllLogs() error
	DeleteLog(idx int) error

	// User, Customer, Auth
	GetCompaniesByName(name string) ([]models.CompanyResponse, error)
	GetUserByUserId(userId string) (models.User, error)
	AddUserMember(user models.User) error
	AddCompany(company models.Company) (models.Company, error)
	AddAuth(auth models.Auth) error
	DeleteAllUserMember() error
	DeleteAllCompany() error
	DeleteAllAuth() error

	// User
	AddUser(user models.User) (models.User, error)
	DeleteUser(user models.User) (models.User, error)
	GetAllUsers() ([]models.User, error)
	GetUserById(id string) (models.User, error)
	GetUserByEmail(id string) (models.User, error)
	GetUsersPage(paging models.Pagination) (models.UserPage, error)
	GetCompaniesPage(paging models.Pagination) (models.CompanyPage, error)

	// User Email Authentication
	GetAllUserEmailAuth() (objs []models.UserEmailAuth, err error)
	GetUserEmailAuthByIdAndEmail(id, email string) (userEmailAuth models.UserEmailAuth, err error)
	GetUserEmailAuthByIdAndStore(id, store string) (userEmailAuth models.UserEmailAuth, err error)
	AddUserEmailAuth(obj models.UserEmailAuth) (models.UserEmailAuth, error)
	DeleteUserEmailAuth(obj models.UserEmailAuth) (models.UserEmailAuth, error)
	DeleteUserEmailAuthByUserId(id string) ([]models.UserEmailAuth, error)
	UpdateUserEmailAuth(obj models.UserEmailAuth) (models.UserEmailAuth, error)
}
