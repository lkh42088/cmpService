package mariadblayer

import "nubes/common/models"

type DBLayer interface {
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
}

