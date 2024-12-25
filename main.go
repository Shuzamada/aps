package main

import (
	"flag"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	isStepMode := flag.Bool("step", false, "Run in step-by-step mode")
	bufferSize := flag.Int("buffer", 5, "Buffer size")
	managerCount := flag.Int("managers", 3, "Number of event managers")
	managerLoad := flag.Int("maxload", 3, "Manager max load")
	lambda := flag.Float64("lambda", 1.0, "Poisson arrival rate (average arrivals per second)")
	stepInterval := flag.Float64("interval", 1.0, "Time interval for one step in seconds")
	flag.Parse()

	system := NewSystem(*bufferSize, *managerCount, *managerLoad, *lambda, *stepInterval)
	if *isStepMode {
		system.RunStepMode()
	} else {
		system.RunAutoMode()
	}
}
