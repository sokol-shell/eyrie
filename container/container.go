package container

import (
	"fmt"
	"github.com/ivan-ivkovic/falconsnest/internal/configuration"
	"reflect"
)

type Container interface {
	addConfiguration(typ reflect.Type, config configuration.Configuration)
	getConfiguration(typ reflect.Type) configuration.Configuration
	seal()
}

type container struct {
	configurations map[string]configuration.Configuration
	initialized    bool
	isSealed       bool
}

var c *container

func getContainer() Container {
	if c == nil {
		c = &container{
			configurations: make(map[string]configuration.Configuration),
			initialized:    true,
		}
	}

	return c
}

func (c *container) addConfiguration(typ reflect.Type, config configuration.Configuration) {
	c.assertNotSealed()

	key := c.generateKey(typ)
	c.configurations[key] = config
}

func (c *container) getConfiguration(typ reflect.Type) configuration.Configuration {
	key := c.generateKey(typ)
	config, ok := c.configurations[key]
	if ok {
		return config
	}

	panic(newResolveError(fmt.Sprintf("Could not resolve %s. Not found.", typ.Name())))
}

func (c *container) generateKey(typ reflect.Type) string {
	return typ.PkgPath() + "/" + typ.Name()
}

func (c *container) seal() {
	c.isSealed = true
}

func (c *container) assertNotSealed() {
	if c.isSealed {
		panic(newSealedContainerError())
	}
}
