package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// NoCache 是一个 Gin 中间件，用来禁止客户端缓存 HTTP 请求的返回结果.
func NoCache(c *gin.Context) {
	c.Header("Cache-Control", "no-cache, no-store, max-age=0, must-revalidate, value")
	c.Header("Expires", "Thu, 01 Jan 1970 00:00:00 GMT")
	c.Header("Last-Modified", time.Now().UTC().Format(http.TimeFormat))
	c.Next()
}
