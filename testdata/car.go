package testdata

type ICar interface {
	GetVIM() string
	GetExhaustType() string
	GetMileage() float32
}

func NewICar() ICar {
	return Car{}
}

type Car struct {
	engine  IEngine
	exhaust IExhaust
}

func NewCar() Car {
	return Car{}
}

func (c Car) GetVIM() string {
	return "WVWMK516ILPG20083"
}

func (c Car) GetExhaustType() string {
	return c.exhaust.FetchExhaustType()
}

func (c Car) GetMileage() float32 {
	return c.engine.ReadMileage()
}
