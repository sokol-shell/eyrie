package tests

type EmbeddedInterface1 interface {
	Method1()
}

type FinalInterface1 interface {
	EmbeddedInterface1
	Method2()
}

func NewFinalInterface1() FinalInterface1 {
	panic("Not implemented.")
}

type EmbeddedInterface2 interface {
	Method3()
}

type FinalInterface2 interface {
	Method3()
	Method4()
}

func NewFinalInterface2() FinalInterface2 {
	panic("Not implemented.")
}
