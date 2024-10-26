package customvalidator

import (
	"go-boilerplate/pkg/customvalidator/testdata"
	"strings"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/suite"
)

type ValidatorTestSuite struct {
	suite.Suite

	validator *Validator
}

func (s *ValidatorTestSuite) SetupTest() {
	s.validator = New()
}

func TestValidatorTestSuite(t *testing.T) {
	suite.Run(t, new(ValidatorTestSuite))
}

func (s *ValidatorTestSuite) TestValidateSimple() {
	s.Run("should pass", func() {
		person := testdata.NewDefaultPersonSimple()
		s.Nil(s.validator.Validate(&person))
	})

	s.Run("should fail: functionality", func() {
		s.Run("when value is empty struct", func() {
			person := testdata.PersonSimple{}
			s.Error(s.validator.Validate(&person))
		})

		s.Run("when value is nil", func() {
			err := s.validator.Validate(nil)
			s.Error(err)

			_, ok := err.(*validator.InvalidValidationError)
			s.True(ok)
		})

		s.Run("when value is not a struct pointer", func() {
			value := 1
			err := s.validator.Validate(&value)
			s.Error(err)

			_, ok := err.(*validator.InvalidValidationError)
			s.True(ok)
		})
	})

	s.Run("should fail: name", func() {
		person := testdata.NewDefaultPersonSimple()

		s.Run("when name is empty", func() {
			person.Name = ""
			s.Error(s.validator.Validate(&person))
		})

		s.Run("when name is too short", func() {
			person.Name = "Jo"
			s.Error(s.validator.Validate(&person))
		})

		s.Run("when name is too long", func() {
			person.Name = strings.Repeat("A", 101)
			s.Error(s.validator.Validate(&person))
		})

		s.Run("when name contains 'Foo'", func() {
			person.Name = "FooBar"
			s.Error(s.validator.Validate(&person))
		})
	})

	s.Run("should fail: age", func() {
		person := testdata.NewDefaultPersonSimple()

		s.Run("when age is negative", func() {
			person.Age = -1
			s.Error(s.validator.Validate(person))
		})

		s.Run("when age is greater than 130", func() {
			person.Age = 131
			s.Error(s.validator.Validate(person))
		})

		s.Run("when age is greater than 130: has custom error message", func() {
			person.Age = 131
			s.EqualError(s.validator.Validate(&person), "Age must be less than or equal to 130")
		})
	})
}

func (s *ValidatorTestSuite) TestValidateExtended() {
	s.Run("should pass", func() {
		person := testdata.NewDefaultPersonExtended()
		s.Nil(s.validator.Validate(&person))
	})

	s.Run("should fail: address", func() {
		person := testdata.NewDefaultPersonExtended()

		s.Run("when country is 'Unknown'", func() {
			person.Address.Country = "Unknown"
			s.EqualError(s.validator.Validate(&person), "country cannot be 'Unknown'")
		})
	})

	s.Run("should fail: main card", func() {
		s.Run("when card number is nil", func() {
			person := testdata.NewDefaultPersonExtended()
			person.MainCard = nil
			s.Error(s.validator.Validate(&person))
		})

		s.Run("when card name is 'Unknown'", func() {
			person := testdata.NewDefaultPersonExtended()
			person.MainCard.Name = "Unknown"
			s.EqualError(s.validator.Validate(&person), "card name cannot be 'Unknown'")
		})
	})

	s.Run("should fail: hobby list", func() {
		s.Run("when hobby list is empty", func() {
			person := testdata.NewDefaultPersonExtended()
			person.HobbyList = []testdata.Hobby{}
			s.Error(s.validator.Validate(&person))
		})

		s.Run("when hobby name is 'Everything'", func() {
			person := testdata.NewDefaultPersonExtended()
			person.HobbyList[0].Name = "Everything"
			s.EqualError(s.validator.Validate(&person), "hobby name cannot be 'Everything'")
		})
	})

	s.Run("should fail: hobby list of list", func() {
		s.Run("when hobby list of list is empty", func() {
			person := testdata.NewDefaultPersonExtended()
			person.HobbyListOfList = [][]testdata.Hobby{}
			s.Error(s.validator.Validate(&person))
		})

		s.Run("when hobby name is 'Everything'", func() {
			person := testdata.NewDefaultPersonExtended()
			person.HobbyListOfList[0][0].Name = "Everything"
			s.EqualError(s.validator.Validate(&person), "hobby name cannot be 'Everything'")
		})
	})

	s.Run("should fail: nick name list", func() {
		person := testdata.NewDefaultPersonExtended()

		s.Run("when nick name list is empty", func() {
			person.NickNameList = []string{}
			s.Error(s.validator.Validate(&person))
		})
	})
}
