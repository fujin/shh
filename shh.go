package main

import (
	"fmt"
	"github.com/freeformz/shh/mm"
	"github.com/freeformz/shh/pollers"
	"time"
)

func writeOut(measurements chan *mm.Measurement) {
	for measurement := range measurements {
		fmt.Println(measurement)
	}
}

func main() {
	measurements := make(chan *mm.Measurement, 100)
	duration, _ := time.ParseDuration("5s")
	ticks := time.Tick(duration)
	go writeOut(measurements)

	mp := pollers.NewMultiPoller()
	mp.RegisterPoller(pollers.Load{})
	mp.RegisterPoller(pollers.Cpu{})

	for now := range ticks {
		measurements <- &mm.Measurement{now, "tick", []byte("true")}
		go mp.Poll(now, measurements)
	}
}
