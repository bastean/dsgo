package read

import (
	"github.com/bastean/dsgo/pkg/context/user/domain/valueobj"
)

func RandomQuery() *Query {
	id := valueobj.IdWithValidValue()

	return &Query{
		Id: id.Value(),
	}
}
