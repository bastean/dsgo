package user

import (
	"strings"

	"github.com/bastean/dsgo/pkg/context/domain/errors"
	"github.com/go-playground/validator/v10"
)

type Email struct {
	Value string `validate:"email"`
}

func NewEmail(value string) (*Email, error) {
	value = strings.TrimSpace(value)

	valueObj := &Email{
		Value: value,
	}

	err := validator.New(validator.WithRequiredStructEnabled()).Struct(valueObj)

	if err != nil {
		return nil, errors.NewInvalidValue(&errors.Bubble{
			Where: "NewEmail",
			What:  "invalid email format",
			Why: errors.Meta{
				"Email": value,
			},
		})
	}

	return valueObj, nil
}
