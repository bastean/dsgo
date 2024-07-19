package env

import (
	"os"
)

type Discord struct {
	AppId, PublicKey, BotToken, TestGuildId string
}

var Bot = &struct {
	*Discord
}{
	Discord: &Discord{
		AppId:       os.Getenv("DSGO_BOT_DISCORD_APP_ID"),
		PublicKey:   os.Getenv("DSGO_BOT_DISCORD_APP_PUBLIC_KEY"),
		BotToken:    os.Getenv("DSGO_BOT_DISCORD_APP_TOKEN"),
		TestGuildId: os.Getenv("DSGO_BOT_DISCORD_TEST_GUILD_ID"),
	},
}
