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

var Port string

var AppId, BotToken, TestGuildId string

var (
	Services = "services"
	Apps     = "apps"
	Server   = &struct {
		Fiber string
	}{
		Fiber: log.Server("fiber"),
	}
	Bot = &struct {
		Discord string
	}{
		Discord: log.Bot("discord"),
	}
)

func usage() {
	fmt.Printf("Usage: %s [OPTIONS]\n", cli)
	flag.PrintDefaults()
}

func main() {
	log.Info("dsGO v0.0.1-alpha.0")

	flag.StringVar(&Port, "port", env.Server.Fiber.Port, "Fiber Server Port")

	flag.StringVar(&AppId, "app", env.Bot.Discord.AppId, "Discord App Id Token")

	flag.StringVar(&BotToken, "token", env.Bot.Discord.BotToken, "Discord Bot Token")

	flag.StringVar(&TestGuildId, "guild", env.Bot.Discord.TestGuildId, "Discord Test Guild Id")

	flag.Usage = usage

	flag.Parse()

	log.Starting(Services)

	if err := service.Run(); err != nil {
		log.Fatal(err.Error())
	}

	log.Started(Services)

	log.Starting(Apps)

	log.Starting(Server.Fiber)

	go func() {
		if err := server.Run(Port); err != nil {
			log.Fatal(err.Error())
		}
	}()

	log.Started(Server.Fiber)

	log.Info(fmt.Sprintf("%s listening on :%s", Server.Fiber, Port))

	if proxy, ok := env.Server.Fiber.HasProxy(); ok {
		log.Info(fmt.Sprintf("%s proxy listening on :%s", Server.Fiber, proxy))
	}

	log.Starting(Bot.Discord)

	go func() {
		if err := bot.Run(AppId, BotToken, TestGuildId); err != nil {
			log.Fatal(err.Error())
		}
	}()

	log.Started(Bot.Discord)

	log.Started(Apps)

	log.Info("press ctrl+c to exit")

	shutdown := make(chan os.Signal, 1)

	signal.Notify(shutdown, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-shutdown

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	log.Stopping(Apps)

	log.Stopping(Server.Fiber)

	errServer := server.Stop(ctx)

	log.Stopped(Server.Fiber)

	log.Stopping(Bot.Discord)

	errBot := bot.Stop()

	log.Stopped(Bot.Discord)

	log.Stopped(Apps)

	log.Stopping(Services)

	errService := service.Stop()

	log.Stopped(Services)

	if err := errors.Join(errServer, errBot, errService); err != nil {
		log.Error(err.Error())
	}

	<-ctx.Done()

	log.Info("exiting...")
}
