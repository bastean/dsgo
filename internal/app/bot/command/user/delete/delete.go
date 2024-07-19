package delete

import (
	"fmt"

	"github.com/bastean/dsgo/internal/pkg/service/errors"
	"github.com/bastean/dsgo/internal/pkg/service/user"
	"github.com/bwmarrin/discordgo"
)

var SubCommand = &discordgo.ApplicationCommandOption{
	Type:        discordgo.ApplicationCommandOptionSubCommand,
	Name:        "delete",
	Description: "Delete an existing user",
	Options: []*discordgo.ApplicationCommandOption{
		{
			Type:        discordgo.ApplicationCommandOptionString,
			Name:        "name",
			Description: "Name of the user to delete",
			Required:    true,
		},
	},
}

func Run(options []*discordgo.ApplicationCommandInteractionDataOption) (*discordgo.InteractionResponse, error) {
	name := options[0].StringValue()

	err := user.Delete.Run(name)

	if err != nil {
		return nil, errors.BubbleUp(err, "Run")
	}

	return &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("%s deleted", name),
		},
	}, nil
}
