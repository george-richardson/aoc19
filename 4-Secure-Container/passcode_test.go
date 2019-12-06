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
	testValidatePasscode(false, 111111, t)
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

func TestValidatePasscode123444(t *testing.T) {
	testValidatePasscode(false, 123444, t)
}

func TestValidatePasscode111122(t *testing.T) {
	testValidatePasscode(true, 111122, t)
}

func TestValidatePasscode112345(t *testing.T) {
	testValidatePasscode(true, 112345, t)
}

func TestValidatePasscode111345(t *testing.T) {
	testValidatePasscode(false, 111345, t)
}

func TestValidatePasscode666666(t *testing.T) {
	testValidatePasscode(false, 666666, t)
}

func TestValidatePasscode654321(t *testing.T) {
	testValidatePasscode(false, 654321, t)
}

func TestValidatePasscode123356(t *testing.T) {
	testValidatePasscode(true, 123356, t)
}

func TestValidatePasscode120366(t *testing.T) {
	testValidatePasscode(false, 120366, t)
}

func TestValidatePasscode112233(t *testing.T) {
	testValidatePasscode(true, 112233, t)
}

func TestValidatePasscode111233(t *testing.T) {
	testValidatePasscode(true, 111233, t)
}

func TestValidatePasscode588999(t *testing.T) {
	testValidatePasscode(true, 588999, t)
}