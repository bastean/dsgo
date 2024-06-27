package verify

import (
	"github.com/bastean/dsgo/pkg/context/user/domain/valueobj"
)

func RandomCommand() *Command {
	id := valueobj.IdWithValidValue()

	return &Command{
		Id: id.Value(),
	}
}
