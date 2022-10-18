package Prog

import (
	"testing"
)

func TestGetCurrentSensorData(t *testing.T) {
	cases := []struct {
		name   string
		module int
		want   int
	}{
		{"Module 1", 1, 230},
		{"Module 2", 2, 12},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := GetCurrentSensorData(c.module)
			if got != c.want {
				t.Fatalf("got %d, want %d", got, c.want)
			}
		})
	}
}
func TestReduceLightIntensity(t *testing.T) {
	cases := []struct {
		name     string
		isreduce bool
		want     bool
	}{
		{"Change intensity to 50%", true, true},
		{"Change intensity to 20%", false, true},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := Light(c.isreduce, true)
			if got != c.want {
				t.Fatalf("got %t, want %t", got, c.want)
			}
		})
	}
}
