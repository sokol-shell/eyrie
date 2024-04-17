package configuration

import (
	"reflect"
	"sync"
)

type TransientConfiguration[T any] struct {
	implementingType reflect.Type
	constructor      func() T
	resolveSemaphore sync.Mutex
}

func NewTransientConfiguration[T any](implementingType reflect.Type, constructor func() T) *TransientConfiguration[T] {
	return &TransientConfiguration[T]{
		implementingType: implementingType,
		constructor:      constructor,
	}
}

func (tc *TransientConfiguration[T]) GetOrCreateInstance() (any, error) {
	tc.resolveSemaphore.Lock()
	defer tc.resolveSemaphore.Unlock()

	var instance T = tc.constructor()

	return any(instance), nil
}
