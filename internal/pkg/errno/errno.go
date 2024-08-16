package errno

import "fmt"

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

func (errno *Errno) Decode() (int, string, string) {
	return errno.HTTP, errno.Code, errno.Message
}
