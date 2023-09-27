package configuration

import "reflect"

type TransientConfiguration struct {
	interfaceType    reflect.Type
	implementingType reflect.Type
}

func NewTransientConfiguration(interfaceType reflect.Type, implementingType reflect.Type) TransientConfiguration {
	return TransientConfiguration{
		interfaceType:    interfaceType,
		implementingType: implementingType,
	}
}

func (tc TransientConfiguration) GetInterfaceType() reflect.Type {
	return tc.interfaceType
}

func (tc TransientConfiguration) GetLifestyle() Lifestyle {
	return Transient
}
