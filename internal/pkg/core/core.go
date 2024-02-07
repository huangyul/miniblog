package core

import (
	"github.com/gin-gonic/gin"
	"miniblog/internal/pkg/errno"
	"net/http"
)

type ErrResponse struct {
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

func WriteResponse(c *gin.Context, err error, data any) {
	if err != nil {
		httpCode, code, message := errno.Decode(err)
		c.JSON(httpCode, ErrResponse{
			Code:    code,
			Message: message,
		})
		return
	}

	c.JSON(http.StatusOK, data)
}
