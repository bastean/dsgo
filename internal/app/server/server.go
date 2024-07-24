package server

import (
	"context"
	"embed"
	"fmt"

	"github.com/bastean/dsgo/internal/app/server/middleware"
	"github.com/bastean/dsgo/internal/app/server/router"
	"github.com/bastean/dsgo/internal/pkg/service/env"
	"github.com/bastean/dsgo/internal/pkg/service/errors"
	"github.com/bastean/dsgo/internal/pkg/service/logger/log"
	"github.com/gofiber/fiber/v2"
)

var (
	Server = &struct {
		Fiber string
	}{
		Fiber: log.Server("fiber"),
	}
)

//go:embed static
var Files embed.FS

var App *fiber.App

func Up() error {
	log.Starting(Server.Fiber)

	App = fiber.New(fiber.Config{
		AppName:      "dsGO",
		ErrorHandler: middleware.Error,
	})

	router.Routing(App, &Files)

	if err := App.Listen(":" + env.ServerFiberPort); err != nil {
		log.CannotBeStarted(Server.Fiber)
		return errors.BubbleUp(err, "Up")
	}

	log.Started(Server.Fiber)

	log.Info(fmt.Sprintf("%s listening on :%s", Server.Fiber, env.ServerFiberPort))

	if proxy, ok := env.HasServerFiberProxy(); ok {
		log.Info(fmt.Sprintf("%s proxy listening on :%s", Server.Fiber, proxy))
	}

	return nil
}

func Down(ctx context.Context) error {
	log.Stopping(Server.Fiber)

	if err := App.ShutdownWithContext(ctx); err != nil {
		log.CannotBeStopped(Server.Fiber)
		return errors.BubbleUp(err, "Down")
	}

	log.Stopped(Server.Fiber)

	return nil
}
