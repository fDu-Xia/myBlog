package middleware

import (
	"github.com/gin-gonic/gin"
	"myBlog/internal/known"
	"myBlog/internal/pkg/core"
	"myBlog/internal/pkg/errno"
	"myBlog/pkg/token"
)

func Authn() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 解析 JWT Token
		username, err := token.ParseRequest(c)
		if err != nil {
			core.WriteResponse(c, errno.ErrTokenInvalid, nil)
			c.Abort()

			return
		}

		c.Set(known.XUsernameKey, username)
		c.Next()
	}
}
