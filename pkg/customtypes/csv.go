package customtypes

import (
	"encoding/json"
	"strings"
)

const (
	DefaultComma = ","
)

type CSV struct {
	Source string
	Comma  string
}

func NewCSV(source string) *CSV {
	v := &CSV{}
	v.fillDefaults().WithSource(source)
	return v
}

func (ct *CSV) fillDefaults() *CSV {
	if ct.Comma == "" {
		ct.Comma = DefaultComma
	}
	return ct
}

func (ct *CSV) WithComma(comma string) *CSV {
	ct.Comma = comma
	return ct
}

func (ct *CSV) WithSource(source string) *CSV {
	ct.Source = source
	return ct
}

func (ct CSV) ToSlice() (strList []string) {
	if ct.Source == "" {
		return
	}

	strList = strings.Split(ct.Source, ct.Comma)
	return
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (ct *CSV) UnmarshalJSON(buf []byte) error {
	var strVal string
	if err := json.Unmarshal(buf, &strVal); err != nil {
		return err
	}

	ct.fillDefaults().WithSource(strVal)
	return nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
func (ct *CSV) UnmarshalText(text []byte) error {
	ct.fillDefaults().WithSource(string(text))
	return nil
}

// UnmarshalParam implements the echo.BindUnmarshaler interface.
func (ct *CSV) UnmarshalParam(param string) error {
	ct.fillDefaults().WithSource(param)
	return nil
}
