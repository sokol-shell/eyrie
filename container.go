package falconsnest

import (
	"reflect"
)

type Container struct {
}

func NewContainer() Container {
	return Container{}
}

func Register[I any, S any]() Registrar[I, S] {
	var i [0]I
	var s [0]S
	var it = reflect.TypeOf(i).Elem()
	var st = reflect.TypeOf(s).Elem()
	var ik = it.Kind()
	var sk = st.Kind()

	if ik != reflect.Interface || sk != reflect.Struct {
		panic(newRegistrationError("Interface and struct expected as type parameters."))
	}

	if !st.Implements(it) {
		panic(newRegistrationError("Struct type does not implement the interface type."))
	}

	return newRegistrar[I, S]()
}
