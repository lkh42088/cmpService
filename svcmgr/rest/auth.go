package rest

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(c *gin.Context, jwtKey[]byte) (jwt.MapClaims, bool, string) {
	ck, err := c.Request.Cookie("token")
	fmt.Println("AuthMiddleware:", ck)
	if err != nil {
		fmt.Println("AuthMiddleware error:", err)
		return nil, false, ""
	}
	tokenString :=ck.Value
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok :=token.Method.(*jwt.SigningMethodHMAC); !ok {
			fmt.Println("AuthMiddleware : failed")
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return jwtKey, nil
	})
	//fmt.Println("AuthMiddleware token:", token)
	//fmt.Println("size (", len(token.Raw),") raw:", token.Raw)
	//fmt.Println("signature:", token.Signature)

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println("AuthMiddleware : Success")
		return claims, true, token.Raw
	}
	fmt.Println("AuthMiddleware : nil false")
	return nil, false, ""
}
