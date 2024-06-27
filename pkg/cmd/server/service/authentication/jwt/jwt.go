package jwt

import (
	"github.com/bastean/dsgo/pkg/cmd/server/service/env"
	"github.com/bastean/dsgo/pkg/context/shared/infrastructure/authentications"
)

type Payload = authentications.Payload

var (
	jwt      = authentications.NewJWT(env.JWT.SecretKey)
	Generate = jwt.Generate
	Validate = jwt.Validate
)
