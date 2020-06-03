package log

import (
	"cmpService/common/config"
	db2 "cmpService/common/db"
	"cmpService/common/lib"
	"cmpService/common/mariadblayer"
	"cmpService/common/models"
	config2 "cmpService/svcmgr/config"
	"errors"
	"fmt"
)

func SetMariaDBForLog() (db *mariadblayer.DBORM, err error) {
	cfg := config2.ReadConfig(config2.SvcmgrConfigPath)
	dbconfig, err := config.NewDBConfig("mysql",
		cfg.MariaUser, cfg.MariaPassword, cfg.MariaDb,
		cfg.MariaIp, 3306)
	if err != nil {
		fmt.Println("[SetMariaDB] Error:", err)
		return
	}
	dataSource := db2.GetDataSourceName(dbconfig)
	db, err = mariadblayer.NewDBORM(dbconfig.DBDriver, dataSource)
	if err != nil {
		fmt.Println("[SetMariaDB] Error:", err)
		return
	}
	return db, err
}

func AutoAddLog(log models.DeviceLog) error {
	//lib.LogWarn("Start AutoAddLog().\n")

	go func() error {
		db, _ := SetMariaDBForLog()
		defer db.Close()
		if log.DeviceCode == "" || log.WorkCode == 0 || log.RegisterId == "" {
			return errors.New(lib.RestAbnormalParam)
		}

		// Find register name
		user, err := db.GetUserByUserId(log.RegisterId)
		if err != nil {
			return err
		}
		log.RegisterName = user.Name
		//lib.LogWarn("Success to get user name.\n")

		// Add log
		if err := db.AddLog(log); err != nil {
			return err
		}
		return nil
	}()

	lib.LogInfo("[AutoAddLog] Log message stored(workCode=%d).\n", log.WorkCode)
	return nil
}

func GetDevice(deviceType string, deviceCode string) interface{} {
	db, _ := SetMariaDBForLog()
	defer db.Close()

	var device interface{}
	switch deviceType {
	case "server":
		device, _ = db.GetDeviceServer(deviceCode)
	case "network":
		device, _ = db.GetDeviceNetwork(deviceCode)
	case "part":
		device, _ = db.GetDevicePart(deviceCode)
	}

	return device
}

func ConvertFieldName(field string) string {
	var convField string

	switch field {
	case "Model":
		convField = "모델 코드"
	case "Contents":
		convField = "기타 사항"
	case "Customer":
		convField = "고객사명"
	case "Manufacture":
		convField = "제조사명"
	case "DeviceType":
		convField = "장비 구분"
	case "WarehousingDate":
		convField = "입고일"
	case "RentDate":
		convField = "임대 기간"
	case "Ownership":
		convField = "소유권"
	case "OwnershipDiv":
		convField = "소유 구분"
	case "OwnerCompany":
		convField = "소유 회사"
	case "HwSn":
		convField = "HW S/N"
	case "IDC":
		convField = "IDC"
	case "Rack":
		convField = "Rack"
	case "Cost":
		convField = "장비 원가"
	case "Purpos":
		convField = "용도"
	case "MonitoringFlag":
		convField = "모니터링 여부"
	case "MonitoringMethod":
		convField = "모니터링 방식"
	case "Ip":
		convField = "IP"
	case "Size":
		convField = "크기"
	case "Spla":
		convField = "SPLA"
	case "Cpu":
		convField = "CPU"
	case "Memory":
		convField = "MEMORY"
	case "HDD":
		convField = "HDD"
	case "RackTag":
		convField = "Rack 태그"
	case "RackLoc":
		convField = "Rack 내 위치 번호"
	case "FirmwareVersion":
		convField = "펌웨어 버전"
	case "Warranty":
		convField = "보증 기간"
	case "RackCode":
		convField = "Rack 사이즈 코드"
	}

	return convField
}
