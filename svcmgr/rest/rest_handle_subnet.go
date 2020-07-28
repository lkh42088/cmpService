package rest

import (
	"cmpService/common/lib"
	"cmpService/common/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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

	// Parse params
	rowsPerPage, err := strconv.Atoi(c.Param("rows"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": lib.RestAbnormalParam})
		return
	}
	offset, err := strconv.Atoi(c.Param("offset"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": lib.RestAbnormalParam})
		return
	}
	orderBy := c.Param("orderby")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": lib.RestAbnormalParam})
		return
	}
	order := c.Param("order")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": lib.RestAbnormalParam})
		return
	}

	page := models.Pagination{
		TotalCount:  0,
		RowsPerPage: rowsPerPage,
		Offset:      offset,
		OrderBy:     orderBy,
		Order:       order,
	}

	data, err := h.db.GetSubnets(page)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	//fmt.Printf("%+v\n", data.Subnet)
	fmt.Printf("%+v\n", data.Page)
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
