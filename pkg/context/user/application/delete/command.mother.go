package delete

import (
	"github.com/bastean/dsgo/pkg/context/user/domain/valueobj"
)

func RandomCommand() *Command {
	id := valueobj.IdWithValidValue()
	password := valueobj.PasswordWithValidValue()

	return &Command{
		Id:       id.Value(),
		Password: password.Value(),
	}
}
