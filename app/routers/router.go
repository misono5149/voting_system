package routers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRoutes() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(cors.Default())

	v1 := r.Group("/")
	v1.Use()
	{
		AddRoutesVoter(v1)
		AddRoutesAdministrator(v1)
	}
	return r
}
