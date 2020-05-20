package mariadblayer

import (
	"cmpService/common/models"
)

type MariaDBLayer interface {
	// Code
	GetAllCodes() ([]models.Code, error)
	GetCodeList(code string, subCode string)([]models.Code, error)
	AddCode(code models.Code) (models.Code, error)
	DeleteCode(code models.Code) (models.Code, error)
	DeleteCodes() error
	// SubCode
	GetAllSubCodes() ([]models.SubCode, error)
	AddSubCode(subCode models.SubCode) (models.SubCode, error)
	DeleteSubCode(subCode models.SubCode) (models.SubCode, error)
	DeleteSubCodes() error
	// Devices
	GetAllDevicesServer(deviceType string, outFlag int) ([]models.DeviceServer, error)
	GetAllDevicesNetwork(deviceType string, outFlag int) ([]models.DeviceNetwork, error)
	GetAllDevicesPart(deviceType string, outFlag int) ([]models.DevicePart, error)
	GetDeviceServer(deviceType string, idx int) ([]models.DeviceServer, error)
	GetDeviceNetwork(deviceType string, idx int) ([]models.DeviceNetwork, error)
	GetDevicePart(deviceType string, idx int) ([]models.DevicePart, error)
	GetDevicesServerForPage(creteria models.PageCreteria) (models.DeviceServerPage, error)
	GetDevicesNetworkForPage(creteria models.PageCreteria) (models.DeviceNetworkPage, error)
	GetDevicesPartForPage(creteria models.PageCreteria) (models.DevicePartPage, error)
	GetDeviceWithCondition(device string, field string, condition string) (interface{}, error)
	AddDeviceServer(server models.DeviceServer)(models.DeviceServer, error)
	AddDeviceNetwork(network models.DeviceNetwork)(models.DeviceNetwork, error)
	AddDevicePart(part models.DevicePart)(models.DevicePart, error)
	AddDevice(data interface{}, tableName string) error
	DeleteAllDevicesServer() error
	DeleteAllDevicesNetwork() error
	DeleteAllDevicesPart() error
	DeleteDeviceServer(server models.DeviceServer) (models.DeviceServer, error)
	DeleteDeviceNetwork(network models.DeviceNetwork) (models.DeviceNetwork, error)
	DeleteDevicePart(part models.DevicePart) (models.DevicePart, error)

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

	// User
	AddUser(user models.User) (models.User, error)
	DeleteUser(user models.User) (models.User, error)
	GetAllUsers() ([]models.User, error)
	GetUserById(id string) (models.User, error)
	GetUserByEmail(id string) (models.User, error)
}

