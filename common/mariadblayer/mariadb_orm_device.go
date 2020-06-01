package mariadblayer

import (
	"cmpService/common/models"
	"errors"
)

const outFlagField = "out_flag"

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

func (db *DBORM) GetDeviceWithJoin(device string, field string, condition string) (
	interface{}, error) {
	dbField := ConvertToColumn(field)
	where := GetWhereString(dbField)
	var dc interface{}
	if GetTableConfig(&dc, device) == false {
		return nil, errors.New("[Error] Need to device selection.\n")
	}

	manufacture, deviceType, tableName := GetDeviceQuery(device)
	var selectString string
	var sizeQueryString string
	if device == "part" {
		selectString = PageSelectQuery
		sizeQueryString = ""
	} else {
		selectString = SizeSelectQuery+","+PageSelectQuery
		sizeQueryString = SizeJoinQuery
	}
	return dc, db.
		//Debug().
		Select(selectString).
		Table(tableName).
		Joins(manufacture).
		Joins(ModelJoinQuery).
		Joins(deviceType).
		Joins(OwnershipJoinQuery).
		Joins(OwnershipDivJoinQuery).
		Joins(IdcJoinQuery).
		Joins(RackJoinQuery).
		Joins(sizeQueryString).
		Joins(CompanyLeftJoinQuery).
		Joins(OwnerCompanyLeftJoinQuery).
		Where(where, condition).
		Find(dc).Error

}

func (db *DBORM) GetDeviceWithoutJoin(device string, code string) (
	interface{}, error) {
	var dc interface{}
	where := GetWhereString(defaultFieldName)
	if GetTableConfig(&dc, device) == false {
		return nil, errors.New("[Error] Need to device selection.\n")
	}

	_, _, tableName := GetDeviceQuery(device)
	return dc, db.Table(tableName).Where(where, code).Find(dc).Error
}

func (db *DBORM) GetDevicesServerForSearch(dc models.DeviceServer) (ds []models.DeviceServerResponse, err error) {
	return ds, db.
		Debug().
		Select(ServerSelectQuery).
		Table(ServerRawTable).
		Joins(ManufactureServerNoAliasJoinQuery).
		Joins(ModelServerNoAliasJoinQuery).
		Joins(DeviceTypeServerNoAliasJoinQuery).
		Joins(OwnershipServerNoAliasJoinQuery).
		Joins(OwnershipDivServerNoAliasJoinQuery).
		Joins(IdcServerNoAliasJoinQuery).
		Joins(RackServerNoAliasJoinQuery).
		Joins(SizeServerNoAliasJoinQuery).
		Joins(CompanyServerNoAliasLeftJoinQuery).
		Joins(OwnerCompanyServerNoAliasLeftJoinQuery).
		Where(&dc).Find(&ds).Error
}

func (db *DBORM) GetDevicesNetworkForSearch(dc models.DeviceNetwork) (ds []models.DeviceNetworkResponse, err error) {
	return ds, db.
		Debug().
		Select(NetworkSelectQuery).
		Table(NetworkRawTable).
		Joins(ManufactureNetworkNoAliasJoinQuery).
		Joins(ModelNetworkNoAliasJoinQuery).
		Joins(DeviceTypeNetworkNoAliasJoinQuery).
		Joins(OwnershipNetworkNoAliasJoinQuery).
		Joins(OwnershipDivNetworkNoAliasJoinQuery).
		Joins(IdcNetworkNoAliasJoinQuery).
		Joins(RackNetworkNoAliasJoinQuery).
		Joins(SizeNetworkNoAliasJoinQuery).
		Joins(CompanyNetworkNoAliasLeftJoinQuery).
		Joins(OwnerCompanyNetworkNoAliasLeftJoinQuery).
		Where(&dc).Find(&ds).Error
}

func (db *DBORM) GetDevicesPartForSearch(dc models.DevicePart) (ds []models.DevicePartResponse, err error) {
	return ds, db.
		Debug().
		Select(PartSelectQuery).
		Table(PartRawTable).
		Joins(ManufacturePartNoAliasJoinQuery).
		Joins(ModelPartNoAliasJoinQuery).
		Joins(DeviceTypePartNoAliasJoinQuery).
		Joins(OwnershipPartNoAliasJoinQuery).
		Joins(OwnershipDivPartNoAliasJoinQuery).
		Joins(IdcPartNoAliasJoinQuery).
		Joins(RackPartNoAliasJoinQuery).
		Joins(CompanyPartNoAliasLeftJoinQuery).
		Joins(OwnerCompanyPartNoAliasLeftJoinQuery).
		Where(&dc).Find(&ds).Error
}

func (db *DBORM) GetLastDeviceCodeInServer() (ds models.DeviceServer, err error) {
	return ds, db.Order("device_code DESC").Last(&ds).Error
}

func (db *DBORM) GetLastDeviceCodeInNetwork() (ds models.DeviceNetwork, err error) {
	return ds, db.Order("device_code DESC").Last(&ds).Error
}

func (db *DBORM) GetLastDeviceCodeInPart() (ds models.DevicePart, err error) {
	return ds, db.Order("device_code DESC").Last(&ds).Error
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

func (db *DBORM) AddDevice(data interface{}, device string)  error {
	return db.Table(device).Create(data).Error
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

func (db *DBORM) DeleteDeviceServer(sd models.DeviceServer) (models.DeviceServer, error) {
	return sd, db.Delete(&sd).Error
}

func (db *DBORM) DeleteDeviceNetwork(nd models.DeviceNetwork) (models.DeviceNetwork, error) {
	return nd, db.Delete(&nd).Error
}

func (db *DBORM) DeleteDevicePart(pd models.DevicePart) (models.DevicePart, error) {
	return pd, db.Delete(&pd).Error
}

func (db *DBORM) UpdateDevice(data interface{}, device string, idx string) error {
	return db.Table(device).Where("device_code = ?", idx).Update(data).Error
}

// Update OutFlag
func (db *DBORM) UpdateOutFlag(codes []string, device string, flag int) error {
	return db.Debug().Table(device).Where("device_code IN (?)", codes).Update(outFlagField, flag).Error
}

func GetWhereString(field string) string {
	return field + " = ?"
}

func GetTableConfig(data *interface{}, device string) bool {
	switch device {
	case "server":
		*data = &[]models.DeviceServerResponse{}
	case "network":
		*data = &[]models.DeviceNetworkResponse{}
	case "part":
		*data = &[]models.DevicePartResponse{}
	default:
		return false

	}
	return true
}

func GetDeviceQuery(device string) (string, string, string) {
	var manufacture string
	var deviceType string
	var tableName string
	switch device {
	case "server":
		manufacture = ManufactureServerJoinQuery
		deviceType = DeviceTypeServerJoinQuery
		tableName = ServerTable
	case "network":
		manufacture = ManufactureNetworkJoinQuery
		deviceType = DeviceTypeNetworkJoinQuery
		tableName = NetworkTable
	case "part":
		manufacture = ManufacturePartJoinQuery
		deviceType = DeviceTypePartJoinQuery
		tableName = PartTable
	}
	return manufacture, deviceType, tableName
}
