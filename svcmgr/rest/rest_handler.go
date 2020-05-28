package rest

import (
	"cmpService/common/lib"
	"cmpService/common/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

const defaultField = "device_code"
const ServerTableName = "device_server_tb"
const NetworkTableName = "device_network_tb"
const PartTableName = "device_part_tb"

func (h *Handler) GetCodes(c *gin.Context) {
	fmt.Println("Getcodes")
	if h.db == nil {
		return
	}
	codes, err := h.db.GetAllCodes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(codes)

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
	for i, v := range codes {
		fmt.Println(i, v.Name)
		list = append(list, v.Name)
	}

	//fmt.Println(list)
	c.JSON(http.StatusOK, list)
}

func (h *Handler) AddCode(c *gin.Context) {
	if h.db == nil {
		return
	}
	var code models.Code
	err := c.ShouldBindJSON(&code)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":lib.RestFailConvertData})
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
		c.JSON(http.StatusBadRequest, gin.H{"error":lib.RestFailConvertData})
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
		c.JSON(http.StatusBadRequest, gin.H{"error":lib.RestAbnormalParam})
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
		c.JSON(http.StatusBadRequest, gin.H{"error":lib.RestAbnormalParam})
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
	fmt.Println("GetDevicesByList")
	if h.db == nil {
		return
	}
	deviceType := c.Param("type")

	f := c.Param("outFlag")
	outFlag, err := strconv.Atoi(f)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":lib.RestAbnormalParam})
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

	fmt.Println("type : ", deviceType, ", outFlag : ", outFlag)

	if deviceType == "server" {
		c.JSON(http.StatusOK, devicesServer)
	} else if deviceType == "network" {
		c.JSON(http.StatusOK, devicesNetwork)
	} else if deviceType == "part" {
		c.JSON(http.StatusOK, devicesPart)
	}
}

func (h *Handler) GetDevicesByIdx(c *gin.Context) {
	fmt.Println("GetDevicesByIdx")
	if h.db == nil {
		return
	}
	deviceType := c.Param("type")

	f := c.Param("idx")
	idx, err := strconv.Atoi(f)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":lib.RestAbnormalParam})
		return
	}

	devicesServer, err := h.db.GetDeviceServer(deviceType, idx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	devicesNetwork, err := h.db.GetDeviceNetwork(deviceType, idx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	devicesPart, err := h.db.GetDevicePart(deviceType, idx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("type : ", deviceType, ", idx : ", idx)

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
	//fmt.Println("[###] %v", devices)
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

// Mornitoring
func (h *Handler) GetDevicesMonitoring(c *gin.Context) {

}

func (h *Handler) AddDevicesMonitoring(c *gin.Context) {
	var msg DeviceMonitoringRequest
	err := c.ShouldBindJSON(&msg)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":lib.RestFailConvertData})
		return
	}
	c.JSON(http.StatusOK, msg)
}

// Add Device
func (h *Handler) AddDevice(c *gin.Context) {
	var dc interface{}
	device := c.Param("type")
	switch device {
	case "server":
		dc = new(models.DeviceServer)
	case "network":
		dc = new(models.DeviceNetwork)
	case "part":
		dc = new(models.DevicePart)
	}
	tableName := GetDeviceTable(c.Param("type"))

	err := c.ShouldBindJSON(&dc)
	//fmt.Println(dc)		// data check
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}

	err = h.db.AddDevice(dc, tableName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "OK")
}

// Update Device outFlag
func (h *Handler) UpdateOutFlag(c *gin.Context) {
	tableName := GetDeviceTable(c.Param("type"))

	flag, _ := strconv.Atoi(c.Param("outFlag"))
	values, err := JsonUnmarshal(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":lib.RestFailConvertData})
		return
	}
	//fmt.Println(values)

	data := values["idx"].(string)

	err = h.db.UpdateOutFlag(data, tableName, flag)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "OK")
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

func MakeDeviceCode(h *Handler, dc interface{}) (string, error) {
	data, dbErr := h.db.GetLastDeviceCode(dc)
	if dbErr != nil {
		return "", dbErr
	}
	code := data.(models.DeviceServer).DeviceCode
	prefix := code[:2]
	num, err := strconv.Atoi(code[2:])
	if err != nil {
		return "", err
	}
	num++
	return fmt.Sprintf("%s%5d", prefix, num), nil
}