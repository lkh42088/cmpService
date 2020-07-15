package log

import (
	"cmpService/common/lib"
	"cmpService/common/models"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
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
	fmt.Printf("😡 new %+v\n", newElem) //todo
	fmt.Printf("😡 old %+v\n", oldElem) //todo



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

				changeInfo.NewStatus, changeInfo.OldStatus = SetLogValue(
					GetCodeLogValue(changeInfo.NewStatus, newElem.Field(i).Type().Field(j).Name),
					GetCodeLogValue(changeInfo.OldStatus, newElem.Field(i).Type().Field(j).Name))

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


			changeInfo.NewStatus, changeInfo.OldStatus = SetLogValue(
				GetCodeLogValue(changeInfo.NewStatus, newElem.Type().Field(i).Name),
				GetCodeLogValue(changeInfo.OldStatus, newElem.Type().Field(i).Name))
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
	fmt.Println(reflect.TypeOf(new).Kind().String()) //todo
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

func GetCodeLogValue(val string, field string) string {
	var returnVal string
	comma := ","
	Code := models.Code{}
	SubCode := models.SubCode{}
	db, _ := SetMariaDBForLog()
	defer db.Close()

	switch field {
	case "Model", "Rack":
		SubCode, _ = db.GetSubCodeByIdx(val)
		returnVal = SubCode.Name
	case "Manufacture", "DeviceType", "Ownership", "OwnershipDiv", "IDC", "RackCode", "Size":
		Code, _ = db.GetCodeByIdx(val)
		returnVal = Code.Name
	/*case "Customer":
		returnVal = "고객사명"
	case "OwnerCompany":
		returnVal = "소유 회사"
	case "MonitoringFlag":
		returnVal = "모니터링 여부"
	case "MonitoringMethod":
		returnVal = "모니터링 방식"
	case "Size":
		returnVal = "크기"*/
	case "Spla":
		if strings.Contains(val, "|") {
			splaArray := strings.Split(val, "|")

			for i := 0; i < len(splaArray); i++ {
				if len(splaArray[i]) != 0 {
					if i == 0 {
						comma = ""
					} else {
						comma = ","
					}

					Code, _ = db.GetCodeByIdx(splaArray[i])
					returnVal = returnVal + comma +  Code.Name
				}
			}
		}
	default:
		returnVal = val
	}

	return returnVal
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
