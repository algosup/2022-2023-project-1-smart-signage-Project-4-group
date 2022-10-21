package main

import (
	"machine"
	"time"
)

func GetCurrentSensorData() (bool, bool) { // Function that get the state of the different power sensor
	sensor1 := machine.PA4
	sensor2 := machine.PA5
	sensor1.Configure(machine.PinConfig{Mode: machine.PinInput})
	sensor2.Configure(machine.PinConfig{Mode: machine.PinInput})
	out1 := (sensor1.Get())
	out2 := (sensor2.Get())
	return out1, out2
}

func Light(isReduce bool, isOn bool) { // Function used to switch the light ON/OFF and control the intensity
	rate := time.Second / 100000000
	leds := machine.PC13
	leds.Configure(machine.PinConfig{Mode: machine.PinOutput})
	if isOn {
		if isReduce {
			for {
				leds.High()
				time.Sleep(rate)
				leds.Low()
				time.Sleep(rate)
			}
		} else {
			leds.High()
		}
	} else {
		leds.Low()
	}
}
