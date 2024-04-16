package errno

import "net/http"

var (
	OK = &Errno{HTTP: 200, Code: "", Message: ""}

	InternalServerError = &Errno{HTTP: 500, Code: "internalerror", Message: "internal server error."}

	ErrPageNotFound = &Errno{HTTP: http.StatusNotFound, Code: "ResourceNotFound", Message: "page not found"}
)
