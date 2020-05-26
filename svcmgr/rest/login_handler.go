package rest

import (
	"cmpService/common/lib"
	"cmpService/common/messages"
	"cmpService/common/models"
	"cmpService/svcmgr/errors"
	"cmpService/svcmgr/utils"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

var jwtKey = []byte("secret")

type Claims struct {
	models.User
	jwt.StandardClaims
}

var smtpServer = utils.SmtpServer {
	Host:     "smtp.gmail.com",
	Port:     "587",
	User:     "nubesbh@gmail.com",
	Password: "tycp zngl ehop smvy",
}

var svcmgrAddress = "localhost"

func (h *Handler) checkUserExists(user models.User) bool {
	getuser, err := h.db.GetUserById(user.UserId)
	if err != nil {
		lib.LogWarnln(err)
		return false
	} else if user.UserId == getuser.UserId {
		return true
	}
	return false
}

func (h *Handler) RegisterUser(c *gin.Context) {
	var user models.User
	c.Bind(&user)
	fmt.Println("RegisterUser: ", user)
	exists := h.checkUserExists(user)
	fmt.Println("exists:", exists)

	valErr := utils.ValidateUser(user, errors.ValidationErrors)
	if exists == true {
		valErr = append(valErr, "ID already exists")
	}
	fmt.Println("error:", valErr)
	if len(valErr) > 0 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"success":false, "errors":valErr})
		return
	}
	models.HashPassword(&user)
	adduser, err := h.db.AddUser(user)
	if err != nil {
		fmt.Println("err:", err)
		c.JSON(http.StatusUnprocessableEntity, gin.H{"success":false, "errors":err})
		return
	}
	fmt.Println("Add user:", adduser)
	c.JSON(http.StatusOK, gin.H{"success":true, "msg":"User created successfully"})
}

func (h *Handler) LoginUserByEmail(c *gin.Context) {
	var loginMsg messages.UserLoginMessage
	c.Bind(&loginMsg)

	fmt.Println("LoginUser2:", loginMsg)
	user, err := h.db.GetUserByEmail(loginMsg.Email)
	fmt.Println("LoginUser2:", user)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"success":false, "errors":err})
		return
	}
	match := models.CheckPasswordHash(loginMsg.Password, user.Password)
	if !match {
		c.JSON(http.StatusUnauthorized, gin.H{"success":false, "errors":"incorrect credentials"})
		return
	}

	expirationTime := time.Now().Add(30 * time.Minute)
	claims := &Claims{
		User: models.User{
			UserId: user.UserId,
			Email: user.Email,
			Name: user.Name,
		},
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	errors.HandleErr(c, err)
	http.SetCookie(c.Writer, &http.Cookie{
		Name: "token",
		Value: tokenString,
		Expires: expirationTime,
	})
	fmt.Println(tokenString)
	c.JSON(http.StatusOK,
		gin.H{"success":true, "msg":"loggged in successfully", "user":claims.User,
			"token":tokenString})
}

func (h *Handler) LoginUserById(c *gin.Context) {
	var loginMsg messages.UserLoginMessage
	c.Bind(&loginMsg)

	user, err := h.db.GetUserById(loginMsg.Id)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"success":false, "errors":err})
		return
	}
	match := models.CheckPasswordHash(loginMsg.Password, user.Password)
	if !match {
		c.JSON(http.StatusUnauthorized, gin.H{"success":false,
			"errors":"incorrect credentials"})
		return
	}

	// Email 인증 여부 Check
	if user.EmailAuth {
		// 이메일 발송
		h.sendAuthMail(c, user.Name, user.Email)
		c.JSON(messages.StatusEmailAuthConfirm,
			gin.H{"success":false,
				"msg":messages.RestStatusText(messages.StatusEmailAuthConfirm)})
		return
	} else if user.GroupEmailAuth {
		if loginMsg.Email != "" {
			// 그룹에 등록된 이메일지 Check
			// userId, userEmail
			checkGroupEmailAuth()

			// 이메일 발송
			h.sendAuthMail(c, user.Name, user.Email)
			c.JSON(messages.StatusEmailAuthConfirm,
				gin.H{"success":false,
					"msg":messages.RestStatusText(messages.StatusEmailAuthConfirm)})
		} else {
			// 이메일 수신 후 발송
			c.JSON(messages.StatusInputEmailAuth,
				gin.H{"success":false,
					"msg":messages.RestStatusText(messages.StatusInputEmailAuth)})
		}
		return
	}

	// Token을 발급해서 응답한다.
	responseWithToken(c, user.UserId, user.Email, user.Name)
}

func checkGroupEmailAuth() {

}

func responseWithToken(c *gin.Context, userId, userEmail, userName string) {
	expirationTime := time.Now().Add(1 * time.Minute)
	claims := &Claims{
		User: models.User{
			UserId: userId,
			Email: userEmail,
			Name: userName,
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
		Name: "token",
		Value: tokenString,
		Expires: expirationTime,
	})
	userInfo, err := convertUser2Msg(claims.User)
	c.JSON(http.StatusOK,
		gin.H{"success":true, "msg":http.StatusText(http.StatusOK), "user":userInfo})
}

func convertUser2Msg(user models.User) (info messages.UserInfo, err error) {
	info.Id = user.UserId
	info.Name = user.Name
	info.Email = user.Email
	info.EmailAuthFlag = user.EmailAuth
	info.EmailAuthGroupFlag = user.GroupEmailAuth
	return info, err
}

func (h *Handler) sendAuthMail(c *gin.Context, username, email string) {

	// 1. Generate UUID
	uuid, err := utils.NewUUID()
	if err != nil {
		fmt.Println("Failed to generate UUID!!")
		c.JSON(http.StatusNotAcceptable, gin.H{"success":false, "msg":"UUID 생성에 실패함!"})
		return
	}

	fmt.Println("Email Authentication process....")

	// 2. Get from DB
	uniqId := username + email
	userEmailAuth, err := h.db.GetUserEmailAuthByUniqId(uniqId)
	fmt.Println("err:", err, " userEmailAuth:", userEmailAuth)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"success":false, "msg":"이 계정은 인증 DB에 존재하지 않음."})
		return
	}
	userEmailAuth.EmailAuthConfirm = false
	userEmailAuth.EmailAuthStore = uuid

	// 3. Update to DB
	h.db.UpdateUserEmailAuth(userEmailAuth)

	// 4. Send email
	emailmsg := utils.MailMsg {
		To: email,
		Header: "콘텐츠브릿지 로그인 Email 인증",
		ServerIp: svcmgrAddress,
		Uuid:   uuid,
		UserId: username,
		Text:   "계정에 대한 이메일 인증을 위해서 아래 URL을 클릭하시기 바랍니다.",
	}
	err = utils.SendMail(smtpServer, emailmsg)
	fmt.Println("SendMail: err ", err)
}

func (h *Handler) Logout(c *gin.Context) {
	http.SetCookie(c.Writer, &http.Cookie{
		Name:       "token",
		MaxAge:     -1,
	})
	c.JSON(http.StatusNoContent, gin.H{"success":true, "msg":"logged out in successfully"})
}

func (h *Handler) GetSession(c *gin.Context) {
	user, isAuthenticated, token := AuthMiddleware(c, jwtKey)
	if !isAuthenticated {
		c.JSON(http.StatusUnauthorized,
			gin.H{"success":false, "msg":"unauthorized"})
		return
	}

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
		uniqId := username + email
		userAuth, err := h.db.GetUserEmailAuthByUniqId(uniqId)
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
		ServerIp:    svcmgrAddress,
		Uuid:   uuid,
		UserId: username,
		Text:   "계정에 대한 이메일 인증을 위해서 아래 URL을 클릭하시기 바랍니다.",
	}
	err = utils.SendMail(smtpServer, emailmsg)
	fmt.Println("SendMail: err ", err)

	c.JSON(http.StatusOK, gin.H{"success":true, "user":user})
}


func (h *Handler) EmailConfirm(c *gin.Context) {
	secret := c.Param("secret")
	revId := c.Param("id")
	revEmail := c.Param("email")
	fmt.Println("EmailConfirm: Secret - ", secret)
	fmt.Println("EmailConfirm: id - ", revId)
	fmt.Println("EmailConfirm: email - ", revEmail)

	// Get from DB
	uniqId := revId + revEmail
	userAuth, err := h.db.GetUserEmailAuthByUniqId(uniqId)
	fmt.Println("err:", err, " userAuth:", userAuth)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"success":false, "msg":"이 계정은 인증 DB에 존재하지 않음."})
		return
	}
	if userAuth.EmailAuthStore == secret {
		c.JSON(http.StatusOK, gin.H{"success":true, "msg":""})
	}
	c.JSON(messages.StatusFailedEmailAuth, gin.H{"success":false,
		"msg":messages.RestStatusText(messages.StatusFailedEmailAuth)})
}