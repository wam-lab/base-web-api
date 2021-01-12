package router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/wam-lab/base-web-api/internal/api/v1"
	"github.com/wam-lab/base-web-api/internal/middleware"
)

func InitUserRouter(g *gin.RouterGroup) {
	auth := g.Group("/user").Use(middleware.JwtAuth())
	{
		auth.POST("/info", v1.UserInfo)
	}
}