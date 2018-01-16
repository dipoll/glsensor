package destinations

import (
	"errors"

	"github.com/dipoll/glsensor/sensors"
)

var senders map[string]func() Sender = make(map[string]func() Sender)

type MetricValue struct {
	M        *sensors.Measurement
	Name     string
	Region   string
	Location string
}

type Sender interface {
	Send(s *MetricValue) error
}

type Router struct {
	destinations []Sender
}

func (d *Router) Send(s *MetricValue) error {
	for _, d := range d.destinations {
		err := d.Send(s)
		if err != nil {
			return err
		}
	}
	return nil
}

//Register function dedicated for
//registering available destinations
func Register(name string, f func() Sender) error {
	if _, ok := senders[name]; ok {
		return errors.New("\"" + name + " destination is already registered")
	}
	senders[name] = f
	return nil
}

func GetSenderByName(name string) (Sender, error) {
	fun, ok := senders[name]
	if !ok {
		return nil, errors.New("No desination sender found with name: " + name)
	}
	return fun(), nil
}
