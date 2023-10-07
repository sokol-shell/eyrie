package configuration

type ConfigurationError struct {
	Code ConfigurationErrorCode
}

func newConfigurationError(code ConfigurationErrorCode) *ConfigurationError {
	return &ConfigurationError{Code: code}
}

func (ce *ConfigurationError) Error() string {
	return "Configuration error."
}

type ConfigurationErrorCode int

const (
	RecursiveConstruction ConfigurationErrorCode = iota
)
