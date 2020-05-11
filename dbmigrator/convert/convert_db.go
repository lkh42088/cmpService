package convert

import (
	"fmt"
	"nubes/common/db"
	"nubes/common/mariadblayer"
	"nubes/common/models"
	"nubes/dbmigrator/cbmodels"
	"nubes/dbmigrator/config"
	"nubes/dbmigrator/mysqllayer"
	"strconv"
	"strings"
	"time"
)

func RunConvertDb() {
	// Old Database: Mysql
	oldConfig := config.GetOldDatabaseConfig()
	oldOptions := db.GetDataSourceName(oldConfig)
	oldDb, err := mysqllayer.NewCBORM(oldConfig.DBDriver, oldOptions)
	if err != nil {
		fmt.Println("oldConfig Error:", err)
		return
	}
	defer oldDb.Close()

	// New Database: Mariadb
	newConfig := config.GetNewDatabaseConfig()
	newOptions := db.GetDataSourceName(newConfig)
	newDb, err := mariadblayer.NewDBORM(newConfig.DBDriver, newOptions)
	if err != nil {
		fmt.Println("newConfig Error:", err)
		return
	}
	defer newDb.Close()

	ConvertItem(oldDb, newDb)
	ConvertItemSub(oldDb, newDb)
	ConvertDeviceServer(oldDb, newDb)
	ConvertDeviceNetwork(oldDb, newDb)
	ConvertDevicePart(oldDb, newDb)
}

func ConvertItem(odb *mysqllayer.CBORM, ndb *mariadblayer.DBORM) {
	olds, err := odb.GetAllItems()
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	for num, old := range olds {
		new := GetCodeByItem(old)
		fmt.Println(num, ":", old, "-->", new)
		ndb.AddCode(new)
	}
}

func ConvertItemSub(odb *mysqllayer.CBORM, ndb *mariadblayer.DBORM) {
	olds, err := odb.GetAllSubItems()
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	for num, old := range olds {
		new := GetSubCodeByItemSub(old)
		fmt.Println(num, ":", old, "-->", new)
		ndb.AddSubCode(new)
	}
}

func ConvertDeviceServer(odb *mysqllayer.CBORM, ndb *mariadblayer.DBORM) {
	olds, err := odb.GetAllDevicesServerFromOldDB()
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	for i, old := range olds {
		new := GetServerTbByDevice(old)
		if i % 100 == 0 {
			time.Sleep(time.Millisecond * 100)
		}
		//fmt.Println(num, ":", old, "-->", new)
		ndb.AddDeviceServer(new)
	}
}

func ConvertDeviceNetwork(odb *mysqllayer.CBORM, ndb *mariadblayer.DBORM) {
	olds, err := odb.GetAllDevicesNetworkFromOldDB()
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	for i, old := range olds {
		new := GetNetworkTbByDevice(old)
		if i % 100 == 0 {
			time.Sleep(time.Millisecond * 100)
		}
		//fmt.Println(num, ":", old, "-->", new)
		ndb.AddDeviceNetwork(new)
	}
}

func ConvertDevicePart(odb *mysqllayer.CBORM, ndb *mariadblayer.DBORM) {
	olds, err := odb.GetAllDevicesPartFromOldDB()
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	for i, old := range olds {
		new := GetPartTbByDevice(old)
		if i % 100 == 0 {
			time.Sleep(time.Millisecond * 100)
		}
		//fmt.Println(num, ":", old, "-->", new)
		ndb.AddDevicePart(new)
	}
}

func DeleteDeviceTb() {
	// New Database: Mariadb
	newConfig := config.GetNewDatabaseConfig()
	newOptions := db.GetDataSourceName(newConfig)
	newDb, err := mariadblayer.NewDBORM(newConfig.DBDriver, newOptions)
	if err != nil {
		fmt.Println("newConfig Error:", err)
		return
	}
	defer newDb.Close()

	newDb.DeleteAllDevicesPart()
	newDb.DeleteAllDevicesNetwork()
	newDb.DeleteAllDevicesServer()
}

func GetCodeByItem(item cbmodels.Item) (code models.Code) {
	code.CodeID = item.ItemID
	code.Type = item.Table
	code.SubType = item.Column
	code.Name = item.Item
	code.Order = item.Desc
	return code
}

func GetSubCodeByItemSub(subitem cbmodels.SubItem) (subcode models.SubCode) {
	subcode.ID = subitem.SubItemID
	subcode.Name = subitem.SubItem
	subcode.CodeID = subitem.ItemID
	subcode.Order = subitem.Desc
	return subcode
}

const TimeFormat = "2006-01-02 15:04:05"

func convInt(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return num
}

func convStr(i int) string {
	str := strconv.Itoa(i)
	if str == "" {
		return "0"
	}
	return strings.TrimSpace(str)
}

func sepIdcRack(s string, num int) int {
	str := strings.Split(s, "|")
	if len(str) >= num {
		i, _ := strconv.Atoi(str[num-1])
		return i
	}
	return 0
}

func sepIps(s string) string {
	result := ""
	str := strings.Split(s, ";")
	if len(str) == 0 {
		return "|"
	}
	for _, v := range str {
		result += v + "|"
	}
	return result
}

func GetServerTbByDevice(device cbmodels.ServerDevice)(sd models.DeviceServer) {
	sd.Idx = uint(device.CbDeviceID)
	sd.OutFlag = false
	sd.Num = device.WrNum
	sd.CommentCnt = device.WrComment
	sd.CommentLastDate, _ = time.Parse(TimeFormat, device.WrLast)
	sd.Option = device.WrOption
	sd.Hit = device.WrHit
	sd.RegisterId = device.MbId
	sd.Password = device.WrPassword
	sd.RegisterName = device.WrName
	sd.RegisterEmail = device.WrEmail
	sd.RegisterDate = device.WrDatetime
	sd.DeviceCode = device.Wr1
	sd.Model = convInt(device.WrSubject)
	sd.Contents = device.WrContent
	sd.Customer = convInt(device.WrTrackback)
	sd.Manufacture = convInt(device.WrLink1)
	sd.DeviceType = convInt(device.WrLink2)
	sd.WarehousingDate = device.WrLink1Hit
	sd.RentDate = device.Wr8
	sd.Ownership = device.Wr5
	sd.OwnerCompany = device.Wr7
	sd.HwSn = device.Wr9
	sd.IDC = sepIdcRack(device.Wr10, 1)
	sd.Rack = sepIdcRack(device.Wr10, 2)
	sd.Cost = device.Wr12
	sd.Purpos = device.Wr13
	sd.Ip = sepIps(device.WrHomepage)
	sd.Size = convInt(device.Wr6)
	sd.Spla = convInt(strings.Trim(device.Wr11, ";"))
	sd.Cpu = device.Wr2
	sd.Memory = device.Wr3
	sd.Hdd = device.Wr4
	sd.MonitoringFlag = 0
	sd.MonitoringMethod = 0

	return sd
}

func GetNetworkTbByDevice(device cbmodels.NetworkDevice)(nd models.DeviceNetwork) {
	nd.Idx = uint(device.CbDeviceID)
	nd.OutFlag = false
	nd.Num = device.WrNum
	nd.CommentCnt = device.WrComment
	nd.CommentLastDate, _ = time.Parse(TimeFormat, device.WrLast)
	nd.Option = device.WrOption
	nd.Hit = device.WrHit
	nd.RegisterId = device.MbId
	nd.Password = device.WrPassword
	nd.RegisterName = device.WrName
	nd.RegisterEmail = device.WrEmail
	nd.RegisterDate = device.WrDatetime
	nd.DeviceCode = device.Wr1
	nd.Model = convInt(device.WrSubject)
	nd.Contents = device.WrContent
	nd.Customer = convInt(device.WrTrackback)
	nd.Manufacture = convInt(device.WrLink1)
	nd.DeviceType = convInt(device.WrLink2)
	nd.WarehousingDate = device.WrLink1Hit
	nd.RentDate = device.Wr8
	nd.Ownership = device.Wr5
	nd.OwnerCompany = device.Wr7
	nd.HwSn = device.Wr9
	nd.IDC = sepIdcRack(device.Wr10, 1)
	nd.Rack = sepIdcRack(device.Wr10, 2)
	nd.Cost = device.Wr12
	nd.Purpos = device.Wr13
	nd.Ip = sepIps(device.WrHomepage)
	nd.Size = convInt(device.Wr6)
	nd.FirmwareVersion = device.Wr2
	nd.MonitoringFlag = 0
	nd.MonitoringMethod = 0

	return nd
}

func GetPartTbByDevice(device cbmodels.PartDevice)(pd models.DevicePart) {
	pd.Idx = uint(device.CbDeviceID)
	pd.OutFlag = false
	pd.Num = device.WrNum
	pd.CommentCnt = device.WrComment
	pd.CommentLastDate, _ = time.Parse(TimeFormat, device.WrLast)
	pd.Option = device.WrOption
	pd.Hit = device.WrHit
	pd.RegisterId = device.MbId
	pd.Password = device.WrPassword
	pd.RegisterName = device.WrName
	pd.RegisterEmail = device.WrEmail
	pd.RegisterDate = device.WrDatetime
	pd.DeviceCode = device.Wr1
	pd.Model = convInt(device.WrSubject)
	pd.Contents = device.WrContent
	pd.Customer = convInt(device.WrTrackback)
	pd.Manufacture = convInt(device.WrLink1)
	pd.DeviceType = convInt(device.WrLink2)
	pd.WarehousingDate = device.WrLink1Hit
	pd.RentDate = device.Wr8
	pd.Ownership = device.Wr5
	pd.OwnerCompany = device.Wr7
	pd.HwSn = device.Wr9
	pd.IDC = sepIdcRack(device.Wr10, 1)
	pd.Rack = sepIdcRack(device.Wr10, 2)
	pd.Cost = device.Wr12
	pd.Purpos = device.Wr13
	pd.Warranty = device.Wr2
	pd.MonitoringFlag = 0
	pd.MonitoringMethod = 0

	return pd
}



