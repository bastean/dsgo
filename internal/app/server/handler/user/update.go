package user

import (
	"github.com/bastean/dsgo/internal/app/server/util/reply"
	"github.com/bastean/dsgo/internal/pkg/service/errors"
	"github.com/bastean/dsgo/internal/pkg/service/user"
	"github.com/gofiber/fiber/v2"
)

func Update() fiber.Handler {
	return func(c *fiber.Ctx) error {
		primitive := new(user.Primitive)

		err := c.BodyParser(primitive)

		if err != nil {
			return errors.BubbleUp(err, "Update")
		}

		err = user.Update.Run(primitive)

		if err != nil {
			return errors.BubbleUp(err, "Update")
		}

		err = c.Status(fiber.StatusCreated).JSON(reply.JSON(true, "Updated", reply.Payload{}))

		if err != nil {
			return errors.BubbleUp(err, "Update")
		}

		return nil
	}
}
