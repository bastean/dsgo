package user

import (
	"github.com/bastean/dsgo/pkg/context/application/user/create"
	"github.com/bastean/dsgo/pkg/context/application/user/delete"
	"github.com/bastean/dsgo/pkg/context/application/user/read"
	"github.com/bastean/dsgo/pkg/context/application/user/update"
	"github.com/bastean/dsgo/pkg/context/domain/repository"
	"github.com/bastean/dsgo/pkg/context/domain/usecase"
)

var (
	Create usecase.Create
	Read   usecase.Read
	Update usecase.Update
	Delete usecase.Delete
)

func Start(repository repository.User) {
	Create = create.New(repository)

	Read = read.New(repository)

	Update = update.New(repository)

	Delete = delete.New(repository)
}
