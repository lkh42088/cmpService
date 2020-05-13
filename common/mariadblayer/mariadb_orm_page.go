package mariadblayer

import (
	"math"
	"nubes/common/models"
)

// order direction
// 0 : ASC
// 1 : DESC
// default : idx DESC
func Orderby(order_field string, direction int) string {
	if order_field == "" || direction > 1 || direction < 0 {
		return "idx DESC"
	}
	orderby := order_field
	if direction == 0 {
		orderby += " ASC"
	} else {
		orderby += " DESC"
	}
	return orderby
}

func TotalPage(count int, size int) int {
	return int(math.Ceil(float64(count) / float64(size)))
}

func CombineCondition(outFlag string) string {
	return "out_flag = '" + outFlag + "'"
}

func (db *DBORM) GetDevicesServerForPage(cri models.PageCreteria) (
	server models.DeviceServerPage, err error) {

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
