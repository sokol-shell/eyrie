//go:build test

package tests

type IInterface1 interface {
	MethodA()
	MethodB()
}

type EmbeddedImplementation struct{}

func (i EmbeddedImplementation) MethodA() {}

type FullImplementation struct {
	EmbeddedImplementation
}

func NewFullImplementation() FullImplementation {
	return FullImplementation{}
}

func (i FullImplementation) MethodB() {}
