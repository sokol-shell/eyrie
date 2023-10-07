package tests

type IInterface interface {
	Method1()
}

func NewIInterface() IInterface {
	return Implementation{}
}

type Implementation struct {
	memoryAddress *int
}

func NewImplementation() Implementation {
	return Implementation{
		memoryAddress: new(int),
	}
}

func (s Implementation) Method1() {
	panic("Not implemented.")
}
