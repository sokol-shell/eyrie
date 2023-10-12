package tests

import (
	"github.com/ivan-ivkovic/falconsnest/container"
	"testing"
)

func Test_ContainerSuccessfullyInstantiatesTwoTransientInstances(t *testing.T) {
	container.Register[IInterface, Implementation](NewImplementation).AsTransient()

	result1 := container.Resolve[IInterface]()
	result2 := container.Resolve[IInterface]()

	if result1.(Implementation).memoryAddress == result2.(Implementation).memoryAddress {
		t.Fatalf("Transient configuration produced only one object.")
	}
}

func Test_ContainerSuccessfullyInstantiatesOnlyOneSingletonInstance(t *testing.T) {
	container.Register[IInterface, Implementation](NewImplementation).AsSingleton()

	result1 := container.Resolve[IInterface]()
	result2 := container.Resolve[IInterface]()

	if result1.(Implementation).memoryAddress != result2.(Implementation).memoryAddress {
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
	container.Register[IInterface, Implementation](NewImplementation).AsTransient()

	result := container.Resolve[IInterface]()

	_, ok := result.(Implementation)
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

func Test_ContainerSuccessfullyRegistersAndResolvesInterfacesAndEmbeddedStructs(t *testing.T) {
	container.Register[IInterface1, FullImplementation](NewFullImplementation).AsSingleton()

	result := container.Resolve[IInterface1]()

	_, ok := result.(FullImplementation)
	if !ok {
		t.Fatalf("Resolve did not return an object of the expected type.")
	}
}

func Test_ContainerSuccessfullyResolvesComplexDependencies(t *testing.T) {
	container.Register[ICar, *Car](NewCar).AsTransient()
	container.Register[IExhaust, *Exhaust](NewExhaust).AsTransient()
	container.Register[IEngine, *Engine](NewEngine).AsTransient()

	car := container.Resolve[ICar]()

	if car.GetExhaustType() != "DUAL" || car.GetEngineMileage() != 156896.226 {
		t.Fatalf("Container did not properly resolve sub-dependencies.")
	}
}

func Test_RecursiveSingletonDependenciesDontCauseStackOverflow(t *testing.T) {
	var expectedErrorMessage = "RecursiveConstructionError: Recursive construction occurred while resolving ifirst."

	container.Register[ifirst, *first](newFirst).AsSingleton()
	container.Register[isecond, *second](newSecond).AsSingleton()
	container.Register[ithird, *third](newThird).AsSingleton()

	defer catchPanic(t, expectedErrorMessage)
	container.Resolve[ifirst]()
}

func Test_RecursiveTransientDependenciesDontCauseStackOverflow(t *testing.T) {
	var expectedErrorMessage = "RecursiveConstructionError: Recursive construction occurred while resolving ifirst."

	container.Register[ifirst, *first](newFirst).AsTransient()
	container.Register[isecond, *second](newSecond).AsTransient()
	container.Register[ithird, *third](newThird).AsTransient()

	defer catchPanic(t, expectedErrorMessage)
	container.Resolve[ifirst]()
}

func Test_CannotRegisterInterfaceToAStruct(t *testing.T) {
	var expectedErrorMessage = "RegistrationError: Interface and struct expected as type parameters."

	defer catchPanic(t, expectedErrorMessage)
	container.Register[Implementation, IInterface](NewIInterface)
}

func Test_CannotRegisterPointerToAStruct(t *testing.T) {
	var expectedErrorMessage = "RegistrationError: Interface and struct expected as type parameters."

	defer catchPanic(t, expectedErrorMessage)
	container.Register[*PointerStruct, Implementation](NewImplementation)
}

func Test_CannotRegisterPointerToAPointer(t *testing.T) {
	var expectedErrorMessage = "RegistrationError: Interface and struct expected as type parameters."

	defer catchPanic(t, expectedErrorMessage)
	container.Register[*PointerStruct, *PointerStruct](NewPointerStruct)
}

func Test_CannotRegisterIfStructDoesNotImplementInterface(t *testing.T) {
	var expectedErrorMessage = "RegistrationError: implementation does not implement IInterface."

	defer catchPanic(t, expectedErrorMessage)
	container.Register[IInterface, implementation](newImplementation)
}

func Test_CannotRegisterInterfaceWhichImplementsEmbeddedInterface1(t *testing.T) {
	var expectedErrorMessage = "RegistrationError: Interface and struct expected as type parameters."

	defer catchPanic(t, expectedErrorMessage)
	container.Register[EmbeddedInterface1, FinalInterface1](NewFinalInterface1)
}

func Test_CannotRegisterInterfaceWhichImplementsEmbeddedInterface2(t *testing.T) {
	var expectedErrorMessage = "RegistrationError: Interface and struct expected as type parameters."

	defer catchPanic(t, expectedErrorMessage)
	container.Register[EmbeddedInterface2, FinalInterface2](NewFinalInterface2)
}

func Test_CannotResolveTypeWhichWasNotRegistered(t *testing.T) {
	var expectedErrorMessage = "ResolveError: Could not resolve EmbeddedInterface1. Not found."

	defer catchPanic(t, expectedErrorMessage)
	container.Resolve[EmbeddedInterface1]()
}

func Test_CannotRegisterAfterFirstResolveWasDone(t *testing.T) {
	var expectedErrorMessage = "SealedContainerError: Cannot register a new type to a sealed container."

	container.Register[IInterface, Implementation](NewImplementation).AsSingleton()
	container.Seal()

	defer catchPanic(t, expectedErrorMessage)
	container.Register[iinterface, implementation](newImplementation).AsSingleton()
}

func catchPanic(t *testing.T, expectedErrorMessage string) {
	t.Helper()
	if err := recover(); err != nil {
		switch err.(type) {
		case *container.RegistrationError:
			var actualErrorMessage = err.(*container.RegistrationError).Error()
			if expectedErrorMessage != actualErrorMessage {
				t.Errorf("expected value: %v\nactual value: %v\n", expectedErrorMessage, actualErrorMessage)
			}
		case *container.ResolveError:
			var actualErrorMessage = err.(*container.ResolveError).Error()
			if expectedErrorMessage != actualErrorMessage {
				t.Errorf("expected value: %v\nactual value: %v\n", expectedErrorMessage, actualErrorMessage)
			}
		case *container.SealedContainerError:
			var actualErrorMessage = err.(*container.SealedContainerError).Error()
			if expectedErrorMessage != actualErrorMessage {
				t.Errorf("expected value: %v\nactual value: %v\n", expectedErrorMessage, actualErrorMessage)
			}
		case *container.RecursiveConstructionError:
			var actualErrorMessage = err.(*container.RecursiveConstructionError).Error()
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
