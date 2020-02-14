package errors

import (
	"github.com/gin-gonic/gin"
)

type ErrorMessage struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Reason  string `json:"reason"`
}

func HttpError(ctx *gin.Context, code int, err error) {
	errorMessage := ErrorMessage{
		Code:   code,
		Reason: err.Error(),
	}

	switch code {
	case 400:
		errorMessage.Message = "Bad request"
	case 401:
		errorMessage.Message = "Unauthorized"
	case 403:
		errorMessage.Message = "Forbidden"
	case 404:
		errorMessage.Message = "Not found"
	case 405:
		errorMessage.Message = "Method Not Allowed"
	case 409:
		errorMessage.Message = "Conflict"
	case 500:
		errorMessage.Message = "Internal error"
	default:
		errorMessage.Message = "Errors"
	}

	ctx.JSON(code, errorMessage)
	ctx.Abort()
}
