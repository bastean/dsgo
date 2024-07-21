package server

import (
	"context"
	"embed"

	"github.com/bastean/dsgo/internal/app/server/middleware"
	"github.com/bastean/dsgo/internal/app/server/router"
	"github.com/bastean/dsgo/internal/pkg/service/errors"
	"github.com/gofiber/fiber/v2"
)

//go:embed static
var Files embed.FS

var Server *fiber.App

func Run(port string) error {
	Server := fiber.New(fiber.Config{
		AppName:      "dsGO",
		ErrorHandler: middleware.Error,
	})

	router.Routing(Server, &Files)

	if err := Server.Listen(":" + port); err != nil {
		return errors.BubbleUp(err, "Run")
	}

	return nil
}

func Stop(ctx context.Context) error {
	if err := Server.ShutdownWithContext(ctx); err != nil {
		return errors.BubbleUp(err, "Stop")
	}

	return nil
}
