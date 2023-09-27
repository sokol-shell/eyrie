package configuration

import "reflect"

type Configuration interface {
	GetInterfaceType() reflect.Type
	GetLifestyle() Lifestyle
}
