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
/*
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
*/

type LED struct {
    pin machine.Pin
    on  bool
}

func New() LED {
    leds := machine.PC13
    leds.Configure(machine.PinConfig{
        Mode: machine.PinOutput,
    })
    l := LED{
        pin: leds,
        on:  false,
    }
    return &l
}

// On indicates if the LED is turned on.
func (lLED) On() bool {
    return l.on
}

// Set sets the LED to the value.
// true will turn the LED on
// false will turn the LED off
func (l LED) Set(value bool) {
    l.on = value
    if l.on {
        l.pin.High()
        return
    } else {
        l.pin.Low()
    }
}

func (lLED) Blink() {
    rate := time.Second / 500
    for i := 0; i < 1000; i++ {
        l.High()
        l.Set(true)
        time.Sleep(rate)
        l.Low()
        l.Set(false)
        time.Sleep(rate)
    }
}
