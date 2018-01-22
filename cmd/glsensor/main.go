package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/dipoll/glsensor"
)

var (
	infile = flag.String("conf", "", "Specify configuration file path")
)

func main() {
	flag.Parse()
	f, err := os.Open(*infile)
	if err != nil {
		log.Fatal("Could not open configuration file")
	}
	conf, err := glsensor.ReadConfiguration(f)
	if err != nil {
		log.Fatal("Could not read configuration: ", err)
	}
	s := glsensor.NewServer([]*glsensor.DeviceConf{&conf})
	fc := make(chan int)
	go func(c chan int) {
		s.Start()
		c <- 1
	}(fc)
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, os.Interrupt, syscall.SIGTERM)
	for {
		select {
		case sig := <-sc:
			log.Println(sig)
			s.Shutdown()
		case <-fc:
			os.Exit(0)
		}
	}
}
