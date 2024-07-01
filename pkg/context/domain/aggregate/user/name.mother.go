package user

import (
	"github.com/bastean/dsgo/pkg/context/domain/errors"
	"github.com/bastean/dsgo/pkg/context/domain/service"
)

func NameWithValidValue() *Name {
	value, err := NewName(service.Create.Regex(`^[A-Za-z0-9]{2,20}$`))

	if err != nil {
		errors.Panic(err.Error(), "NameWithValidValue")
	}

	return value
}

func NameWithInvalidLength() (string, error) {
	value := "x"

	_, err := NewName(value)

	return value, err
}

func NameWithInvalidAlphanumeric() (string, error) {
	value := "<></>"

	_, err := NewName(value)

	return value, err
}
