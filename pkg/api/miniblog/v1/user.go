package v1

type CreateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

// LoginRequest 指定 ’POST /login‘ 请求参数
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// LoginResponse `POST /login` 返回参数
type LoginRespose struct {
	Token string `json:"token"`
}
