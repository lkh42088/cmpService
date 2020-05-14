package mariadblayer

import (
	"errors"
	"fmt"
	"nubes/common/models"
)

func (db *DBORM) GetAllDevicesServer(deviceType string, outFlag int) (devices []models.DeviceServer,
	err error) {
	return devices, db.Where("out_flag=?", outFlag).Find(&devices).Error
}

func (db *DBORM) GetAllDevicesNetwork(deviceType string, outFlag int) (devices []models.DeviceNetwork,
	err error) {
	return devices, db.Where("out_flag=?", outFlag).Find(&devices).Error
}

func (db *DBORM) GetAllDevicesPart(deviceType string, out_flag int) (devices []models.DevicePart,
	err error) {
	return devices, db.Where("out_flag=?", out_flag).Find(&devices).Error
}

func (db *DBORM) GetAllDevicesComment(deviceType string, out_flag int) (devices []models.DeviceComment,
	err error) {
	return devices, db.Where("out_flag=?", out_flag).Find(&devices).Error
}

func (db *DBORM) GetDeviceServer(deviceType string, idx int) (device []models.DeviceServer,
	err error) {
	return device, db.Where("idx=?", idx).Find(&device).Error
}

func (db *DBORM) GetDeviceNetwork(deviceType string, idx int) (device []models.DeviceNetwork,
	err error) {
	return device, db.Where("idx=?", idx).Find(&device).Error
}

func (db *DBORM) GetDevicePart(deviceType string, idx int) (device []models.DevicePart,
	err error) {
	return device, db.Where("idx=?", idx).Find(&device).Error
}

func (db *DBORM) GetDeviceComment(deviceType string, idx int) (device []models.DeviceComment,
	err error) {
	return device, db.Where("idx=?", idx).Find(&device).Error
}

func (db *DBORM) AddDeviceServer(device models.DeviceServer) (models.DeviceServer, error) {
	return device, db.Create(&device).Error
}

func (db *DBORM) AddDeviceNetwork(device models.DeviceNetwork) (models.DeviceNetwork, error) {
	return device, db.Create(&device).Error
}

func (db *DBORM) AddDevicePart(device models.DevicePart) (models.DevicePart, error) {
	return device, db.Create(&device).Error
}

func (db *DBORM) AddDeviceComment(comment models.DeviceComment) (models.DeviceComment, error) {
	return comment, db.Create(&comment).Error
}

func (db *DBORM) DeleteAllDevicesServer() error {
	return db.Delete(&models.DeviceServer{}).Error
}

func (db *DBORM) DeleteAllDevicesPart() error {
	return db.Delete(&models.DevicePart{}).Error
}

func (db *DBORM) DeleteAllDevicesNetwork() error {
	return db.Delete(&models.DeviceNetwork{}).Error
}

func (db *DBORM) DeleteAllDevicesComment() error {
	return db.Delete(&models.DeviceComment{}).Error
}

func (db *DBORM) DeleteDeviceServer(sd models.DeviceServer) (models.DeviceServer, error) {
	return sd, db.Delete(&sd).Error
}

func (db *DBORM) DeleteDeviceNetwork(nd models.DeviceNetwork) (models.DeviceNetwork, error) {
	return nd, db.Delete(&nd).Error
}

func (db *DBORM) DeleteDevicePart(pd models.DevicePart) (models.DevicePart, error) {
	return pd, db.Delete(&pd).Error
}

func (db *DBORM) DeleteDeviceComment(dc models.DeviceComment) (models.DeviceComment, error) {
	return dc, db.Delete(&dc).Error
}

func (db *DBORM) GetDeviceWithCondition(device string, field string, condition string) (
	interface{}, error) {
	dbField := ConvertToColumn(field)
	where := GetWhereString(dbField)
	fmt.Println(where)
	var dc interface{}
	if GetTableConfig(&dc, device) == false {
		return nil, errors.New("[Error] Need to device selection.")
	}
	return dc, db.Where(where, condition).Find(dc).Error

}

func (db *DBORM) AddDevice(data interface{}, device string) error {
	return db.Table(device).Create(data).Error
}

func GetWhereString(field string) string {
	return field + " = ?"
}

func GetTableConfig(data *interface{}, device string) bool {
	switch device {
	case "server":
		*data = &[]models.DeviceServer{}
	case "network":
		*data = &[]models.DeviceNetwork{}
	case "part":
		*data = &[]models.DevicePart{}
	default:
		return false

	}
	return true
}
