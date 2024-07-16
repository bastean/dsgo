package bot

import (
	"fmt"

	"github.com/bastean/dsgo/internal/pkg/service/errors"
	"github.com/bastean/dsgo/internal/pkg/service/logger/log"
	"github.com/bwmarrin/discordgo"
)

var Discord *discordgo.Session

func Run(app, token, guild string) error {
	Discord, err := discordgo.New("Bot " + token)

	if err != nil {
		return errors.BubbleUp(err, "Run")
	}

	Discord.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		log.Info(fmt.Sprintf("bot logged in as: %s#%s", s.State.User.Username, s.State.User.Discriminator))
	})

	Discord.ApplicationCommandBulkOverwrite(app, guild,
		[]*discordgo.ApplicationCommand{
			{
				Name:        "hello",
				Description: "Print Hello World",
			},
		},
	)

	Discord.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		data := i.ApplicationCommandData()

		switch data.Name {
		case "hello":
			err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: "Hello, World!",
				},
			})

			if err != nil {
				log.Error(err.Error())
			}
		}
	})

	if err := Discord.Open(); err != nil {
		return errors.BubbleUp(err, "Run")
	}

	return nil
}

func Stop() error {
	if err := Discord.Close(); err != nil {
		return errors.BubbleUp(err, "Stop")
	}

	return nil
}
