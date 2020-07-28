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

func (h *Handler) includeEmailAuthToUserDetails(users []models.UserDetail) (newusers []models.UserDetail, err error) {
	for _, user := range users {
		fmt.Println(">>>> userId: ", user.UserId)
		if user.GroupEmailAuth {
			user.GroupEmailAuthList, err = h.db.GetLoginAuthsByUserId(user.UserId)
			if err != nil {
				fmt.Println("List1 : error ", err)
			} else {
				fmt.Println("List1 : ", user.GroupEmailAuthList)
			}
		}
		user.ParticipateInAccountList, err = h.db.GetLoginAuthsByAuthUserId(user.UserId)
		if err != nil {
			fmt.Println("List2 : error ", err)
		} else {
			fmt.Println("List2 : ", user.ParticipateInAccountList)
		}
		var list []models.LoginAuth
		if user.EmailAuth {
			for _, item := range user.ParticipateInAccountList {
				if item.UserId == user.UserId && item.AuthUserId == user.UserId {
					continue
				}
				list = append(list, item)
			}
			user.ParticipateInAccountList = list
		}
		newusers = append(newusers, user)
	}
	return newusers, err
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
	fmt.Println("2. get email group list:")
	users.Users, _ = h.includeEmailAuthToUserDetails(users.Users)
	users.Page.String()
	fmt.Println("OK users:", len(users.Users))
	c.JSON(http.StatusOK, users)
}

func (h *Handler) GetUsersPageWithSearchParam(c *gin.Context) {
	var msg models.PageRequestMsg
	c.Bind(&msg)
	fmt.Printf("GetUsersPageWithSearchParam() msg %v\n", msg)

	page := models.Pagination{
		TotalCount:  0,
		RowsPerPage: msg.RowsPerPage,
		Offset:      msg.Offset,
		OrderBy:     msg.OrderBy,
		Order:       msg.Order,
	}

	var users models.UserPage
	var query string
	var err error
	if msg.Param.Type != "" && msg.Param.Content != "" {
		switch msg.Param.Type {
		case "name":
			query = "user_name like '%" + msg.Param.Content + "%'"
			users, err = h.db.GetUsersPageBySearch(page, query)
		case "cpName":
			query = "c.cp_name like '%" + msg.Param.Content + "%'"
			users, err = h.db.GetUsersPageBySearch(page, query)
		default:
			query = "user_id like '%" + msg.Param.Content + "%'"
			users, err = h.db.GetUsersPageBySearch(page, query)
		}
	}
	if err != nil {
		fmt.Printf("get error %s\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}
	users.Users, _ = h.includeEmailAuthToUserDetails(users.Users)
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

	fmt.Println("User: ", user)
	models.HashPassword(&user)
	adduser, err := h.db.AddUser(user)
	fmt.Println("Add User: ", adduser)
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

func (h *Handler) ModifyUser(c *gin.Context) {
	var msg messages.UserRegisterMessage
	c.Bind(&msg)

	fmt.Println("Register Message: ", msg)
	msg.String()

	oldUser, err := h.db.GetUserById(msg.Id)
	if err != nil {
		fmt.Println("ModifyUser: error 0.")
		c.JSON(http.StatusUnprocessableEntity, gin.H{"success": false, "errors": err})
		return
	}
	fmt.Printf("oldUser: %v\n", oldUser)

	user, emailAuthList := msg.Translate()
	user.Idx = oldUser.Idx
	if user.CompanyIdx == 0 {
		user.CompanyIdx = oldUser.CompanyIdx
	}
	if len(msg.Password) > 6 {
		models.HashPassword(&user)
	} else {
		user.Password = oldUser.Password
	}

	fmt.Println("User: ", user)
	updateUser, err := h.db.UpdateUser(user)
	fmt.Println("update user: ", updateUser)
	if oldUser.GroupEmailAuth {
		h.db.DeleteLoginAuthsByUserIdx(user.Idx)
	}
	if updateUser.GroupEmailAuth && len(emailAuthList) > 0 {
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"success": false, "errors": err})
			fmt.Println("ModifyUser: error 4.", err)
			return
		}
		for _, loginAuth := range emailAuthList {
			loginAuth.UserIdx = updateUser.Idx
			loginAuth, err := h.db.AddLoginAuth(loginAuth)
			if err != nil {
				fmt.Println("Failed to add loginAuth!")
			}
			fmt.Println("Add loginAuth: ", loginAuth)
		}
	}
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"success": false, "errors": err})
		fmt.Println("ModifyUser: error 5.", err)
		return
	}
	fmt.Println("Modify user:", updateUser)
	c.JSON(http.StatusOK, gin.H{"success": true, "msg": updateUser})
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
	fmt.Println("CheckDuplicatedUser: message ", userMsg)
	exists := h.CheckUserExists(userMsg.Id)
	if exists {
		fmt.Println("CheckDuplicatedUser: fail")
		c.JSON(http.StatusUnprocessableEntity, gin.H{"success": false, "errors": ""})
		return
	}
	fmt.Println("CheckDuplicatedUser: success")
	c.JSON(http.StatusOK, gin.H{"success": true, "msg": "존재하지 않는 ID 입니다."})
}

func deleteUser(h *Handler, idx uint) bool {
	var user models.User
	_, err := h.db.DeleteLoginAuthsByUserIdx(idx)
	if err != nil {
		fmt.Println("deleteLoginAuthsByUserIdx err: ", err)
		return false
	}

	user.Idx = idx
	_, err = h.db.DeleteUser(user)
	if err != nil {
		fmt.Println("deleteUser err: ", err)
		return false
	}
	return true
}

func (h *Handler) UnRegisterUser(c *gin.Context) {
	var msg messages.DeleteUserMessage
	c.Bind(&msg)
	fmt.Println("UnRegister Message: ", msg)
	for _, idx := range msg.IdxList {
		deleteUser(h, uint(idx))
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "msg": "User created successfully"})
}

func (h *Handler) UnRegisterUserBackup(c *gin.Context) {
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
