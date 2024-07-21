package user

import (
	"github.com/bastean/dsgo/internal/app/server/util/reply"
	"github.com/bastean/dsgo/internal/pkg/service/errors"
	"github.com/bastean/dsgo/internal/pkg/service/user"
	"github.com/gofiber/fiber/v2"
)

func Read() fiber.Handler {
	return func(c *fiber.Ctx) error {
		primitive := &struct {
			Name string
		}{}

		err := c.BodyParser(primitive)

		if err != nil {
			return errors.BubbleUp(err, "Read")
		}

		user, err := user.Read.Run(primitive.Name)

		if err != nil {
			return errors.BubbleUp(err, "Read")
		}

		err = c.Status(fiber.StatusCreated).JSON(reply.JSON(true, "user found", reply.Payload{
			"name": user.Name,
			"role": user.Role,
		}))

		if err != nil {
			return errors.BubbleUp(err, "Read")
		}

		return nil
	}
}
