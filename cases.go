package main

import (
	"time"
)

func TimeRules() { // Function used to shut down the light during 1AM to 6AM
	for {
		if time.Now().Hour() >= 6 || time.Now().Hour() <= 1 {
			Light(false, true)
		} else {
			Light(false, false)
		}
	}
}
