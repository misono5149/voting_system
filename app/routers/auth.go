package routers

import (
	"voting_system/app/controllers"

	"github.com/gin-gonic/gin"
)

func AddRoutesAuth(r *gin.RouterGroup) {
	r.POST("/auth/signup", controllers.SignUp)    // 회원가입
	r.POST("/auth/login", controllers.Login)      // 로그인
	r.GET("/auth/exists/:id", controllers.Exists) // 아이디 중복확인 api 라든가 나중에 체크
	r.POST("/auth/logout", controllers.Logout)    // 로그아웃
}
