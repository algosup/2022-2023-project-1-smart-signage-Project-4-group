package main

import (
	"machine"
	"time"
)

func main() {
	machine.UART1.Configure(machine.UARTConfig{BaudRate: 9600, TX: machine.PA2, RX: machine.PA3})
	_, join := machine.UART1.Write([]byte("AT+JOIN\r\n"))
	rate2 := time.Second * 10
	time.Sleep(rate2)
	//_, read := machine.UART1.Read(join)
	if join != nil {
		time.Sleep(rate2)
		machine.UART1.Write([]byte("AT+MSG='Joined'\r\n"))
		println(join)
	}

	for i := 0; i < 100; i++ {
		zqsedrftgy(false)
	}
	for i := 0; i < 100; i++ {
		zqsedrftgy(true)
	}

}
func zqsedrftgy(CustomerOn bool) {
	if CustomerOn {
		//_, join := machine.UART1.Write([]byte("AT+JOIN\r\n"))

		//println(join)
		Light(true, true)

		//_, msg := machine.UART1.Write([]byte("AT+MSG='allumé'\r\n"))
		//_, read := machine.UART1.Read(msg)
		//println(msg)
		//println(read)
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
