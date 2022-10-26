package main

import (
	"machine"
	"time"
)

func main() {
	uart := machine.UART2
	uart.Configure(machine.UARTConfig{BaudRate: 9600, TX: machine.A2, RX: machine.A3})
	_, err := uart.Write([]byte("AT+JOIN\r\n"))

	if err != nil {
		println("Error: " + err.Error())

		// return ""
	}
	msg1 := ""
	for {
		if uart.Buffered() > 0 {
			rb, err := uart.ReadByte()
			if err != nil {
				println("Error: " + err.Error())
				continue
				// return ""
			}
			msg1 += string(rb)
			if msg1[len(msg1)-1] == '\n' {
				break
			}
			uart.Write([]byte("AT+MSG\"Hello\"\r\n"))
		}

	}

	minLight()
	/*for {
		for i := 0; i < 5000; i++ {
			pres(false)
		}
		for i := 0; i < 5000; i++ {
			pres(true)
		}
	}*/

}

func pres(blink bool) {
	if blink {
		Light(true, true)
	} else {
		Light(false, true)
	}

}

func Off() {
	leds := machine.PA12
	leds.Configure(machine.PinConfig{Mode: machine.PinOutput})
	leds.Low()
}

func On() {
	leds := machine.PA12
	leds.Configure(machine.PinConfig{Mode: machine.PinOutput})
	leds.High()
}

func minLight() {
	for {
		On()
		time.Sleep(time.Second / 1000)
		Off()
		time.Sleep(time.Second / 100)
	}
}
func Light(isReduce bool, isOn bool) {
	rate := time.Second / 5000
	leds := machine.PA12
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
