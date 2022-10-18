package main

import (
	"machine"
	"time"
)

func main() {
	rate := time.Second / 100000000
	leds := machine.PC13
	leds.Configure(machine.PinConfig{Mode: machine.PinOutput})
	for i := 0; i < 100000000; i++ {
		leds.High()
		// time.Sleep(rate)
		// leds.Low()
		// time.Sleep(rate)
	}
	for {
		leds.High()
		time.Sleep(rate)
		leds.Low()
		time.Sleep(rate)
	}
}
