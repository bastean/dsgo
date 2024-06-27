package user

import (
	"github.com/bastean/dsgo/pkg/context/shared/domain/messages"
	"github.com/bastean/dsgo/pkg/context/user/application/create"
	"github.com/bastean/dsgo/pkg/context/user/application/delete"
	"github.com/bastean/dsgo/pkg/context/user/application/login"
	"github.com/bastean/dsgo/pkg/context/user/application/read"
	"github.com/bastean/dsgo/pkg/context/user/application/update"
	"github.com/bastean/dsgo/pkg/context/user/application/verify"
	"github.com/bastean/dsgo/pkg/context/user/domain/model"
)

var (
	Create *create.Handler
	Read   *read.Handler
	Update *update.Handler
	Delete *delete.Handler
	Verify *verify.Handler
	Login  *login.Handler
)

func Init(repository model.Repository, broker messages.Broker, hashing model.Hashing) {
	Create = NewCreate(repository, broker)

	Read = NewRead(repository)

	Update = NewUpdate(repository, hashing)

	Delete = NewDelete(repository, hashing)

	Verify = NewVerify(repository)

	Login = NewLogin(repository, hashing)
}
