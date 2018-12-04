package controllers

import (
	"time"

	"github.com/gin-gonic/gin"
	"gopkg.in/dgrijalva/jwt-go.v3"
)

const JWT_KEY = "secure_as_hell_key"
const createdFormat = "2006-01-02 15:04:05"

type login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	var loginVals login

	// Create the token
	token := jwt.New(jwt.GetSigningMethod("HS256"))
	claims := token.Claims.(jwt.MapClaims)

	/*	// add extra things to the jwt claims
		for key, value := range ginJwt.PayloadFunc(loginVals.Username) {
		claims[key] = value
	}*/

	expire := time.Now().Add(time.Hour)
	claims["id"] = loginVals.Username
	claims["exp"] = expire.Unix()
	claims["orig_iat"] = time.Now().Format(createdFormat)

	tokenString, err := token.SignedString([]byte(JWT_KEY))

	if err != nil {
		panic(err)
	}

	c.JSON(200, gin.H{
		"auth_token":  tokenString,
		"expire_time": expire.Format(createdFormat),
		"is_success":  200,
	})
}
