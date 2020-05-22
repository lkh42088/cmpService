package rest

import (
	"cmpService/common/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) GetDevicesForPage(c *gin.Context) {
	//fmt.Println("GetDevicesForPage")
	if h.db == nil {
		return
	}

	// Parse params
	size, _ := strconv.Atoi(c.Param("size"))
	cnt, _ := strconv.Atoi(c.Param("checkcnt"))
	dir, _ := strconv.Atoi(c.Param("dir"))
	page := models.PageCreteria{
		DeviceType : c.Param("type"),
		OrderKey: c.Param("order"),
		Size: size,
		OutFlag: c.Param("outFlag"),
		Direction: dir,
		CheckCnt: cnt,	// current row counter
	}

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
		c.JSON(http.StatusOK, devicePage)
	}
}


