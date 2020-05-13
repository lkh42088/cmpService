package mariadblayer

import (
	"math"
	"nubes/common/models"
	"fmt"
	"strings"
)

const limitNum = 1000
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
		orderby += " ASC"
	} else {
		orderby += " DESC"
	}
	fmt.Println(orderby)
	return orderby
}

func TotalPage(count int, size int) int {
	return int(math.Ceil(float64(count) / float64(size)))
}

func CombineCondition(outFlag string) string {
	return "out_flag = '" + outFlag + "'"
}

func SetLimitNum(cri models.PageCreteria) models.PageCreteria {
	// basic rule : 1000 data per 1 time
	cri.Size = limitNum
	return cri
}

func (db *DBORM) GetDevicesServerForPage(cri models.PageCreteria) (
	server models.DeviceServerPage, err error) {

	SetLimitNum(cri)
	beginNum := cri.Size * (cri.CurPage - 1)
	db.Model(&server.Devices).Count(&cri.Count)
	orderby := Orderby(cri.OrderKey, cri.Direction)
	cri.TotalPage = TotalPage(cri.Count, cri.Size)
	err = db.
		Order(orderby).
		Limit(cri.Size).
		Offset(beginNum).
		Where(CombineCondition(cri.OutFlag)).
		Find(&server.Devices).Error
	server.Page = cri
	return server, err
}

func (db *DBORM) GetDevicesNetworkForPage(cri models.PageCreteria) (
	network models.DeviceNetworkPage, err error) {

	beginNum := cri.Size * (cri.CurPage - 1)
	db.Model(&network.Devices).Count(&cri.Count)
	orderby := Orderby(cri.OrderKey, cri.Direction)
	cri.TotalPage = TotalPage(cri.Count, cri.Size)
	err = db.
		Order(orderby).
		Limit(cri.Size).
		Offset(beginNum).
		Where(CombineCondition(cri.OutFlag)).
		Find(&network.Devices).Error
	network.Page = cri
	return network, err
}

func (db *DBORM) GetDevicesPartForPage(cri models.PageCreteria) (
	part models.DevicePartPage, err error) {

	beginNum := cri.Size * (cri.CurPage - 1)
	db.Model(&part.Devices).Count(&cri.Count)
	orderby := Orderby(cri.OrderKey, cri.Direction)
	cri.TotalPage = TotalPage(cri.Count, cri.Size)
	err = db.
		Order(orderby).
		Limit(cri.Size).
		Offset(beginNum).
		Where(CombineCondition(cri.OutFlag)).
		Find(&part.Devices).Error
	part.Page = cri
	return part, err
}

func ConvertToColumn(field string) string {
	col := strings.ToLower(field)

	switch col {
	case "idx":
		col = "idx"
	case "outflag":
		col = "out_flag"
	case "num":
		col = "num"
	case "commentcnt":
		col = "comment_cnt"
	case "commentlastdate":
		col = "comment_last_date"
	case "option":
		col = "option"
	case "hit":
		col = "hit"
	case "registerid":
		col = "register_id"
	case "password":
		col = "register_password"
	case "registername":
		col = "register_name"
	case "registeremail":
		col = "register_email"
	case "registerdate":
		col = "register_date"
	case "devicecode":
		col = "device_code"
	case "model":
		col = "model_cd"
	case "contents":
		col = "contents"
	case "customer":
		col = "customer_cd"
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
	}
	return col
}