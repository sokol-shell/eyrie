package falconsnest_test

import (
	"falconsnest"
	"falconsnest/testdata"
	"testing"
)

func Test_Container(t *testing.T) {
	falconsnest.Register[testdata.ICar, testdata.Car](testdata.NewCar).AsTransient()
	falconsnest.Register[testdata.IEngine, testdata.Engine](testdata.NewEngine).AsTransient()
	falconsnest.Register[testdata.IExhaust, testdata.Exhaust](testdata.NewExhaust).AsTransient()
	falconsnest.Resolve[testdata.IExhaust]()
	falconsnest.Resolve[testdata.IEngine]()
	falconsnest.Resolve[testdata.ICar]()
}

func Test_CannotRegisterInterfaceToAStruct(t *testing.T) {
	var expectedErrorMessage = "Interface and struct expected as type parameters."

	defer checkPanic(t, expectedErrorMessage)
	falconsnest.Register[testdata.Car, testdata.ICar](testdata.NewICar)
}

func Test_CannotRegisterIfStructDoesNotImplementInterface(t *testing.T) {
	var expectedErrorMessage = "Exhaust does not implement ICar."

	defer checkPanic(t, expectedErrorMessage)
	falconsnest.Register[testdata.ICar, testdata.Exhaust](testdata.NewExhaust)
}

func Test_CannotRegisterInterfaceWhichImplementsEmbeddedInterface1(t *testing.T) {
	var expectedErrorMessage = "Interface and struct expected as type parameters."

	defer checkPanic(t, expectedErrorMessage)
	falconsnest.Register[testdata.EmbeddedInterface1, testdata.FinalInterface1](testdata.NewFinalInterface1)
}

func Test_CannotRegisterInterfaceWhichImplementsEmbeddedInterface2(t *testing.T) {
	var expectedErrorMessage = "Interface and struct expected as type parameters."

	defer checkPanic(t, expectedErrorMessage)
	falconsnest.Register[testdata.EmbeddedInterface2, testdata.FinalInterface2](testdata.NewFinalInterface2)
}

func checkPanic(t *testing.T, expectedErrorMessage string) {
	t.Helper()
	if err := recover(); err != nil {
		switch err.(type) {
		case *falconsnest.RegistrationError:
			var actualErrorMessage = err.(*falconsnest.RegistrationError).Msg
			if expectedErrorMessage != actualErrorMessage {
				t.Errorf("expected value: %v\nactual value: %v\n", expectedErrorMessage, actualErrorMessage)
			}
		default:
			t.Errorf("Container did not return an error of the expected type.")
		}
	} else {
		t.Errorf("Container did not panic.")
	}
}
