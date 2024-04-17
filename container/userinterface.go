package container

import (
	"fmt"
	"reflect"

	"github.com/ivan-ivkovic/eyrie/internal/configuration"
)

func Register[I any, S any](constructor func() S) Registrar[I, S] {
	var i [0]I
	var s [0]S
	var it = reflect.TypeOf(i).Elem()
	var st = reflect.TypeOf(s).Elem()
	var ik = it.Kind()
	var sk = st.Kind()

	if ik != reflect.Interface || (sk != reflect.Struct && sk != reflect.Pointer) {
		panic(newRegistrationError("Interface and struct expected as type parameters."))
	}

	if !st.Implements(it) {
		var msg = fmt.Sprintf("%s does not implement %s.", st.Name(), it.Name())
		panic(newRegistrationError(msg))
	}

	return newRegistrar[I, S](getContainer(), constructor, it, st)
}

func Resolve[I any]() I {
	var i [0]I
	var interfaceType = reflect.TypeOf(i).Elem()

	config := getContainer().getConfiguration(interfaceType)
	instance, err := config.GetOrCreateInstance()
	if err != nil {
		switch code := err.(*configuration.ConfigurationError).Code; code {
		case configuration.RecursiveConstruction:
			msg := fmt.Sprintf("Recursive construction occurred while resolving %s.", interfaceType.Name())
			panic(newRecursiveConstructionError(msg))
		default:
			panic("Not reachable.")
		}
	}

	return instance.(I)
}

func Seal() {
	c := getContainer()
	c.seal()
}

func Clear() {
	c = nil
}
