package router

import (
	"github.com/gin-gonic/gin"
	"github.com/wam-lab/base-web-api/common/conno"
	"github.com/wam-lab/base-web-api/internal/global"
	"github.com/wam-lab/base-web-api/internal/middleware"
	"time"
)

func Router() *gin.Engine {
	r := gin.New()
	if global.Config.GetString("mode") == conno.PRO {
		gin.SetMode(gin.ReleaseMode)
	}

	r.Use(
		middleware.LoggerWithZap(global.Log, time.RFC3339, true),
		middleware.RecoveryWithZap(global.Log, true),
		middleware.Cors(),
	)

	apiGroup := r.Group("/api/v1")
	InitAuthRouter(apiGroup)
	return r
}
