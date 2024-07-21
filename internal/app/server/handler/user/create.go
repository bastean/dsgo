package user

import (
	"github.com/bastean/dsgo/internal/app/server/util/reply"
	"github.com/bastean/dsgo/internal/pkg/service/errors"
	"github.com/bastean/dsgo/internal/pkg/service/user"
	"github.com/gofiber/fiber/v2"
)

func Create() fiber.Handler {
	return func(c *fiber.Ctx) error {
		primitive := new(user.Primitive)

		err := c.BodyParser(primitive)

		if err != nil {
			return errors.BubbleUp(err, "Create")
		}

		err = user.Create.Run(primitive)

		if err != nil {
			return errors.BubbleUp(err, "Create")
		}

		err = c.Status(fiber.StatusCreated).JSON(reply.JSON(true, "Created", reply.Payload{}))

		if err != nil {
			return errors.BubbleUp(err, "Create")
		}

		return nil
	}
}
