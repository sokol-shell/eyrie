//go:build test

package tests

type iinterface interface {
	Method()
}

type implementation struct{}

func newImplementation() implementation {
	return implementation{}
}

func (s implementation) Method() {}
