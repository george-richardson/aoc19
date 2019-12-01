package main

import (
	"testing"
)

func fuelRequiredTest(expected int, mass int, t *testing.T) {
	actual := fuelRequired(mass)
	if actual != expected {
		t.Errorf("'%v' does not match expected '%v' for input mass '%v'", actual, expected, mass)
	}
}

func TestMass12(t *testing.T) {
	fuelRequiredTest(2, 12, t)
}

func TestMass14(t *testing.T) {
	fuelRequiredTest(2, 12, t)
}

func TestMass1969(t *testing.T) {
	fuelRequiredTest(654, 1969, t)
}

func TestMass100756(t *testing.T) {
	fuelRequiredTest(33583, 100756, t)
}
