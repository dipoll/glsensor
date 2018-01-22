package glsensor

import (
	"encoding/json"
	"strings"
	"testing"
)

const testConf = `{	

	"name" : "DeviceName",
	"location" : "DeviceLocation",
	"region" : "DeviceRegion",
	"sensors" : [
		{
			"type" : "DS18B20",
			"name" : "DS18B20-01",
			"id" : "0216249a9eee"
		},
		{
			"type" : "DS18B20",
			"name" : "DS18B20-02",
			"id" : "01159047f8ff"
		},		
		{
			"type" : "DHT11",
			"name" : "KY-015-01",
			"pin" : 17
		}
	],
	"destinations" : [
		{
			"type" : "console",
			"name" : "Basic Console"
		
		},
		{       
			"name" : "Primary Influx storage",
			"type": "Influx",
			"server": "InfluxServiceStorage",
			"port" : 8086,
			"user" : "testuser",
			"password" : "test",
			"collection" : "testCollection"

		}
	]	
}`

const noSensor string = `{
	"type" : "NonExisting",
	"name" : "DS18B20-02",
	"id" : "01159047f8ff"
}	`

var cfg DeviceConf
var server Server

func TestDeviceStructureReading(t *testing.T) {
	cfg, _ = ReadConfiguration(strings.NewReader(testConf))
	if cfg.Name != "DeviceName" {
		t.Fatal("Bad parsing of the name")
	}
}

func TestSensorStructReading(t *testing.T) {
	var sensor SensorConf
	dec := json.NewDecoder(strings.NewReader(noSensor))
	err := dec.Decode(&sensor)
	if err == nil {
		t.Fatal("Fails reading of non existing handler for a sensor")
	}
}
