package rest

import (
	"cmpService/common/lib"
	"cmpService/common/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) GetCompaniesByName(c *gin.Context) {
	if h.db == nil {
		return
	}
	name := c.Param("name")
	customers, err := h.db.GetCompaniesByName(name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, customers)
}

func (h *Handler) GetUsersPage(c *gin.Context) {
	fmt.Println("GetUserPage...")

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
	fmt.Println("1. page:")
	page.String()
	users, err := h.db.GetUsersPage(page)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	fmt.Println("2. page:")
	users.Page.String()
	fmt.Println("OK users:", len(users.Users))
	c.JSON(http.StatusOK, users)
}
