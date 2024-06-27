package login

import (
	"github.com/bastean/dsgo/pkg/context/user/domain/valueobj"
)

func RandomQuery() *Query {
	email := valueobj.EmailWithValidValue()
	password := valueobj.PasswordWithValidValue()

	return &Query{
		Email:    email.Value(),
		Password: password.Value(),
	}
}
