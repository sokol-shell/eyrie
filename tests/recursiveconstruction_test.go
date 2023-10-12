package tests

import "github.com/ivan-ivkovic/falconsnest/container"

type ifirst interface {
	first()
}

type first struct {
	s isecond
}

func newFirst() *first {
	return &first{
		s: container.Resolve[isecond](),
	}
}

func (f *first) first() {}

type isecond interface {
	second()
}

type second struct {
	t ithird
}

func newSecond() *second {
	return &second{
		t: container.Resolve[ithird](),
	}
}

func (s *second) second() {}

type ithird interface {
	third()
}

type third struct {
	f ifirst
}

func newThird() *third {
	return &third{
		f: container.Resolve[ifirst](),
	}
}

func (t *third) third() {}
