package core

import (
	"github.com/gin-gonic/gin"
	"myBlog/internal/pkg/errno"
	"net/http"
)

// ErrResponse 定义了发生错误时的返回消息.
type ErrResponse struct {
	// Code 指定了业务错误码.
	Code string `json:"code"`

	// Message 包含了可以直接对外展示的错误信息.
	Message string `json:"message"`
}

func WriteResponse(c *gin.Context, err *errno.Errno, data interface{}) {
	if err != nil {
		hCode, code, message := err.Decode()
		c.JSON(hCode, ErrResponse{
			Code:    code,
			Message: message,
		})
		return
	}
	c.JSON(http.StatusOK, data)
}
