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

	var actual *errors.InvalidValue

	suite.ErrorAs(err, &actual)

	expected := &errors.InvalidValue{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "NewRole",
		What:  "role must be only one of these values: administrator, moderator, contributor",
		Why: errors.Meta{
			"Role": value,
		},
	}}

	suite.EqualError(expected, actual.Error())
}

func TestUnitRoleValueObjectSuite(t *testing.T) {
	suite.Run(t, new(RoleValueObjectTestSuite))
}
