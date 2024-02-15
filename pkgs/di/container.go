package di

import (
	"fmt"
	"log"
	"reflect"
)

type Container struct {
	definitions map[string]*Definition
	values      map[string]any
}

func NewContainer() *Container {
	return &Container{
		definitions: make(map[string]*Definition),
		values:      make(map[string]any),
	}
}

// ProvideConstructor and attach a constructor to a tag
func (c *Container) ProvideConstructor(tag *Tag, constructor any) *Container {
	definition := NewDefinition(tag, constructor, nil)
	c.definitions[tag.Id()] = definition
	return c
}

func (c *Container) Set(tagger Tagger, value any) *Container {
	tag := tagger()
	log.Print("Providing value: ", tag.Id())
	definition := NewDefinition(tag, nil, value)
	c.definitions[tag.Id()] = definition
	return c
}

// Call a function by injecting its arguments
func (c *Container) Call(function any) ([]reflect.Value, error) {
	t := reflect.TypeOf(function)
	v := reflect.ValueOf(function)

	args := make([]reflect.Value, t.NumIn())
	for i := 0; i < t.NumIn(); i++ {
		tag := NewTag(t.In(i).Elem())
		arg, err := c.resolve(tag)
		if err != nil {
			return nil, err
		}

		args[i] = arg
	}

	return v.Call(args), nil
}

func (c *Container) resolve(tag *Tag) (reflect.Value, error) {
	id := tag.Id()

	def, found := c.definitions[id]
	if !found {
		return reflect.ValueOf(nil), fmt.Errorf("di: service %s not found in definitions", id)
	}

	if def.value != nil {
		return reflect.ValueOf(def.value), nil
	}

	output, err := c.Call(def.value)
	if err != nil {
		return reflect.ValueOf(nil), err
	}

	return output[0], nil
}

func Resolve[S any](container *Container) (*S, error) {
	var s [0]S
	tag := NewTag(reflect.TypeOf(s).Elem())

	def, found := container.definitions[tag.Id()]
	if !found {
		return nil, fmt.Errorf("di: defition not found for %s", tag.Id())
	}

	log.Print(def.value)
	if def.value != nil {
		return def.value.(*S), nil
	}

	var service S
	resolved, err := container.resolve(tag)
	if err != nil {
		return nil, err
	}

	makeService := func(sptr any) {
		service := reflect.ValueOf(sptr).Elem()

		service.Set(resolved)
	}

	makeService(&service)

	container.values[tag.Id()] = service

	return &service, nil
}
