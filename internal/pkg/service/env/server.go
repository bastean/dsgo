package env

import (
	"os"
)

type Fiber struct {
	URL, Port string
}

func (fiber *Fiber) HasProxy() (string, bool) {
	proxy := os.Getenv("DSGO_DEV_AIR_PROXY_PORT")

	if proxy != "" && proxy != fiber.Port {
		return proxy, true
	}

	return "", false
}

var Server = &struct {
	*Fiber
}{
	Fiber: &Fiber{
		URL:  os.Getenv("DSGO_SERVER_FIBER_URL"),
		Port: os.Getenv("DSGO_SERVER_FIBER_PORT"),
	},
}
