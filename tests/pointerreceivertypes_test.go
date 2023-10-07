package tests

type IPointerInterface interface {
	PointerReceiverMethod()
}

type PointerStruct struct{}

func NewPointerStruct() *PointerStruct {
	return &PointerStruct{}
}

func (ps *PointerStruct) PointerReceiverMethod() {
	panic("Not implemented.")
}
