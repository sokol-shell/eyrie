package container

import (
	"github.com/ivan-ivkovic/falconsnest/internal/configuration"
	"reflect"
)

type Registrar[I any, S any] interface {
	AsSingleton()
	AsTransient()
}

type registrar[I any, S any] struct {
	container        Container
	constructor      func() S
	interfaceType    reflect.Type
	implementingType reflect.Type
}

func newRegistrar[I any, S any](
	c Container,
	constructor func() S,
	interfaceType reflect.Type,
	implementingType reflect.Type,
) Registrar[I, S] {
	return &registrar[I, S]{
		container:        c,
		constructor:      constructor,
		interfaceType:    interfaceType,
		implementingType: implementingType,
	}
}

func (r *registrar[I, S]) AsSingleton() {
	config := configuration.NewSingletonConfiguration[S](r.implementingType, r.constructor)
	r.container.addConfiguration(r.interfaceType, config)
}

func (r *registrar[I, S]) AsTransient() {
	config := configuration.NewTransientConfiguration[S](r.implementingType, r.constructor)
	r.container.addConfiguration(r.interfaceType, config)
}
