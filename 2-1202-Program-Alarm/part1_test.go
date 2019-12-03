package main

import (
	"reflect"
	"testing"
)

func executeProgramStepTest(expectedAnswer int, expectedState []int, expectedHalt bool, program []int, address int, t *testing.T) {
	originalProgram := program
	actual, halting, err := executeProgramStep(&program, address)
	if err != nil {
		t.Error(err)
	}

	if actual != expectedAnswer {
		t.Errorf("Output '%v' does not match expected '%v' for input program '%v' starting at address '%v'", actual, expectedAnswer, originalProgram, address)
	}

	if !reflect.DeepEqual(program, expectedState) {
		t.Errorf("Output '%v' does not match expected '%v' for input program '%v' starting at address '%v'", program, expectedState, originalProgram, address)
	}

	if halting != expectedHalt {
		t.Errorf("Output '%v' does not match expected '%v' for input program '%v' starting at address '%v'", halting, expectedHalt, originalProgram, address)
	}
}

func executeProgramTest(expected string, program string, t *testing.T) {
	actual, err := executeProgram(program)
	if err != nil {
		t.Error(err)
	}

	if actual != expected {
		t.Errorf("Output '%v' does not match expected '%v' for input program '%v'", actual, expected, program)
	}
}

func TestStep1(t *testing.T) {
	executeProgramStepTest(2, []int{2, 0, 0, 0, 99}, false, []int{1, 0, 0, 0, 99}, 0, t)
}

func TestStep2(t *testing.T) {
	executeProgramStepTest(99, []int{1, 0, 0, 0, 99}, true, []int{1, 0, 0, 0, 99}, 4, t)
}

func TestStep3(t *testing.T) {
	executeProgramStepTest(6, []int{2, 3, 0, 6, 99}, false, []int{2, 3, 0, 3, 99}, 0, t)
}

func TestStep4(t *testing.T) {
	executeProgramStepTest(9801, []int{2, 4, 4, 5, 99, 9801}, false, []int{2, 4, 4, 5, 99, 0}, 0, t)
}

func TestStep5(t *testing.T) {
	executeProgramStepTest(99, []int{99}, true, []int{99}, 0, t)
}

func TestProgram1(t *testing.T) {
	executeProgramTest("30,1,1,4,2,5,6,0,99", "1,1,1,4,99,5,6,0,99", t)
}
