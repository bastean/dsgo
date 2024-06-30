package user

import (
	"github.com/bastean/dsgo/pkg/context/domain/errors"
	"github.com/bastean/dsgo/pkg/context/domain/service"
)

func EmailWithValidValue() *Email {
	value, err := NewEmail(service.Create.Email())

	if err != nil {
		errors.Panic(err.Error(), "EmailWithValidValue")
	}

	return value
}

func EmailWithInvalidValue() (string, error) {
	value := "x"

	_, err := NewEmail(value)

	return value, err
}
