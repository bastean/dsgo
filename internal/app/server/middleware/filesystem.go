package middleware

import (
	"embed"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
)

func FileSystem(files *embed.FS) fiber.Handler {
	return filesystem.New(filesystem.Config{
		Root:   http.FS(files),
		Browse: true,
	})
}
