package falconsnest

type Registrar[I any, S any] interface {
	AsSingleton()
	AsTransient()
}

func newRegistrar[I any, S any]() Registrar[I, S] {
	return registrar[I, S]{}
}

type registrar[I any, S any] struct {
}

func (r registrar[I, S]) AsSingleton() {

}

func (r registrar[I, S]) AsTransient() {

}
