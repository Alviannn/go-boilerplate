package customtypes

import "encoding/json"

type Nullable[T comparable] struct {
	value T

	IsNull    bool
	IsDefined bool
}

func (ct *Nullable[T]) UnmarshalJSON(data []byte) error {
	ct.IsDefined = true
	if string(data) == "null" {
		ct.IsNull = true
		return nil
	}
	return json.Unmarshal(data, &ct.value)
}

func (ct *Nullable[T]) GetPtr() *T {
	if ct.IsNull {
		return nil
	}
	return &ct.value
}
