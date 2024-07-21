package user_test

import (
	"testing"

	"github.com/bastean/dsgo/pkg/context/domain/aggregate/user"
	"github.com/bastean/dsgo/pkg/context/domain/errors"
	"github.com/stretchr/testify/suite"
)

type NameValueObjectTestSuite struct {
	suite.Suite
}

func (suite *NameValueObjectTestSuite) SetupTest() {}

func (suite *NameValueObjectTestSuite) TestWithInvalidLength() {
	value, err := user.NameWithInvalidLength()

	var actual *errors.ErrInvalidValue

	suite.ErrorAs(err, &actual)

	expected := &errors.ErrInvalidValue{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "NewName",
		What:  "Name must be between " + "2" + " to " + "20" + " characters and be alphanumeric only",
		Why: errors.Meta{
			"Name": value,
		},
	}}

	suite.EqualError(expected, actual.Error())
}

func (suite *NameValueObjectTestSuite) TestWithInvalidAlphanumeric() {
	value, err := user.NameWithInvalidAlphanumeric()

	var actual *errors.ErrInvalidValue

	suite.ErrorAs(err, &actual)

	expected := &errors.ErrInvalidValue{Bubble: &errors.Bubble{
		When:  actual.When,
		Where: "NewName",
		What:  "Name must be between " + "2" + " to " + "20" + " characters and be alphanumeric only",
		Why: errors.Meta{
			"Name": value,
		},
	}}

	suite.EqualError(expected, actual.Error())
}

func TestUnitNameValueObjectSuite(t *testing.T) {
	suite.Run(t, new(NameValueObjectTestSuite))
}
