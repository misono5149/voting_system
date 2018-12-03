package middleware

import (
	"net/http"
	"time"
	"voting_system/app/services"
	"voting_system/app/services/models"

	ginJwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
)

type Login struct {
	StudentId string `json:"student_id"`
	Password  string `json:"password"`
}

const JWT_KEY = "apple"

func RenderData(data interface{}, ctx *gin.Context) {
	meta := gin.H{
		"status_code": http.StatusOK,
		"error":       false,
	}

	body := gin.H{
		"data":   data,
		"errors": false,
		"meta":   meta,
	}

	RenderResponse(&body, http.StatusOK, ctx)
}

func RenderErrorSingle(message string, status int, ctx *gin.Context) {
	meta := gin.H{
		"status_code": status,
		"error":       true,
	}

	body := gin.H{
		"data":   nil,
		"errors": message,
		"meta":   meta,
	}

	RenderResponse(&body, status, ctx)
}

func RenderResponse(body *gin.H, status int, ctx *gin.Context) {
	ctx.IndentedJSON(status, body)
}

func GinJWTMiddlewareHandler() *ginJwt.GinJWTMiddleware {
	authMiddleware := &ginJwt.GinJWTMiddleware{
		Realm:      "voting api",
		Key:        []byte(JWT_KEY),
		Timeout:    time.Hour,
		MaxRefresh: time.Hour,
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVals Login
			if err := c.ShouldBind(&loginVals); err != nil {
				return "", ginJwt.ErrMissingLoginValues
			}
			userId := loginVals.StudentId
			password := loginVals.Password

			voter := services.GetVoterInfo(userId, password)

			if voter != (models.Voter{}) {
				return &models.Voter{
					StudentId: userId,
					Name:      voter.Name,
				}, nil
			}
			return nil, ginJwt.ErrFailedAuthentication
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},

		TokenLookup:   "header:Authorization",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	}
	return authMiddleware
}
