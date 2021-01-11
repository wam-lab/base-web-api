package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/wam-lab/base-web-api/common/conno"
	"github.com/wam-lab/base-web-api/common/errno"
	"github.com/wam-lab/base-web-api/common/jwt"
	"github.com/wam-lab/base-web-api/internal/global"
	"github.com/wam-lab/base-web-api/internal/global/response"
)

func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.Request.Header.Get(conno.TokenHeader)
		if tokenStr == "" {
			response.NeedLogin(c, errno.ErrUserUnlogin)
			c.Abort()
			return
		}

		token, claims, err := jwt.ParseTokenClaims(tokenStr, global.Config.GetString("Jwt.Secret"))
		if err != nil {
			response.Forbidden(c, errno.ErrInvalidToken)
			c.Abort()
			return
		}

		if !token.Valid {
			response.NeedLogin(c, errno.ErrExpiredToken)
			c.Abort()
			return
		}

		c.Set("claims", claims)
		c.Next()
	}
}
