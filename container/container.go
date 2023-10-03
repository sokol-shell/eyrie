package container

import (
	"falconsnest/internal/configuration"
	"fmt"
	"reflect"
)

type Container interface {
	addConfiguration(typ reflect.Type, config configuration.Configuration)
	getConfiguration(typ reflect.Type) configuration.Configuration
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
	key := c.generateKey(typ)
	c.configurations[key] = config
}

func (c container) getConfiguration(typ reflect.Type) configuration.Configuration {
	key := c.generateKey(typ)
	config, ok := c.configurations[key]
	if ok {
		return config
	}

	panic(newResolveError(fmt.Sprintf("Could not resolve %s. Not found.", typ.Name())))
}

func (c container) generateKey(typ reflect.Type) string {
	return typ.PkgPath() + "/" + typ.Name()
}

func Register[I any, S any](constructor func() S) Registrar[I, S] {
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

	return newRegistrar[I, S](GetContainer(), constructor, it, st)
}

func Resolve[I any]() I {
	var i [0]I
	var interfaceType = reflect.TypeOf(i).Elem()

	config := GetContainer().getConfiguration(interfaceType)
	instance := config.GetOrCreateInstance()
	return instance.(I)
}
