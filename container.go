package falconsnest

type Container struct {
}

func NewContainer() Container {
	return Container{}
}

func RegisterType[I any, S any]() Registrar[I, S] {
	return newRegistrar[I, S]()
}
