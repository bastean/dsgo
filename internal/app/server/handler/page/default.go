package page

import (
	"github.com/bastean/dsgo/internal/pkg/service/errors"
	"github.com/gofiber/fiber/v2"
)

func Default() fiber.Handler {
	return func(c *fiber.Ctx) error {
		if err := c.Redirect("/", fiber.StatusFound); err != nil {
			return errors.BubbleUp(err, "Default")
		}

		return nil
	}
}
