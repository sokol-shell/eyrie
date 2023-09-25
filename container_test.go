package falconsnest_test

import (
	"falconsnest"
	"falconsnest/testdata"
	"testing"
)

func Test_Container(t *testing.T) {
	falconsnest.Register[testdata.ICar, testdata.Car]().AsSingleton()
}

func Test_CannotRegisterInterfaceToAStruct(t *testing.T) {
	var expectedErrorMessage = "Interface and struct expected as type parameters."

	defer checkPanic(t, expectedErrorMessage)
	falconsnest.Register[testdata.Car, testdata.ICar]().AsSingleton()
}

func Test_CannotRegisterIfStructDoesNotImplementInterface(t *testing.T) {
	var expectedErrorMessage = "Exhaust does not implement ICar."

	defer checkPanic(t, expectedErrorMessage)
	falconsnest.Register[testdata.ICar, testdata.Exhaust]().AsSingleton()
}

func Test_CannotRegisterInterfaceWhichImplementsEmbeddedInterface(t *testing.T) {
	var expectedErrorMessage = "Interface and struct expected as type parameters."

	defer checkPanic(t, expectedErrorMessage)
	falconsnest.Register[testdata.Embeddedinterface1, testdata.Finalinterface1]().AsSingleton()
}

func Test_CannotRegisterInterfaceWhichImplementsEmbeddedInterface2(t *testing.T) {
	var expectedErrorMessage = "Interface and struct expected as type parameters."

	defer checkPanic(t, expectedErrorMessage)
	falconsnest.Register[testdata.Embeddedinterface2, testdata.Finalinterface2]().AsSingleton()
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
