package mariadblayer

import (
	"nubes/common/models"
)

type MariaDBLayer interface {
	// Code
	GetAllCodes() ([]models.Code, error)
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
	GetAllDevicesComment(deviceType string, outFlag int) ([]models.DeviceComment, error)
	GetDeviceServer(deviceType string, idx int) ([]models.DeviceServer, error)
	GetDeviceNetwork(deviceType string, idx int) ([]models.DeviceNetwork, error)
	GetDevicePart(deviceType string, idx int) ([]models.DevicePart, error)
	GetDeviceComment(deviceType string, idx int) ([]models.DeviceComment, error)
	GetDevicesServerForPage(creteria models.PageCreteria) (models.DeviceServerPage, error)
	GetDevicesNetworkForPage(creteria models.PageCreteria) (models.DeviceNetworkPage, error)
	GetDevicesPartForPage(creteria models.PageCreteria) (models.DevicePartPage, error)
	GetDeviceWithCondition(device string, field string, condition string) (interface{}, error)
	AddDeviceServer(server models.DeviceServer)(models.DeviceServer, error)
	AddDeviceNetwork(network models.DeviceNetwork)(models.DeviceNetwork, error)
	AddDevicePart(part models.DevicePart)(models.DevicePart, error)
	AddDeviceComment(comment models.DeviceComment)(models.DeviceComment, error)
	DeleteAllDevicesServer() error
	DeleteAllDevicesNetwork() error
	DeleteAllDevicesPart() error
	DeleteAllDevicesComment() error
	DeleteDeviceServer(server models.DeviceServer) (models.DeviceServer, error)
	DeleteDeviceNetwork(network models.DeviceNetwork) (models.DeviceNetwork, error)
	DeleteDevicePart(part models.DevicePart) (models.DevicePart, error)
	DeleteDeviceComment(comment models.DeviceComment) (models.DeviceComment, error)

	// User
	AddUser(user models.User) (models.User, error)
	DeleteUser(user models.User) (models.User, error)
	GetAllUsers() ([]models.User, error)
	GetUserById(id string) (models.User, error)
	GetUserByEmail(id string) (models.User, error)
}

