package rest

import (
	"cmpService/common/lib"
	"cmpService/common/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strconv"
)

/*
func (h *Handler) GetDevicesForPage(c *gin.Context) {
	if h.db == nil {
		return
	}

	// Parse params
	size, err := strconv.Atoi(c.Param("size"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":lib.RestAbnormalParam})
		return
	}
	cnt, err := strconv.Atoi(c.Param("checkcnt"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":lib.RestAbnormalParam})
		return
	}
	dir, err := strconv.Atoi(c.Param("dir"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":lib.RestAbnormalParam})
		return
	}
	page := models.PageCreteria{
		DeviceType : c.Param("type"),
		OrderKey: c.Param("order"),
		Size: size,
		OutFlag: c.Param("outFlag"),
		Direction: dir,
		CheckCnt: cnt,	// current row counter
	}
	fmt.Println("1. page:");
	page.String()

	switch page.DeviceType {
	case "server":
		devicePage, err := h.db.GetDevicesServerWithJoin(page)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		//for i, v := range devicePage.Devices {
		//	fmt.Printf("%d %v\n", i+1, v)
		//}
		fmt.Println("2. page:");
		devicePage.Page.String()
		c.JSON(http.StatusOK, devicePage)
	case "network":
		devicePage, err := h.db.GetDevicesNetworkWithJoin(page)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		//for i, v := range devicePage.Devices {
		//	fmt.Printf("%d %v\n", i+1, v)
		//}
		fmt.Println("2. page: .");
		devicePage.Page.String()
		c.JSON(http.StatusOK, devicePage)
	case "part":
		devicePage, err := h.db.GetDevicesPartWithJoin(page)
		if err != nil {

			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		//for i, v := range devicePage.Devices {
		//	fmt.Printf("%d %v\n", i+1, v)
		//}
		fmt.Println("2. page: ..");
		devicePage.Page.String()
		c.JSON(http.StatusOK, devicePage)
	}
}
*/

func (h *Handler) GetDevicesTypeCount(c *gin.Context) {
	if h.db == nil {
		return
	}

	deviceType := c.Param("type")

	mapDevice, err := JsonUnmarshal(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": lib.RestFailConvertData})
		return
	}

	convertData := ConvertDeviceData(mapDevice, deviceType, mapDevice["deviceCode"].(string))
	if convertData == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": lib.RestAbnormalParam})
		return
	}

	/*0 : 반입, 1 : 반출*/
	var outFlag string
	if mapDevice["operatingFlag"].(bool) {
		outFlag = "0"
	}

	if mapDevice["carryingFlag"].(bool) {
		if outFlag != "" {
			outFlag = outFlag + ", 1"
		} else {
			outFlag = "1"
		}
	}

	// 0 : false
	var rentPeriod string
	if mapDevice["rentPeriod"].(bool) {
		rentPeriod = "1"
	} else {
		rentPeriod = "0"
	}

	page := models.PageCreteria{
		DeviceType:     c.Param("type"),
		OutFlag:        outFlag,
		RentPeriodFlag: rentPeriod,
	}

	switch page.DeviceType {
	case "server":
		devicePage, err := h.db.GetDevicesTypeCountServerWithJoin(page,
			*convertData.(*models.DeviceServer))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, devicePage)
	case "network":
		devicePage, err := h.db.GetDevicesTypeCountNetworkWithJoin(page,
			*convertData.(*models.DeviceNetwork))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, devicePage)
	case "part":
		devicePage, err := h.db.GetDevicesTypeCountPartWithJoin(page,
			*convertData.(*models.DevicePart))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, devicePage)
	}
}

func (h *Handler) GetDevicesForPageSearchTest(c *gin.Context) {
	bodyByte, err := ioutil.ReadAll(c.Request.Body)
	fmt.Printf("%+v\n", bodyByte)

	mapTest, err := JsonUnmarshal(c.Request.Body)
	fmt.Printf("%+v [%+v]\n", mapTest, err)

	mapDevice, err := JsonUnmarshal(c.Request.Body)
	fmt.Printf("%+v [%+v]\n", mapDevice, err)
}

func (h *Handler) GetDevicesForPage(c *gin.Context) {
	if h.db == nil {
		return
	}

	// Parse params
	row, err := strconv.Atoi(c.Param("row"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": lib.RestAbnormalParam})
		return
	}
	curpage, err := strconv.Atoi(c.Param("page"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": lib.RestAbnormalParam})
		return
	}
	dir, err := strconv.Atoi(c.Param("dir"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": lib.RestAbnormalParam})
		return
	}
	offsetPage, err := strconv.Atoi(c.Param("offsetPage"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": lib.RestAbnormalParam})
		return
	}
	page := models.PageCreteria{
		DeviceType: c.Param("type"),
		OrderKey:   c.Param("order"),
		Row:        row,
		OutFlag:    c.Param("outFlag"),
		Direction:  dir,
		Page:       curpage,
		OffsetPage: offsetPage,
	}
	/*fmt.Println("1. page:");
	page.String()*/

	switch page.DeviceType {
	case "server":
		devicePage, err := h.db.GetDevicesServerWithJoin(page)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		//for i, v := range devicePage.Devices {
		//	fmt.Printf("%d %v\n", i+1, v)
		//}
		devicePage.Page.String()
		c.JSON(http.StatusOK, devicePage)
	case "network":
		devicePage, err := h.db.GetDevicesNetworkWithJoin(page)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		/*for i, v := range devicePage.Devices {
			fmt.Printf("%d %v\n", i+1, v)
		}*/
		devicePage.Page.String()
		c.JSON(http.StatusOK, devicePage)
	case "part":
		devicePage, err := h.db.GetDevicesPartWithJoin(page)
		if err != nil {

			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		//for i, v := range devicePage.Devices {
		//	fmt.Printf("%d %v\n", i+1, v)
		//}
		devicePage.Page.String()
		c.JSON(http.StatusOK, devicePage)
	}
}

func (h *Handler) GetDevicesForPageSearch(c *gin.Context) {
	if h.db == nil {
		return
	}
	deviceType := c.Param("type")

	// Parse params
	row, err := strconv.Atoi(c.Param("row"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": lib.RestAbnormalParam})
		return
	}
	curpage, err := strconv.Atoi(c.Param("page"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": lib.RestAbnormalParam})
		return
	}
	dir, err := strconv.Atoi(c.Param("dir"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": lib.RestAbnormalParam})
		return
	}
	offsetPage, err := strconv.Atoi(c.Param("offsetPage"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": lib.RestAbnormalParam})
		return
	}

	mapDevice, err := JsonUnmarshal(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": lib.RestFailConvertData})
		return
	}

	convertData := ConvertDeviceData(mapDevice, deviceType, mapDevice["deviceCode"].(string))
	if convertData == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": lib.RestAbnormalParam})
		return
	}

	/*	convertData = ConvertDeviceData(mapDevice, deviceType, mapDevice["customer"].(string))
		if convertData == nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": lib.RestAbnormalParam})
			return
		}

		fmt.Println("★★★★★★★★★★★★★★★★ : ", mapDevice["customer"].(string))*/

	/*0 : 반입, 1 : 반출*/
	var outFlag string
	if mapDevice["operatingFlag"].(bool) {
		outFlag = "0"
	}

	if mapDevice["carryingFlag"].(bool) {
		if outFlag != "" {
			outFlag = outFlag + ", 1"
		} else {
			outFlag = "1"
		}
	}

	// 0 : false
	var rentPeriod string
	if mapDevice["rentPeriod"].(bool) {
		rentPeriod = "1"
	} else {
		rentPeriod = "0"
	}

	page := models.PageCreteria{
		DeviceType:     c.Param("type"),
		OrderKey:       c.Param("order"),
		Row:            row,
		OutFlag:        outFlag,
		Direction:      dir,
		Page:           curpage,
		OffsetPage:     offsetPage,
		RentPeriodFlag: rentPeriod,
	}

	switch deviceType {
	case "server":
		devicePage, err := h.db.GetDevicesServerSearchWithJoin(page, *convertData.(*models.DeviceServer))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		/*for i, v := range devicePage.Devices {
			fmt.Printf("%d %v\n", i+1, v)
		}*/
		devicePage.Page.String()
		c.JSON(http.StatusOK, devicePage)
	case "network":
		devicePage, err := h.db.GetDevicesNetworkSearchWithJoin(page, *convertData.(*models.DeviceNetwork))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		devicePage.Page.String()
		c.JSON(http.StatusOK, devicePage)
	case "part":
		//devicePage, err := h.db.GetDevicesPartWithJoin(page)
		devicePage, err := h.db.GetDevicesPartSearchWithJoin(page, *convertData.(*models.DevicePart))
		if err != nil {

			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		devicePage.Page.String()
		c.JSON(http.StatusOK, devicePage)
	}
}
