package log

import (
	"cmpService/common/lib"
	"cmpService/common/models"
	"errors"
	"time"
)

func DeviceRegLog(dc interface{}, deviceType string) error {
	var log models.DeviceLog
	switch deviceType {
	case "server":
		ds, ok := dc.(*models.DeviceServer)
		if !ok {
			return errors.New("Can't data parse.\n")
		}
		log = models.DeviceLog{
			DeviceCode: ds.DeviceCode,
			WorkCode:   lib.RegisterDevice,
			LogLevel:   lib.LevelInfo,
			RegisterId: ds.RegisterId,
			RegisterDate: time.Now(),
		}
		lib.LogInfo("%v\n", log)
	case "network":
		dn, ok := dc.(*models.DeviceNetwork)
		if !ok {
			return errors.New("Can't data parse.\n")
		}
		log = models.DeviceLog{
			DeviceCode: dn.DeviceCode,
			WorkCode:   lib.RegisterDevice,
			LogLevel:   lib.LevelInfo,
			RegisterId: dn.RegisterId,
			RegisterDate: time.Now(),
		}
		lib.LogInfo("%v\n", log)
	case "part":
		dp, ok := dc.(*models.DevicePart)
		if !ok {
			return errors.New("Can't data parse.\n")
		}
		log = models.DeviceLog{
			DeviceCode: dp.DeviceCode,
			WorkCode:   lib.RegisterDevice,
			LogLevel:   lib.LevelInfo,
			RegisterId: dp.RegisterId,
			RegisterDate: time.Now(),
		}
		lib.LogInfo("%v\n", log)
	default:
		return errors.New("Device type is invalid.\n")
	}

	err := AutoAddLog(log)
	if err != nil {
		return errors.New("[RegisterDeviceLog] Failed to insert log message in DB")
	}
	return nil
}
