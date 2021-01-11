package router

import (
	"github.com/gin-gonic/gin"
	"github.com/wam-lab/base-web-api/common/conno"
	"github.com/wam-lab/base-web-api/internal/global"
	"github.com/wam-lab/base-web-api/internal/middleware"
)

func Router() *gin.Engine {
	var r *gin.Engine
	c := global.Config
	if c.GetString("mode") != conno.PRO {
		r = gin.Default()
	} else {
		// TODO new production gin engine
		r = gin.Default()
	}

	r.Use(middleware.Cors())

	apiGroup := r.Group("/api/v1")
	InitAuthRouter(apiGroup)
	return r
}
