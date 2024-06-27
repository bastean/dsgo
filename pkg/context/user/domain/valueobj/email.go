package valueobj

import (
	"github.com/bastean/dsgo/pkg/context/shared/domain/models"
	"github.com/bastean/dsgo/pkg/context/shared/domain/valueobjs"
)

func NewEmail(email string) (models.ValueObject[string], error) {
	return valueobjs.NewEmail(email)
}
