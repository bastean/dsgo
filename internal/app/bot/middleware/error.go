package middleware

import (
	"github.com/bastean/dsgo/internal/pkg/service/errors"
	"github.com/bastean/dsgo/internal/pkg/service/logger/log"
	"github.com/bwmarrin/discordgo"
)

func Error(err error) *discordgo.InteractionResponse {
	var errInvalidValue *errors.ErrInvalidValue
	var errAlreadyExist *errors.ErrAlreadyExist
	var errNotExist *errors.ErrNotExist
	var errFailure *errors.ErrFailure
	var errInternal *errors.ErrInternal

	content := ""

	switch {
	case errors.As(err, &errInvalidValue):
		content = errInvalidValue.What
	case errors.As(err, &errAlreadyExist):
		content = errAlreadyExist.What
	case errors.As(err, &errNotExist):
		content = errNotExist.What
	case errors.As(err, &errFailure):
		content = errFailure.What
	case errors.As(err, &errInternal):
		content = "internal bot error"
		fallthrough
	default:
		log.Error(err.Error())
	}

	return &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Flags:   discordgo.MessageFlagsEphemeral,
			Content: content,
		},
	}
}
