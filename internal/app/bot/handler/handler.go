package handler

import (
	"github.com/bwmarrin/discordgo"
)

func Events(session *discordgo.Session) {
	session.AddHandler(Ready)
	session.AddHandler(InteractionCreate)
}
