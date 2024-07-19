package handler

import (
	"fmt"

	"github.com/bastean/dsgo/internal/pkg/service/logger/log"
	"github.com/bwmarrin/discordgo"
)

func Ready(session *discordgo.Session, event *discordgo.Ready) {
	log.Info(fmt.Sprintf("%s logged in as: %s#%s", log.Bot("discord"), session.State.User.Username, session.State.User.Discriminator))
}
