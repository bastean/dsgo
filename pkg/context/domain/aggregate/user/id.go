package user

import (
	"strings"

	"github.com/bastean/dsgo/pkg/context/domain/errors"
	"github.com/go-playground/validator/v10"
)

type Id struct {
	Value string `validate:"uuid4"`
}

func NewId(value string) (*Id, error) {
	value = strings.TrimSpace(value)

	valueObj := &Id{
		Value: value,
	}

	err := validator.New(validator.WithRequiredStructEnabled()).Struct(valueObj)

	if err != nil {
		return nil, errors.NewInvalidValue(&errors.Bubble{
			Where: "NewId",
			What:  "invalid uuid4 format",
			Why: errors.Meta{
				"Id": value,
			},
		})
	}

	return valueObj, nil
}
