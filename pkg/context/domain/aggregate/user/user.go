package user

import (
	"github.com/bastean/dsgo/pkg/context/domain/errors"
)

type User struct {
	*Id
	*Email
	*Username
	*Verified
}

type Primitive struct {
	Id, Email, Username string
	Verified            bool
}

func create(primitive *Primitive) (*User, error) {
	idVO, errId := NewId(primitive.Id)
	emailVO, errEmail := NewEmail(primitive.Email)
	usernameVO, errUsername := NewUsername(primitive.Username)
	verifiedVO, errVerified := NewVerified(primitive.Verified)

	err := errors.Join(errId, errEmail, errUsername, errVerified)

	if err != nil {
		return nil, errors.BubbleUp(err, "create")
	}

	return &User{
		Id:       idVO,
		Email:    emailVO,
		Username: usernameVO,
		Verified: verifiedVO,
	}, nil
}

func (user *User) ToPrimitives() *Primitive {
	return &Primitive{
		Id:       user.Id.Value,
		Email:    user.Email.Value,
		Username: user.Username.Value,
		Verified: user.Verified.Value,
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
	primitive.Verified = false

	user, err := create(primitive)

	if err != nil {
		return nil, errors.BubbleUp(err, "NewUser")
	}

	return user, nil
}
