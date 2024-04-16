package core

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ErrResponse 定义了发生错误时返回的信息
type ErrResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func WriteResponse(c *gin.Context, err error, data any) {
	if err != nil {

		return
	}
	c.JSON(http.StatusOK, data)
}
