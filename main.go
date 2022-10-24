package main

import (
	"fmt"
	"machine"
	"time"
)

func main() {
	CustomerOn := true
	if CustomerOn {
		Light(true, true)
	} else {
		Light(false, false)
	}
}

func Light(isReduce bool, isOn bool) {
	machine.UART0.Configure(machine.UARTConfig{BaudRate: 9600, TX: machine.D1, RX: machine.D0})
	_, err := machine.UART0.Write([]byte("AT+JOIN\r\n"))
	_, env := machine.UART0.Write([]byte("AT+MSG='LED allumé'\r\n"))
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

func Sleep(duration int) {
	time.Sleep(time.Second * time.Duration(duration))
}

func printBuf(b []byte) {
	for _, val := range b {
		fmt.Printf("%x ", val)
	}
}

/*func main2() {

	options := serial.OpenOptions{
		PortName:              "/dev/ttyX",
		BaudRate:              9600,
		DataBits:              8,
		StopBits:              1,
		MinimumReadSize:       0,
		InterCharacterTimeout: 50,
	}

	port, err := serial.Open(options)
	if err != nil {
		log.Printf("port.Read: %v", err)
		return
	}

	// Make sure to close it later.
	defer port.Close()

	var s string = `AT+MSG="Bonjo"\r`
	b := []byte(s)

	n, err := port.Write(b)
	if err != nil {
		log.Printf("port.Write: %v", err)
	}

	log.Println("Written bytes: ", n)

	//Sleep(1)

	res := make([]byte, 64)
	n, err = port.Read(res)
	if err != nil && err != io.EOF {
		log.Printf("port.Read: %v", err)
	}

	log.Println("READ bytes: ", n)

	printBuf(res)

}
*/
