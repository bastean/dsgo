package user

import (
	"github.com/bastean/dsgo/internal/app/bot/command/user/create"
	"github.com/bastean/dsgo/internal/app/bot/command/user/delete"
	"github.com/bastean/dsgo/internal/app/bot/command/user/read"
	"github.com/bastean/dsgo/internal/app/bot/command/user/update"
	"github.com/bastean/dsgo/internal/app/bot/router"
	"github.com/bwmarrin/discordgo"
)

var Command = &discordgo.ApplicationCommand{
	Type:        discordgo.ChatApplicationCommand,
	Name:        "user",
	Description: "Perform CRUD operations",
	Options: []*discordgo.ApplicationCommandOption{
		create.SubCommand,
		read.SubCommand,
		update.SubCommand,
		delete.SubCommand,
	},
}

var Routing = router.Command{
	create.SubCommand.Name: create.Run,
	read.SubCommand.Name:   read.Run,
	update.SubCommand.Name: update.Run,
	delete.SubCommand.Name: delete.Run,
}

func Router(subcommand []*discordgo.ApplicationCommandInteractionDataOption) (*discordgo.InteractionResponse, error) {
	return Routing[subcommand[0].Name](subcommand[0].Options)
}
