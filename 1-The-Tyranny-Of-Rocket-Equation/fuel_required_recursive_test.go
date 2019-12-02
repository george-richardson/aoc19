package main

import (
	"testing"
)

func fuelRequiredRecursiveTest(expected int, mass int, t *testing.T) {
	actual := fuelRequiredRecursive(mass)
	if actual != expected {
		t.Errorf("'%v' does not match expected '%v' for input mass '%v'", actual, expected, mass)
	}
}

func TestMass14Recursive(t *testing.T) {
	fuelRequiredRecursiveTest(2, 14, t)
}

func TestMass1969Recursive(t *testing.T) {
	fuelRequiredRecursiveTest(966, 1969, t)
}

func TestMass100756Recursive(t *testing.T) {
	fuelRequiredRecursiveTest(50346, 100756, t)
}
