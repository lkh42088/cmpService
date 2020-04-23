package mariadblayer

import (
	"fmt"
	db2 "nubes/common/db"
	"nubes/common/models"
	"testing"
	"time"
)

func TestDeviceAddEntry(t *testing.T) {
	config := getTestConfig()
	options := db2.GetDataSourceName(config)
	fmt.Println("conf:", config)
	fmt.Println("options:", options)
	db, err := NewDBORM(config.DBDriver, options)
	if err != nil {
		return
	}
	device := models.Device{
		MenuType: "device",
		OutFlag: false,
		Num: 3,
		CommentCnt: 3,
		CommentLastDate: time.Now(),
		Option: "option",
		Hit: 3,
		RegisterId: "yjs_id",
		Password: "password",
		RegisterName: "yjs",
		RegisterEmail: "yjs@nubes-bridge.com",
		RegisterDate: time.Now(),
		DeviceCode: "CBSNUBES03",
		Model: 1,
		Contents: "텍스트 내용",
		Customer: 2,
		Manufacture: 3,
		DeviceType: 4,
		WarehousingDate: "20200423",
		RentDate: "20200423|20200512",
		Ownership: "1",
		OwnerCompany: "소유업체명~~",
		HwSn: "hw/sn",
		IDC: 5,
		Rack: 6,
		Cost: "원가",
		Purpos: "용도",
		Ip: "255.255.255.1",
		Size: 7,
		Spla: 8,
		Cpu: "cpu",
		Memory: "memory",
		Hdd: "hdd",
		FirmwareVersion: "만료",
		Warranty: "warranty",
		MonitoringFlag: 0,
		MonitoringMethod: 9,
	}
	device, err = db.AddDevice(device)
	fmt.Println("device: ", device, "err:", err)
}
