package mariadblayer

import (
	"cmpService/common/models"
	"fmt"
	"testing"
	"time"
)

func TestDeviceAddEntry_Server(t *testing.T) {
	db, err := getTestDb()
	if err != nil {
		return
	}

	device := models.DeviceServer{
		DeviceCommon: models.DeviceCommon{
			OutFlag:          false,
			CommentCnt:       0,
			CommentLastDate:  time.Now(),
			RegisterId:       "nhj_id",
			RegisterDate:     time.Now(),
			DeviceCode:       "CBSNUBES01_1",
			Model:            1,
			Contents:         "텍스트 내용",
			Customer:         "NB",
			Manufacture:      3,
			DeviceType:       4,
			WarehousingDate:  "20200423",
			RentDate:         "20200423|20200512",
			Ownership:        "1",
			OwnerCompany:     "소유업체명~~",
			HwSn:             "hw/sn",
			IDC:              5,
			Rack:             6,
			Cost:             "원가",
			Purpose:          "용도",
			MonitoringFlag:   false,
			MonitoringMethod: 9,
		},
		Ip:     "255.255.255.2",
		Size:   7,
		Spla:   "8",
		Cpu:    "cpu",
		Memory: "memory",
		Hdd:    "hdd",
	}
	device, err = db.AddDeviceServer(device)
	fmt.Println("collectdevice: ", device, "err:", err)

	device = models.DeviceServer{
		DeviceCommon: models.DeviceCommon{
			OutFlag:          false,
			CommentCnt:       0,
			CommentLastDate:  time.Now(),
			RegisterId:       "yjs_id",
			RegisterDate:     time.Now(),
			DeviceCode:       "CBSNUBES02_1",
			Model:            1,
			Contents:         "텍스트 내용~",
			Customer:         "NB",
			Manufacture:      3,
			DeviceType:       4,
			WarehousingDate:  "20200423",
			RentDate:         "20200423|20200512",
			Ownership:        "1",
			OwnerCompany:     "소유업체명~~UU",
			HwSn:             "hw/sn",
			IDC:              5,
			Rack:             6,
			Cost:             "원가",
			Purpose:          "용도",
			MonitoringFlag:   false,
			MonitoringMethod: 9,
		},
		Ip:     "255.255.255.2",
		Size:   7,
		Spla:   "8",
		Cpu:    "cpu",
		Memory: "memory",
		Hdd:    "hdd",
	}
	device, err = db.AddDeviceServer(device)
	fmt.Println("collectdevice: ", device, "err:", err)

	device = models.DeviceServer{
		DeviceCommon: models.DeviceCommon{
			OutFlag:          false,
			CommentCnt:       0,
			CommentLastDate:  time.Now(),
			RegisterId:       "pms_id",
			RegisterDate:     time.Now(),
			DeviceCode:       "CBSNUBES03_1",
			Model:            1,
			Contents:         "텍스트 내용~",
			Customer:         "NB",
			Manufacture:      3,
			DeviceType:       4,
			WarehousingDate:  "20200423",
			RentDate:         "20200423|20200512",
			Ownership:        "1",
			OwnerCompany:     "소유업체명~~YAY",
			HwSn:             "hw/sn",
			IDC:              5,
			Rack:             6,
			Cost:             "원가",
			Purpose:          "용도",
			MonitoringFlag:   false,
			MonitoringMethod: 9,
		},
		Ip:     "255.255.255.2",
		Size:   7,
		Spla:   "8",
		Cpu:    "cpu",
		Memory: "memory",
		Hdd:    "hdd",
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
		DeviceCommon: models.DeviceCommon{
			OutFlag:          false,
			CommentCnt:       0,
			CommentLastDate:  time.Now(),
			RegisterId:       "EE_id",
			RegisterDate:     time.Now(),
			DeviceCode:       "CBSNETWORK03",
			Model:            1,
			Contents:         "텍스트 내용_NET3",
			Customer:         "NB",
			Manufacture:      3,
			DeviceType:       4,
			WarehousingDate:  "20200423",
			RentDate:         "20200423|20200512",
			Ownership:        "1",
			OwnerCompany:     "EE 소유업체명~~",
			HwSn:             "hw/sn",
			IDC:              5,
			Rack:             6,
			Cost:             "원가",
			Purpose:          "용도",
			MonitoringFlag:   false,
			MonitoringMethod: 9,
		},
		Ip:              "255.255.255.22",
		Size:            7,
		FirmwareVersion: "만료",
	}
	device, err = db.AddDeviceNetwork(device)
	fmt.Println("collectdevice: ", device, "err:", err)

	device = models.DeviceNetwork{
		DeviceCommon: models.DeviceCommon{
			OutFlag:          false,
			CommentCnt:       0,
			CommentLastDate:  time.Now(),
			RegisterId:       "JDG_id",
			RegisterDate:     time.Now(),
			DeviceCode:       "CBSNETWORK04",
			Model:            1,
			Contents:         "텍스트 내용_NET4",
			Customer:         "NB",
			Manufacture:      3,
			DeviceType:       4,
			WarehousingDate:  "20200423",
			RentDate:         "20200423|20200512",
			Ownership:        "1",
			OwnerCompany:     "소유업체명★",
			HwSn:             "hw/sn",
			IDC:              5,
			Rack:             6,
			Cost:             "JDG원TO가",
			Purpose:          "JDG용TO도",
			MonitoringFlag:   false,
			MonitoringMethod: 9,
		},
		Ip:              "255.255.255.12",
		Size:            7,
		FirmwareVersion: "만료",
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
		DeviceCommon: models.DeviceCommon{
			OutFlag:          false,
			CommentCnt:       0,
			CommentLastDate:  time.Now(),
			RegisterId:       "jjh_id",
			RegisterDate:     time.Now(),
			DeviceCode:       "CBSPART01",
			Model:            1,
			Contents:         "텍스트 내용_PART1",
			Customer:         "NB",
			Manufacture:      3,
			DeviceType:       4,
			WarehousingDate:  "20200423",
			RentDate:         "20200423|20200512",
			Ownership:        "1",
			OwnerCompany:     "소유업체명~~YOYO",
			HwSn:             "hw/sn",
			IDC:              5,
			Rack:             6,
			Cost:             "원가",
			Purpose:          "용도",
			MonitoringFlag:   false,
			MonitoringMethod: 9,
		},
		Warranty: "warranty",
	}
	device, err = db.AddDevicePart(device)
	fmt.Println("collectdevice: ", device, "err:", err)
}

func TestDBORM_GetLastDeviceCode(t *testing.T) {
	db, err := getTestLkhDb()
	if err != nil {
		return
	}
	fmt.Println(db.GetLastDeviceCodeInServer())
	fmt.Println(db.GetLastDeviceCodeInNetwork())
	fmt.Println(db.GetLastDeviceCodeInPart())
}
