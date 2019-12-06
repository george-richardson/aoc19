package main

import (
	"reflect"
	"testing"
)

func testColumnSplit(expected []int, input int, t *testing.T) {
	actual := columnSplit(input)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("got '%v' expected '%v'", actual, expected)
	}
}

func TestColumnSplit1(t *testing.T) {
	testColumnSplit([]int {1, 2, 3}, 321, t)
}

func testValidatePasscode(expected bool, input int, t *testing.T) {
	actual := validatePasscode(input)
	if actual != expected {
		t.Errorf("got '%v' expected '%v'", actual, expected)
	}
}

func TestValidatePasscode111111(t *testing.T) {
	testValidatePasscode(true, 111111, t)
}

func TestValidatePasscode122345(t *testing.T) {
	testValidatePasscode(true, 122345, t)
}

func TestValidatePasscode223450(t *testing.T) {
	testValidatePasscode(false, 223450, t)
}

func TestValidatePasscode123789(t *testing.T) {
	testValidatePasscode(false, 123789, t)
}