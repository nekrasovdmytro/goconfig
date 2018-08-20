package goconfig

type Configurator struct {
	Data interface{}
	loader ConfigLoader
}

type ConfigLoader interface {
	load(filename string) interface{}
}

func (c *Configurator) setLoader(loader ConfigLoader) {
	c.loader = loader
}

func (c *Configurator) load(filename string)  {
	c.Data = c.loader.load(filename)
}

func (c *Configurator) getData() interface{}  {
	return c.Data
}
