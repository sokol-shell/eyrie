package falconsnest

import (
	"falconsnest/internal/configuration"
	"reflect"
)

type Registrar[I any, S any] interface {
	AsSingleton()
	AsTransient()
}

type registrar[I any, S any] struct {
	container        Container
	interfaceType    reflect.Type
	implementingType reflect.Type
}

func newRegistrar[I any, S any](c Container, interfaceType reflect.Type, implementingType reflect.Type) Registrar[I, S] {
	return registrar[I, S]{
		container:        c,
		interfaceType:    interfaceType,
		implementingType: implementingType,
	}
}

func (r registrar[I, S]) AsSingleton() {
	config := configuration.NewSingletonConfiguration(r.interfaceType, r.implementingType)
	r.container.addConfiguration(r.interfaceType, config)
}

func (r registrar[I, S]) AsTransient() {
	config := configuration.NewTransientConfiguration(r.interfaceType, r.implementingType)
	r.container.addConfiguration(r.interfaceType, config)
}
