package router

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/wam-lab/base-web-api/internal/api/v1"
	"github.com/wam-lab/base-web-api/internal/middleware"
)

func InitAuthRouter(g *gin.RouterGroup) {
	auth := g.Group("/auth")
	{
		auth.POST("/login", v1.Login)
		auth.POST("/register", v1.Register)
	}

	auth.Use(middleware.JwtAuth())
	{
		auth.POST("/refresh", v1.RefreshToken)
	}
}
