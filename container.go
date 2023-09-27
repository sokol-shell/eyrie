package falconsnest

import (
	"falconsnest/internal/configuration"
	"fmt"
	"reflect"
)

type Container interface {
	addConfiguration(typ reflect.Type, config configuration.Configuration)
}

type container struct {
	configurations map[string]configuration.Configuration
	initialized    bool
}

var c container

func GetContainer() Container {
	if !c.initialized {
		c = container{
			configurations: make(map[string]configuration.Configuration),
			initialized:    true,
		}
	}

	return c
}

func (c container) addConfiguration(typ reflect.Type, config configuration.Configuration) {
	key := typ.PkgPath() + "/" + typ.Name()
	c.configurations[key] = config
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

	return newRegistrar[I, S](GetContainer(), it, st)
}
