package errno

import (
	"errors"
	"fmt"
)

type Errno struct {
	HTTP    int
	Code    string
	Message string
}

func (errno *Errno) Error() string {
	return errno.Message
}

func (errno *Errno) SetMessage(format string, args ...interface{}) *Errno {
	errno.Message = fmt.Sprintf(format, args...)
	return errno
}

func Decode(err error) (int, string, string) {
	if err == nil {
		return OK.HTTP, OK.Code, OK.Message
	}

	var typed *Errno
	switch {
	case errors.As(err, &typed):
		return typed.HTTP, typed.Code, typed.Message
	default:
	}

	// 默认返回未知错误码和错误信息. 该错误代表服务端出错
	return InternalServerError.HTTP, InternalServerError.Code, err.Error()
}
