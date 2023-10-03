package configuration

import "reflect"

type TransientConfiguration[T any] struct {
	implementingType reflect.Type
	constructor      func() T
}

func NewTransientConfiguration[T any](implementingType reflect.Type, constructor func() T) TransientConfiguration[T] {
	return TransientConfiguration[T]{
		implementingType: implementingType,
		constructor:      constructor,
	}
}

func (tc TransientConfiguration[T]) GetOrCreateInstance() any {
	var instance T = tc.constructor()
	return any(instance)
}
