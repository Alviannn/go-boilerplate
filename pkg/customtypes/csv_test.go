package customtypes_test

import (
	"encoding/json"
	"fmt"
	"go-boilerplate/pkg/customtypes"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/suite"
)

type CSVTestSuite struct {
	suite.Suite
}

func TestCSVTestSuite(t *testing.T) {
	suite.Run(t, new(CSVTestSuite))
}

func (s *CSVTestSuite) TestToSlice() {
	s.Run("Success: src is empty", func() {
		csv := customtypes.NewCSV("")

		actual := csv.ToSlice()

		s.Empty(actual, "slice must be empty")
	})

	s.Run("Success: src not empty with default comma", func() {
		csv := customtypes.NewCSV("1,2,3,4,5")

		actual := csv.ToSlice()
		expected := []string{"1", "2", "3", "4", "5"}

		s.Equal(expected, actual, "slice value must match expected")
	})

	s.Run("Success: src not empty with custom comma", func() {
		csv := customtypes.NewCSV("1;2;3;4;5").WithComma(";")

		actual := csv.ToSlice()
		expected := []string{"1", "2", "3", "4", "5"}

		s.Equal(expected, actual, "slice value must match expected")
	})

	s.Run("Failed: spaces around values should not be trimmed", func() {
		csv := customtypes.NewCSV("1, 2, 3, 4, 5").WithComma(",")

		actual := csv.ToSlice()
		expected := []string{"1", " 2", " 3", " 4", " 5"}

		s.Equal(expected, actual, "slice value must match expected")
	})
}

func (s *CSVTestSuite) TestUnmarshalJSON() {
	s.Run("Success: unmarshal JSON", func() {
		type Person struct {
			Name    string          `json:"name"`
			Hobbies customtypes.CSV `json:"hobbies"`
		}

		var person Person
		err := json.Unmarshal([]byte(`{"name":"John, Jane","hobbies":"reading,swimming"}`), &person)

		s.NoError(err)
		s.Equal("John, Jane", person.Name)
		s.Equal([]string{"reading", "swimming"}, person.Hobbies.ToSlice())
	})
}

func (s *CSVTestSuite) TestUnmarshalText() {
	s.Run("Success: unmarshal text", func() {
		var csv customtypes.CSV
		err := csv.UnmarshalText([]byte("reading,swimming"))

		s.NoError(err)
		s.Equal([]string{"reading", "swimming"}, csv.ToSlice())
	})
}

func (s *CSVTestSuite) TestEchoBind() {
	s.Run("Success: bind JSON", func() {
		type Person struct {
			Name    string          `json:"name"`
			Hobbies customtypes.CSV `json:"hobbies"`
		}

		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(`{"name":"John, Jane","hobbies":"reading,swimming"}`))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		ctx := echo.New().NewContext(req, nil)

		var person Person
		err := ctx.Bind(&person)

		s.NoError(err)
		s.Equal("John, Jane", person.Name)
		s.Equal([]string{"reading", "swimming"}, person.Hobbies.ToSlice())
	})

	s.Run("Success: bind query params", func() {
		type GetPersonRequest struct {
			Name    string          `query:"name"`
			Hobbies customtypes.CSV `query:"hobbies"`
		}

		values := url.Values{}
		values.Set("name", "John")
		values.Set("hobbies", "reading,swimming")

		req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/?%s", values.Encode()), nil)

		ctx := echo.New().NewContext(req, nil)

		var reqDTO GetPersonRequest
		err := ctx.Bind(&reqDTO)

		s.NoError(err)
		s.Equal("John", reqDTO.Name)
		s.Equal([]string{"reading", "swimming"}, reqDTO.Hobbies.ToSlice())
	})
}
