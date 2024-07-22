package log

import (
	"fmt"
	"strings"

	"github.com/bastean/dsgo/pkg/context/infrastructure/record/log"
	"github.com/common-nighthawk/go-figure"
	"github.com/fatih/color"
)

var (
	Blue    = color.New(color.FgBlue, color.Bold).Sprint
	Cyan    = color.New(color.FgCyan, color.Bold).Sprint
	Green   = color.New(color.FgGreen, color.Bold).Sprint
	Magenta = color.New(color.FgMagenta, color.Bold).Sprint
	Red     = color.New(color.FgRed, color.Bold).Sprint
)

var (
	Log   = new(log.Log)
	Debug = func(message string) { Log.Debug(Cyan(message)) }
	Error = func(message string) { Log.Error(Red(message)) }
	Fatal = func(message string) { Log.Fatal(Red(message)) }
	Info  = func(message string) { Log.Info(Blue(message)) }
)

var (
	Ok = func(message string) { Log.Info(Green(message)) }
)

func Logo() {
	figureDs := figure.NewFigure("ds", "speed", true).Slicify()
	figureGo := figure.NewFigure("GO", "speed", true).Slicify()

	width := 0
	fixedWidth := 0

	for _, line := range figureDs {
		width = len(line)

		if width > fixedWidth {
			fixedWidth = width
		}
	}

	for i, line := range figureDs {
		width = len(line)

		if width < fixedWidth {
			line += strings.Repeat(" ", (fixedWidth - width))
		}

		fmt.Println(Magenta(line), Cyan(figureGo[i]))
	}

	fmt.Println()
}

func Service(service string) string {
	return fmt.Sprintf("service:%s", service)
}

func Module(module string) string {
	return fmt.Sprintf("module:%s", module)
}

func Server(app string) string {
	return fmt.Sprintf("server:%s", app)
}

func Bot(app string) string {
	return fmt.Sprintf("bot:%s", app)
}

func Starting(service string) {
	Info(fmt.Sprintf("starting %s...", service))
}

func Started(service string) {
	Ok(fmt.Sprintf("%s started", service))
}

func CannotBeStarted(service string) {
	Error(fmt.Sprintf("%s cannot be started", service))
}

func Stopping(service string) {
	Info(fmt.Sprintf("stopping %s...", service))
}

func Stopped(service string) {
	Ok(fmt.Sprintf("%s stopped", service))
}

func CannotBeStopped(service string) {
	Error(fmt.Sprintf("%s cannot be stopped", service))
}

func EstablishingConnectionWith(service string) {
	Info(fmt.Sprintf("establishing connection with %s...", service))
}

func ConnectionEstablishedWith(service string) {
	Ok(fmt.Sprintf("connection established with %s", service))
}

func ConnectionFailedWith(service string) {
	Error(fmt.Sprintf("connection failed with %s", service))
}

func ClosingConnectionWith(service string) {
	Info(fmt.Sprintf("closing connection with %s...", service))
}

func ConnectionClosedWith(service string) {
	Ok(fmt.Sprintf("connection closed with %s", service))
}

func DisconnectionFailedWith(service string) {
	Error(fmt.Sprintf("disconnection failed with %s", service))
}
