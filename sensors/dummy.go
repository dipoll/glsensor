package sensors

type Dummy struct {
	Name string
	Type string
}

func (c *Dummy) Measure() ([]Measurement, error) {
	return []Measurement{Measurement{V: "100.1"}}, nil
}

type DS18B20 struct {
	Dummy
}

type DHT11 struct {
	Dummy
}

func init() {
	Register("Dummy", func() Measurer { return &Dummy{} })
	Register("DS18B20", func() Measurer { return &DS18B20{} })
	Register("DHT11", func() Measurer { return &DHT11{} })
}
