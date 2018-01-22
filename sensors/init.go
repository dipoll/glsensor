package sensors

import (
	"errors"
	"time"
)

// Measurer interface represents a simple
// sensor that gets metrics
type Measurer interface {
	Measure() ([]Measurement, error)
}

// Metric represnts metric information
type Metric struct {
	Name string
	Unit string
}

// Measurement is actual data from the sensor
type Measurement struct {
	T    time.Time
	Info *Metric
	V    string
}

var sensorRegister map[string]func() Measurer = make(map[string]func() Measurer)

// Register registers initializer function for particular sensor
func Register(typeName string, f func() Measurer) {
	sensorRegister[typeName] = f
}

// GetSensorByName returns the initialization function of particular type
func GetSensorByName(name string) (Measurer, error) {
	fun, ok := sensorRegister[name]
	if !ok {
		return nil, errors.New("Sensor's type is not found in the registry")
	}
	return fun(), nil
}
