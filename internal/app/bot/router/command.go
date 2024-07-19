package router

import (
	"github.com/bwmarrin/discordgo"
)

type Command = map[string]func([]*discordgo.ApplicationCommandInteractionDataOption) (*discordgo.InteractionResponse, error)
