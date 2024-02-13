package di

import (
	"fmt"
	"reflect"
)

type Service struct {
	reflect.Type
}

func NewService(Type reflect.Type) *Service {
	return &Service{
		Type: Type,
	}
}

func (s *Service) Id() string {
	return fmt.Sprintf("%s/%s", s.PkgPath(), s.Name())
}
