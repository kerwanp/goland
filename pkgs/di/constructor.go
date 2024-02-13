package di

import "reflect"

type Constructor struct {
	Type  reflect.Type
	Value reflect.Value
}

func NewConstructor(constructor any) *Constructor {
	return &Constructor{
		Type:  reflect.TypeOf(constructor),
		Value: reflect.ValueOf(constructor),
	}
}
