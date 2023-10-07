package configuration

import "reflect"

type TransientConfiguration[T any] struct {
	implementingType reflect.Type
	constructor      func() T
	resolving        bool
}

func NewTransientConfiguration[T any](implementingType reflect.Type, constructor func() T) *TransientConfiguration[T] {
	return &TransientConfiguration[T]{
		implementingType: implementingType,
		constructor:      constructor,
	}
}

func (tc *TransientConfiguration[T]) GetOrCreateInstance() (any, error) {
	if tc.resolving {
		return nil, newConfigurationError(RecursiveConstruction)
	}
	tc.resolving = true
	defer func() { tc.resolving = false }()

	var instance T = tc.constructor()

	return any(instance), nil
}
