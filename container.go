package falconsnest

import (
	"fmt"
	"reflect"
)

type Container interface {
}

type container struct {
}

var c container

func GetContainer() Container {
	return c
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
		var msg = fmt.Sprintf("%s does not implement %s.", st.Name(), it.Name())
		panic(newRegistrationError(msg))
	}

	return newRegistrar[I, S]()
}
