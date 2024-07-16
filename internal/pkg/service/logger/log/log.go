package log

import (
	"fmt"

	"github.com/bastean/dsgo/pkg/context/infrastructure/record/log"
)

var (
	Log   = new(log.Log)
	Debug = Log.Debug
	Error = Log.Error
	Fatal = Log.Fatal
	Info  = Log.Info
)

func Service(service string) string {
	return fmt.Sprintf("service:%s", service)
}

func Module(module string) string {
	return fmt.Sprintf("module:%s", module)
}

func Server(server string) string {
	return fmt.Sprintf("server:%s", server)
}

func Bot(bot string) string {
	return fmt.Sprintf("bot:%s", bot)
}

func Starting(service string) {
	Info(fmt.Sprintf("starting %s...", service))
}

func Started(service string) {
	Info(fmt.Sprintf("%s started", service))
}

func Stopping(service string) {
	Info(fmt.Sprintf("stopping %s...", service))
}

func Stopped(service string) {
	Info(fmt.Sprintf("%s stopped", service))
}

func EstablishingConnectionWith(service string) {
	Info(fmt.Sprintf("establishing connection with %s...", service))
}

func ConnectionEstablishedWith(service string) {
	Info(fmt.Sprintf("connection established with %s", service))
}

func ConnectionFailedWith(service string) {
	Info(fmt.Sprintf("connection failed with %s", service))
}

func ClosingConnectionWith(service string) {
	Info(fmt.Sprintf("closing connection with %s...", service))
}

func ConnectionClosedWith(service string) {
	Info(fmt.Sprintf("connection closed with %s", service))
}
