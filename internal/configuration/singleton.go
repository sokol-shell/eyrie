package configuration

import "reflect"

type SingletonConfiguration struct {
	interfaceType    reflect.Type
	implementingType reflect.Type
}

func NewSingletonConfiguration(interfaceType reflect.Type, implementingType reflect.Type) SingletonConfiguration {
	return SingletonConfiguration{
		interfaceType:    interfaceType,
		implementingType: implementingType,
	}
}

func (sc SingletonConfiguration) GetInterfaceType() reflect.Type {
	return sc.interfaceType
}

func (sc SingletonConfiguration) GetLifestyle() Lifestyle {
	return Singleton
}
