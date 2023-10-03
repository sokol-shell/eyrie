package falconsnest

import "fmt"

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
