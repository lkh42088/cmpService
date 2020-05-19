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
	errors.HandleErr(c, err)
	http.SetCookie(c.Writer, &http.Cookie{
		Name: "token",
		Value: tokenString,
		Expires: expirationTime,
	})
	fmt.Println("Success: token - ",tokenString)
	c.JSON(http.StatusOK,
		gin.H{"success":true, "msg":"loggged in successfully", "user":claims.User,
			"token":tokenString})
}

func (h *Handler) GetSession(c *gin.Context) {
	user, isAuthenticated := AuthMiddleware(c, jwtKey)
	if !isAuthenticated {
		c.JSON(http.StatusUnauthorized,
			gin.H{"success":false, "msg":"unauthorized"})
	}
	c.JSON(http.StatusOK, gin.H{"success":true, "user":user})
}

