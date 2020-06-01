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

func AutoAddLog(log models.DeviceLog) error{
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
