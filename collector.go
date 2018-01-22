package glsensor

import (
	"log"
	"sync"

	"github.com/dipoll/glsensor/destinations"
	"github.com/dipoll/glsensor/sensors"
)

//Server Main service structure
type Server struct {
	configuration       []*DeviceConf
	destinations        *destinations.Router
	currentMeasurements []destinations.MetricValue
	mtx                 sync.RWMutex
	readyToSend         chan *destinations.MetricValue
	halt                chan bool
}

//NewServer creates a new collector server
func NewServer(conf []*DeviceConf) *Server {
	s := Server{
		configuration: conf,
		readyToSend:   make(chan *destinations.MetricValue),
		halt:          make(chan bool),
	}
	return &s
}

//Start runs listen and forward data to senders
func (m *Server) Start() {
	log.Println("Starting server")
	m.forwardToSender()
}

func (m *Server) forwardToSender() bool {
	for {
		select {
		case msg := <-m.readyToSend:
			m.destinations.Send(msg)
		case shutDown := <-m.halt:
			if shutDown {
				log.Println("Got shutdown signal, stopping ...")
				return true
			}
		default:
			continue
		}
	}
}

//Shutdown sends signal to stop handling of destinations
func (m *Server) Shutdown() {
	log.Println("Shutdown the server!")
	m.halt <- true
}

// CollectAll runs all measurers to collect data
// into the memory and sends all notifications to the
// destinations
func (m *Server) CollectAll() error {
	for _, device := range m.configuration {
		err := m.collectFromDevice(device)
		if err != nil {
			return err
		}
	}
	return nil
}

func (m *Server) collectFromDevice(d *DeviceConf) error {
	fin := make(chan int, len(d.Sensors))
	for _, sconf := range d.Sensors {
		go func(sens sensors.Measurer, c chan int) {
			defer func() {
				if r := recover(); r != nil {
					c <- 0
				}
			}()
			measurements, err := sens.Measure()
			if err != nil {
				return
			}
			for _, meas := range measurements {
				go m.addRetreivedValue(meas, d, false)
			}

			c <- 1
		}(sconf.Sensor, fin)
	}
	for i := 0; i < len(d.Sensors); i++ {
		<-fin
	}
	close(fin)
	return nil
}

//Added retrieved value
func (m *Server) addRetreivedValue(meas sensors.Measurement, d *DeviceConf, notify bool) {
	m.mtx.Lock()
	rval := destinations.MetricValue{M: &meas, Name: d.Name, Location: d.Location, Region: d.Region}
	m.currentMeasurements = append(m.currentMeasurements, rval)
	m.mtx.Unlock()
	if notify {
		m.readyToSend <- &rval
	}
}
