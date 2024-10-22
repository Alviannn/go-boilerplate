package customvalidator

import "reflect"

type helper struct{}

func (helper) isStruct(value any) bool {
	valueRef := reflect.ValueOf(value)
	return reflect.Indirect(valueRef).Kind() == reflect.Struct
}

func (helper) isSliceOrArray(value any) bool {
	valueRef := reflect.ValueOf(value)
	indirectValueRef := reflect.Indirect(valueRef)

	return indirectValueRef.Kind() == reflect.Slice ||
		indirectValueRef.Kind() == reflect.Array
}
