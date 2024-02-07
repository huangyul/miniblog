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

// Error 返回 Message
func (e *Errno) Error() string {
	return e.Message
}

// SetMessage 设置 Errno 的 Message
func (e *Errno) SetMessage(format string, values ...any) *Errno {
	e.Message = fmt.Sprintf(format, values)
	return e
}

// Decode 尝试从 error 中解析出 errno
func Decode(err error) (int, string, string) {
	if err == nil {
		return OK.HTTP, OK.Message, OK.Code
	}

	var typed *Errno
	ok := errors.As(err, &typed)
	if !ok {
		return InternalServerErr.HTTP, InternalServerErr.Message, InternalServerErr.Code
	}

	return typed.HTTP, typed.Message, typed.Code

}
