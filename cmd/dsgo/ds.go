package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bastean/dsgo/internal/app/bot"
	"github.com/bastean/dsgo/internal/app/server"
	"github.com/bastean/dsgo/internal/pkg/service"
	"github.com/bastean/dsgo/internal/pkg/service/env"
	"github.com/bastean/dsgo/internal/pkg/service/logger/log"
)

const cli = "dsgo"

var (
	Services = "services"
	Apps     = "apps"
)

func usage() {
	fmt.Printf("Usage: %s [OPTIONS]\n\n", cli)
	fmt.Print("Example of interoperability between a Web App and a Discord Bot using a layered architecture.\n\n")
	flag.PrintDefaults()
}

func main() {
	flag.StringVar(&env.DatabaseSQLiteName, "database", env.DatabaseSQLiteName, "SQLite database file path (default \"In-Memory\")")

	flag.StringVar(&env.ServerFiberPort, "port", env.ServerFiberPort, "Fiber Server Port (optional)")

	flag.StringVar(&env.BotDiscordAppId, "app", env.BotDiscordAppId, "Discord App Id Token (required)")

	flag.StringVar(&env.BotDiscordToken, "token", env.BotDiscordToken, "Discord Bot Token (required)")

	flag.StringVar(&env.BotDiscordTestGuildId, "guild", env.BotDiscordTestGuildId, "Discord Test Guild Id (optional)")

	flag.Usage = usage

	flag.Parse()

	log.Logo()

	log.Starting(Services)

	if err := service.Up(); err != nil {
		log.Fatal(err.Error())
	}

	log.Started(Services)

	log.Starting(Apps)

	go func() {
		if err := server.Up(); err != nil {
			log.Fatal(err.Error())
		}
	}()

	go func() {
		if err := bot.Up(); err != nil {
			log.Fatal(err.Error())
		}
	}()

	log.Started(Apps)

	log.Info("press ctrl+c to exit")

	shutdown := make(chan os.Signal, 1)

	signal.Notify(shutdown, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-shutdown

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	log.Stopping(Apps)

	errServer := server.Down(ctx)

	errBot := bot.Down()

	if err := errors.Join(errServer, errBot); err != nil {
		log.Error(err.Error())
	}

	log.Stopped(Apps)

	log.Stopping(Services)

	errService := service.Down()

	if errService != nil {
		log.Error(errService.Error())
	}

	log.Stopped(Services)

	<-ctx.Done()

	log.Info("exiting...")
}
