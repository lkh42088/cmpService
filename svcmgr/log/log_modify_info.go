package log

import (
	"cmpService/common/lib"
	"cmpService/common/models"
	"errors"
	"fmt"
	"reflect"
	"time"
)

func DeviceInfoModify(dc interface{}, deviceType string, code string) error {
	var log models.DeviceLog
	var field string
	var oldStatus string
	var newStatus string
	oldDevice := GetDevice(deviceType, code)
	//fmt.Printf("%+v\n", oldDevice)

	// reflect.ValueOf(dc).Elem().Field(0) => DeviceCommonResponse
	// reflect.ValueOf(dc).Elem().Field(1) => Ip
	// ......
	newElem := reflect.ValueOf(dc).Elem()
	oldElem := reflect.ValueOf(oldDevice)
	fmt.Printf("%+v\n", oldElem)

	for i := 0; i < newElem.NumField(); i++ {
		//fmt.Println(newElem.Field(0).NumField(), newElem.Field(i).Type())
		if newElem.Field(i).Type().String() == "models.DeviceCommon" {
			for j := 0; j < newElem.Field(i).NumField(); j++ {
				if newElem.Field(i).Field(j).String() == "" {
					continue
				}
				if reflect.DeepEqual(newElem.Field(i).Field(j), oldElem.Field(i).Field(j)) { //todo
					fmt.Printf("%+v\n", oldElem.Field(i).Field(j))
					continue
				}
				field = newElem.Field(i).Type().Field(j).Name  //todo
				oldStatus = oldElem.Field(i).Field(j).String() //todo
				newStatus = newElem.Field(i).Field(j).String()
				fmt.Println(newStatus, oldStatus, field)
			}
		} else {
			if newElem.Field(i).String() == "" {
				continue
			}
			if reflect.DeepEqual(newElem.Field(i), oldElem.Field(i)) {
				continue
			}
			field = newElem.Type().Field(i).Name  //todo
			oldStatus = oldElem.Field(i).String() //todo
			newStatus = newElem.Field(i).String()
			fmt.Println(newStatus, oldStatus, field)
		}
		switch deviceType {
		case "server":
			ds, ok := oldDevice.(models.DeviceServer)
			if !ok {
				return errors.New("Can't data parse.\n")
			}
			log = models.DeviceLog{
				DeviceCode:   ds.DeviceCode,
				WorkCode:     lib.ChangeInformation,
				LogLevel:     lib.LevelInfo,
				RegisterId:   ds.RegisterId,
				Field:        field,
				OldStatus:    oldStatus,
				NewStatus:    newStatus,
				RegisterDate: time.Now(),
			}
			lib.LogWarn("%v\n", log)
		case "network":
			dn, ok := oldDevice.(models.DeviceNetwork)
			if !ok {
				return errors.New("Can't data parse.\n")
			}
			fmt.Println(dn)
			log = models.DeviceLog{
				DeviceCode:   dn.DeviceCode,
				WorkCode:     lib.ChangeInformation,
				LogLevel:     lib.LevelInfo,
				RegisterId:   dn.RegisterId,
				Field:        field,
				OldStatus:    oldStatus,
				NewStatus:    newStatus,
				RegisterDate: time.Now(),
			}
			lib.LogWarn("%v\n", log)
		case "part":
			dp, ok := oldDevice.(models.DevicePart)
			if !ok {
				return errors.New("Can't data parse.\n")
			}
			log = models.DeviceLog{
				DeviceCode:   dp.DeviceCode,
				WorkCode:     lib.ChangeInformation,
				LogLevel:     lib.LevelInfo,
				RegisterId:   dp.RegisterId,
				Field:        field,
				OldStatus:    oldStatus,
				NewStatus:    newStatus,
				RegisterDate: time.Now(),
			}
			lib.LogWarn("%v\n", log)
		default:
			return errors.New("Device type is invalid.\n")
		}

		err := AutoAddLog(log)
		if err != nil {
			return fmt.Errorf("[RegisterDeviceLog] error : %s\n", err)
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
			ds, ok := dc.(models.DeviceServerResponse)
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
			lib.LogInfo("%v\n", log)
		case "network":
			dn, ok := dc.(models.DeviceNetworkResponse)
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
			lib.LogInfo("%v\n", log)
		case "part":
			dp, ok := dc.(models.DevicePart)
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
			lib.LogInfo("%v\n", log)
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
