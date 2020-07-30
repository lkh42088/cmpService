package rest

import (
	"cmpService/common/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func (h *Handler) AddSubnet(c *gin.Context) {
	if h.db == nil {
		return
	}

	var subnet models.SubnetMgmt
	c.ShouldBindJSON(&subnet)
	//fmt.Printf("%+v\n", subnet)

	err := h.db.AddSubnet(subnet)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "OK")
}

func (h *Handler) GetSubnets(c *gin.Context) {
	if h.db == nil {
		return
	}

	var val models.PageRequestForSearch
	c.ShouldBindJSON(&val)
	//fmt.Printf("val : %+v\n", val)
	page := models.PageRequestForSearch{
		RowsPerPage: val.RowsPerPage,
		Offset:      val.Offset,
		OrderBy:     val.OrderBy,
		Order:       val.Order,
		SearchParam: val.SearchParam,
	}

	data, err := h.db.GetSubnets(page)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	//fmt.Printf("%+v\n", data)
	c.JSON(http.StatusOK, data)
}

func (h *Handler) UpdateSubnet(c *gin.Context) {
	if h.db == nil {
		return
	}

	var subnet models.SubnetMgmt
	c.ShouldBindJSON(&subnet)
	//fmt.Printf("%+v\n", subnet)

	err := h.db.UpdateSubnet(subnet)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "Update Subnet OK")
}

func (h *Handler) DeleteSubnets(c *gin.Context) {
	idx := c.Param("idx")
	fmt.Printf("%+v\n", idx)
	idxArray := strings.Split(idx, ",")

	err := h.db.DeleteSubnets(idxArray)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "Delete Subnet OK")
}
