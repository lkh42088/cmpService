package mariadblayer

import (
	"cmpService/common/models"
	"errors"
	_ "fmt"
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

func (db *DBORM) GetDeviceServer(code string) (device models.DeviceServer, err error) {
	return device, db.Where(models.DeviceServer{
		DeviceCommon: models.DeviceCommon{DeviceCode: code},
	}).Find(&device).Error
}

func (db *DBORM) GetDeviceNetwork(code string) (device models.DeviceNetwork,
	err error) {
	return device, db.Where(models.DeviceNetwork{
		DeviceCommon: models.DeviceCommon{DeviceCode: code},
	}).Find(&device).Error
}

func (db *DBORM) GetDevicePart(code string) (device models.DevicePart, err error) {
	return device, db.Where(models.DevicePart{
		DeviceCommon: models.DeviceCommon{DeviceCode: code},
	}).Find(&device).Error
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
		selectString = SizeSelectQuery + "," + PageSelectQuery
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
		Joins(CompanyJoinQuery).
		Joins(OwnerCompanyJoinQuery).
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

func (db *DBORM) GetDeviceWithSplaJoin(spla []string) (codes []models.Code, err error) {
	return codes, db.
		//Debug().
		Table(CodeRawTable).
		Where(models.Code{SubType: "spla_cd"}).
		Where("c_idx IN (?)", spla).
		Find(&codes).Error
}

func (db *DBORM) GetDevicesServerForSearch(dc models.DeviceServer) (ds []models.DeviceServerResponse, err error) {
	return ds, db.
		//Debug().
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
		Joins(CompanyServerNoAliasJoinQuery).
		Joins(OwnerCompanyServerNoAliasJoinQuery).
		Where(&dc).Find(&ds).Error
}

func (db *DBORM) GetDevicesNetworkForSearch(dc models.DeviceNetwork) (ds []models.DeviceNetworkResponse, err error) {
	return ds, db.
		//Debug().
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
		Joins(CompanyNetworkNoAliasJoinQuery).
		Joins(OwnerCompanyNetworkNoAliasJoinQuery).
		Where(&dc).Find(&ds).Error
}

func (db *DBORM) GetDevicesPartForSearch(dc models.DevicePart) (ds []models.DevicePartResponse, err error) {
	return ds, db.
		//Debug().
		Select(PartSelectQuery).
		Table(PartRawTable).
		Joins(ManufacturePartNoAliasJoinQuery).
		Joins(ModelPartNoAliasJoinQuery).
		Joins(DeviceTypePartNoAliasJoinQuery).
		Joins(OwnershipPartNoAliasJoinQuery).
		Joins(OwnershipDivPartNoAliasJoinQuery).
		Joins(IdcPartNoAliasJoinQuery).
		Joins(RackPartNoAliasJoinQuery).
		Joins(CompanyPartNoAliasJoinQuery).
		Joins(OwnerCompanyPartNoAliasJoinQuery).
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

func (db *DBORM) AddDevice(data interface{}, device string) error {
	return db.Table(GetTableName(device)).Create(data).Error
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

func (db *DBORM) UpdateDevice(data interface{}, device string, deviceCode string) (
	interface{}, error) {
	var err error
	switch device {
	case "device_server_tb":
		err = db.
			//Debug().
			Table(device).
			Where("device_code = ?", deviceCode).
			Update(map[string]interface{}{
			// remove unnecessary data
			//"device_idx": data.(*models.DeviceServer).Idx,
			//"comment_cnt": data.(*models.DeviceServer).CommentCnt,
			//"comment_last_date": data.(*models.DeviceServer).CommentLastDate,
			//"register_id": data.(*models.DeviceServer).RegisterId,
			//"register_date": data.(*models.DeviceServer).RegisterDate,
			"out_flag": data.(*models.DeviceServer).OutFlag,
			"device_code": data.(*models.DeviceServer).DeviceCode,
			"model_cd": data.(*models.DeviceServer).Model,
			"contents": data.(*models.DeviceServer).Contents,
			"user_id": data.(*models.DeviceServer).Customer,
			"manufacture_cd": data.(*models.DeviceServer).Manufacture,
			"device_type_cd": data.(*models.DeviceServer).DeviceType,
			"warehousing_date": data.(*models.DeviceServer).WarehousingDate,
			"rent_date": data.(*models.DeviceServer).RentDate,
			"ownership_cd": data.(*models.DeviceServer).Ownership,
			"ownership_div_cd": data.(*models.DeviceServer).OwnershipDiv,
			"owner_company": data.(*models.DeviceServer).OwnerCompany,
			"hw_sn": data.(*models.DeviceServer).HwSn,
			"idc_cd": data.(*models.DeviceServer).IDC,
			"rack_cd": data.(*models.DeviceServer).Rack,
			"cost": data.(*models.DeviceServer).Cost,
			"purpose": data.(*models.DeviceServer).Purpose,
			"monitoring_flag": data.(*models.DeviceServer).MonitoringFlag,
			"monitoring_method": data.(*models.DeviceServer).MonitoringMethod,
			"ip": data.(*models.DeviceServer).Ip,
			"size_cd": data.(*models.DeviceServer).Size,
			"spla_cd": data.(*models.DeviceServer).Spla,
			"cpu": data.(*models.DeviceServer).Cpu,
			"memory": data.(*models.DeviceServer).Memory,
			"hdd": data.(*models.DeviceServer).Hdd,
			"rack_tag": data.(*models.DeviceServer).RackTag,
			"rack_loc": data.(*models.DeviceServer).RackLoc,
		}).Error
	case "device_network_tb":
		err = db.
			//Debug().
			Table(device).
			Where("device_code = ?", deviceCode).
			Update(map[string]interface{}{
			// remove unnecessary data
			//"device_idx": data.(*models.DeviceNetwork).Idx,
			//"comment_cnt": data.(*models.DeviceNetwork).CommentCnt,
			//"comment_last_date": data.(*models.DeviceNetwork).CommentLastDate,
			//"register_id": data.(*models.DeviceNetwork).RegisterId,
			//"register_date": data.(*models.DeviceNetwork).RegisterDate,
			"out_flag": data.(*models.DeviceNetwork).OutFlag,
			"device_code": data.(*models.DeviceNetwork).DeviceCode,
			"model_cd": data.(*models.DeviceNetwork).Model,
			"contents": data.(*models.DeviceNetwork).Contents,
			"user_id": data.(*models.DeviceNetwork).Customer,
			"manufacture_cd": data.(*models.DeviceNetwork).Manufacture,
			"device_type_cd": data.(*models.DeviceNetwork).DeviceType,
			"warehousing_date": data.(*models.DeviceNetwork).WarehousingDate,
			"rent_date": data.(*models.DeviceNetwork).RentDate,
			"ownership_cd": data.(*models.DeviceNetwork).Ownership,
			"ownership_div_cd": data.(*models.DeviceNetwork).OwnershipDiv,
			"owner_company": data.(*models.DeviceNetwork).OwnerCompany,
			"hw_sn": data.(*models.DeviceNetwork).HwSn,
			"idc_cd": data.(*models.DeviceNetwork).IDC,
			"rack_cd": data.(*models.DeviceNetwork).Rack,
			"cost": data.(*models.DeviceNetwork).Cost,
			"purpose": data.(*models.DeviceNetwork).Purpose,
			"monitoring_flag": data.(*models.DeviceNetwork).MonitoringFlag,
			"monitoring_method": data.(*models.DeviceNetwork).MonitoringMethod,
			"ip": data.(*models.DeviceNetwork).Ip,
			"size_cd": data.(*models.DeviceNetwork).Size,
			"firmware_version": data.(*models.DeviceNetwork).FirmwareVersion,
			"rack_tag": data.(*models.DeviceNetwork).RackTag,
			"rack_loc": data.(*models.DeviceNetwork).RackLoc,
		}).Error
	case "device_part_tb":
		err = db.
			//Debug().
			Table(device).
			Where("device_code = ?", deviceCode).
			Update(map[string]interface{}{
			// remove unnecessary data
			//"device_idx": data.(*models.DevicePart).Idx,
			//"comment_cnt": data.(*models.DevicePart).CommentCnt,
			//"comment_last_date": data.(*models.DevicePart).CommentLastDate,
			//"register_id": data.(*models.DevicePart).RegisterId,
			//"register_date": data.(*models.DevicePart).RegisterDate,
			"out_flag": data.(*models.DevicePart).OutFlag,
			"device_code": data.(*models.DevicePart).DeviceCode,
			"model_cd": data.(*models.DevicePart).Model,
			"contents": data.(*models.DevicePart).Contents,
			"user_id": data.(*models.DevicePart).Customer,
			"manufacture_cd": data.(*models.DevicePart).Manufacture,
			"device_type_cd": data.(*models.DevicePart).DeviceType,
			"warehousing_date": data.(*models.DevicePart).WarehousingDate,
			"rent_date": data.(*models.DevicePart).RentDate,
			"ownership_cd": data.(*models.DevicePart).Ownership,
			"ownership_div_cd": data.(*models.DevicePart).OwnershipDiv,
			"owner_company": data.(*models.DevicePart).OwnerCompany,
			"hw_sn": data.(*models.DevicePart).HwSn,
			"idc_cd": data.(*models.DevicePart).IDC,
			"rack_cd": data.(*models.DevicePart).Rack,
			"cost": data.(*models.DevicePart).Cost,
			"purpose": data.(*models.DevicePart).Purpose,
			"monitoring_flag": data.(*models.DevicePart).MonitoringFlag,
			"monitoring_method": data.(*models.DevicePart).MonitoringMethod,
			"warranty": data.(*models.DevicePart).Warranty,
			"rack_code_cd": data.(*models.DevicePart).RackCode,
		}).Error
	}

	return data, err

}

// Update OutFlag
func (db *DBORM) UpdateOutFlag(codes []string, device string, flag int) error {
	return db.Table(device).Where("device_code IN (?)", codes).Update(outFlagField, flag).Error
}

//log search
func (db *DBORM) GetDeviceLogs(code string) (logs []models.DeviceLog, err error) {
	return logs, db.Where(models.DeviceLog{DeviceCode: code}).Find(&logs).Error
}

/*func (db *DBORM) GetDeviceLogInServer(code string) (logs []models.DeviceLog, err error) {
	return logs, db.Where(models.DeviceLog{DeviceCode: code}).Find(&logs).Error
}

func (db *DBORM) GetDeviceLogInNetwork(code string) (comments []models.DeviceComment, err error) {
	return comments, db.Where(models.DeviceLog{DeviceCode: code}).Find(&comments).Error
}

func (db *DBORM) GetDeviceLogInPart(code string) (comments []models.DeviceComment, err error) {
	return comments, db.Where(models.DeviceLog{DeviceCode: code}).Find(&comments).Error
}*/

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

func GetTableName(device string) string {
	switch device {
	case "server":
		return ServerRawTable
	case "network":
		return NetworkRawTable
	case "part":
		return PartRawTable
	default:
		return ""

	}
	return ""
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
