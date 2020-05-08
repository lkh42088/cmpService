package rest

import (
	"fmt"
	"net/http"
	"nubes/common/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

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

func (h *Handler) AddCode(c *gin.Context) {
	if h.db == nil {
		return
	}
	var code models.Code
	err := c.ShouldBindJSON(&code)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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
	/*	fmt.Println("GetDevicesByList1 : ",p)
		deviceType, err := strconv.Atoi(p)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
			return
		}
		fmt.Println("GetDevicesByList2")*/

	f := c.Param("outFlag")
	outFlag, err := strconv.Atoi(f)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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
	} else if string(deviceType) == "network" {
		c.JSON(http.StatusOK, devicesNetwork)
	} else if string(deviceType) == "part" {
		c.JSON(http.StatusOK, devicesPart)
	}
}

func (h *Handler) GetDevicesByIdx(c *gin.Context) {
	fmt.Println("GetDevicesByIdx")
	if h.db == nil {
		return
	}
	deviceType := c.Param("type")
	/*	fmt.Println("GetDevicesByList1 : ",p)
		deviceType, err := strconv.Atoi(p)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
			return
		}
		fmt.Println("GetDevicesByList2")*/

	f := c.Param("idx")
	idx, err := strconv.Atoi(f)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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
	} else if string(deviceType) == "network" {
		c.JSON(http.StatusOK, devicesNetwork)
	} else if string(deviceType) == "part" {
		c.JSON(http.StatusOK, devicesPart)
	}
}

// Mornitoring
func (h *Handler) GetDevicesMonitoring(c *gin.Context) {

}

func (h *Handler) AddDevicesMonitoring(c *gin.Context) {
	var msg DeviceMonitoringRequest
	err := c.ShouldBindJSON(&msg)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, msg)
}

