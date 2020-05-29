package mariadblayer

import (
	"cmpService/common/lib"
	"cmpService/common/models"
	"math"
	"strings"
)

//const limitNum = 1000
// order direction
// 0 : ASC
// 1 : DESC
// default : DEVICE_CODE DESC
func Orderby(order_field string, direction int) string {
	if order_field == "" || direction > 1 || direction < 0 {
		return "DEVICE_CODE DESC"
	}
	orderby := ConvertToColumn(order_field)
	if direction == 0 {
		orderby = "d." + orderby + " ASC"
	} else {
		orderby = "d." + orderby + " ASC"
	}
	//fmt.Println(orderby)
	return orderby
}

func TotalPage(count int, size int) int {
	return int(math.Ceil(float64(count) / float64(size)))
}

func CombineCondition(outFlag string) string {
	return "out_flag = '" + outFlag + "'"
}

// NB specific code
func SetThousandCount(cri *models.PageCreteria) {
	if cri.CheckCnt <= cri.Size {
		if cri.CheckCnt > cri.Count {
			return
		}
		cri.CheckCnt = 0
	} else {
		cri.CheckCnt = ((cri.CheckCnt - 1) / cri.Size) * cri.Size
	}
}

func (db *DBORM) GetDevicesServerForPage(cri models.PageCreteria) (
	server models.DeviceServerPage, err error) {

	db.Model(&server.Devices).Count(&cri.Count)
	orderby := Orderby(cri.OrderKey, cri.Direction)
	SetThousandCount(&cri)
	err = db.
		Order(orderby).
		Limit(cri.Size).
		Offset(cri.CheckCnt).
		Where(CombineCondition(cri.OutFlag)).
		Find(&server.Devices).Error
	if err != nil {
		lib.LogWarn("[Error] %s\n", err)
	}
	server.Page = cri
	return server, err
}

func (db *DBORM) GetDevicesServerWithJoin(cri models.PageCreteria) (
	server models.DeviceServerPage, err error) {

	db.Model(&models.DeviceServer{}).Count(&cri.Count)
	orderBy := Orderby(cri.OrderKey, cri.Direction)
	SetThousandCount(&cri)

	err = db.
		//Debug().
		Select(SizeSelectQuery+","+PageSelectQuery).
		Model(&models.DeviceServer{}).
		Table(ServerTable).
		Order(orderBy).
		Limit(cri.Size).
		Offset(cri.CheckCnt).
		Where(CombineCondition(cri.OutFlag)).
		Joins(ManufactureServerJoinQuery).
		Joins(ModelJoinQuery).
		Joins(DeviceTypeServerJoinQuery).
		Joins(OwnershipJoinQuery).
		Joins(OwnershipDivJoinQuery).
		Joins(IdcJoinQuery).
		Joins(RackJoinQuery).
		Joins(SizeJoinQuery).
		Joins(CompanyLeftJoinQuery).
		Find(&server.Devices).Error

	if err != nil {
		lib.LogWarn("[Error] %s\n", err)
	}

	server.Page = cri
	return server, err
}

func (db *DBORM) GetDevicesNetworkForPage(cri models.PageCreteria) (
	network models.DeviceNetworkPage, err error) {

	db.Model(&network.Devices).Count(&cri.Count)
	orderby := Orderby(cri.OrderKey, cri.Direction)
	SetThousandCount(&cri)
	err = db.
		Order(orderby).
		Limit(cri.Size).
		Offset(cri.CheckCnt).
		Where(CombineCondition(cri.OutFlag)).
		Find(&network.Devices).Error
	if err != nil {
		lib.LogWarn("[Error] %s\n", err)
	}
	network.Page = cri
	return network, err
}

func (db *DBORM) GetDevicesNetworkWithJoin(cri models.PageCreteria) (
	network models.DeviceNetworkPage, err error) {

	db.Model(&models.DeviceNetwork{}).Count(&cri.Count)
	orderBy := Orderby(cri.OrderKey, cri.Direction)
	SetThousandCount(&cri)

	err = db.
		//Debug().
		Select(SizeSelectQuery+","+PageSelectQuery).
		Model(&models.DeviceNetwork{}).
		Table(NetworkTable).
		Order(orderBy).
		Limit(cri.Size).
		Offset(cri.CheckCnt).
		Where(CombineCondition(cri.OutFlag)).
		Joins(ManufactureNetworkJoinQuery).
		Joins(ModelJoinQuery).
		Joins(DeviceTypeNetworkJoinQuery).
		Joins(OwnershipJoinQuery).
		Joins(OwnershipDivJoinQuery).
		Joins(IdcJoinQuery).
		Joins(RackJoinQuery).
		Joins(SizeJoinQuery).
		Joins(CompanyLeftJoinQuery).
		Find(&network.Devices).Error

	if err != nil {
		lib.LogWarn("[Error] %s\n", err)
	}

	network.Page = cri
	return network, err
}

func (db *DBORM) GetDevicesPartForPage(cri models.PageCreteria) (
	part models.DevicePartPage, err error) {

	db.Model(&part.Devices).Count(&cri.Count)
	orderby := Orderby(cri.OrderKey, cri.Direction)
	SetThousandCount(&cri)
	err = db.
		Order(orderby).
		Limit(cri.Size).
		Offset(cri.CheckCnt).
		Where(CombineCondition(cri.OutFlag)).
		Find(&part.Devices).Error
	if err != nil {
		lib.LogWarn("[Error] %s\n", err)
	}
	part.Page = cri
	return part, err
}

func (db *DBORM) GetDevicesPartWithJoin(cri models.PageCreteria) (
	part models.DevicePartPage, err error) {

	db.Model(&models.DevicePart{}).Count(&cri.Count)
	orderBy := Orderby(cri.OrderKey, cri.Direction)
	SetThousandCount(&cri)

	err = db.
		//Debug().
		Select(PageSelectQuery).
		Model(&models.DeviceNetwork{}).
		Table(PartTable).
		Order(orderBy).
		Limit(cri.Size).
		Offset(cri.CheckCnt).
		Where(CombineCondition(cri.OutFlag)).
		Joins(ManufacturePartJoinQuery).
		Joins(ModelJoinQuery).
		Joins(DeviceTypePartJoinQuery).
		Joins(OwnershipJoinQuery).
		Joins(OwnershipDivJoinQuery).
		Joins(IdcJoinQuery).
		Joins(RackJoinQuery).
		Joins(CompanyLeftJoinQuery).
		Find(&part.Devices).Error

	if err != nil {
		lib.LogWarn("[Error] %s\n", err)
	}

	part.Page = cri
	return part, err
}

func ConvertToColumn(field string) string {
	col := strings.ToLower(field)

	switch col {
	case "idx":
		col = "device_idx"
	case "outflag":
		col = "out_flag"
	case "num":
		col = "num"
	case "commentcnt":
		col = "comment_cnt"
	case "commentlastdate":
		col = "comment_last_date"
	case "registerid":
		col = "register_id"
	case "registerdate":
		col = "register_date"
	case "devicecode":
		col = "device_code"
	case "model":
		col = "model_cd"
	case "contents":
		col = "contents"
	case "customer":
		col = "user_id"
	case "manufacture":
		col = "manufacture_cd"
	case "devicetype":
		col = "device_type_cd"
	case "warehousingdate":
		col = "warehousing_date"
	case "rentdate":
		col = "rent_date"
	case "ownership":
		col = "ownership_cd"
	case "ownercompany":
		col = "owner_company"
	case "hwsn":
		col = "hw_sn"
	case "idc":
		col = "idc_cd"
	case "rack":
		col = "rack_cd"
	case "cost":
		col = "cost"
	case "purpos":
		col = "purpos"
	case "ip":
		col = "ip"
	case "size":
		col = "size_cd"
	case "spla":
		col = "spla_cd"
	case "cpu":
		col = "cpu"
	case "memory":
		col = "memory"
	case "hdd":
		col = "hdd"
	case "monitoringflag":
		col = "monitoring_flag"
	case "monitoringmethod":
		col = "mornitoring_method"
	case "firmwareversion":
		col = "firmware_version"
	case "warranty":
		col = "warranty"
	case "rackcode":
		col = "rack_code_cd"
	case "racktag":
		col = "rack_tag"
	case "rackloc":
		col = "rack_loc"
	}
	return col
}