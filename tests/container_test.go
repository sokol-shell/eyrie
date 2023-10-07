package tests

import (
	"falconsnest/container"
	"testing"
)

func Test_ContainerSuccessfullyInstantiatesTwoTransientInstances(t *testing.T) {
	container.Register[IInterface, Struct](NewStruct).AsTransient()

	result1 := container.Resolve[IInterface]()
	result2 := container.Resolve[IInterface]()

	if result1.(Struct).memoryAddress == result2.(Struct).memoryAddress {
		t.Fatalf("Transient configuration produced only one object.")
	}
}

func Test_ContainerSuccessfullyInstantiatesOnlyOneSingletonInstance(t *testing.T) {
	container.Register[IInterface, Struct](NewStruct).AsSingleton()

	result1 := container.Resolve[IInterface]()
	result2 := container.Resolve[IInterface]()

	if result1.(Struct).memoryAddress != result2.(Struct).memoryAddress {
		t.Fatalf("Singleton configuration produced two different objects.")
	}
}

func Test_ContainerSuccessfullyRegistersAndResolvesTypesWithPointerReceivers(t *testing.T) {
	container.Register[IPointerInterface, *PointerStruct](NewPointerStruct).AsSingleton()

	result := container.Resolve[IPointerInterface]()

	_, ok := result.(*PointerStruct)
	if !ok {
		t.Fatalf("Resolve did not return an object of the expected type.")
	}
}

func Test_ContainerSuccessfullyRegistersAndResolvesExportedTypesAndConstructors(t *testing.T) {
	container.Register[IInterface, Struct](NewStruct).AsTransient()

	result := container.Resolve[IInterface]()

	_, ok := result.(Struct)
	if !ok {
		t.Fatalf("Resolve did not return an object of the expected type.")
	}
}

func Test_ContainerSuccessfullyRegistersAndResolvesUnexportedTypesAndConstructors(t *testing.T) {
	container.Register[iinterface, implementation](newImplementation).AsTransient()

	result := container.Resolve[iinterface]()

	_, ok := result.(implementation)
	if !ok {
		t.Fatalf("Resolve did not return an object of the expected type.")
	}
}

func Test_CannotRegisterInterfaceToAStruct(t *testing.T) {
	var expectedErrorMessage = "Interface and struct expected as type parameters."

	defer catchPanic(t, expectedErrorMessage)
	container.Register[Struct, IInterface](NewIInterface)
}

func Test_CannotRegisterPointerToAStruct(t *testing.T) {
	var expectedErrorMessage = "Interface and struct expected as type parameters."

	defer catchPanic(t, expectedErrorMessage)
	container.Register[PointerStruct, Struct](NewStruct)
}

func Test_CannotRegisterIfStructDoesNotImplementInterface(t *testing.T) {
	var expectedErrorMessage = "Exhaust does not implement ICar."

	defer catchPanic(t, expectedErrorMessage)
	container.Register[IInterface, Struct](NewStruct)
}

func Test_CannotRegisterInterfaceWhichImplementsEmbeddedInterface1(t *testing.T) {
	var expectedErrorMessage = "Interface and struct expected as type parameters."

	defer catchPanic(t, expectedErrorMessage)
	container.Register[EmbeddedInterface1, FinalInterface1](NewFinalInterface1)
}

func Test_CannotRegisterInterfaceWhichImplementsEmbeddedInterface2(t *testing.T) {
	var expectedErrorMessage = "Interface and struct expected as type parameters."

	defer catchPanic(t, expectedErrorMessage)
	container.Register[EmbeddedInterface2, FinalInterface2](NewFinalInterface2)
}

func catchPanic(t *testing.T, expectedErrorMessage string) {
	t.Helper()
	if err := recover(); err != nil {
		switch err.(type) {
		case *container.RegistrationError:
			var actualErrorMessage = err.(*container.RegistrationError).Msg
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
