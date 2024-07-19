package env

import (
	"os"
)

type Security struct {
	AllowedHosts string
}

type Gin struct {
	URL, Port, Mode string
	Security        *Security
}

func (gin *Gin) HasProxy() (string, bool) {
	proxy := os.Getenv("DSGO_DEV_AIR_PROXY_PORT")

	if proxy != "" && proxy != gin.Port {
		return proxy, true
	}

	return "", false
}

var Server = &struct {
	*Gin
}{
	Gin: &Gin{
		URL:  os.Getenv("DSGO_SERVER_GIN_URL"),
		Port: os.Getenv("DSGO_SERVER_GIN_PORT"),
		Mode: os.Getenv("DSGO_SERVER_GIN_MODE"),
		Security: &Security{
			AllowedHosts: os.Getenv("DSGO_SERVER_GIN_ALLOWED_HOSTS"),
		},
	},
}
