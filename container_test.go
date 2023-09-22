package falconsnest

import (
	"falconsnest/testdata"
	"testing"
)

func Test_Container(t *testing.T) {
	Register[testdata.ICar, testdata.Car]().AsSingleton()
}

func Test_CannotRegisterInterfaceToAStruct(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			var actualErrorMessage = err.(*RegistrationError).Msg
			var expectedErrorMessage = "Interface and struct expected as type parameters."
			if expectedErrorMessage != actualErrorMessage {
				t.Errorf("expected value: %v\nactual value: %v\n", expectedErrorMessage, actualErrorMessage)
			}
		} else {
			t.Errorf("Container did not panic.")
		}
	}()

	Register[testdata.Car, testdata.ICar]().AsSingleton()
}

func Test_CannotRegisterIfStructDoesNotImplementInterface(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			switch err.(type) {
			case *RegistrationError:
				var actualErrorMessage = err.(*RegistrationError).Msg
				var expectedErrorMessage = "Struct type does not implement the interface type."
				if expectedErrorMessage != actualErrorMessage {
					t.Errorf("expected value: %v\nactual value: %v\n", expectedErrorMessage, actualErrorMessage)
				}
			default:
				t.Errorf("Container did not return an error of the expected type.")
			}
		} else {
			t.Errorf("Container did not panic.")
		}
	}()

	Register[testdata.ICar, testdata.Exhaust]().AsSingleton()
}
