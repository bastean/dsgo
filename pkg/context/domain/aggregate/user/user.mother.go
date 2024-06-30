package user

import (
	"github.com/bastean/dsgo/pkg/context/domain/errors"
)

func Random() *User {
	id := IdWithValidValue()
	email := EmailWithValidValue()
	username := UsernameWithValidValue()

	user, err := New(&Primitive{
		Id:       id.Value,
		Email:    email.Value,
		Username: username.Value,
	})

	if err != nil {
		errors.Panic(err.Error(), "RandomUser")
	}

	return user
}

func RandomPrimitive() *Primitive {
	id := IdWithValidValue()
	email := EmailWithValidValue()
	username := UsernameWithValidValue()

	user, err := New(&Primitive{
		Id:       id.Value,
		Email:    email.Value,
		Username: username.Value,
	})

	if err != nil {
		errors.Panic(err.Error(), "RandomUser")
	}

	return user.ToPrimitives()
}
