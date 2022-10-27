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

		}
		uart.Write([]byte("AT+MSG=\"Joined\"\r\n"))
	}

	for {
		On()
		uart.Write([]byte("AT+MSG=\"Allumé\"\r\n"))
		time.Sleep(time.Second * 10)
		minLight(365 * time.Second)
		uart.Write([]byte("AT+MSG=\"baisse\"\r\n"))
		Off()
		uart.Write([]byte("AT+MSG=\"eteinte\"\r\n"))
		time.Sleep(time.Second * 105)
	}

}

func Off() {
	leds := machine.PC13
	leds.Configure(machine.PinConfig{Mode: machine.PinOutput})
	leds.Low()
}

func On() {
	leds := machine.PC13
	leds.Configure(machine.PinConfig{Mode: machine.PinOutput})
	leds.High()
}

func minLight(dur time.Duration) {
	start := time.Now()
	for {
		if time.Now().Sub(start) > dur {
			break
		}

		On()
		time.Sleep(time.Second / 1000)
		Off()
		time.Sleep(time.Second / 100)

	}
}
