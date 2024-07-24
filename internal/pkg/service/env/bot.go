package env

import (
	"os"
)

var (
	BotDiscordAppId       = os.Getenv("DSGO_BOT_DISCORD_APP_ID")
	BotDiscordPublicKey   = os.Getenv("DSGO_BOT_DISCORD_APP_PUBLIC_KEY")
	BotDiscordToken       = os.Getenv("DSGO_BOT_DISCORD_APP_TOKEN")
	BotDiscordTestGuildId = os.Getenv("DSGO_BOT_DISCORD_TEST_GUILD_ID")
)
