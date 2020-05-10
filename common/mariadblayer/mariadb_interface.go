package mariadblayer

import "nubes/common/models"

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
	GetDeviceServer(deviceType string, idx int) ([]models.DeviceServer, error)
	GetDeviceNetwork(deviceType string, idx int) ([]models.DeviceNetwork, error)
	GetDevicePart(deviceType string, idx int) ([]models.DevicePart, error)
	AddDeviceServer(server models.DeviceServer)(models.DeviceServer, error)
	AddDeviceNetwork(server models.DeviceNetwork)(models.DeviceNetwork, error)
	AddDevicePart(server models.DevicePart)(models.DevicePart, error)
	DeleteAllDevicesServer() error
	DeleteAllDevicesNetwork() error
	DeleteAllDevicesPart() error
	DeleteDeviceServer(server models.DeviceServer) (models.DeviceServer, error)
	DeleteDeviceNetwork(server models.DeviceNetwork) (models.DeviceNetwork, error)
	DeleteDevicePart(server models.DevicePart) (models.DevicePart, error)

	// User
	AddUser(user models.User) (models.User, error)
	DeleteUser(user models.User) (models.User, error)
	GetAllUsers() ([]models.User, error)
	GetUserById(id string) (models.User, error)
	GetUserByEmail(id string) (models.User, error)
}

