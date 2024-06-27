package user

import (
	"github.com/bastean/dsgo/pkg/cmd/server/service/communication/rabbitmq"
	"github.com/bastean/dsgo/pkg/context/shared/domain/messages"
)

var QueueSendConfirmation = &messages.Queue{
	Name: messages.NewRecipientName(&messages.RecipientNameComponents{
		Service: "user",
		Entity:  "user",
		Action:  "send confirmation",
		Event:   "created",
		Status:  "succeeded",
	}),
	Bindings: []string{rabbitmq.Binding.Event.CreatedSucceeded},
}