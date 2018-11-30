package routers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
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
