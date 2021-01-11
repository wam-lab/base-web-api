package response

import (
	"github.com/gin-gonic/gin"
	"github.com/wam-lab/base-web-api/common/errno"
	"net/http"
)

func Json(c *gin.Context, result errno.Result) {
	c.JSON(http.StatusOK, result)
}

func NeedLogin(c *gin.Context, result errno.Result) {
	c.JSON(http.StatusUnauthorized, result)
}

func Forbidden(c *gin.Context, result errno.Result) {
	c.JSON(http.StatusForbidden, result)
}

func JsonWithTrace(c *gin.Context, result errno.Result, trace string) {
	traceId := c.Value(trace).(string)
	c.JSON(http.StatusOK, result.WithTrace(traceId))
}
