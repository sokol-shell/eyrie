package testdata

type IExhaust interface {
	FetchExhaustType() string
}

type Exhaust struct {
}

func (e Exhaust) FetchExhaustType() string {
	return "DUAL"
}
