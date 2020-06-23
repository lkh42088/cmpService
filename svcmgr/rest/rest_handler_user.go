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

func getPagination(c *gin.Context) (p models.Pagination, err error) {
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
	p = models.Pagination{
		TotalCount:  0,
		RowsPerPage: rowsPerPage,
		Offset:      offset,
		OrderBy:     orderBy,
		Order:       order,
	}
	return p, err
}

func (h *Handler) GetCompaniesPage(c *gin.Context) {
	page, err := getPagination(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": lib.RestAbnormalParam})
		return
	}

	fmt.Println("1. page:")
	page.String()
	companies, err := h.db.GetCompaniesPage(page)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	fmt.Println("2. page:")
	companies.Page.String()
	fmt.Println("OK users:", len(companies.Companies))
	c.JSON(http.StatusOK, companies)
}

func (h *Handler) checkCompanyExists(name string) bool {
	company, err := h.db.GetCompanyByName(name)
	if err != nil {
		lib.LogWarnln(err)
		return false
	} else if name == company.Name {
		return true
	}
	return true
}

func (h *Handler) AddCompany(c *gin.Context) {
	var companyMsg models.Company
	c.Bind(&companyMsg)
	fmt.Println("recv company: ", companyMsg)
	exists := h.checkCompanyExists(companyMsg.Name)
	fmt.Println("exists: ", exists)

	addCompany, err := h.db.AddCompany(companyMsg)
	if err != nil {
		fmt.Println("err:", err)
		c.JSON(http.StatusUnprocessableEntity, gin.H{"success": false, "errors": err})
		return
	}
	fmt.Println("Add company: ", addCompany)
	c.JSON(http.StatusOK, gin.H{"success": true, "msg": addCompany})
}

func (h *Handler) CheckDuplicatedCompany(c *gin.Context) {
	var companyMsg models.Company
	c.Bind(&companyMsg)
	fmt.Println("recv company: ", companyMsg)
	fmt.Println("recv company name: ", companyMsg.Name)
	exists := h.checkCompanyExists(companyMsg.Name)
	if exists == true {
		fmt.Println("It exists: ", exists)
		c.JSON(http.StatusUnprocessableEntity, gin.H{"success": false, "errors": "It has been exists!"})
		return
	}
	fmt.Println("It does not exists: ", exists)
	c.JSON(http.StatusOK, gin.H{"success": true, "msg": ""})
}
