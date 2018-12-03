package routers

import (
	"voting_system/app/middleware"

	"github.com/gin-gonic/gin"
)

func AddRoutesAuth(r *gin.RouterGroup) {
	// 선거 관리
	r.POST("/login", middleware.GinJWTMiddlewareHandler().LoginHandler)
}
