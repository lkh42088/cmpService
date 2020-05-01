package mariadblayer

import (
	"fmt"
	"nubes/common/models"
	"testing"
	"time"
)

func TestDeviceAddEntry_Server(t *testing.T) {
	db, err := getTestDb()
	if err != nil {
		return
	}
	device := models.DeviceServer{
		OutFlag: false,
		Num: 1,
		CommentCnt: 0,
		CommentLastDate: time.Now(),
		Option: "option",
		Hit: 1,
		RegisterId: "nhj_id",
		Password: "password",
		RegisterName: "nhj",
		RegisterEmail: "nhj@nubes-bridge.com",
		RegisterDate: time.Now(),
		DeviceCode: "CBSNUBES01",
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
		Ip: "255.255.255.2",
		Size: 7,
		Spla: 8,
		Cpu: "cpu",
		Memory: "memory",
		Hdd: "hdd",
		MonitoringFlag: 0,
		MonitoringMethod: 9,
	}
	device, err = db.AddDeviceServer(device)
	fmt.Println("collectdevice: ", device, "err:", err)

	device = models.DeviceServer{
		OutFlag: false,
		Num: 2,
		CommentCnt: 0,
		CommentLastDate: time.Now(),
		Option: "option",
		Hit: 2,
		RegisterId: "yjs_id",
		Password: "password",
		RegisterName: "yjs",
		RegisterEmail: "yjs@nubes-bridge.com",
		RegisterDate: time.Now(),
		DeviceCode: "CBSNUBES02",
		Model: 1,
		Contents: "텍스트 내용~",
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
		Ip: "255.255.255.2",
		Size: 7,
		Spla: 8,
		Cpu: "cpu",
		Memory: "memory",
		Hdd: "hdd",
		MonitoringFlag: 0,
		MonitoringMethod: 9,
	}
	device, err = db.AddDeviceServer(device)
	fmt.Println("collectdevice: ", device, "err:", err)

	device = models.DeviceServer{
		OutFlag: false,
		Num: 3,
		CommentCnt: 0,
		CommentLastDate: time.Now(),
		Option: "option",
		Hit: 3,
		RegisterId: "pms_id",
		Password: "password",
		RegisterName: "pms",
		RegisterEmail: "pms@nubes-bridge.com",
		RegisterDate: time.Now(),
		DeviceCode: "CBSNUBES03",
		Model: 1,
		Contents: "텍스트 내용~",
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
		Ip: "255.255.255.2",
		Size: 7,
		Spla: 8,
		Cpu: "cpu",
		Memory: "memory",
		Hdd: "hdd",
		MonitoringFlag: 0,
		MonitoringMethod: 9,
	}
	device, err = db.AddDeviceServer(device)
	fmt.Println("collectdevice: ", device, "err:", err)

}


func TestDeviceAddEntry_Network(t *testing.T) {
	db, err := getTestDb()
	if err != nil {
		return
	}
	device := models.DeviceNetwork{
		OutFlag: false,
		Num: 2,
		CommentCnt: 0,
		CommentLastDate: time.Now(),
		Option: "option",
		Hit: 2,
		RegisterId: "pms_id",
		Password: "password",
		RegisterName: "pms",
		RegisterEmail: "pms@nubes-bridge.com",
		RegisterDate: time.Now(),
		DeviceCode: "CBSNETWORK02",
		Model: 1,
		Contents: "텍스트 내용_NET2",
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
		Ip: "255.255.255.2",
		Size: 7,
		FirmwareVersion: "만료",
		MonitoringFlag: 0,
		MonitoringMethod: 9,
	}
	device, err = db.AddDeviceNetwork(device)
	fmt.Println("collectdevice: ", device, "err:", err)
}


func TestDeviceAddEntry_Part(t *testing.T) {
	db, err := getTestDb()
	if err != nil {
		return
	}
	device := models.DevicePart{
		OutFlag: false,
		Num: 1,
		CommentCnt: 0,
		CommentLastDate: time.Now(),
		Option: "option",
		Hit: 1,
		RegisterId: "jjh_id",
		Password: "password",
		RegisterName: "jjh",
		RegisterEmail: "jjh@nubes-bridge.com",
		RegisterDate: time.Now(),
		DeviceCode: "CBSPART01",
		Model: 1,
		Contents: "텍스트 내용_PART1",
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
		Warranty: "warranty",
		MonitoringFlag: 0,
		MonitoringMethod: 9,
	}
	device, err = db.AddDevicePart(device)
	fmt.Println("collectdevice: ", device, "err:", err)
}