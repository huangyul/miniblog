package core

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/huangyul/miniblog/internal/pkg/errno"
)

// ErrResponse 定义了发生错误时返回的信息
type ErrResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type Response struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

func WriteResponse(c *gin.Context, err error, data any) {
	if err != nil {
		hcode, code, message := errno.Decode(err)
		c.JSON(hcode, ErrResponse{
			Code:    code,
			Message: message,
		})
		return
	}
	c.JSON(http.StatusOK, data)
}
