package customerror

import (
	"encoding/json"

	"github.com/rs/zerolog"
)

type ErrorJSON struct {
	src                *Error
	Message            string   `json:"message"`
	SourceErrorMessage string   `json:"sourceErrorMessage"`
	StackLine          string   `json:"stackLine,omitempty"`
	Stack              []string `json:"stack,omitempty"`
}

func (ej ErrorJSON) MarshalZerologObject(e *zerolog.Event) {
	var (
		ctx    = ej.src.Context
		rawMap = make(map[string]any)
	)

	if ctx != nil {
		e.Ctx(ctx)
	}

	jsonBuf, _ := json.Marshal(ej)
	_ = json.Unmarshal(jsonBuf, &rawMap)

	for key, value := range rawMap {
		e.Interface(key, value)
	}
}
