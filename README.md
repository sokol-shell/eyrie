# Go Dependency Injection Framework

Small, lightweight dependency injection framework for Go.

### Usage
```go
type Interface interface {
	Method()
}

type Implementation struct {}

func NewImplementation() *Implementation {
    return &Implementation{}
}
func (i *Implementation) Method() {}
```

Register singleton configuration:
```go
container.Register[Interface, *Implementation](NewImplementation).AsSingleton()
```

Register transient configuration:
```go
container.Register[Interface, *Implementation](NewImplementation).AsTransient()
```

Type resolving:
```go
var i Interface = container.Resolve[Interface]()
```