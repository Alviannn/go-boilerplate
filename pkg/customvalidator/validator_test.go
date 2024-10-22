package customvalidator

import (
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/suite"
)

type (
	PersonSimple struct {
		Name string `json:"name" validate:"required,min=3,max=100"`
		Age  int    `json:"age" validate:"gte=0,lte=130"`
	}
)

func (m *PersonSimple) Validate() (err error) {
	if strings.Contains(m.Name, "Foo") {
		err = errors.New("name cannot contain 'Foo'")
		return
	}
	return
}

func (m *PersonSimple) ChangeValidationMessage(fieldErr validator.FieldError) (errorMessage string) {
	structFieldName := fieldErr.StructField()
	failedTag := fieldErr.Tag()

	if structFieldName == "Age" && failedTag == "lte" {
		errorMessage = fmt.Sprintf(
			"%s must be less than or equal to %s",
			structFieldName,
			fieldErr.Param(),
		)
	}
	return
}

type ValidatorTestSuite struct {
	suite.Suite

	validator             *Validator
	defaultPersonSimple   PersonSimple
	defaultPersonExtended PersonExtended
}

func (s *ValidatorTestSuite) SetupTest() {
	s.validator = New()

	s.defaultPersonSimple = PersonSimple{
		Name: "John",
		Age:  30,
	}

	s.defaultPersonExtended = PersonExtended{
		Name: "John",
		Age:  30,

		Address: Address{
			Country: "USA",
			City:    "New York",
		},

		MainCard: &Card{
			Name:   "Visa",
			Number: "123456789",
		},

		HobbyList: []Hobby{
			{
				Name:  "Dancing",
				Score: 2,
			},
			{
				Name:  "Traveling",
				Score: 5,
			},
		},

		HobbyListOfList: [][]Hobby{
			{
				{
					Name:  "Dancing",
					Score: 2,
				},
				{
					Name:  "Traveling",
					Score: 5,
				},
			},
			{
				{
					Name:  "Cooking",
					Score: 3,
				},
				{
					Name:  "Reading",
					Score: 4,
				},
			},
		},
	}
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
			person := PersonSimple{}
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

type (
	PersonExtended struct {
		Name string `validate:"required,min=3,max=100"`
		Age  int    `validate:"gte=0,lte=130"`

		Address         Address   `validate:"required"`
		MainCard        *Card     `validate:"required"`
		HobbyList       []Hobby   `validate:"required,min=2,dive"`
		HobbyListOfList [][]Hobby `validate:"required,min=2,dive"`
	}

	Address struct {
		Country string `validate:"required,min=3"`
		City    string `validate:"required,min=3"`
	}

	Card struct {
		Name   string `validate:"required,min=3,max=100"`
		Number string `validate:"required,min=3,max=100"`
	}

	Hobby struct {
		Name  string `validate:"required,min=3,max=100"`
		Score int    `validate:"gte=0,lte=5"`
	}
)

func (m *Address) Validate() (err error) {
	if m.Country == "Unknown" {
		err = errors.New("country cannot be 'Unknown'")
		return
	}
	return
}

func (m *Hobby) Validate() (err error) {
	if m.Name == "Everything" {
		err = errors.New("hobby name cannot be 'Everything'")
		return
	}
	return
}

func (m *Card) Validate() (err error) {
	if m.Name == "Unknown" {
		err = errors.New("card name cannot be 'Unknown'")
		return
	}
	return
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
			person.HobbyList = []Hobby{}
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
			person.HobbyListOfList = [][]Hobby{}
			s.Error(s.validator.Validate(&person))
		})

		s.Run("when hobby name is 'Everything'", func() {
			person := s.defaultPersonExtended
			person.HobbyListOfList[0][0].Name = "Everything"
			s.EqualError(s.validator.Validate(&person), "hobby name cannot be 'Everything'")
		})
	})
}
