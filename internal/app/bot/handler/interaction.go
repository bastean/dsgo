package handler

import (
	"github.com/bastean/dsgo/internal/app/bot/command"
	"github.com/bastean/dsgo/internal/app/bot/middleware"
	"github.com/bastean/dsgo/internal/pkg/service/logger/log"
	"github.com/bwmarrin/discordgo"
)

func InteractionCreate(session *discordgo.Session, event *discordgo.InteractionCreate) {
	request := event.ApplicationCommandData()

	response, err := command.Routing[request.Name](request.Options)

	if err != nil {
		response = middleware.ErrorHandler(err)
	}

	err = session.InteractionRespond(event.Interaction, response)

	if err != nil {
		log.Error(err.Error())
	}
}
