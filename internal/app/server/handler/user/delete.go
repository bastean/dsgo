package user

import (
	"github.com/bastean/dsgo/internal/app/server/util/reply"
	"github.com/bastean/dsgo/internal/pkg/service/errors"
	"github.com/bastean/dsgo/internal/pkg/service/user"
	"github.com/gofiber/fiber/v2"
)

func Delete() fiber.Handler {
	return func(c *fiber.Ctx) error {
		primitive := &struct {
			Name string
		}{}

		err := c.BodyParser(primitive)

		if err != nil {
			return errors.BubbleUp(err, "Delete")
		}

		err = user.Delete.Run(primitive.Name)

		if err != nil {
			return errors.BubbleUp(err, "Delete")
		}

		err = c.Status(fiber.StatusCreated).JSON(reply.JSON(true, "user deleted", reply.Payload{}))

		if err != nil {
			return errors.BubbleUp(err, "Delete")
		}

		return nil
	}
}
