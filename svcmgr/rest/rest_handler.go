package rest

import (
	"cmpService/common/lib"
	"cmpService/common/models"
	"cmpService/svcmgr/log"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

const defaultField = "device_code"
const ServerTableName = "device_server_tb"
const NetworkTableName = "device_network_tb"
const PartTableName = "device_part_tb"

func (h *Handler) GetCodes(c *gin.Context) {
	//fmt.Println("Getcodes")
	if h.db == nil {
		return
	}
	codes, err := h.db.GetAllCodes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, codes)
}

func (h *Handler) GetSubCodes(c *gin.Context) {
	if h.db == nil {
		return
	}
	subcodes, err := h.db.GetAllSubCodes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, subcodes)
}

func (h *Handler) GetCodeList(c *gin.Context) {
	if h.db == nil {
		return
	}
	code := c.Param("code")
	subCode := c.Param("subcode")
	codes, err := h.db.GetCodeList(code, subCode)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var list []string
	for _, v := range codes {
		//fmt.Println(i, v.Name)
		list = append(list, v.Name)
	}

	//fmt.Println(list)
	c.JSON(http.StatusOK, list)
}

func (h *Handler) GetSubCodeList(c *gin.Context) {
	if h.db == nil {
		return
	}
	cIdx := strings.Split(c.Param("c_idx"), ",")
	subCodes, err := h.db.GetSubCodeList(cIdx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, subCodes)
}

func (h *Handler) AddCode(c *gin.Context) {
	if h.db == nil {
		return
	}
	var code models.Code
	err := c.ShouldBindJSON(&code)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": lib.RestFailConvertData})
		return
	}
	code, err = h.db.AddCode(code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, code)
}

func (h *Handler) AddSubCode(c *gin.Context) {
	if h.db == nil {
		return
	}
	var subcode models.SubCode
	err := c.ShouldBindJSON(&subcode)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": lib.RestFailConvertData})
		return
	}
	subcode, err = h.db.AddSubCode(subcode)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, subcode)
}

func (h *Handler) DeleteCode(c *gin.Context) {
	if h.db == nil {
		return
	}
	p := c.Param("id")
	id, err := strconv.Atoi(p)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": lib.RestAbnormalParam})
		return
	}
	code := models.Code{
		CodeID: uint(id),
	}
	code, err = h.db.DeleteCode(code)
	c.JSON(http.StatusOK, code)
}

func (h *Handler) DeleteSubCode(c *gin.Context) {
	if h.db == nil {
		return
	}
	p := c.Param("id")
	id, err := strconv.Atoi(p)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": lib.RestAbnormalParam})
		return
	}
	code := models.Code{
		CodeID: uint(id),
	}
	code, err = h.db.DeleteCode(code)
	c.JSON(http.StatusOK, code)
}

func (h *Handler) DeleteCodes(c *gin.Context) {
	if h.db == nil {
		return
	}
	err := h.db.DeleteCodes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, nil)
}

func (h *Handler) DeleteSubCodes(c *gin.Context) {
	if h.db == nil {
		return
	}
	err := h.db.DeleteSubCodes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, nil)
}

// Device
func (h *Handler) GetDevicesByList(c *gin.Context) {
	//fmt.Println("GetDevicesByList")
	if h.db == nil {
		return
	}
	deviceType := c.Param("type")

	f := c.Param("outFlag")
	outFlag, err := strconv.Atoi(f)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": lib.RestAbnormalParam})
		return
	}

	devicesServer, err := h.db.GetAllDevicesServer(deviceType, outFlag)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	devicesNetwork, err := h.db.GetAllDevicesNetwork(deviceType, outFlag)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	devicesPart, err := h.db.GetAllDevicesPart(deviceType, outFlag)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//fmt.Println("type : ", deviceType, ", outFlag : ", outFlag)

	if deviceType == "server" {
		c.JSON(http.StatusOK, devicesServer)
	} else if deviceType == "network" {
		c.JSON(http.StatusOK, devicesNetwork)
	} else if deviceType == "part" {
		c.JSON(http.StatusOK, devicesPart)
	}
}

// Search Devices
// With join
// default field : device_code
func (h *Handler) GetDevicesByCode(c *gin.Context) {
	if h.db == nil {
		return
	}
	deviceType := c.Param("type")
	field := c.Param("field")
	if field == "" {
		field = defaultField
	}
	code := c.Param("value")
	devices, err := h.db.GetDeviceWithJoin(deviceType, field, code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	_ = ConvertSplaString(h, devices, 0, deviceType) // no error check

	//fmt.Println("[###] %+v", devices)
	c.JSON(http.StatusOK, devices)
}

// Without join
func (h *Handler) GetDeviceWithoutJoin(c *gin.Context) {
	if h.db == nil {
		return
	}
	deviceType := c.Param("type")
	code := c.Param("value")
	devices, err := h.db.GetDeviceWithoutJoin(deviceType, code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	//fmt.Println("[###] %v", devices)
	c.JSON(http.StatusOK, devices)
}

// Search device
func (h *Handler) GetDevicesForSearch(c *gin.Context) {
	if h.db == nil {
		return
	}
	device := c.Param("type")
	mapDevice, err := JsonUnmarshal(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	convertData := ConvertDeviceData(mapDevice, device, "")
	if convertData == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": lib.RestAbnormalParam})
		return
	}

	switch device {
	case "server":
		devices, err := h.db.GetDevicesServerForSearch(*convertData.(*models.DeviceServer))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		for i, _ := range devices {
			_ = ConvertSplaString(h, &devices, i, device) // no error check
		}
		c.JSON(http.StatusOK, devices)
	case "network":
		devices, err := h.db.GetDevicesNetworkForSearch(*convertData.(*models.DeviceNetwork))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, devices)
	case "part":
		devices, err := h.db.GetDevicesPartForSearch(*convertData.(*models.DevicePart))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, devices)
	}
	return
}

// Search device log
func (h *Handler) GetDevicesByLog(c *gin.Context) {
	if h.db == nil {
		return
	}
	deviceType := c.Param("type")
	deviceCode := c.Param("value")

	logs, err := h.db.GetDeviceLogs(deviceCode)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	//fmt.Println("[###] %v", comments)
	c.JSON(http.StatusOK, logs)

	switch deviceType {
	case "server":
	case "network":
	case "part":
	}
	return
}

// Mornitoring
func (h *Handler) GetDevicesMonitoring(c *gin.Context) {

}

func (h *Handler) AddDevicesMonitoring(c *gin.Context) {
	var msg DeviceMonitoringRequest
	err := c.ShouldBindJSON(&msg)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": lib.RestFailConvertData})
		return
	}
	c.JSON(http.StatusOK, msg)
}

// Add Device
func (h *Handler) AddDevice(c *gin.Context) {
	device := c.Param("type")

	code, err := MakeDeviceCode(h, device)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	mapDevice, err := JsonUnmarshal(c.Request.Body)
	//fmt.Printf("%+v [%+v]\n", mapDevice, err)
	convertData := ConvertDeviceData(mapDevice, device, code)
	if err != nil || convertData == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": lib.RestAbnormalParam})
		return
	}

	fmt.Printf("%+v\n", convertData)
	err = h.db.AddDevice(convertData, device)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err = log.DeviceRegLog(convertData, device); err != nil {
		lib.LogWarn("%s\n", err)
	}

	c.JSON(http.StatusOK, "OK")
}

// Update Device
func (h *Handler) UpdateDevice(c *gin.Context) {
	device := c.Param("type")
	code := c.Param("deviceCode")
	tableName := GetDeviceTable(device)

	mapDevice, err := JsonUnmarshal(c.Request.Body)
	convertData := ConvertDeviceData(mapDevice, device, code)
	if err != nil || convertData == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": lib.RestAbnormalParam})
		return
	}

	oldDevice := log.GetDevice(device, code)
	data, err := h.db.UpdateDevice(convertData, tableName, code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 변경된 장비 정보 LOG 로 등록
	info := log.CompareInfo{
		NewDevice:  data,
		OldDevice:  oldDevice,
		DeviceType: device,
		DeviceCode: code,
	}

	if err = log.DeviceInfoModify(info); err != nil {
		lib.LogWarn("%s\n", err)
	}

	c.JSON(http.StatusOK, "OK")
}

// Update Device outFlag
func (h *Handler) UpdateOutFlag(c *gin.Context) {
	tableName := GetDeviceTable(c.Param("type"))

	values, err := JsonUnmarshal(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": lib.RestFailConvertData})
		return
	}
	//lib.LogInfo("[values] %s\n", values)

	var flag int
	if val, ok := values["outFlag"]; ok {
		flag, err = strconv.Atoi(val.(string))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": lib.RestAbnormalParam})
			return
		}
	}
	userId := values["userId"].(string)
	data := strings.Split(values["deviceCode"].(string), ",")
	err = h.db.UpdateOutFlag(data, tableName, flag)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err = log.DeviceUpdateOutFlag(data, c.Param("type"), int(flag), userId); err != nil {
		lib.LogWarn("%s\n", err)
	}

	c.JSON(http.StatusOK, "OK")
}
