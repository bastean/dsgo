package valueobj

import (
	"github.com/bastean/dsgo/pkg/context/shared/domain/errors"
	"github.com/bastean/dsgo/pkg/context/shared/domain/models"
	"github.com/bastean/dsgo/pkg/context/shared/domain/services"
)

func IdWithValidValue() models.ValueObject[string] {
	value, err := NewId(services.Create.UUID())

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
