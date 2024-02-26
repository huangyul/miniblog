package core

import (
	"miniblog/internal/pkg/errno"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ErrResponse 带有 err 的相应
type ErrResponse struct {
	Code    string
	Message string
}

func WriteResponse(ctx *gin.Context, err error, data any) {
	if err != nil {
		hCode, code, message := errno.Decode(err)
		ctx.JSON(hCode, ErrResponse{
			Code:    code,
			Message: message,
		})
		return
	}

	ctx.JSON(http.StatusOK, data)
}
