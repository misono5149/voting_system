package routers

import (
	"github.com/gin-gonic/gin"
)

func InitRoutes() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/")
	v1.Use()
	{
		AddRoutesVoter(v1)
		AddRoutesAdministrator(v1)
	}
	return r
}
