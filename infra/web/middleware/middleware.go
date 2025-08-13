package middleware

import (
	"github.com/EDEN-NN/hydra-api/internal/apperrors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		for _, err := range c.Errors {
			switch e := err.Err.(type) {
			case *apperrors.AppError:
				c.JSON(errorToHTTPStatus(e.Code), gin.H{
					"error":   e.Message,
					"code":    e.Code,
					"details": e.Metadata,
				})
			default:
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "internal server error",
					"code":  "INTERNAL",
				})
			}
			return
		}
	}
}


func errorToHTTPStatus(code apperrors.ErrorCode) int {
	switch code {
	case apperrors.ENOTFOUND:
		return http.StatusNotFound
	case apperrors.EINVALID:
		return http.StatusBadRequest
	case apperrors.EUNAUTHORIZED:
		return http.StatusUnauthorized
	case apperrors.ECONFLICT:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
