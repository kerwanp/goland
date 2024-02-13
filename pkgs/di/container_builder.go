package di

import (
	"log"
	"reflect"
)

type ContainerBuilder struct {
	definitions map[string]*Definition
	services    map[string]any
}

func NewContainer() *ContainerBuilder {
	return &ContainerBuilder{
		definitions: make(map[string]*Definition),
		services:    make(map[string]any),
	}
}

func (builder *ContainerBuilder) resolve(id string) []reflect.Value {
	log.Printf("Resolving %s", id)

	def := builder.definitions[id]
	args := make([]reflect.Value, def.constructor.Type.NumIn())

	for i := 0; i < def.constructor.Type.NumIn(); i++ {
		argType := def.constructor.Type.In(i)
		argService := NewService(argType)
		arg := builder.resolve(argService.Id())

		args[i] = arg[0]
	}

	return def.constructor.Value.Call(args)
}

func Register[S any](builder *ContainerBuilder, constructor any) *Definition {
	var s [0]S
	definition := NewDefinition(
		NewService(reflect.TypeOf(s).Elem()),
		NewConstructor(constructor),
	)

	builder.definitions[definition.service.Id()] = definition

	return definition
}

func Resolve[S any](container *ContainerBuilder) *S {
	var s [0]S
	serviceType := NewService(reflect.TypeOf(s).Elem())

	if service := container.services[serviceType.Id()]; service != nil {
		return service.(*S)
	}

	var service S
	resolved := container.resolve(serviceType.Id())[0]

	makeService := func(sptr any) {
		service := reflect.ValueOf(sptr).Elem()

		service.Set(resolved)
	}

	makeService(&service)

	container.services[serviceType.Id()] = service

	return &service
}
