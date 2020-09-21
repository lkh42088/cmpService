package rest

import (
	"bytes"
	"cmpService/common/messages"
	"cmpService/common/models"
	"cmpService/svcmgr/errors"
	"cmpService/svcmgr/utils"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"time"
)

var jwtKey = []byte("secret")

type Claims struct {
	User messages.UserInfo
	jwt.StandardClaims
}

var smtpServer = utils.SmtpServer{
	Host:     "smtp.gmail.com",
	Port:     "587",
	User:     "nubesbh@gmail.com",
	Password: "tycp zngl ehop smvy",
}

var WebServerAddress string

func (h *Handler) LoginUserByEmail(c *gin.Context) {
	var loginMsg messages.UserLoginMessage
	c.Bind(&loginMsg)

	fmt.Println("LoginUser2:", loginMsg)
	user, err := h.db.GetUserByEmail(loginMsg.Email)
	fmt.Println("LoginUser2:", user)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"success": false, "errors": err})
		return
	}
	match := models.CheckPasswordHash(loginMsg.Password, user.Password)
	if !match {
		c.JSON(http.StatusUnauthorized, gin.H{"success": false, "errors": "incorrect credentials"})
		return
	}

	expirationTime := time.Now().Add(30 * time.Minute)
	claims := &Claims{
		User: messages.UserInfo{
			Id:                 user.UserId,
			Email:              user.Email,
			Name:               user.Name,
			EmailAuthFlag:      user.EmailAuth,
			EmailAuthGroupFlag: user.GroupEmailAuth,
		},
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	errors.HandleErr(c, err)
	http.SetCookie(c.Writer, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
		//Secure:		true,
		//SameSite:   http.SameSiteNoneMode,
	})
	fmt.Println(tokenString)
	c.JSON(http.StatusOK,
		gin.H{"success": true, "msg": "loggged in successfully", "user": claims.User,
			"token": tokenString})
}

func (h *Handler) LoginEmailAuthConfirm(c *gin.Context) {
	var loginMsg messages.UserLoginMessage
	c.Bind(&loginMsg)

	fmt.Println(">>>>>> LoginEmailAuthConfirm")
	var restStatus int
	if h.isConfirmEmailAuth(loginMsg.Id, loginMsg.Email) {
		user, err := h.db.GetUserDetailById(loginMsg.Id)
		if err != nil {
			restStatus = http.StatusUnprocessableEntity
			c.JSON(restStatus, gin.H{"success": false, "errors": err})
			return
		}

		// Token을 발급해서 응답한다.
		responseWithToken(c, user, loginMsg.Email)
		return
	}
	restStatus = messages.StatusFailedEmailAuth
	c.JSON(restStatus, gin.H{"success": false, "errors": messages.RestStatusText(restStatus)})
}

func (h *Handler) LoginFrontConfirm(c *gin.Context) {
	var msg messages.UserLoginMessage
	var restStatus int
	c.Bind(&msg)

	fmt.Println(">>>>>> LoginFrontConfirm")
	fmt.Println("message:", msg)
	fmt.Println("id:", msg.Id)
	fmt.Println("email:", msg.Email)
	user, err := h.db.GetUserDetailById(msg.Id)
	if err != nil {
		fmt.Println("error 1:", err)
		msg.Password = ""
		msg.Comment = "계정이 존재하지 않습니다."
		c.JSON(http.StatusUnprocessableEntity, gin.H{"success": false, "errors": err, "msg": msg})
		return
	}
	match := models.CheckPasswordHash(msg.Password, user.Password)
	if !match {
		fmt.Println("error 2")
		restStatus = http.StatusUnauthorized
		msg.Password = ""
		msg.Comment = "패스워드가 잘못 되었습니다."
		c.JSON(restStatus, gin.H{"success": false, "errors": "incorrect credentials", "msg": msg})
		return
	}
	authLogin, err := h.db.GetLoginAuthByUserIdAndTargetEmail(msg.Id, msg.Email)
	if err != nil {
		msg.Password = ""
		msg.Comment = "인증 이메일이 잘못 되었습니다."
		c.JSON(http.StatusUnprocessableEntity, gin.H{"success": false, "errors": err})
	} else if !authLogin.EmailAuthConfirm {
		msg.Password = ""
		msg.Comment = "인증 되지 않았습니다."
		c.JSON(http.StatusUnprocessableEntity, gin.H{"success": false, "errors": err, "msg": msg})
	}

	fmt.Println("[user]", user)

	// Token을 발급해서 응답한다.
	responseWithToken(c, user, "")
}

func (h *Handler) LoginGroupEmail(c *gin.Context) {
	var loginMsg messages.UserLoginMessage
	var restStatus int
	c.Bind(&loginMsg)

	fmt.Println(">>>>>> LoginGroupEmail")
	fmt.Println("message:", loginMsg)
	fmt.Println("id:", loginMsg.Id)
	user, err := h.db.GetUserDetailById(loginMsg.Id)
	if err != nil {
		fmt.Println("error 1:", err)
		c.JSON(http.StatusUnprocessableEntity, gin.H{"success": false, "errors": err})
		return
	}
	match := models.CheckPasswordHash(loginMsg.Password, user.Password)
	if !match {
		fmt.Println("error 2")
		restStatus = http.StatusUnauthorized
		c.JSON(restStatus, gin.H{"success": false, "errors": "incorrect credentials"})
		return
	}

	fmt.Println("[user]", user)

	// Email 인증 여부 Check
	if user.GroupEmailAuth {
		if loginMsg.Email != "" {
			// 그룹에 등록된 이메일인지 Check
			// userId, userEmail
			if h.checkGroupEmailAuth(loginMsg.Id, loginMsg.Email) == false {
				fmt.Println("error 4:")
				restStatus = messages.StatusFailedNotHaveAuthEmail
				c.JSON(restStatus, gin.H{"success": false, "msg": messages.RestStatusText(restStatus)})
				return
			}

			fmt.Println("email auth group...")
			// 이메일 발송
			fmt.Println("send 2")
			err = h.sendAuthMail(c, loginMsg.Id, loginMsg.Email)
			if err != nil {
				return
			}
			fmt.Println("error 5:")
			restStatus = messages.StatusSentEmailAuth
			c.JSON(restStatus, gin.H{"success": false, "msg": messages.RestStatusText(restStatus)})
		} else {
			// 이메일 수신 후 발송
			fmt.Println("error 6:")
			restStatus = messages.StatusInputEmailAuth
			c.JSON(restStatus, gin.H{"success": false, "msg": messages.RestStatusText(restStatus)})
		}
		return
	}

	fmt.Println("success..")
	// Token을 발급해서 응답한다.
	responseWithToken(c, user, "")
}

func (h *Handler) CheckPassword(c *gin.Context) {
	msg, err := JsonUnmarshal(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": ""})
		return
	}
	fmt.Printf("msg: %v\n", msg)
	id := msg["id"].(string)
	password := msg["password"].(string)
	user, err := h.db.GetUserById(id)
	match := models.CheckPasswordHash(password, user.Password)
	if !match {
		c.JSON(http.StatusUnauthorized, gin.H{"success": false, "errors": ""})
		return
	}
}

func (h *Handler) LoginUserById(c *gin.Context) {
	var msg messages.UserLoginMessage
	var restStatus int
	c.Bind(&msg)

	fmt.Println("message:", msg)
	fmt.Println("id:", msg.Id)
	userx, err := h.db.GetUserById(msg.Id)
	if err != nil {
		fmt.Println("[LoginUserById] error 0:", err)
		c.JSON(http.StatusUnprocessableEntity, gin.H{"success": false, "errors": err})
		return
	}
	fmt.Println("userx: ", userx)

	user, err := h.db.GetUserDetailById(msg.Id)
	if err != nil {
		fmt.Println("[LoginUserById] error 1:", err)
		c.JSON(http.StatusUnprocessableEntity, gin.H{"success": false, "errors": err})
		return
	}
	match := models.CheckPasswordHash(msg.Password, user.Password)
	if !match {
		fmt.Println("error 2")
		restStatus = http.StatusUnauthorized
		msg.Comment = "Incorrect credentials"
		msg.Password = ""
		msg.Result = restStatus
		c.JSON(restStatus, gin.H{"success": false, "errors": "incorrect credentials"})
		return
	}

	fmt.Println("[user]", user)

	// Email 인증 여부 Check
	if user.EmailAuth {
		// 본인 이메일 발송
		fmt.Println("send 1")
		err = h.sendAuthMailForMyselfAuth(c, user.UserId, user.Email)
		if err != nil {
			fmt.Println("Failed to send email")
			return
		}
		fmt.Println("error 3:")
		msg.Comment = "In self email auth, sent authenticated email."
		msg.Password = ""
		msg.Email = user.Email
		msg.Result = messages.StatusSentEmailAuth
		c.JSON(http.StatusOK, gin.H{"success": false, "msg": msg})
		return
	} else if user.GroupEmailAuth {
		if msg.Email != "" {
			// 그룹에 등록된 이메일인지 Check
			// userId, userEmail
			if h.checkGroupEmailAuth(msg.Id, msg.Email) == false {
				fmt.Println("error 4:")
				restStatus = messages.StatusFailedNotHaveAuthEmail
				c.JSON(restStatus, gin.H{"success": false, "msg": messages.RestStatusText(restStatus)})
				return
			}

			fmt.Println("email auth group...")
			// 이메일 발송
			fmt.Println("send 2")
			err = h.sendAuthMail(c, msg.Id, msg.Email)
			if err != nil {
				fmt.Println("error 5:")
				restStatus = messages.StatusFailedNotHaveAuthEmail
				c.JSON(restStatus, gin.H{"success": false, "msg": messages.RestStatusText(restStatus)})
				return
			}
			msg.Comment = "In group email auth, sent authenticated email."
			msg.Password = ""
			msg.Email = msg.Email
			msg.Result = messages.StatusSentEmailAuth
			c.JSON(http.StatusOK, gin.H{"success": false, "msg": msg})
		} else {
			// 이메일 수신 후 발송
			fmt.Println("error 6:")
			msg.Comment = "In group email auth, you must input email for authentication."
			msg.Password = ""
			msg.Result = messages.StatusInputEmailAuth
			c.JSON(http.StatusOK, gin.H{"success": false, "msg": msg})
		}
		return
	}

	fmt.Println("success..")
	// Token을 발급해서 응답한다.
	responseWithToken(c, user, "")
}

func (h *Handler) LoginSendEmail(c *gin.Context) {
	var msg messages.UserLoginMessage
	var restStatus int
	c.Bind(&msg)

	fmt.Println("message:", msg)
	userx, err := h.db.GetUserById(msg.Id)
	if err != nil {
		fmt.Println("[LoginUserById] error 0:", err)
		c.JSON(http.StatusUnprocessableEntity, gin.H{"success": false, "errors": err})
		return
	}
	fmt.Println("userx: ", userx)

	user, err := h.db.GetUserDetailById(msg.Id)
	if err != nil {
		fmt.Println("[LoginUserById] error 1:", err)
		c.JSON(http.StatusUnprocessableEntity, gin.H{"success": false, "errors": err})
		return
	}
	match := models.CheckPasswordHash(msg.Password, user.Password)
	if !match {
		fmt.Println("error 2")
		restStatus = http.StatusUnauthorized
		msg.Comment = "Incorrect credentials"
		msg.Password = ""
		msg.Result = restStatus
		c.JSON(restStatus, gin.H{"success": false, "errors": "incorrect credentials"})
		return
	}

	fmt.Println("[user]", user)

	// Email 인증 여부 Check
	if msg.Email != "" {
		// 그룹에 등록된 이메일인지 Check
		// userId, userEmail
		if h.checkGroupEmailAuth(msg.Id, msg.Email) == false {
			fmt.Println("error 4:")
			restStatus = messages.StatusFailedNotHaveAuthEmail
			msg.Comment = "In group email auth, your email dose not exist!"
			msg.Password = ""
			msg.Result = restStatus
			c.JSON(restStatus, gin.H{"success": false, "msg": msg})
			return
		}

		fmt.Println("email auth group...")
		// 이메일 발송
		fmt.Println("send 2")
		err = h.sendAuthMail(c, msg.Id, msg.Email)
		if err != nil {
			restStatus = messages.StatusFailedNotHaveAuthEmail
			msg.Comment = "In group email auth, your email dose not exist!"
			msg.Password = ""
			msg.Result = restStatus
			c.JSON(restStatus, gin.H{"success": false, "msg": msg})
			return
		}
		fmt.Println("error 5:")
		msg.Comment = "In group email auth, sent authenticated email."
		msg.Password = ""
		msg.Result = messages.StatusSentEmailAuth
		c.JSON(http.StatusOK, gin.H{"success": false, "msg": msg})
		return
	}

	// 이메일 수신 후 발송
	fmt.Println("error 6:")
	msg.Comment = "In group email auth, you must input email for authentication."
	msg.Password = ""
	msg.Result = messages.StatusInputEmailAuth
	c.JSON(http.StatusOK, gin.H{"success": false, "msg": msg})
}

func (h *Handler) isConfirmEmailAuth(userId, userEmail string) bool {
	// 1. Get from DB
	userEmailAuth, err := h.db.GetUserEmailAuthByIdAndEmail(userId, userEmail)
	if err != nil {
		fmt.Println("err:", err, " userEmailAuth:", userEmailAuth)
		return false
	}
	if userEmailAuth.EmailAuthConfirm {
		return true
	}
	return false
}

func (h *Handler) checkGroupEmailAuth(userId, userEmail string) bool {
	// 1. Get from DB
	userEmailAuth, err := h.db.GetLoginAuthByUserIdAndTargetEmail(userId, userEmail)
	if err != nil {
		fmt.Println("err:", err, " userEmailAuth:", userEmailAuth)
		return false
	}
	return true
}

func (h *Handler) checkGroupEmailAuthOld(userId, userEmail string) bool {
	// 1. Get from DB
	userEmailAuth, err := h.db.GetUserEmailAuthByIdAndEmail(userId, userEmail)
	if err != nil {
		fmt.Println("err:", err, " userEmailAuth:", userEmailAuth)
		return false
	}
	return true
}

func responseWithToken(c *gin.Context, user models.UserDetail, authEmail string) {
	expirationTime := time.Now().Add(60 * 24 * time.Minute)
	claims := &Claims{
		User: messages.UserInfo{
			Id:                 user.UserId,
			Name:               user.Name,
			Email:              user.Email,
			Level:              user.AuthLevel,
			CpIdx:              user.CompanyIdx,
			CpName:             user.CompanyName,
			EmailAuthFlag:      user.EmailAuth,
			EmailAuthGroupFlag: user.GroupEmailAuth,
			AuthEmail:          authEmail,
			Avata:				user.Avata,
		},
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	fmt.Printf(">>>> GEN TOKEN: size(%d) token(%s)\n", len(tokenString), tokenString)
	errors.HandleErr(c, err)
	http.SetCookie(c.Writer, &http.Cookie{
		Name:    	"token",
		Value:   	tokenString,
		Expires: 	expirationTime,
		//Secure:		true,
		//SameSite:   http.SameSiteNoneMode,
	})
	var msg messages.UserLoginMessage
	msg.Id = user.UserId
	msg.Result = http.StatusOK
	msg.Comment = "Login success!"
	c.JSON(http.StatusOK,
		gin.H{"success": true, "msg": msg, "user": claims.User})
}

func (h *Handler) sendAuthMailForMyselfAuth(c *gin.Context, id, email string) error {

	// 1. Generate UUID
	uuid, err := utils.NewUUID()
	if err != nil {
		fmt.Println("Failed to generate UUID!!")
		c.JSON(http.StatusNotAcceptable, gin.H{"success": false, "msg": "UUID 생성에 실패함!"})
		return err
	}

	fmt.Println("Email Authentication process....")
	fmt.Println("id:", id)
	fmt.Println("email:", email)

	// 2. Get from DB
	userEmailAuth, err := h.db.GetLoginAuthByMySelfAuth(id)
	fmt.Println("err:", err, " userEmailAuth:", userEmailAuth)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"success": false, "msg": "이 계정은 인증 DB에 존재하지 않음."})
		return err
	}
	userEmailAuth.EmailAuthConfirm = false
	userEmailAuth.EmailAuthStore = uuid

	fmt.Printf("..userEmailAuth: %v\n", userEmailAuth)
	fmt.Printf("EamilAuthStore: %s\n", userEmailAuth.EmailAuthStore)
	fmt.Printf("EmailAuthConfirm: %t\n", userEmailAuth.EmailAuthConfirm)

	// 3. Update to DB
	userEmailAuth, err = h.db.UpdateLoginAuth(userEmailAuth)
	fmt.Printf("Update, EamilAuthStore: %s\n", userEmailAuth.EmailAuthStore)
	fmt.Printf("Update, EmailAuthConfirm: %t\n", userEmailAuth.EmailAuthConfirm)
	if err != nil {
		fmt.Println("Failed to update user email auth!:", err)
		c.JSON(http.StatusNotAcceptable, gin.H{"success": false, "msg": "이 계정은 인증 DB에 업데이트 하는 것을 실패했다."})
		return err
	}

	// 4. Send email
	emailmsg := utils.MailMsg{
		To:       email,
		Header:   "콘텐츠브릿지 로그인 Email 인증",
		ServerIp: WebServerAddress,
		Uuid:     uuid,
		UserId:   id,
		TargetId: id,
		Text:     "계정에 대한 이메일 인증을 위해서 아래 URL을 클릭하시기 바랍니다.",
	}
	err = utils.SendMail(smtpServer, emailmsg)
	fmt.Println("SendMail: err ", err)
	return nil
}

func (h *Handler) sendAuthMail(c *gin.Context, id, email string) error {

	// 1. Generate UUID
	uuid, err := utils.NewUUID()
	if err != nil {
		fmt.Println("Failed to generate UUID!!")
		c.JSON(http.StatusNotAcceptable, gin.H{"success": false, "msg": "UUID 생성에 실패함!"})
		return err
	}

	fmt.Println("Email Authentication process....")
	fmt.Println("id:", id)
	fmt.Println("email:", email)

	// 2. Get from DB
	userEmailAuth, err := h.db.GetLoginAuthByUserIdAndTargetEmail(id, email)
	fmt.Println("err:", err, " userEmailAuth:", userEmailAuth)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"success": false, "msg": "이 계정은 인증 DB에 존재하지 않음."})
		return err
	}
	userEmailAuth.EmailAuthConfirm = false
	userEmailAuth.EmailAuthStore = uuid

	fmt.Println("userEmailAuth:", userEmailAuth)

	// 3. Update to DB
	userEmailAuth, err = h.db.UpdateLoginAuth(userEmailAuth)
	if err != nil {
		fmt.Println("Failed to update user email auth!:", err)
		c.JSON(http.StatusNotAcceptable, gin.H{"success": false, "msg": "이 계정은 인증 DB에 업데이트 하는 것을 실패했다."})
		return err
	}

	// 4. Send email
	emailmsg := utils.MailMsg{
		To:       email,
		Header:   "콘텐츠브릿지 로그인 Email 인증",
		ServerIp: WebServerAddress,
		Uuid:     uuid,
		UserId:   id,
		TargetId: userEmailAuth.AuthUserId,
		Text:     "계정에 대한 이메일 인증을 위해서 아래 URL을 클릭하시기 바랍니다.",
	}
	err = utils.SendMail(smtpServer, emailmsg)
	fmt.Println("SendMail: err ", err)
	return nil
}

func (h *Handler) sendAuthMailold(c *gin.Context, id, email string) error {

	// 1. Generate UUID
	uuid, err := utils.NewUUID()
	if err != nil {
		fmt.Println("Failed to generate UUID!!")
		c.JSON(http.StatusNotAcceptable, gin.H{"success": false, "msg": "UUID 생성에 실패함!"})
		return err
	}

	fmt.Println("Email Authentication process....")
	fmt.Println("id:", id)
	fmt.Println("email:", email)

	// 2. Get from DB
	userEmailAuth, err := h.db.GetUserEmailAuthByIdAndEmail(id, email)
	fmt.Println("err:", err, " userEmailAuth:", userEmailAuth)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"success": false, "msg": "이 계정은 인증 DB에 존재하지 않음."})
		return err
	}
	userEmailAuth.EmailAuthConfirm = false
	userEmailAuth.EmailAuthStore = uuid

	fmt.Println("userEmailAuth:", userEmailAuth)

	// 3. Update to DB
	userEmailAuth, err = h.db.UpdateUserEmailAuth(userEmailAuth)
	if err != nil {
		fmt.Println("Failed to update user email auth!:", err)
		c.JSON(http.StatusNotAcceptable, gin.H{"success": false, "msg": "이 계정은 인증 DB에 업데이트 하는 것을 실패했다."})
		return err
	}

	// 4. Send email
	emailmsg := utils.MailMsg{
		To:       email,
		Header:   "콘텐츠브릿지 로그인 Email 인증",
		ServerIp: WebServerAddress,
		Uuid:     uuid,
		UserId:   id,
		Text:     "계정에 대한 이메일 인증을 위해서 아래 URL을 클릭하시기 바랍니다.",
	}
	err = utils.SendMail(smtpServer, emailmsg)
	fmt.Println("SendMail: err ", err)
	return nil
}

func (h *Handler) Logout(c *gin.Context) {
	//fmt.Println("User Logout....")
	http.SetCookie(c.Writer, &http.Cookie{
		Name:   "token",
		MaxAge: -1,
		//SameSite: http.SameSiteNoneMode,
	})
	c.JSON(http.StatusNoContent, gin.H{"success": true, "msg": "logged out in successfully"})
}

func (h *Handler) GetSession(c *gin.Context) {
	user, isAuthenticated, _ := AuthMiddleware(c, jwtKey)
	if !isAuthenticated {
		c.JSON(http.StatusUnauthorized,
			gin.H{"success": false, "msg": "unauthorized"})
		return
	}

	/*
		fmt.Println(">>> USER: ", user)
		var username, email string
		var haveEmailAuth, haveGroupEmailAuth bool
		for i, data := range user {
			fmt.Println("[", i, "]", data)
			if i == "username" {
				username = fmt.Sprintf("%v", data)
			} else if i == "email" {
				email = fmt.Sprintf("%v", data)
			} else if i == "haveEmailAuth" {
				haveEmailAuth = data.(interface{}).(bool)
			} else if i == "haveGroupEmailAuth" {
				haveGroupEmailAuth = data.(interface{}).(bool)
			}
		}

		// Generate UUID
		uuid, err := utils.NewUUID()
		if err != nil {
			fmt.Println("Failed to generate UUID!!")
			c.JSON(http.StatusNotAcceptable, gin.H{"success":false, "msg":"UUID 생성에 실패함!"})
			return
		}
		fmt.Printf("uuid(%d) :(%s)\n", len(uuid), uuid)
		fmt.Println("username: ", username)
		fmt.Println("email: ", email)
		fmt.Println("haveEmailAuth: ", haveEmailAuth)
		fmt.Println("haveGroupEmailAuth: ", haveGroupEmailAuth)
		fmt.Println("token(", len(token), ") : ", token)

		if haveEmailAuth || haveGroupEmailAuth {
			fmt.Println("Email Authentication process....")
			// Get from DB
			userAuth, err := h.db.GetUserEmailAuthByIdAndEmail(username, email)
			fmt.Println("err:", err, " userAuth:", userAuth)
			if err != nil {
				c.JSON(http.StatusNotAcceptable, gin.H{"success":false, "msg":"이 계정은 인증 DB에 존재하지 않음."})
				return
			}
			userAuth.EmailAuthConfirm = false
			userAuth.EmailAuthStore = uuid

			// Update to DB
			h.db.UpdateUserEmailAuth(userAuth)
		}

		//// Send email
		emailmsg := utils.MailMsg {
			To: email,
			Header: "콘텐츠브릿지 로그인 Email 인증",
			ServerIp:    WebServerAddress,
			Uuid:   uuid,
			UserId: username,
			Text:   "계정에 대한 이메일 인증을 위해서 아래 URL을 클릭하시기 바랍니다.",
		}
		err = utils.SendMail(smtpServer, emailmsg)
		fmt.Println("SendMail: err ", err)
	*/

	c.JSON(http.StatusOK, gin.H{"success": true, "user": user})
}

func (h *Handler) EmailConfirm(c *gin.Context) {
	msg, err := JsonUnmarshal(c.Request.Body)
	if err != nil {
		fmt.Println("EmailConfirm error:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": http.StatusText(http.StatusBadRequest)})
		return
	}
	fmt.Println("msg:", len(msg))
	fmt.Println("msg:", msg)

	userId := msg["id"].(string)
	targetId := msg["target"].(string)
	secret := msg["secret"].(string)

	fmt.Println("EmailConfirm: id - ", userId)
	fmt.Println("EmailConfirm: target - ", targetId)
	fmt.Println("EmailConfirm: Secret - ", secret)

	// Get from DB
	userAuth, err := h.db.GetLoginAuthByUserIdAndTargetId(userId, targetId)
	fmt.Println("err:", err, " userAuth:", userAuth)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"success": false, "msg": "이 계정은 인증 DB에 존재하지 않음."})
		return
	}

	fmt.Printf("..userEmailAuth: %v\n", userAuth)
	fmt.Printf("EamilAuthStore: %s\n", userAuth.EmailAuthStore)
	fmt.Printf("EmailAuthConfirm: %t\n", userAuth.EmailAuthConfirm)

	if userAuth.EmailAuthStore != secret {
		c.JSON(messages.StatusFailedEmailAuth, gin.H{"success": false,
			"msg": messages.RestStatusText(messages.StatusFailedEmailAuth)})
		return
	}

	userAuth.EmailAuthConfirm = true
	fmt.Printf(">>> EmailAuthConfirm: %t\n", userAuth.EmailAuthConfirm)
	h.db.UpdateLoginAuth(userAuth)

	c.JSON(http.StatusOK, gin.H{"success": true, "msg": ""})
}

// Get ReCAPTCHA code
func (h *Handler) GetCaptcha(c *gin.Context) {
	if h.db == nil {
		return
	}

	var humanKey models.GoogleKey
	c.ShouldBindJSON(&humanKey)
	//fmt.Printf("humanKey : %+v\n", humanKey.HumanKey)

	requestBody := bytes.NewBufferString(humanKey.HumanKey)
	urlString := models.GoogleVerifyUrl
	urlString += "?secret=" + models.SecretKey
	urlString += "&response=" + humanKey.HumanKey
	//fmt.Println(urlString)

	response, err := http.Post(urlString,
		"text/plain", requestBody)
	if err != nil {
		//fmt.Printf("err1: %+v\n", err)
		c.JSON(http.StatusBadRequest, err)
		return
	}
	defer response.Body.Close()

	if response.Body != nil {
		responseBody, err := ioutil.ReadAll(response.Body)
		if err != nil {
			//fmt.Printf("err2: %+v\n", err)
			c.JSON(http.StatusUnauthorized, err)
			return
		}

		//fmt.Printf("response : %+v\n", string(responseBody))
		c.JSON(http.StatusOK, string(responseBody))
	}
}
