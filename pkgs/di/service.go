package di

import (
	"fmt"
	"reflect"
)

type Tag struct {
	reflect.Type
}

func NewTag(Type reflect.Type) *Tag {
	return &Tag{
		Type: Type,
	}
}

func (s *Tag) Id() string {
	return fmt.Sprintf("%s/%s", s.PkgPath(), s.Name())
}

func T[I any]() *Tag {
	var t [0]I
	return NewTag(reflect.TypeOf(t).Elem())
}

type Tagger func() *Tag
