package configuration

import (
	"reflect"
	"sync"
)

type SingletonConfiguration[T any] struct {
	implementingType reflect.Type
	constructor      func() T
	instance         *T
	resolveSemaphore sync.Mutex
}

func NewSingletonConfiguration[T any](implementingType reflect.Type, constructor func() T) *SingletonConfiguration[T] {
	return &SingletonConfiguration[T]{
		implementingType: implementingType,
		constructor:      constructor,
	}
}

func (sc *SingletonConfiguration[T]) GetOrCreateInstance() (any, error) {
	if sc.instance == nil {
		sc.resolveSemaphore.Lock()
		defer sc.resolveSemaphore.Unlock()

		if sc.instance == nil {
			var instance T = sc.constructor()
			sc.instance = &instance
		}
	}

	return any(*sc.instance), nil
}
