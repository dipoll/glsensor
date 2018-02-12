package destinations

import (
	"errors"
	"sync"

	"github.com/dipoll/glsensor/sensors"
)

//var senders map[string]func() Sender = make(map[string]func() Sender)
var senders = make(map[string]func() Sender)

//DestinationNotFound error for failed lookup
var DestinationNotFound = errors.New("No desination sender found")

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
	destinations []Sender `json:"destinations"`
}

func (d *Router) Send(s *MetricValue) error {
	var mut sync.RWMutex
	if len(d.destinations) < 1 {
		return errors.New("No destinations assigned!")
	}
	for _, d := range d.destinations {
		mut.Lock()
		defer mut.Unlock()
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

//GetSenderbyName returns hadler from senders maps and error
//if handler is not found DestinationNotFount is returned
func GetSenderByName(name string) (Sender, error) {
	fun, ok := senders[name]
	if !ok {
		return nil, DestinationNotFound
	}
	return fun(), nil
}
