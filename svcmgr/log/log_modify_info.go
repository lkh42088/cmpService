package log

import (
	"cmpService/common/lib"
	"cmpService/common/models"
	"errors"
	"reflect"
	"time"
)

func DeviceInfoModify(dc interface{}, deviceType string) error {
	var log models.DeviceLog
	for i := 0; i < reflect.TypeOf(dc).NumField(); i++ {
		switch deviceType {
		case "server":
			ds, ok := dc.(*models.DeviceServer)
			if !ok {
				return errors.New("Can't data parse.\n")
			}
			log = models.DeviceLog{
				DeviceCode:   ds.DeviceCode,
				WorkCode:     lib.RegisterDevice,
				LogLevel:     lib.LevelInfo,
				RegisterId:   ds.RegisterId,
				Field: "",
				OldStatus: "",
				NewStatus: "",
				RegisterDate: time.Now(),
			}
			lib.LogInfo("%s\n", log)
		case "network":
			dn, ok := dc.(*models.DeviceNetwork)
			if !ok {
				return errors.New("Can't data parse.\n")
			}
			log = models.DeviceLog{
				DeviceCode:   dn.DeviceCode,
				WorkCode:     lib.RegisterDevice,
				LogLevel:     lib.LevelInfo,
				RegisterId:   dn.RegisterId,
				RegisterDate: time.Now(),
			}
			lib.LogInfo("%s\n", log)
		case "part":
			dp, ok := dc.(*models.DevicePart)
			if !ok {
				return errors.New("Can't data parse.\n")
			}
			log = models.DeviceLog{
				DeviceCode:   dp.DeviceCode,
				WorkCode:     lib.RegisterDevice,
				LogLevel:     lib.LevelInfo,
				RegisterId:   dp.RegisterId,
				RegisterDate: time.Now(),
			}
			lib.LogInfo("%s\n", log)
		default:
			return errors.New("Device type is invalid.\n")
		}

		err := AutoAddLog(log)
		if err != nil {
			return errors.New("[RegisterDeviceLog] Failed to insert log message in DB")
		}
	}
	return nil
}

func DeviceUpdateOutFlag(data []string, deviceType string, outFlag int, userId string) error {
	var log models.DeviceLog
	var flag int
	if outFlag == 0 {
		flag = lib.ImportDevice
	} else {
		flag = lib.ExportDevice
	}
	for _, v := range data {
		dc := GetDevice(deviceType, v)
		switch deviceType {
		case "server":
			ds, ok := dc.(*models.DeviceServerResponse)
			if !ok {
				return errors.New("Can't data parse.\n")
			}
			log = models.DeviceLog{
				DeviceCode:   ds.DeviceCode,
				WorkCode:     flag,
				LogLevel:     lib.LevelInfo,
				RegisterId:   userId,
				RegisterDate: time.Now(),
			}
			lib.LogInfo("%s\n", log)
		case "network":
			dn, ok := dc.(*models.DeviceNetworkResponse)
			if !ok {
				return errors.New("Can't data parse.\n")
			}
			log = models.DeviceLog{
				DeviceCode:   dn.DeviceCode,
				WorkCode:     flag,
				LogLevel:     lib.LevelInfo,
				RegisterId:   userId,
				RegisterDate: time.Now(),
			}
			lib.LogInfo("%s\n", log)
		case "part":
			dp, ok := dc.(*models.DevicePartResponse)
			if !ok {
				return errors.New("Can't data parse.\n")
			}
			log = models.DeviceLog{
				DeviceCode:   dp.DeviceCode,
				WorkCode:     flag,
				LogLevel:     lib.LevelInfo,
				RegisterId:   userId,
				RegisterDate: time.Now(),
			}
			lib.LogInfo("%s\n", log)
		default:
			return errors.New("Device type is invalid.\n")
		}

		err := AutoAddLog(log)
		if err != nil {
			return errors.New("[RegisterDeviceLog] Failed to insert log message in DB")
		}
	}

	return nil
}
