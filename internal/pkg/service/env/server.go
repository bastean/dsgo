package env

import (
	"os"
)

var (
	ServerFiberURL  = os.Getenv("DSGO_SERVER_FIBER_URL")
	ServerFiberPort = os.Getenv("DSGO_SERVER_FIBER_PORT")
)

func HasServerFiberProxy() (string, bool) {
	proxy := os.Getenv("DSGO_DEV_AIR_PROXY_PORT")

	if proxy != "" && proxy != ServerFiberPort {
		return proxy, true
	}

	return "", false
}
