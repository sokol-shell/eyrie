package falconsnest

type Car interface {
	GetVIM() string
	GetExhaustType() string
	GetMileage() float32
}

type car struct {
	engine  Engine
	exhaust Exhaust
}

func (c car) GetVIM() string {
	return "WVWMK516ILPG20083"
}

func (c car) GetExhaustType() string {
	return c.exhaust.FetchExhaustType()
}

func (c car) GetMileage() float32 {
	return c.engine.ReadMileage()
}
