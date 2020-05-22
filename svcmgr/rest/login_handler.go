package rest

import (
	"cmpService/common/lib"
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
	getuser, err := h.db.GetUserById(user.ID)
	if err != nil {
		lib.LogWarnln(err)
		return false
	} else if user.ID == getuser.ID {
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
	var loginUser models.User
	c.Bind(&loginUser)

	fmt.Println("LoginUser2:", loginUser)
	user, err := h.db.GetUserByEmail(loginUser.Email)
	fmt.Println("LoginUser2:", user.String())
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"success":false, "errors":err})
		return
	}
	match := models.CheckPasswordHash(loginUser.Password, user.Password)
	if !match {
		c.JSON(http.StatusUnauthorized, gin.H{"success":false, "errors":"incorrect credentials"})
		return
	}

	expirationTime := time.Now().Add(30 * time.Minute)
	claims := &Claims{
		User: models.User{
			ID: user.ID,
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
	var loginUser models.User
	c.Bind(&loginUser)

	fmt.Println("LoginUser:", loginUser)
	user, err := h.db.GetUserById(loginUser.ID)
	fmt.Println("in DB, LoginUser:", user)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"success":false, "errors":err})
		return
	}
	match := models.CheckPasswordHash(loginUser.Password, user.Password)
	if !match {
		c.JSON(http.StatusUnauthorized, gin.H{"success":false, "errors":"incorrect credentials"})
		return
	}

	expirationTime := time.Now().Add(1 * time.Minute)
	claims := &Claims{
		User: models.User{
			ID: user.ID,
			Email: user.Email,
			Name: user.Name,
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
	fmt.Println("Success: token - ",tokenString)
	c.JSON(http.StatusOK,
		gin.H{"success":true, "msg":"logged in successfully", "user":claims.User,
			"token":tokenString})
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
	uuid := c.Param("secret")
	fmt.Println("EmailConfirm: Secret - ", uuid)

	user, isAuthenticated, token := AuthMiddleware(c, jwtKey)
	if !isAuthenticated {
		c.JSON(http.StatusUnauthorized,
			gin.H{"success":false, "msg":"unauthorized"})
		return
	}

	fmt.Println(">>> USER: ", user)
	fmt.Println(">>> token: ", token)
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

	if haveEmailAuth || haveGroupEmailAuth {
		// Get from DB
		uniqId := username + email
		userAuth, err := h.db.GetUserEmailAuthByUniqId(uniqId)
		fmt.Println("err:", err, " userAuth:", userAuth)
		if err != nil {
			c.JSON(http.StatusNotAcceptable, gin.H{"success":false, "msg":"이 계정은 인증 DB에 존재하지 않음."})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"success":true, "msg":""})
}