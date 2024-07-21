package user

import (
	"strings"

	"github.com/bastean/dsgo/pkg/context/domain/errors"
	"github.com/go-playground/validator/v10"
)

const NameMinCharactersLength = "2"
const NameMaxCharactersLength = "20"

type Name struct {
	Value string `validate:"gte=2,lte=20,alphanum"`
}

func NewName(value string) (*Name, error) {
	value = strings.TrimSpace(value)

	valueObj := &Name{
		Value: value,
	}

	err := validator.New(validator.WithRequiredStructEnabled()).Struct(valueObj)

	if err != nil {
		return nil, errors.NewInvalidValue(&errors.Bubble{
			Where: "NewName",
			What:  "Name must be between " + NameMinCharactersLength + " to " + NameMaxCharactersLength + " characters and be alphanumeric only",
			Why: errors.Meta{
				"Name": value,
			},
		})
	}

	return valueObj, nil
}
