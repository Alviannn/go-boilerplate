package helpers

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type MapperHelperTestSuite struct {
	suite.Suite
}

func (s *MapperHelperTestSuite) TestMapperSlice() {
	s.Run("when mapping slice: int64 to string", func() {
		numList := []int64{1, 2, 3, 4, 5}

		expected := []string{"1", "2", "3", "4", "5"}
		actual := MapperSlice(numList, func(v int64) string {
			return fmt.Sprint(v)
		})

		s.Equal(expected, actual)
	})

	s.Run("when mapping slice: get person names", func() {
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

		expected := []string{"John", "Jane", "Bob"}
		actual := MapperSlice(personList, func(v person) string {
			return v.Name
		})

		s.Equal(expected, actual)
	})
}

func TestMapperHelperTestSuite(t *testing.T) {
	suite.Run(t, new(MapperHelperTestSuite))
}
