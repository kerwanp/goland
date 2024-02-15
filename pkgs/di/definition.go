package di

type Definition struct {
	tag         *Tag
	constructor any
	value       any
}

func NewDefinition(tag *Tag, constructor any, value any) *Definition {
	return &Definition{
		tag:         tag,
		constructor: constructor,
		value:       value,
	}
}
