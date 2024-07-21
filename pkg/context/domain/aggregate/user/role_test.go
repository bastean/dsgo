package user_test

import (
	"testing"

	"github.com/bastean/dsgo/pkg/context/domain/aggregate/user"
	"github.com/bastean/dsgo/pkg/context/domain/errors"
	"github.com/stretchr/testify/suite"
)

type RoleValueObjectTestSuite struct {
	suite.Suite
}

func (suite *RoleValueObjectTestSuite) SetupTest() {}

func (suite *RoleValueObjectTestSuite) TestWithInvalidValue() {
	value, err := user.RoleWithInvalidValue()

	var actual *errors.ErrInvalidValue

	suite.ErrorAs(err, &actual)

	expected := &errors.ErrInvalidValue{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "NewRole",
		What:  "Role must be only one of these values: Administrator, Moderator, Contributor",
		Why: errors.Meta{
			"Role": value,
		},
	}}

	suite.EqualError(expected, actual.Error())
}

func TestUnitRoleValueObjectSuite(t *testing.T) {
	suite.Run(t, new(RoleValueObjectTestSuite))
}
