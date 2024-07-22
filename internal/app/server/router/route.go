package router

import (
	"github.com/bastean/dsgo/internal/app/server/handler/page"
	"github.com/bastean/dsgo/internal/app/server/handler/user"
	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
	public := app.Group("/")

	public.Get("/", page.Home())
	public.Put("/", user.Create())
	public.Post("/", user.Read())
	public.Patch("/", user.Update())
	public.Delete("/", user.Delete())

	app.Use(page.Default())
}
