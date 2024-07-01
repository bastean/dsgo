package user

import (
	"strings"

	"github.com/bastean/dsgo/pkg/context/domain/errors"
	"github.com/go-playground/validator/v10"
)

var RoleOneOf = []string{"administrator", "moderator", "contributor"}

type Role struct {
	Value string `validate:"oneof=administrator moderator contributor"`
}

func NewRole(value string) (*Role, error) {
	value = strings.TrimSpace(value)

	valueObj := &Role{
		Value: value,
	}

	err := validator.New(validator.WithRequiredStructEnabled()).Struct(valueObj)

	if err != nil {
		return nil, errors.NewInvalidValue(&errors.Bubble{
			Where: "NewRole",
			What:  "role must be only one of these values: " + strings.Join(RoleOneOf, ", "),
			Why: errors.Meta{
				"Role": value,
			},
		})
	}

	return valueObj, nil
}
