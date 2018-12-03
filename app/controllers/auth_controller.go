package controllers

import (
	"net/http"
	"time"
	"voting_system/app/middleware"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

const JWT_KEY = "apple"
const createdFormat = "2006-01-02 15:04:05"

func Login(c *gin.Context) {
	var loginVals middleware.Login
	c.ShouldBindJSON(&loginVals)

	// Create the token
	token := jwt.New(jwt.GetSigningMethod("HS256"))
	claims := token.Claims.(jwt.MapClaims)

	expire := time.Now().Add(time.Hour)
	claims["id"] = loginVals.StudentId
	claims["exp"] = expire.Format(createdFormat)
	claims["orig_iat"] = time.Now().Unix()

	tokenString, err := token.SignedString([]byte(JWT_KEY))

	if err != nil {
		middleware.RenderErrorSingle("Create JWT Token faild", http.StatusUnauthorized, c)
		return
	}

	c.JSON(200, gin.H{
		"is_success":  200,
		"expire_time": claims["exp"],
		"auth_token":  tokenString,
	})
	/*
		RenderData(gin.H{
			//"login":       claims["id"],
			"expire_time": claims["exp"],
			//"iat":         claims["orig_iat"],
			"token": tokenString,
		}, c)
	*/
}
