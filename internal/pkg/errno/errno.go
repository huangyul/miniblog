package errno

import "fmt"

type Errno struct {
	HTTP    int
	Code    string
	Message string
}

// Error 返回 errno 的 message
func (e *Errno) Error() string {
	return e.Message
}

// 设置 message
func (e *Errno) SetMessage(format string, args ...any) *Errno {
	e.Message = fmt.Sprintf(format, args...)
	return e
}

func Decode(err error) (int, string, string) {
	if err == nil {
		return OK.HTTP, OK.Code, OK.Message
	}

	switch typed := err.(type) {
	case *Errno:
		return typed.HTTP, typed.Code, typed.Message

	}

	return InternalServerError.HTTP, InternalServerError.Code, err.Error()
}
