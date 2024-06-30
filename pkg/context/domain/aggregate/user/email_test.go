package user_test

import (
	"testing"

	"github.com/bastean/dsgo/pkg/context/domain/aggregate/user"
	"github.com/bastean/dsgo/pkg/context/domain/errors"
	"github.com/stretchr/testify/suite"
)

type EmailValueObjectTestSuite struct {
	suite.Suite
}

func (suite *EmailValueObjectTestSuite) SetupTest() {}

func (suite *EmailValueObjectTestSuite) TestWithInvalidValue() {
	value, err := user.EmailWithInvalidValue()

	var actual *errors.InvalidValue

	suite.ErrorAs(err, &actual)

	expected := &errors.InvalidValue{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "NewEmail",
		What:  "invalid email format",
		Why: errors.Meta{
			"Email": value,
		},
	}}

	suite.EqualError(expected, actual.Error())
}

func TestUnitEmailValueObjectSuite(t *testing.T) {
	suite.Run(t, new(EmailValueObjectTestSuite))
}
