package user

import (
	"strings"

	"github.com/bastean/dsgo/pkg/context/domain/errors"
	"github.com/go-playground/validator/v10"
)

var RoleOneOf = []string{"Administrator", "Moderator", "Contributor"}

type Role struct {
	Value string `validate:"oneof=administrator moderator contributor"`
}

func NewRole(value string) (*Role, error) {
	value = strings.TrimSpace(value)

	value = strings.ToLower(value)

	valueObj := &Role{
		Value: value,
	}

	err := validator.New(validator.WithRequiredStructEnabled()).Struct(valueObj)

	if err != nil {
		return nil, errors.NewInvalidValue(&errors.Bubble{
			Where: "NewRole",
			What:  "Role must be only one of these values: " + strings.Join(RoleOneOf, ", "),
			Why: errors.Meta{
				"Role": value,
			},
		})
	}

	return valueObj, nil
}
