package page

import (
	"github.com/bastean/dsgo/internal/app/server/component/page/home"
	"github.com/bastean/dsgo/internal/pkg/service/errors"
	"github.com/gofiber/fiber/v2"
)

func Home() fiber.Handler {
	return func(c *fiber.Ctx) error {
		if err := home.Page().Render(c.Context(), c.Type("html").Response().BodyWriter()); err != nil {
			return errors.BubbleUp(err, "Home")
		}

		return nil
	}
}
