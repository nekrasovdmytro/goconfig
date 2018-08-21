package goconfig

type Configurator struct {
	Data map[string]string
	loader ConfigLoader
}

type ConfigLoader interface {
	load(filename string) map[string]string
}

func (c *Configurator) SetLoader(loader ConfigLoader) {
	c.loader = loader
}

func (c *Configurator) Load(filename string)  {
	c.Data = c.loader.load(filename)
}

func (c *Configurator) GetData() interface{}  {
	return c.Data
}
