package main

import (
	"time"
)

func TimeRules() {
	for {
		if time.Now().Hour() >= 6 && time.Now().Hour() <= 1 {
			Light(false, true)
		} else {
			Light(false, false)
		}
	}
}
