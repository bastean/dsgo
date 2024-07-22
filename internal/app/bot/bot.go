package bot

import (
	"github.com/bastean/dsgo/internal/app/bot/command"
	"github.com/bastean/dsgo/internal/app/bot/handler"
	"github.com/bastean/dsgo/internal/pkg/service/errors"
	"github.com/bastean/dsgo/internal/pkg/service/logger/log"
	"github.com/bwmarrin/discordgo"
)

var (
	Bot = &struct {
		Discord string
	}{
		Discord: log.Bot("discord"),
	}
)

var (
	err     error
	Session *discordgo.Session
)

func Run(app, token, guild string) error {
	log.Starting(Bot.Discord)

	Session, err = discordgo.New("Bot " + token)

	if err != nil {
		return errors.BubbleUp(err, "Run")
	}

	_, err = Session.ApplicationCommandBulkOverwrite(app, guild, command.Commands)

	if err != nil {
		return errors.BubbleUp(err, "Run")
	}

	handler.Events(Session)

	if err := Session.Open(); err != nil {
		return errors.BubbleUp(err, "Run")
	}

	log.Started(Bot.Discord)

	return nil
}

func Stop() error {
	log.Stopping(Bot.Discord)

	if err := Session.Close(); err != nil {
		return errors.BubbleUp(err, "Stop")
	}

	log.Stopped(Bot.Discord)

	return nil
}
