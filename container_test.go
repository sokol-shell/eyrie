package falconsnest

import (
	"falconsnest/testdata"
	"testing"
)

func Test_Container(t *testing.T) {
	RegisterType[testdata.ICar, testdata.Car]().AsSingleton()
}
