package main

import (
	"machine"
	"time"
)

func main() {
	CustomerOn := true
	machine.UART1.Configure(machine.UARTConfig{BaudRate: 9600, TX: machine.PA2, RX: machine.PA3})
	_, join := machine.UART1.Write([]byte("AT+JOIN\r\n"))
	_, msg := machine.UART1.Write([]byte("AT+MSG='allumé'\r\n"))

	if CustomerOn {

		println(join)
		Light(true, true)

		println(msg)
	} else {
		Light(false, false)

	}
}

func Light(isReduce bool, isOn bool) {
	rate := time.Second
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

func Sleep(duration int) {
	time.Sleep(time.Second * time.Duration(duration))
}

// func printBuf(b []byte) {
// 	for _, val := range b {
// 		fmt.Printf("%x ", val)
// 	}
// }
