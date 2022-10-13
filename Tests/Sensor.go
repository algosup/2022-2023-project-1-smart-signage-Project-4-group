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

// Having intensity change by a percentage given in input
func ReduceLightIntensity(intensity int) bool {
	var time int
	if intensity < 100 && intensity > 0{
		time = 50 * (intensity / 100)
		// put a machine.blink(time)
		time += 0
		return true
	} else {
		return false
	}
}