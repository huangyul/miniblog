package errno

import "net/http"

var (
	OK = &Errno{HTTP: http.StatusOK, Code: "", Message: ""}

	InternalServerError = &Errno{HTTP: http.StatusInternalServerError, Code: "InternalError", Message: "Internal server error."}

	ErrPageNotFound = &Errno{HTTP: http.StatusNotFound, Code: "ResourceNotFound.PageNotFound", Message: "Page not found"}

	ErrBind = &Errno{HTTP: http.StatusBadRequest, Code: "BindError", Message: "Error"}

	ErrInvalidParameter = &Errno{
		HTTP:    http.StatusBadRequest,
		Code:    "InvalidParameter",
		Message: "parameter verification failed",
	}
)
