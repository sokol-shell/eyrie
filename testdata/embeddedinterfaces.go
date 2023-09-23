package testdata

type Embeddedinterface1 interface {
	Method1()
}

type Finalinterface1 interface {
	Embeddedinterface1
	Method2()
}

type Embeddedinterface2 interface {
	Method3()
}

type Finalinterface2 interface {
	Method3()
	Method4()
}
