package rest

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"nubes/common/models"
	"strconv"
)

func (h *Handler) GetDevicesForPage(c *gin.Context) {
	fmt.Println("GetDevicesForPage")
	if h.db == nil {
		return
	}

	// Parse params
	size, _ := strconv.Atoi(c.Param("size"))
	curpage, _ := strconv.Atoi(c.Param("page"))
	dir, _ := strconv.Atoi(c.Param("dir"))
	page := models.PageCreteria{
		DeviceType : c.Param("type"),
		OrderKey: c.Param("order"),
		Size: size,
		OutFlag: c.Param("outFlag"),
		CurPage: curpage,
		Direction: dir,
	}

	switch page.DeviceType {
	case "server":
		devicePage, err := h.db.GetDevicesServerForPage(page)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		for _, v := range devicePage.Devices {
			fmt.Printf("%v\n", v)
		}
		c.JSON(http.StatusOK, devicePage)
	case "network":
		devicePage, err := h.db.GetDevicesNetworkForPage(page)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		for _, v := range devicePage.Devices {
			fmt.Printf("%v\n", v)
		}
		c.JSON(http.StatusOK, devicePage)
	case "part":
		devicePage, err := h.db.GetDevicesPartForPage(page)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		for _, v := range devicePage.Devices {
			fmt.Printf("%v\n", v)
		}
		c.JSON(http.StatusOK, devicePage)
	}
}

