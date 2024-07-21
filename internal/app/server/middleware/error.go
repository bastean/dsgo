package middleware

import (
	"github.com/bastean/dsgo/internal/app/server/util/reply"
	"github.com/bastean/dsgo/internal/pkg/service/errors"
	"github.com/bastean/dsgo/internal/pkg/service/logger/log"
	"github.com/gofiber/fiber/v2"
)

func Error(c *fiber.Ctx, err error) error {
	var errInvalidValue *errors.ErrInvalidValue
	var errAlreadyExist *errors.ErrAlreadyExist
	var errNotExist *errors.ErrNotExist
	var errFailure *errors.ErrFailure
	var errInternal *errors.ErrInternal

	switch {
	case errors.As(err, &errInvalidValue):
		c.Status(fiber.StatusUnprocessableEntity).JSON(reply.JSON(false, errInvalidValue.What, errInvalidValue.Why))
	case errors.As(err, &errAlreadyExist):
		c.Status(fiber.StatusConflict).JSON(reply.JSON(false, errAlreadyExist.What, errAlreadyExist.Why))
	case errors.As(err, &errNotExist):
		c.Status(fiber.StatusNotFound).JSON(reply.JSON(false, errNotExist.What, errNotExist.Why))
	case errors.As(err, &errFailure):
		c.Status(fiber.StatusBadRequest).JSON(reply.JSON(false, errFailure.What, errFailure.Why))
	case errors.As(err, &errInternal):
		c.Status(fiber.StatusInternalServerError).JSON(reply.JSON(false, "internal server error", reply.Payload{}))
		fallthrough
	default:
		log.Error(err.Error())
	}

	return nil
}
