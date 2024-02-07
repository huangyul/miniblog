package errno

import "net/http"

var (
	// OK 表示成功正常的返回
	OK = &Errno{http.StatusOK, "ok", "ok"}

	// InternalServerErr 表示未知错误
	InternalServerErr = &Errno{http.StatusInternalServerError, "500 server error", "server error"}

	// NotFoundErr 表示网页找不到
	NotFoundErr = &Errno{http.StatusNotFound, "not found", "not found"}
)
