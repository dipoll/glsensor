package glsensor

import (
	"encoding/json"
	"io"

	"github.com/dipoll/glsensor/destinations"
	"github.com/dipoll/glsensor/sensors"
)

/*
SensorConf structure is configuration object to be read from json
*/
type SensorConf struct {
	Sensor sensors.Measurer
}

/*
DestinationConf represents configuration data handler
for destinations
*/
type DestinationConf struct {
	Destination destinations.Sender
}

/*
UnmarshalJSON does reading of json structure by
applying particular type of the sensor's handler
*/
func (s *SensorConf) UnmarshalJSON(data []byte) error {
	type Sens struct {
		Type string `json:"Type,omitempty"`
	}
	var _type Sens

	err := json.Unmarshal(data, &_type)
	if err != nil {
		return err
	}
	cfg, err := sensors.GetSensorByName(_type.Type)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, &cfg)
	if err != nil {
		return err
	}
	s.Sensor = cfg
	return nil
}

/*
UnmarshalJSON does reading of json structure by
applying particular type of the sensor's handler
*/
func (d *DestinationConf) UnmarshalJSON(data []byte) error {
	type SenderInstance struct {
		Type string `json:"Type,omitempty"`
	}
	var _type SenderInstance

	err := json.Unmarshal(data, &_type)
	if err != nil {
		return err
	}
	cfg, err := destinations.GetSenderByName(_type.Type)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &cfg)
	if err != nil {
		return err
	}
	d.Destination = cfg
	return nil
}

/*
DeviceConf structure is configuration object to be read from json
Is a placehoslder for full configuration
*/
type DeviceConf struct {
	Name         string
	Location     string
	Region       string
	Sensors      []SensorConf
	Destinations []DestinationConf
}

/*
ReadConfiguration reads full configuration from json string
and returns initialized SensorConf
*/
func ReadConfiguration(r io.Reader) (DeviceConf, error) {
	var dc DeviceConf
	dec := json.NewDecoder(r)
	err := dec.Decode(&dc)

	return dc, err
}
