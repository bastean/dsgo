package bot

import (
	"github.com/bastean/dsgo/internal/app/bot/command"
	"github.com/bastean/dsgo/internal/app/bot/handler"
	"github.com/bastean/dsgo/internal/pkg/service/errors"
	"github.com/bwmarrin/discordgo"
)

var (
	err     error
	Session *discordgo.Session
)

func Run(app, token, guild string) error {
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

	return nil
}

func Stop() error {
	if err := Session.Close(); err != nil {
		return errors.BubbleUp(err, "Stop")
	}

	return nil
}
