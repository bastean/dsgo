package create

import (
	"fmt"

	"github.com/bastean/dsgo/internal/pkg/service/errors"
	"github.com/bastean/dsgo/internal/pkg/service/user"
	"github.com/bwmarrin/discordgo"
)

var SubCommand = &discordgo.ApplicationCommandOption{
	Type:        discordgo.ApplicationCommandOptionSubCommand,
	Name:        "create",
	Description: "Create",
	Options: []*discordgo.ApplicationCommandOption{
		{
			Type:        discordgo.ApplicationCommandOptionString,
			Name:        "name",
			Description: "Name which must be unique",
			Required:    true,
		},
		{
			Type:        discordgo.ApplicationCommandOptionString,
			Name:        "role",
			Description: "Role which can be: administrator, moderator or contributor",
			Required:    true,
		},
	},
}

func Run(options []*discordgo.ApplicationCommandInteractionDataOption) (*discordgo.InteractionResponse, error) {
	name := options[0].StringValue()
	role := options[1].StringValue()

	err := user.Create.Run(&user.Primitive{
		Name: name,
		Role: role,
	})

	if err != nil {
		return nil, errors.BubbleUp(err, "Run")
	}

	return &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("%s created as %s", name, role),
		},
	}, nil
}
