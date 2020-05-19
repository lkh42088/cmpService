package convert

import (
	"fmt"
	"cmpService/common/db"
	"cmpService/common/mariadblayer"
	"cmpService/common/models"
	"cmpService/dbmigrator/cbmodels"
	"cmpService/dbmigrator/config"
	"cmpService/dbmigrator/mysqllayer"
	"strconv"
	"strings"
	"time"
)

var idx_comment uint = 0

func RunConvertDb() {
	convertInternal(ConvertItem)
	convertInternal(ConvertItemSub)
	convertInternal(ConvertDeviceServer)
	convertInternal(ConvertDeviceNetwork)
	convertInternal(ConvertDevicePart)
}

func convertInternal(convert func(*mysqllayer.CBORM, *mariadblayer.DBORM)) {
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

	convert(oldDb, newDb)
}

func ConvertItem(odb *mysqllayer.CBORM, ndb *mariadblayer.DBORM) {
	olds, err := odb.GetAllItems()
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	for num, old := range olds {
		newDB := GetCodeByItem(old)
		fmt.Println(num, ":", old, "-->", newDB)
		ndb.AddCode(newDB)
	}
}

func ConvertItemSub(odb *mysqllayer.CBORM, ndb *mariadblayer.DBORM) {
	olds, err := odb.GetAllSubItems()
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	for num, old := range olds {
		newDB := GetSubCodeByItemSub(old)
		fmt.Println(num, ":", old, "-->", newDB)
		ndb.AddSubCode(newDB)
	}
}

func ConvertDeviceServer(odb *mysqllayer.CBORM, ndb *mariadblayer.DBORM) {
	olds, err := odb.GetAllDevicesServerFromOldDB()
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	for i, old := range olds {
		// case depth == 0 : device table
		// case depth != 0 : comment table
		if i % 100 == 0 {
			time.Sleep(time.Millisecond * 50)
		}
		sd, dc := GetServerTbByDevice(old)
		if old.WrIsComment == 0 {
			fmt.Println("server:", i, ": dev")
			ndb.AddDeviceServer(sd)
		} else  {
			idx_comment++
			fmt.Println("server:", i, ": comment, ", idx_comment)
			dc.Idx = idx_comment
			ndb.AddComment(dc)
		}
	}
}

func ConvertDeviceNetwork(odb *mysqllayer.CBORM, ndb *mariadblayer.DBORM) {
	olds, err := odb.GetAllDevicesNetworkFromOldDB()
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	for i, old := range olds {
		// case depth == 0 : device table
		// case depth != 0 : comment table
		if i % 100 == 0 {
			time.Sleep(time.Millisecond * 100)
		}
		nd, dc := GetNetworkTbByDevice(old)
		if old.WrIsComment == 0 {
			fmt.Println("network:", i, ": dev")
			ndb.AddDeviceNetwork(nd)
		} else  {
			idx_comment++
			dc.Idx = idx_comment
			fmt.Println("network:", i, ": comment, ", idx_comment)
			ndb.AddComment(dc)
		}
	}
}

func ConvertDevicePart(odb *mysqllayer.CBORM, ndb *mariadblayer.DBORM) {
	olds, err := odb.GetAllDevicesPartFromOldDB()
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	for i, old := range olds {
		// case depth == 0 : device table
		// case depth != 0 : comment table
		if i % 100 == 0 {
			time.Sleep(time.Millisecond * 100)
		}
		pd, dc := GetPartTbByDevice(old)
		if old.WrIsComment == 0 {
			fmt.Println("part:", i, ": dev")
			ndb.AddDevicePart(pd)
		} else  {
			idx_comment++
			dc.Idx = idx_comment
			fmt.Println("part:", i, ": comment, ", idx_comment)
			ndb.AddComment(dc)
		}
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

	newDb.DeleteSubCodes()
	newDb.DeleteCodes()
	newDb.DeleteAllDevicesPart()
	newDb.DeleteAllDevicesNetwork()
	newDb.DeleteAllDevicesServer()
	newDb.DeleteAllComments()
	newDb.DeleteAllLogs()
}

func GetCodeByItem(item cbmodels.Item) (code models.Code) {
	code.CodeID = item.ItemID
	code.Type = convertCodeType(item.Table)
	code.SubType = convertCodeSubtype(item.Column)
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

func convertCodeType(s string) string {
	var codeType string
	switch strings.TrimSpace(s) {
	case "total":
		codeType = "total"
	case "device":
		codeType = "device_server"
	case "ndevice":
		codeType = "device_network"
	case "pdevice":
		codeType = "device_part"
	default:
		codeType = ""
	}
	return codeType
}

func convertCodeSubtype(s string) string {
	var codeSubType string
	switch strings.TrimSpace(s) {
	case "wr_51":
		codeSubType = "ownership_cd_1"
	case "wr_52":
		codeSubType = "ownership_cd_2"
	case "wr_6":
		codeSubType = "size_cd"
	case "wr_101":
		codeSubType = "idc_cd"
	case "wr_11":
		codeSubType = "spla_cd"
	case "wr_link1":
		codeSubType = "manufacture_cd"
	case "wr_link2":
		codeSubType = "device_type_cd"
	case "switch":
		codeSubType = ""
	default:
		codeSubType = ""
	}
	return codeSubType
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

func GetServerTbByDevice(device cbmodels.ServerDevice)(
	sd models.DeviceServer, dc models.DeviceComment) {
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
	sd.Spla = strings.Replace(device.Wr11, ";", "|", -1)
	sd.Cpu = device.Wr2
	sd.Memory = device.Wr3
	sd.Hdd = device.Wr4
	sd.MonitoringFlag = 0
	sd.MonitoringMethod = 0

	// Comment Table
	//dc.Idx = uint(device.CbDeviceID)
	dc.DeviceCode = device.Wr1
	dc.Depth = device.WrIsComment
	dc.Contents = device.WrContent
	dc.RegisterId = device.MbId
	dc.RegisterName = device.WrName
	dc.RegisterDate = device.WrDatetime

	return sd, dc
}

func GetNetworkTbByDevice(device cbmodels.NetworkDevice)(
	nd models.DeviceNetwork, dc models.DeviceComment) {
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

	dc.DeviceCode = device.Wr1
	dc.Depth = device.WrIsComment
	dc.Contents = device.WrContent
	dc.RegisterId = device.MbId
	dc.RegisterName = device.WrName
	dc.RegisterDate = device.WrDatetime

	return nd, dc
}

func GetPartTbByDevice(device cbmodels.PartDevice)(
	pd models.DevicePart, dc models.DeviceComment) {
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

	dc.DeviceCode = device.Wr1
	dc.Depth = device.WrIsComment
	dc.Contents = device.WrContent
	dc.RegisterId = device.MbId
	dc.RegisterName = device.WrName
	dc.RegisterDate = device.WrDatetime

	return pd, dc
}

