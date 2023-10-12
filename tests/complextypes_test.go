package tests

import "github.com/ivan-ivkovic/falconsnest/container"

// #region IEngine -> Engine
type IEngine interface {
	ReadMileage() float32
}

type Engine struct {
}

func NewEngine() *Engine {
	return &Engine{}
}

func (e *Engine) ReadMileage() float32 {
	return 156896.226
}

//#endregion

// #region IExhaust -> Exhaust
type IExhaust interface {
	FetchExhaustType() string
}

type Exhaust struct {
}

func NewExhaust() *Exhaust {
	return &Exhaust{}
}

func (e *Exhaust) FetchExhaustType() string {
	return "DUAL"
}

//#endregion

// #region ICar -> Car
type ICar interface {
	GetVIM() string
	GetExhaustType() string
	GetEngineMileage() float32
}

type Car struct {
	engine  IEngine
	exhaust IExhaust
}

func NewCar() *Car {
	return &Car{
		engine:  container.Resolve[IEngine](),
		exhaust: container.Resolve[IExhaust](),
	}
}

func (c *Car) GetVIM() string {
	return "WVWMK516ILPG20083"
}

func (c *Car) GetExhaustType() string {
	return c.exhaust.FetchExhaustType()
}

func (c *Car) GetEngineMileage() float32 {
	return c.engine.ReadMileage()
}

//#endregion
