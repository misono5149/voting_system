package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

func InitRoutes() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/")
	v1.Use(cors.Default())
	{
		AddRoutesVoter(v1)
		AddRoutesAdministrator(v1)
	}
	return r
}
