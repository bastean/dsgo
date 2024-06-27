package create

import (
	"github.com/bastean/dsgo/pkg/context/user/domain/valueobj"
)

func RandomCommand() *Command {
	id := valueobj.IdWithValidValue()
	email := valueobj.EmailWithValidValue()
	username := valueobj.UsernameWithValidValue()
	password := valueobj.PasswordWithValidValue()

	return &Command{
		Id:       id.Value(),
		Email:    email.Value(),
		Username: username.Value(),
		Password: password.Value(),
	}
}
