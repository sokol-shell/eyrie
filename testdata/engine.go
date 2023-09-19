package testdata

type IEngine interface {
	ReadMileage() float32
}

type Engine struct {
}

func (e Engine) ReadMileage() float32 {
	return 156896.226
}
