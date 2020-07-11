package rest

import (
	"cmpService/common/lib"
	"cmpService/common/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) checkCompanyExists(name string) bool {
	company, err := h.db.GetCompanyByCpName(name)
	if err != nil {
		lib.LogWarnln(err)
		return false
	} else if name == company.Name {
		return true
	}
	return true
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

func (h *Handler) GetCompaniesByCpName(c *gin.Context) {
	if h.db == nil {
		return
	}
	name := c.Param("cpName")
	customers, err := h.db.GetCompaniesByCpName(name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("companies by name: ", customers)
	c.JSON(http.StatusOK, customers)
}

func (h *Handler) GetUserDetailsByCpIdx(c *gin.Context) {
	if h.db == nil {
		return
	}
	cpIdxString:= c.Param("cpIdx")
	fmt.Println("cpIdxString: ", cpIdxString)
	cpIdx, _ := strconv.Atoi(cpIdxString)
	fmt.Println("cpId: ", cpIdx)
	users, err := h.db.GetUserDetailsByCpIdx(cpIdx)
	if err != nil {
		fmt.Println("error: ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("users by cpName: ", users)
	c.JSON(http.StatusOK, users)
}

func (h *Handler) GetCompaniesWithUserByLikeCpName(c *gin.Context) {
	if h.db == nil {
		return
	}
	name := c.Param("cpName")
	customers, err := h.db.GetCompaniesWithUserByLikeCpName(name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("companies by name: ", customers)
	c.JSON(http.StatusOK, customers)
}

func (h *Handler) GetCompanies(c *gin.Context) {
	if h.db == nil {
		return
	}
	customers, err := h.db.GetCompanies()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, customers)
}

func (h *Handler) AddCompany(c *gin.Context) {
	var companyMsg models.CompanyDetail
	c.Bind(&companyMsg)
	fmt.Printf("recv company: %v\n", companyMsg)
	exists := h.checkCompanyExists(companyMsg.Name)
	fmt.Println("exists: ", exists)

	addCompany, err := h.db.AddCompany(companyMsg.Company)
	if err != nil {
		fmt.Println("err:", err)
		c.JSON(http.StatusUnprocessableEntity, gin.H{"success": false, "errors": err})
		return
	}
	var newUser models.User
	newUser.UserId = companyMsg.UserId
	newUser.IsCompanyAccount = true
	newUser.Password = companyMsg.UserPassword
	newUser.AuthLevel = 5
	newUser.CompanyIdx = int(addCompany.Idx)
	newUser.Email = companyMsg.Email
	newUser.Name = companyMsg.Name
	newUser.HP = companyMsg.Tel
	newUser.GroupEmailAuth = false
	newUser.EmailAuth = false
	newUser.Address = companyMsg.Address
	newUser.AddressDetail = companyMsg.AddressDetail
	newUser.Zipcode = companyMsg.Zipcode

	models.HashPassword(&newUser)
	adduser, err := h.db.AddUser(newUser)
	if err != nil {
		fmt.Println("add user: err:", err)
		c.JSON(http.StatusUnprocessableEntity, gin.H{"success": false, "errors": err})
		return
	}

	fmt.Printf("Add company: %v\n", addCompany)
	fmt.Printf("Add user: %v\n", adduser)

	c.JSON(http.StatusOK, gin.H{"success": true, "msg": addCompany})
}

func deleteCompany(h * Handler, idx int) bool {
	var company models.Company
	company.Idx = uint(idx)
	users, err := h.db.GetUserDetailsByCpIdx(idx)
	for _, user := range users {
		deleteUser(h, user.Idx)
	}

	_, err = h.db.DeleteCompany(company)
	if err != nil {
		fmt.Println("delete company: err ", err)
		return false
	}
	return true
}

func (h *Handler) DeleteCompany(c *gin.Context) {
	var msg models.DeleteCompanyMessage
	c.Bind(&msg)
	fmt.Println("recv msg: ", msg)
	for _, idx := range msg.IdxList {
		deleteCompany(h, idx)
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "msg": ""})
}

