package bot

import (
	"github.com/bastean/dsgo/internal/app/bot/command"
	"github.com/bastean/dsgo/internal/app/bot/handler"
	"github.com/bastean/dsgo/internal/pkg/service/env"
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

func Up() error {
	log.Starting(Bot.Discord)

	Session, err = discordgo.New("Bot " + env.BotDiscordToken)

	if err != nil {
		log.CannotBeStarted(Bot.Discord)
		return errors.BubbleUp(err, "Up")
	}

	_, err = Session.ApplicationCommandBulkOverwrite(env.BotDiscordAppId, env.BotDiscordTestGuildId, command.Commands)

	if err != nil {
		log.CannotBeStarted(Bot.Discord)
		return errors.BubbleUp(err, "Up")
	}

	handler.Events(Session)

	if err := Session.Open(); err != nil {
		log.ConnectionFailedWith(Bot.Discord)
		return errors.BubbleUp(err, "Up")
	}

	log.Started(Bot.Discord)

	return nil
}

func Down() error {
	log.Stopping(Bot.Discord)

	if err := Session.Close(); err != nil {
		log.DisconnectionFailedWith(Bot.Discord)
		return errors.BubbleUp(err, "Down")
	}

	log.Stopped(Bot.Discord)

	return nil
}
