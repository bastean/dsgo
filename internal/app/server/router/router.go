package router

import (
	"embed"

	"github.com/bastean/dsgo/internal/app/server/middleware"
	"github.com/gofiber/fiber/v2"
)

func Routing(app *fiber.App, files *embed.FS) {
	app.Use(middleware.Recover)

	app.Use("/public", middleware.FileSystem(files))

	app.Use(middleware.Headers)

	app.Use(middleware.Limiter)

	Routes(app)
}
