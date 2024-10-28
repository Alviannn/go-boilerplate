package helpers

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type SliceHelperTestSuite struct {
	suite.Suite
}

func (s *SliceHelperTestSuite) TestSliceIsIn() {
	var (
		intList = []int{1, 2, 3}
		strList = []string{"a", "b", "c"}
	)

	s.Run("int (pass): slice contains item", func() {
		s.True(SliceIsIn(intList, 1))
	})
	s.Run("int (fail): slice does not contain item", func() {
		s.False(SliceIsIn(intList, 4))
	})

	s.Run("string (pass): slice contains item", func() {
		s.True(SliceIsIn(strList, "a"))
	})
	s.Run("string (fail): slice does not contain item", func() {
		s.False(SliceIsIn(strList, "d"))
	})
}

func (s *SliceHelperTestSuite) TestSliceIsInWithFunc() {
	type person struct {
		Name string
		Age  int
	}

	personList := []person{
		{
			Name: "John",
			Age:  30,
		},
		{
			Name: "Jane",
			Age:  25,
		},
		{
			Name: "Bob",
			Age:  35,
		},
	}

	s.Run("person (pass): when finding by name", func() {
		toFind := person{
			Name: "John",
			Age:  30,
		}
		s.True(SliceIsInWithFunc(personList, toFind, func(expected person, actual person) bool {
			return expected.Name == actual.Name
		}))
	})

	s.Run("person (fail): when finding by name", func() {
		toFind := person{
			Name: "Dave",
			Age:  25,
		}
		s.False(SliceIsInWithFunc(personList, toFind, func(expected person, actual person) bool {
			return expected.Name == actual.Name
		}))
	})
}

func TestSliceHelperSuite(t *testing.T) {
	suite.Run(t, new(SliceHelperTestSuite))
}
