//go:build test

package tests

import "time"

type ISlowConstructor interface {
	MyMethod()
}

type SlowConstructor struct {
	memoryAddress *int
}

func NewSlowConstructor() SlowConstructor {
	time.Sleep(10 * time.Millisecond)
	return SlowConstructor{
		memoryAddress: new(int),
	}
}

func (s SlowConstructor) MyMethod() {
	panic("not implemented")
}
