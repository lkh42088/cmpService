package rest

import (
	"cmpService/common/models"
	"cmpService/dbmigrator/convert"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

func JsonUnmarshal(body io.ReadCloser) (m map[string]interface{}, err error) {
	bodyByte, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, errors.New("Request body is invalid.\n")
	}
	mapData := make(map[string]interface{})
	err = json.Unmarshal(bodyByte,&mapData)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return mapData, nil
}

func GetDeviceTable(device string) string {
	var tableName string

	switch device {
	case "server":
		tableName = ServerTableName
	case "network":
		tableName = NetworkTableName
	case "part":
		tableName = PartTableName
	}
	return tableName
}

func MakeDeviceCode(h *Handler, device string, dc *interface{}) (string, error) {
	var code string
	switch device {
	case "server":
		data, _ := h.db.GetLastDeviceCodeInServer()
		code = data.DeviceCode
	case "network":
		data, _ := h.db.GetLastDeviceCodeInNetwork()
		code = data.DeviceCode
	case "part":
		data, _ := h.db.GetLastDeviceCodeInPart()
		code = data.DeviceCode
	}
	prefix := code[:3]
	num, err := strconv.Atoi(code[3:])
	if err != nil {
		return "", err
	}
	num++
	code = fmt.Sprintf("%s%5d", prefix, num)
	code = strings.Replace(code, " ", "0", -1) // remove space
	//lib.LogWarn("[NEW Code] %s\n", code)
	return code, nil
}

func ConvertDeviceData(device map[string]interface{}, deviceType string, code string) interface{} {
	if device == nil {
		return nil
	}
	switch deviceType {
	case "server":
		dc := new(models.DeviceServer)
		dc.DeviceCommon = ConvertDeviceCommon(device, code)
		if val, ok := device["ip"]; ok {
			dc.Ip = val.(string)
		}
		if val, ok := device["size"]; ok {
			dc.Size, _ = strconv.Atoi(val.(string))
		}
		if val, ok := device["spla"]; ok {
			dc.Spla = val.(string)
		}
		if val, ok := device["cpu"]; ok {
			dc.Cpu = val.(string)
		}
		if val, ok := device["memory"]; ok {
			dc.Memory = val.(string)
		}
		if val, ok := device["hdd"]; ok {
			dc.Hdd = val.(string)
		}
		if val, ok := device["rackTag"]; ok {
			dc.RackTag = val.(string)
		}
		if val, ok := device["rackLoc"]; ok {
			dc.RackLoc, _ = strconv.Atoi(val.(string))
		}
		fmt.Println("device code:", dc.DeviceCode)
		return dc
	case "network":
		dc := new(models.DeviceNetwork)
		dc.DeviceCommon = ConvertDeviceCommon(device, code)
		if val, ok := device["ip"]; ok {
			dc.Ip = val.(string)
		}
		if val, ok := device["size"]; ok {
			dc.Size, _ = strconv.Atoi(val.(string))
		}
		if val, ok := device["firmwareVersion"]; ok {
			dc.FirmwareVersion = val.(string)
		}
		if val, ok := device["rackTag"]; ok {
			dc.RackTag = val.(string)
		}
		if val, ok := device["rackLoc"]; ok {
			dc.RackLoc, _ = strconv.Atoi(val.(string))
		}
		return dc
	case "part":
		dc := new(models.DevicePart)
		dc.DeviceCommon = ConvertDeviceCommon(device, code)
		if val, ok := device["warranty"]; ok {
			dc.Warranty = val.(string)
		}
		if val, ok := device["rackCode"]; ok {
			dc.RackCode, _ = strconv.Atoi(val.(string))
		}
		return dc
	default:
		return nil
	}
	return nil
}

func ConvertDeviceCommon(device map[string]interface{}, code string) models.DeviceCommon {
	if device == nil {
		fmt.Println("test")
		return models.DeviceCommon{}
	}
	fmt.Println(code)

	var commentLastDate time.Time
	var dc models.DeviceCommon

	if val, ok := device["outFlag"]; ok {
		if val != "" {
			if flag, _ := strconv.Atoi(val.(string)); flag == 1 {
				dc.OutFlag = true
			} else {
				dc.OutFlag = false
			}
		}
	}
	if val, ok := device["commentCnt"]; ok {
		dc.CommentCnt, _ = strconv.Atoi(val.(string))
	}

	if val, ok := device["commentLastDate"]; ok {
		commentLastDate, _ = time.Parse(convert.TimeFormat, val.(string))
	}
	dc.CommentLastDate = commentLastDate

	if val, ok := device["registerId"]; ok {
		dc.RegisterId = val.(string)
	}

	dc.RegisterDate = time.Now()
	dc.DeviceCode = code

	if val, ok := device["model"]; ok {
		dc.Model, _ = strconv.Atoi(val.(string))
	}
	if val, ok := device["contents"]; ok {
		dc.Contents = val.(string)
	}
	if val, ok := device["customer"]; ok {
		dc.Customer = val.(string)
	}
	if val, ok := device["manufacture"]; ok {
		dc.Manufacture, _ = strconv.Atoi(val.(string))
	}
	if val, ok := device["deviceType"]; ok {
		dc.DeviceType, _ = strconv.Atoi(val.(string))
	}
	if val, ok := device["warehousingDate"]; ok {
		dc.WarehousingDate = val.(string)
	}
	if val, ok := device["rentDate"]; ok {
		dc.RentDate = val.(string)
	}
	if val, ok := device["ownership"]; ok {
		dc.Ownership = val.(string)
	}
	if val, ok := device["ownershipDiv"]; ok {
		dc.OwnershipDiv = val.(string)
	}
	if val, ok := device["ownerCompany"]; ok {
		dc.OwnerCompany = val.(string)
	}
	if val, ok := device["hwSn"]; ok {
		dc.HwSn = val.(string)
	}
	if val, ok := device["idc"]; ok {
		dc.IDC, _ = strconv.Atoi(val.(string))
	}
	if val, ok := device["rack"]; ok {
		dc.Rack, _ = strconv.Atoi(val.(string))
	}
	if val, ok := device["cost"]; ok {
		dc.Cost = val.(string)
	}
	if val, ok := device["purpos"]; ok {
		dc.Purpos = val.(string)
	}
	if val, ok := device["monitoringFlag"]; ok {
		if val != "" {
			if flag, _ := strconv.Atoi(val.(string)); flag == 1 {
				dc.MonitoringFlag = true
			} else {
				dc.MonitoringFlag = false
			}
		}
	}
	if val, ok := device["monitoringMethod"]; ok {
		dc.MonitoringMethod, _ = strconv.Atoi(val.(string))
	}
	return dc
}