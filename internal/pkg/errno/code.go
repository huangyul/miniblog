package errno

var (
	// OK 成功的请求
	OK = &Errno{HTTP: 200, Code: "", Message: ""}

	InternalServerError = &Errno{HTTP: 500, Code: "", Message: "Internal server error"}

	ErrPageNotFound = &Errno{HTTP: 404, Code: "", Message: "Page not found"}
)
