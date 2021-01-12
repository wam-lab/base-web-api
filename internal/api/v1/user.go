package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/wam-lab/base-web-api/common/errno"
	"github.com/wam-lab/base-web-api/internal/global/response"
)

func UserInfo(c *gin.Context)  {
	response.Json(c, errno.OK.WithData(map[string]interface{}{
		"id": 1,
		"username": "yguilai",
	}))
}
