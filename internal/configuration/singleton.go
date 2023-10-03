package configuration

import "reflect"

type SingletonConfiguration[T any] struct {
	implementingType reflect.Type
	constructor      T
}

func NewSingletonConfiguration[T any](implementingType reflect.Type, constructor func() T) SingletonConfiguration[T] {
	return SingletonConfiguration[T]{
		implementingType: implementingType,
		constructor:      constructor(),
	}
}

func (sc SingletonConfiguration[T]) GetOrCreateInstance() any {
	panic("Not implemented.")
}
