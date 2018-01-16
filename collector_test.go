package glsensor

import (
	"reflect"
	"strings"
	"testing"

	"github.com/dipoll/glsensor/sensors"
)

const testMServerConf = `{	

	"name" : "DeviceName",
	"location" : "DeviceLocation",
	"region" : "DeviceRegion",
	"sensors" : [
		{
			"type" : "Dummy",
			"name" : "Dummy-01",
			"id" : "0216249a9eee"
		}
	],
	"destinations" : [
		{
			"type" : "console",
			"name" : "debug"
		}
	]
}`

func TestCollectorRun(t *testing.T) {
	conf, err := ReadConfiguration(strings.NewReader(testMServerConf))
	if err != nil {
		t.Fatal("Configuration can't be read!", conf)
	}
	s := NewMServer([]*DeviceConf{&conf})
	err = s.CollectAll()
	if err != nil {
		t.Fatal("Fails to collect all metrics", err)
	}
}

func TestNewMServer(t *testing.T) {
	type args struct {
		conf []*DeviceConf
	}
	tests := []struct {
		name string
		args args
		want *MServer
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMServer(tt.args.conf); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMServer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMServer_forwardSender(t *testing.T) {
	tests := []struct {
		name string
		m    *MServer
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.m.forwardToSender()
		})
	}
}

func TestMServer_Shutdown(t *testing.T) {
	tests := []struct {
		name string
		m    *MServer
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.m.Shutdown()
		})
	}
}

func TestMServer_CollectAll(t *testing.T) {
	tests := []struct {
		name    string
		m       *MServer
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.m.CollectAll(); (err != nil) != tt.wantErr {
				t.Errorf("MServer.CollectAll() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMServer_collectFromDevice(t *testing.T) {
	type args struct {
		d *DeviceConf
	}
	tests := []struct {
		name    string
		m       *MServer
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.m.collectFromDevice(tt.args.d); (err != nil) != tt.wantErr {
				t.Errorf("MServer.collectFromDevice() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMServer_addRetreivedValue(t *testing.T) {
	type args struct {
		meas   sensors.Measurement
		d      *DeviceConf
		notify bool
	}
	tests := []struct {
		name string
		m    *MServer
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.m.addRetreivedValue(tt.args.meas, tt.args.d, tt.args.notify)
		})
	}
}
