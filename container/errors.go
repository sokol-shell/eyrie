package container

import "fmt"

// #region RegistrationError
type RegistrationError struct {
	Msg string
}

func newRegistrationError(msg string) *RegistrationError {
	return &RegistrationError{
		Msg: msg,
	}
}

func (re *RegistrationError) Error() string {
	return fmt.Sprintf("RegistrationError: %s", re.Msg)
}

//#enregion

// #region ResolveError
type ResolveError struct {
	Msg string
}

func newResolveError(msg string) *ResolveError {
	return &ResolveError{
		Msg: msg,
	}
}

func (re *ResolveError) Error() string {
	return fmt.Sprintf("ResolveError: %s", re.Msg)
}

//#endregion

// #region SealedContainerError
type SealedContainerError struct{}

func newSealedContainerError() *SealedContainerError {
	return &SealedContainerError{}
}

func (sce *SealedContainerError) Error() string {
	return fmt.Sprintf("SealedContainerError: Cannot register a new type to a sealed container.")
}

//#endregion

// #region RecursiveConstructionError
type RecursiveConstructionError struct {
	Msg string
}

func newRecursiveConstructionError(msg string) *RecursiveConstructionError {
	return &RecursiveConstructionError{
		Msg: msg,
	}
}

func (rce *RecursiveConstructionError) Error() string {
	return fmt.Sprintf("RecursiveConstructionError: %s", rce.Msg)
}

//#endregion
