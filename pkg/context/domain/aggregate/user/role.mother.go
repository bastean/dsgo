package user

import (
	"github.com/bastean/dsgo/pkg/context/domain/errors"
	"github.com/bastean/dsgo/pkg/context/domain/service"
)

func RoleWithValidValue() *Role {
	value, err := NewRole(service.Create.RandomString([]string{"administrator", "moderator", "contributor"}))

	if err != nil {
		errors.Panic(err.Error(), "RoleWithValidValue")
	}

	return value
}

func RoleWithInvalidValue() (string, error) {
	value := "x"

	_, err := NewRole(value)

	return value, err
}
