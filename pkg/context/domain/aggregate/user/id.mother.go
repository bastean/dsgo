package user

import (
	"github.com/bastean/dsgo/pkg/context/domain/errors"
	"github.com/bastean/dsgo/pkg/context/domain/service"
)

func IdWithValidValue() *Id {
	value, err := NewId(service.Create.UUID())

	if err != nil {
		errors.Panic(err.Error(), "IdWithValidValue")
	}

	return value

}

func IdWithInvalidValue() (string, error) {
	value := "x"

	_, err := NewId(value)

	return value, err
}
