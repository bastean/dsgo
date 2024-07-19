package command

import (
	"github.com/bastean/dsgo/internal/app/bot/command/user"
	"github.com/bastean/dsgo/internal/app/bot/router"
	"github.com/bwmarrin/discordgo"
)

var Commands = []*discordgo.ApplicationCommand{
	user.Command,
}

var Routing = router.Command{
	user.Command.Name: user.Router,
}
