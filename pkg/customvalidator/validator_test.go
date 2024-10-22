package customvalidator

import (
	"go-boilerplate/pkg/customvalidator/testdata"
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
)

type ValidatorTestSuite struct {
	suite.Suite

	validator             *Validator
	defaultPersonSimple   testdata.PersonSimple
	defaultPersonExtended testdata.PersonExtended
}

func (s *ValidatorTestSuite) SetupTest() {
	s.validator = New()
	s.defaultPersonSimple = testdata.NewDefaultPersonSimple()
	s.defaultPersonExtended = testdata.NewDefaultPersonExtended()
}

func TestValidatorTestSuite(t *testing.T) {
	suite.Run(t, new(ValidatorTestSuite))
}

func (s *ValidatorTestSuite) TestValidateSimple() {
	s.Run("should pass", func() {
		person := s.defaultPersonSimple
		s.Nil(s.validator.Validate(&person))
	})

	s.Run("should fail: functionality", func() {
		s.Run("when value is nil", func() {
			err := s.validator.Validate(nil)
			s.EqualError(err, "value cannot be nil")
		})

		s.Run("when value is empty struct", func() {
			person := testdata.PersonSimple{}
			s.Error(s.validator.Validate(&person))
		})

		s.Run("when value is not a struct pointer", func() {
			value := 1
			s.EqualError(s.validator.Validate(&value), "value must be a struct or a pointer to a struct")
		})
	})

	s.Run("should fail: name", func() {
		person := s.defaultPersonSimple

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
		person := s.defaultPersonSimple

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
		person := s.defaultPersonExtended
		s.Nil(s.validator.Validate(&person))
	})

	s.Run("should fail: address", func() {
		person := s.defaultPersonExtended

		s.Run("when country is 'Unknown'", func() {
			person.Address.Country = "Unknown"
			s.EqualError(s.validator.Validate(&person), "country cannot be 'Unknown'")
		})
	})

	s.Run("should fail: main card", func() {
		mainCardClone := *s.defaultPersonExtended.MainCard
		person := s.defaultPersonExtended
		person.MainCard = &mainCardClone

		s.Run("when card name is 'Unknown'", func() {
			person.MainCard.Name = "Unknown"
			s.EqualError(s.validator.Validate(&person), "card name cannot be 'Unknown'")
		})
	})

	s.Run("should fail: hobby list", func() {
		s.Run("when hobby list is empty", func() {
			person := s.defaultPersonExtended
			person.HobbyList = []testdata.Hobby{}
			s.Error(s.validator.Validate(&person))
		})

		s.Run("when hobby name is 'Everything'", func() {
			person := s.defaultPersonExtended
			person.HobbyList[0].Name = "Everything"
			s.EqualError(s.validator.Validate(&person), "hobby name cannot be 'Everything'")
		})
	})

	s.Run("should fail: hobby list of list", func() {
		s.Run("when hobby list of list is empty", func() {
			person := s.defaultPersonExtended
			person.HobbyListOfList = [][]testdata.Hobby{}
			s.Error(s.validator.Validate(&person))
		})

		s.Run("when hobby name is 'Everything'", func() {
			person := s.defaultPersonExtended
			person.HobbyListOfList[0][0].Name = "Everything"
			s.EqualError(s.validator.Validate(&person), "hobby name cannot be 'Everything'")
		})
	})
}
