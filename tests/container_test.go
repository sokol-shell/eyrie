//go:build test

package tests

import (
	"sync"
	"testing"

	"github.com/ivan-ivkovic/eyrie/container"
)

func Test_ContainerSuccessfullyInstantiatesTwoTransientInstances(t *testing.T) {
	container.Clear()
	container.Register[IInterface, Implementation](NewImplementation).AsTransient()

	result1 := container.Resolve[IInterface]()
	result2 := container.Resolve[IInterface]()

	if result1.(Implementation).memoryAddress == result2.(Implementation).memoryAddress {
		t.Fatalf("Transient configuration produced only one object.")
	}
}

func Test_ContainerSuccessfullyInstantiatesOnlyOneSingletonInstance(t *testing.T) {
	container.Clear()
	container.Register[IInterface, Implementation](NewImplementation).AsSingleton()

	result1 := container.Resolve[IInterface]()
	result2 := container.Resolve[IInterface]()

	if result1.(Implementation).memoryAddress != result2.(Implementation).memoryAddress {
		t.Fatalf("Singleton configuration produced two different objects.")
	}
}

func Test_ContainerSuccessfullyRegistersAndResolvesTypesWithPointerReceivers(t *testing.T) {
	container.Clear()
	container.Register[IPointerInterface, *PointerStruct](NewPointerStruct).AsSingleton()

	result := container.Resolve[IPointerInterface]()

	_, ok := result.(*PointerStruct)
	if !ok {
		t.Fatalf("Resolve did not return an object of the expected type.")
	}
}

func Test_ContainerSuccessfullyRegistersAndResolvesExportedTypesAndConstructors(t *testing.T) {
	container.Clear()
	container.Register[IInterface, Implementation](NewImplementation).AsTransient()

	result := container.Resolve[IInterface]()

	_, ok := result.(Implementation)
	if !ok {
		t.Fatalf("Resolve did not return an object of the expected type.")
	}
}

func Test_ContainerSuccessfullyRegistersAndResolvesUnexportedTypesAndConstructors(t *testing.T) {
	container.Clear()
	container.Register[iinterface, implementation](newImplementation).AsTransient()

	result := container.Resolve[iinterface]()

	_, ok := result.(implementation)
	if !ok {
		t.Fatalf("Resolve did not return an object of the expected type.")
	}
}

func Test_ContainerSuccessfullyRegistersAndResolvesInterfacesAndEmbeddedStructs(t *testing.T) {
	container.Clear()
	container.Register[IInterface1, FullImplementation](NewFullImplementation).AsSingleton()

	result := container.Resolve[IInterface1]()

	_, ok := result.(FullImplementation)
	if !ok {
		t.Fatalf("Resolve did not return an object of the expected type.")
	}
}

func Test_ContainerSuccessfullyResolvesComplexDependencies(t *testing.T) {
	container.Clear()
	container.Register[ICar, *Car](NewCar).AsTransient()
	container.Register[IExhaust, *Exhaust](NewExhaust).AsTransient()
	container.Register[IEngine, *Engine](NewEngine).AsTransient()

	car := container.Resolve[ICar]()

	if car.GetExhaustType() != "DUAL" || car.GetEngineMileage() != 156896.226 {
		t.Fatalf("Container did not properly resolve sub-dependencies.")
	}
}

func Test_CannotRegisterInterfaceToAStruct(t *testing.T) {
	container.Clear()
	var expectedErrorMessage = "RegistrationError: Interface and struct expected as type parameters."

	defer catchPanic(t, expectedErrorMessage)
	container.Register[Implementation, IInterface](NewIInterface)
}

func Test_CannotRegisterPointerToAStruct(t *testing.T) {
	container.Clear()
	var expectedErrorMessage = "RegistrationError: Interface and struct expected as type parameters."

	defer catchPanic(t, expectedErrorMessage)
	container.Register[*PointerStruct, Implementation](NewImplementation)
}

func Test_CannotRegisterPointerToAPointer(t *testing.T) {
	container.Clear()
	var expectedErrorMessage = "RegistrationError: Interface and struct expected as type parameters."

	defer catchPanic(t, expectedErrorMessage)
	container.Register[*PointerStruct, *PointerStruct](NewPointerStruct)
}

func Test_CannotRegisterIfStructDoesNotImplementInterface(t *testing.T) {
	container.Clear()
	var expectedErrorMessage = "RegistrationError: implementation does not implement IInterface."

	defer catchPanic(t, expectedErrorMessage)
	container.Register[IInterface, implementation](newImplementation)
}

func Test_CannotRegisterInterfaceWhichImplementsEmbeddedInterface1(t *testing.T) {
	container.Clear()
	var expectedErrorMessage = "RegistrationError: Interface and struct expected as type parameters."

	defer catchPanic(t, expectedErrorMessage)
	container.Register[EmbeddedInterface1, FinalInterface1](NewFinalInterface1)
}

func Test_CannotRegisterInterfaceWhichImplementsEmbeddedInterface2(t *testing.T) {
	container.Clear()
	var expectedErrorMessage = "RegistrationError: Interface and struct expected as type parameters."

	defer catchPanic(t, expectedErrorMessage)
	container.Register[EmbeddedInterface2, FinalInterface2](NewFinalInterface2)
}

func Test_CannotResolveTypeWhichWasNotRegistered(t *testing.T) {
	container.Clear()
	var expectedErrorMessage = "ResolveError: Could not resolve EmbeddedInterface1. Not found."

	defer catchPanic(t, expectedErrorMessage)
	container.Resolve[EmbeddedInterface1]()
}

func Test_CannotRegisterAfterSealingTheContainer(t *testing.T) {
	container.Clear()
	var expectedErrorMessage = "SealedContainerError: Cannot register a new type to a sealed container."

	container.Register[IInterface, Implementation](NewImplementation).AsSingleton()
	container.Seal()

	defer catchPanic(t, expectedErrorMessage)
	container.Register[iinterface, implementation](newImplementation).AsSingleton()
}

func Test_ConcurrentResolveOfASingletonResultsInOnlyOneInstance(t *testing.T) {
	container.Clear()
	container.Register[ISlowConstructor, SlowConstructor](NewSlowConstructor).AsSingleton()

	var mu sync.Mutex
	instances := make([]ISlowConstructor, 0)

	numberOfRoutines := 100

	var routinesReady, waitForStart, waitForFinish sync.WaitGroup

	routinesReady.Add(numberOfRoutines)
	waitForStart.Add(1)
	waitForFinish.Add(numberOfRoutines)
	for i := 0; i < numberOfRoutines; i++ {
		go func() {
			routinesReady.Done()
			waitForStart.Wait()
			result := container.Resolve[ISlowConstructor]()

			mu.Lock()
			instances = append(instances, result)
			mu.Unlock()

			waitForFinish.Done()
		}()
	}

	routinesReady.Wait()
	waitForStart.Done()
	waitForFinish.Wait()

	for i := 0; i < numberOfRoutines; i++ {
		for j := i + 1; j < numberOfRoutines; j++ {
			if instances[i].(SlowConstructor).memoryAddress != instances[j].(SlowConstructor).memoryAddress {
				t.Fatalf("Singleton configuration produced two different objects.")
			}
		}
	}
}

func Test_ConcurrentResolveOfATransientConfigurationResultsInDifferentInstances(t *testing.T) {
	container.Clear()
	container.Register[ISlowConstructor, SlowConstructor](NewSlowConstructor).AsTransient()

	var mu sync.Mutex
	instances := make([]ISlowConstructor, 0)

	numberOfRoutines := 100

	var routinesReady, waitForStart, waitForFinish sync.WaitGroup

	routinesReady.Add(numberOfRoutines)
	waitForStart.Add(1)
	waitForFinish.Add(numberOfRoutines)
	for i := 0; i < numberOfRoutines; i++ {
		go func() {
			routinesReady.Done()
			waitForStart.Wait()
			result := container.Resolve[ISlowConstructor]()

			mu.Lock()
			instances = append(instances, result)
			mu.Unlock()

			waitForFinish.Done()
		}()
	}

	routinesReady.Wait()
	waitForStart.Done()
	waitForFinish.Wait()

	for i := 0; i < numberOfRoutines; i++ {
		for j := i + 1; j < numberOfRoutines; j++ {
			if instances[i].(SlowConstructor).memoryAddress == instances[j].(SlowConstructor).memoryAddress {
				t.Fatalf("transient configuration produced two same objects")
			}
		}
	}
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
