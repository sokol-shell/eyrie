package falconsnest

type Exhaust interface {
	FetchExhaustType() string
}

type exhaust struct {
}

func (e exhaust) FetchExhaustType() string {
	return "DUAL"
}
