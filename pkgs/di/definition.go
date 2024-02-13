package di

type Definition struct {
	service     *Service
	constructor *Constructor
}

func NewDefinition(service *Service, constructor *Constructor) *Definition {
	return &Definition{
		service:     service,
		constructor: constructor,
	}
}
