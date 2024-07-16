package middleware

import (
	"net/http"

	"github.com/bastean/dsgo/internal/app/server/service/errors"
	"github.com/bastean/dsgo/internal/app/server/service/logger"
	"github.com/bastean/dsgo/internal/app/server/util/reply"
	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		var invalidValue *errors.ErrInvalidValue
		var alreadyExist *errors.ErrAlreadyExist
		var notExist *errors.ErrNotExist
		var failure *errors.ErrFailure
		var internal *errors.ErrInternal

		for _, err := range c.Errors {
			switch {
			case errors.As(err, &invalidValue):
				c.JSON(http.StatusUnprocessableEntity, reply.JSON(false, invalidValue.What, invalidValue.Why))
			case errors.As(err, &alreadyExist):
				c.JSON(http.StatusConflict, reply.JSON(false, alreadyExist.What, alreadyExist.Why))
			case errors.As(err, &notExist):
				c.JSON(http.StatusNotFound, reply.JSON(false, notExist.What, notExist.Why))
			case errors.As(err, &failure):
				c.JSON(http.StatusBadRequest, reply.JSON(false, failure.What, failure.Why))
			case errors.As(err, &internal):
				c.JSON(http.StatusInternalServerError, reply.JSON(false, "internal server error", reply.Payload{}))
				fallthrough
			default:
				logger.Error(err.Error())
			}
		}
	}
}