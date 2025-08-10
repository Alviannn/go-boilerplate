package helpers

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type SliceHelperTestSuite struct {
	suite.Suite
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
