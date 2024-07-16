package env

import (
	"os"
)

var MySQL = &struct {
	DSN, Name string
}{
	DSN:  os.Getenv("DATABASE_MYSQL_DSN"),
	Name: os.Getenv("DATABASE_MYSQL_NAME"),
}

var SQLite = &struct {
	Name string
}{
	Name: os.Getenv("DATABASE_SQLITE_NAME"),
}

type security struct {
	AllowedHosts string
}

type server struct {
	URL, Port, Mode string
	Security        *security
}

func (server *server) HasProxy() (string, bool) {
	proxy := os.Getenv("DSGO_DEV_AIR_PROXY_PORT")

	if proxy != "" && proxy != server.Port {
		return proxy, true
	}

	return "", false
}

var Server = &server{
	URL:  os.Getenv("DSGO_SERVER_URL"),
	Port: os.Getenv("DSGO_SERVER_GIN_PORT"),
	Mode: os.Getenv("DSGO_SERVER_GIN_MODE"),
	Security: &security{
		AllowedHosts: os.Getenv("DSGO_SERVER_GIN_ALLOWED_HOSTS"),
	},
}

type discord struct {
	AppId, PublicKey, BotToken, TestGuildId string
}

var Discord = &discord{
	AppId:       os.Getenv("DSGO_BOT_DISCORD_APP_ID"),
	PublicKey:   os.Getenv("DSGO_BOT_DISCORD_APP_PUBLIC_KEY"),
	BotToken:    os.Getenv("DSGO_BOT_DISCORD_APP_TOKEN"),
	TestGuildId: os.Getenv("DSGO_BOT_DISCORD_TEST_GUILD_ID"),
}
