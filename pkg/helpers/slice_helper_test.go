package helpers

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type SliceHelperTestSuite struct {
	suite.Suite
}

func (s *SliceHelperTestSuite) TestSliceFind() {
	var (
		intList = []int{1, 2, 3}
		strList = []string{"a", "b", "c"}
	)

	s.Run("int (pass): found item in slice", func() {
		ptr := SliceFind(intList, 1)
		s.NotNil(ptr)
		s.Equal(1, *ptr)
	})
	s.Run("int (fail): slice does not contain item", func() {
		ptr := SliceFind(intList, 4)
		s.Nil(ptr)
	})

	s.Run("string (pass): slice contains item", func() {
		ptr := SliceFind(strList, "b")
		s.NotNil(ptr)
		s.Equal("b", *ptr)
	})
	s.Run("string (fail): slice does not contain item", func() {
		ptr := SliceFind(strList, "d")
		s.Nil(ptr)
	})
}

func (s *SliceHelperTestSuite) TestSliceFindFunc() {
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
		nameToFind := "John"
		ptr := SliceFindFunc(personList, func(current person) bool {
			return current.Name == nameToFind
		})
		s.NotNil(ptr)
		s.Equal(nameToFind, ptr.Name)
	})

	s.Run("person (fail): when finding by name", func() {
		nameToFind := "Dave"
		ptr := SliceFindFunc(personList, func(current person) bool {
			return current.Name == nameToFind
		})
		s.Nil(ptr)
	})
}

func (s *SliceHelperTestSuite) TestSliceMap() {
	type person struct {
		Name string
		Age  int
	}

	s.Run("map (pass): when mapping int64 -> string", func() {
		numList := []int64{1, 2, 3, 4, 5}

		expected := []string{"1", "2", "3", "4", "5"}
		actual := SliceMap(numList, func(v int64) string {
			return fmt.Sprint(v)
		})

		s.Equal(expected, actual)
	})

	s.Run("map (pass): when mapping person -> string", func() {
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

		expected := []string{"John", "Jane", "Bob"}
		actual := SliceMap(personList, func(v person) string {
			return v.Name
		})

		s.Equal(expected, actual)
	})
}

func TestSliceHelperSuite(t *testing.T) {
	suite.Run(t, new(SliceHelperTestSuite))
}
