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
	splitList := strings.Split(ct.Source, ct.GetComma())
	if len(splitList) == 1 && splitList[0] == "" {
		return []string{}
	}
	return splitList
}

func (ct *CSV) GetComma() string {
	if ct.Comma == "" {
		ct.Comma = DefaultComma
	}
	return ct.Comma
}

func (ct *CSV) UnmarshalJSON(buf []byte) error {
	ct.Source = string(buf)
	return nil
}

func (ct *CSV) UnmarshalText(text []byte) error {
	ct.Source = string(text)
	return nil
}

func (ct *CSV) UnmarshalParam(param string) error {
	ct.Source = param
	return nil
}
