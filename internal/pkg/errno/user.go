package errno

var (
	ErrUserAlreadyExist = &Errno{
		HTTP:    400,
		Code:    "FailedOperation.UserAlreadyExit",
		Message: "User already exist.",
	}
)
