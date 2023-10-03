package testdata

type IExhaust interface {
	FetchExhaustType() string
}

type Exhaust struct {
}

func NewExhaust() Exhaust {
	return Exhaust{}
}

func (e Exhaust) FetchExhaustType() string {
	return "DUAL"
}
