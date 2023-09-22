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
