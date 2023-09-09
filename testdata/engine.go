package falconsnest

type Engine interface {
	ReadMileage() float32
}

type engine struct {
}

func (e engine) ReadMileage() float32 {
	return 156896.226
}
