package user

import (
	"github.com/bastean/dsgo/pkg/context/domain/errors"
)

type User struct {
	*Name
	*Role
}

type Primitive struct {
	Name, Role string
}

func create(primitive *Primitive) (*User, error) {
	nameVO, errName := NewName(primitive.Name)
	roleVO, errRole := NewRole(primitive.Role)

	err := errors.Join(errName, errRole)

	if err != nil {
		return nil, errors.BubbleUp(err, "create")
	}

	return &User{
		Name: nameVO,
		Role: roleVO,
	}, nil
}

func (user *User) ToPrimitives() *Primitive {
	return &Primitive{
		Name: user.Name.Value,
		Role: user.Role.Value,
	}
}

func FromPrimitives(primitive *Primitive) (*User, error) {
	user, err := create(primitive)

	if err != nil {
		return nil, errors.BubbleUp(err, "FromPrimitives")
	}

	return user, nil
}

func New(primitive *Primitive) (*User, error) {
	user, err := create(primitive)

	if err != nil {
		return nil, errors.BubbleUp(err, "NewUser")
	}

	return user, nil
}
