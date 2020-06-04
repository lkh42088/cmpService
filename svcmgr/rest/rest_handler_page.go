package rest

import (
	"cmpService/common/lib"
	"cmpService/common/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

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


