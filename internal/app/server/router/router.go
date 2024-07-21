package router

import (
	"embed"

	"github.com/bastean/dsgo/internal/app/server/middleware"
	"github.com/gofiber/fiber/v2"
)

func Routing(server *fiber.App, files *embed.FS) {
	server.Use(middleware.Recover)

	server.Use("/public", middleware.FileSystem(files))

	server.Use(middleware.Headers)

	server.Use(middleware.Limiter)

	Routes(server)
}
