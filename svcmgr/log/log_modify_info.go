package log

import (
	"cmpService/common/lib"
	"cmpService/common/models"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"time"
)

type CompareInfo struct {
	NewDevice  interface{}
	OldDevice  interface{}
	DeviceType string
	DeviceCode string
}

type ChangeInfo struct {
	Field     string
	OldStatus string
	NewStatus string
}

func DeviceInfoModify(info CompareInfo) error {
	changeInfo := ChangeInfo{}
	// reflect.ValueOf(dc).Elem().Field(0) => DeviceCommonResponse
	// reflect.ValueOf(dc).Elem().Field(1) => Ip
	// ......
	newElem := reflect.ValueOf(info.NewDevice).Elem()
	oldElem := reflect.ValueOf(info.OldDevice)
	//fmt.Printf("new %+v\n", newElem) //todo

	for i := 0; i < newElem.NumField(); i++ {
		// nested struct check
		if newElem.Field(i).Type().String() == "models.DeviceCommon" {
			for j := 0; j < newElem.Field(i).NumField(); j++ {
				// empty value remove & find change value
				if newElem.Field(i).Field(j).String() == "" ||
					newElem.Field(i).Field(j).Interface() == oldElem.Field(i).Field(j).Interface() {
					continue
				}

				changeInfo.NewStatus, changeInfo.OldStatus = SetLogValue(
					newElem.Field(i).Field(j).Interface(),
					oldElem.Field(i).Field(j).Interface())
				changeInfo.Field = ConvertFieldName(newElem.Field(i).Type().Field(j).Name)
				if changeInfo.Field == "" {
					continue
				}
				StoreLog(info, changeInfo)
				changeInfo = ChangeInfo{} // init struct
			}
		} else {
			// extra field check
			// empty value remove & find change value
			if newElem.Field(i).String() == "" ||
				newElem.Field(i).Interface() == oldElem.Field(i).Interface() {
				continue
			}
			changeInfo.NewStatus, changeInfo.OldStatus = SetLogValue(
				newElem.Field(i).Interface(),
				oldElem.Field(i).Interface())
			changeInfo.Field = ConvertFieldName(newElem.Type().Field(i).Name)
		}
		StoreLog(info, changeInfo)
		changeInfo = ChangeInfo{} // init struct
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

func SetLogValue(new interface{}, old interface{}) (newVal string, oldVal string) {
	if new == nil || old == nil {
		return "", ""
	}
	fmt.Printf(reflect.TypeOf(new).Kind().String()) //todo
	switch reflect.TypeOf(new).Kind() {
	case reflect.Int:
		newVal = strconv.Itoa(int(reflect.ValueOf(new).Int()))
		oldVal = strconv.Itoa(int(reflect.ValueOf(old).Int()))
	case reflect.String:
		newVal = reflect.ValueOf(new).String()
		oldVal = reflect.ValueOf(old).String()
	case reflect.Bool:
		newVal = strconv.FormatBool(reflect.ValueOf(new).Bool())
		oldVal = strconv.FormatBool(reflect.ValueOf(old).Bool())
	default:
		return "", ""
	}
	return newVal, oldVal
}

func StoreLog(info CompareInfo, v ChangeInfo) error {
	var log models.DeviceLog

	// empty value check
	if v.Field == "" {
		return nil
	}

	switch info.DeviceType {
	case "server":
		ds, ok := info.OldDevice.(models.DeviceServer)
		if !ok {
			return errors.New("Can't data parse.\n")
		}
		log = models.DeviceLog{
			DeviceCode:   ds.DeviceCode, //todo
			WorkCode:     lib.ChangeInformation,
			LogLevel:     lib.LevelInfo,
			RegisterId:   ds.RegisterId, //todo
			Field:        v.Field,
			OldStatus:    v.OldStatus,
			NewStatus:    v.NewStatus,
			RegisterDate: time.Now(),
		}
		lib.LogWarn("%+v\n", log)
	case "network":
		dn, ok := info.OldDevice.(models.DeviceNetwork)
		if !ok {
			return errors.New("Can't data parse.\n")
		}
		fmt.Println(dn)
		log = models.DeviceLog{
			DeviceCode:   dn.DeviceCode,
			WorkCode:     lib.ChangeInformation,
			LogLevel:     lib.LevelInfo,
			RegisterId:   dn.RegisterId,
			Field:        v.Field,
			OldStatus:    v.OldStatus,
			NewStatus:    v.NewStatus,
			RegisterDate: time.Now(),
		}
		lib.LogWarn("%+v\n", log)
	case "part":
		dp, ok := info.OldDevice.(models.DevicePart)
		if !ok {
			return errors.New("Can't data parse.\n")
		}
		log = models.DeviceLog{
			DeviceCode:   dp.DeviceCode,
			WorkCode:     lib.ChangeInformation,
			LogLevel:     lib.LevelInfo,
			RegisterId:   dp.RegisterId,
			Field:        v.Field,
			OldStatus:    v.OldStatus,
			NewStatus:    v.NewStatus,
			RegisterDate: time.Now(),
		}
		lib.LogWarn("%+v\n", log)
	default:
		return errors.New("Device type is invalid.\n")
	}

	err := AutoAddLog(log)
	if err != nil {
		return fmt.Errorf("[RegisterDeviceLog] error : %s\n", err)
	}
	return nil
}
