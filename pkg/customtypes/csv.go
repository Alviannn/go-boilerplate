package customtypes

import (
	"encoding"
	"encoding/json"
	"strings"

	"github.com/labstack/echo/v4"
)

const (
	DefaultComma = ","
)

type CSV struct {
	json.Unmarshaler
	encoding.TextUnmarshaler
	echo.BindUnmarshaler

	Source string
	Comma  string
}

func (ct *CSV) ToSlice() []string {
	currentComma := ct.Comma
	if currentComma == "" {
		currentComma = DefaultComma
	}

	splitList := strings.Split(ct.Source, currentComma)
	if len(splitList) == 1 && splitList[0] == "" {
		return []string{}
	}
	return splitList
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (ct *CSV) UnmarshalJSON(buf []byte) error {
	ct.Source = string(buf)
	return nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (ct *CSV) UnmarshalText(text []byte) error {
	ct.Source = string(text)
	return nil
}

// UnmarshalParam implements the echo.BindUnmarshaler interface.
func (ct *CSV) UnmarshalParam(param string) error {
	ct.Source = param
	return nil
}
