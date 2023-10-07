package tests

type IInterface interface{}

type Struct struct {
	object *int
}

func NewStruct() Struct {
	return Struct{
		object: new(int),
	}
}
