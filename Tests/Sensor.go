package Prog

func GetCurrentSensorData(module int) int {
	var out int
	if module == 1 {
		out = ReadByPin(1)
	} else {
		out = ReadByPin(2)
	}
	return out
}
func ReadByPin(pin int) int {
	var out int
	if pin == 1 {
		// put a machine.pin(module's pin
		out = 230
	} else {
		// put a machine.pin(module's pin)
		out = 12
	}
	return out
}
