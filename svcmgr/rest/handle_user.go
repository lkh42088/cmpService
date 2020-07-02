package rest

import (
	"cmpService/common/lib"
	"cmpService/common/messages"
	"cmpService/common/models"
	"cmpService/svcmgr/errors"
	"cmpService/svcmgr/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) CheckUserExists(userId string) bool {
	user, err := h.db.GetUserById(userId)
	if err != nil {
		lib.LogWarnln(err)
		return false
	} else if userId == user.UserId {
		return true
	}
	return false
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

func (h *Handler) RegisterUser(c *gin.Context) {
	var userMsg messages.UserRegisterMessage
	c.Bind(&userMsg)

	fmt.Println("Register Message: ", userMsg)
	userMsg.String()

	exists := h.CheckUserExists(userMsg.Id)
	fmt.Println("exists:", exists)

	//c.JSON(http.StatusUnprocessableEntity, gin.H{"success": false, "errors": "test"})

	valErr := utils.ValidateUserbyMsg(userMsg, errors.ValidationErrors)
	if exists == true {
		valErr = append(valErr, "ID already exists")
	}
	fmt.Println("error:", valErr)
	if len(valErr) > 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"success": false, "errors": valErr})
		return
	}

	user, emailAuthList := userMsg.Translate()
	if userMsg.CpIdx > 0 {
		user.CompanyIdx = userMsg.CpIdx
	} else if userMsg.CpName != "" {
		// get company by name
		company, err := h.db.GetCompanyByCpName(userMsg.CpName)
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"success": false, "errors": valErr})
			return
		}
		user.CompanyIdx = int(company.Idx)
	} else {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"success": false, "errors": valErr})
		return
	}

	models.HashPassword(&user)
	adduser, err := h.db.AddUser(user)
	if len(emailAuthList) > 0 {
		for _, loginAuth := range emailAuthList {
			loginAuth.UserIdx = adduser.Idx
			loginAuth, err := h.db.AddLoginAuth(loginAuth)
			if err != nil {
				fmt.Println("Failed to add loginAuth!")
			}
			fmt.Println("Add loginAuth: ", loginAuth)
		}
	}
	if err != nil {
		fmt.Println("err:", err)
		c.JSON(http.StatusUnprocessableEntity, gin.H{"success": false, "errors": err})
		return
	}
	fmt.Println("Add user:", adduser)
	c.JSON(http.StatusOK, gin.H{"success": true, "msg": adduser})
}

func (h *Handler) RegisterUserBackup(c *gin.Context) {
	var userMsg messages.UserRegisterMessage
	c.Bind(&userMsg)
	fmt.Println("Register Message: ", userMsg)
	exists := h.CheckUserExists(userMsg.Id)
	fmt.Println("exists:", exists)

	valErr := utils.ValidateUserbyMsg(userMsg, errors.ValidationErrors)
	if exists == true {
		valErr = append(valErr, "ID already exists")
	}
	fmt.Println("error:", valErr)
	if len(valErr) > 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"success": false, "errors": valErr})
		return
	}

	user, emailAuthList := userMsg.Convert()
	if userMsg.CpIdx > 0 {
		user.CompanyIdx = userMsg.CpIdx
	} else if userMsg.CpName != "" {
		// get company by name
		company, err := h.db.GetCompanyByCpName(userMsg.CpName)
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"success": false, "errors": valErr})
			return
		}
		user.CompanyIdx = int(company.Idx)
	} else {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"success": false, "errors": valErr})
		return
	}

	models.HashPassword(&user)
	adduser, err := h.db.AddUser(user)
	if len(emailAuthList) > 0 {
		for _, emailAuth := range emailAuthList {
			emailAuth.UserIdx = adduser.Idx
			emailAuth, err := h.db.AddUserEmailAuth(emailAuth)
			if err != nil {
				fmt.Println("Failed to add emailAuth!")
			}
			fmt.Println("Add emailAuth: ", emailAuth)
		}
	}
	if err != nil {
		fmt.Println("err:", err)
		c.JSON(http.StatusUnprocessableEntity, gin.H{"success": false, "errors": err})
		return
	}
	fmt.Println("Add user:", adduser)
	c.JSON(http.StatusOK, gin.H{"success": true, "msg": adduser})
}

func (h *Handler) CheckDuplicatedUser(c *gin.Context) {
	var userMsg messages.UserRegisterMessage
	c.Bind(&userMsg)
	fmt.Println("Register Message: ", userMsg)
	exists := h.CheckUserExists(userMsg.Id)
	if exists {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"success": false, "errors": ""})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "msg": ""})
}

func (h *Handler) UnRegisterUser(c *gin.Context) {
	var userMsg messages.UserRegisterMessage
	c.Bind(&userMsg)
	fmt.Println("UnRegister Message: ", userMsg)
	user, err := h.db.GetUserById(userMsg.Id)
	if err != nil || user.UserId != userMsg.Id {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"success": false, "errors": "Id dose not exist"})
		return
	}

	if user.EmailAuth || user.GroupEmailAuth {
		emailAuthList, err := h.db.DeleteUserEmailAuthByUserId(user.UserId)
		if err != nil {
			fmt.Println("err 1:", err)
			c.JSON(http.StatusUnprocessableEntity, gin.H{"success": false, "errors": err})
			return
		}
		fmt.Println("delete emailAuthList:", emailAuthList)
	}

	adduser, err := h.db.DeleteUser(user)
	if err != nil {
		fmt.Println("err 2:", err)
		c.JSON(http.StatusUnprocessableEntity, gin.H{"success": false, "errors": err})
		return
	}

	fmt.Println("Delete user:", adduser)
	c.JSON(http.StatusOK, gin.H{"success": true, "msg": "User created successfully"})
}
