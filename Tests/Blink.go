package Prog

import (
	"machine"
	"time"
)

func BlinkLed(rate time.Duration) {
	leds := machine.PC13
	leds.Configure(machine.PinConfig{Mode: machine.PinOutput})
	for {
		leds.High()
		time.Sleep(rate)
		leds.Low()
		time.Sleep(rate)
	}
}

func LowIntensityLed() {
	rate := time.Second / 100000000
	BlinkLed(rate)
}
