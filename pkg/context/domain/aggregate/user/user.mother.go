package user

import (
	"github.com/bastean/dsgo/pkg/context/domain/errors"
)

func Random() *User {
	name := NameWithValidValue()
	role := RoleWithValidValue()

	user, err := New(&Primitive{

		Name: name.Value,
		Role: role.Value,
	})

	if err != nil {
		errors.Panic(err.Error(), "RandomUser")
	}

	return user
}

func RandomPrimitive() *Primitive {
	name := NameWithValidValue()
	role := RoleWithValidValue()

	user, err := New(&Primitive{
		Name: name.Value,
		Role: role.Value,
	})

	if err != nil {
		errors.Panic(err.Error(), "RandomPrimitive")
	}

	return user.ToPrimitives()
}
