package configuration

type Configuration interface {
	GetOrCreateInstance() any
}
