package mariadblayer

import (
	"nubes/common/models"
)

func (db *DBORM) GetAllDevices() (devices []models.Device, err error) {
	return devices, db.Find(&devices).Error
}

func (db *DBORM) AddDevice(device models.Device) (models.Device, error) {
	return device, db.Create(&device).Error
}
