package main

import (
	"reflect"
	"testing"
)

// Test manhattan distance
func testManhattanDistance(expected int, p1, p2 point, t *testing.T) {
	actual := p1.distanceTo(p2)
	if expected != actual {
		t.Errorf("calculated manhattan distance '%v' does not match expected '%v'", expected, actual)
	}
}

func TestManhattan1(t *testing.T) {
	testManhattanDistance(0, point{x:1, y:1}, point{x:1, y:1}, t)
}

func TestManhattan2(t *testing.T) {
	testManhattanDistance(1, point{x:1, y:1}, point{x:1, y:2}, t)
}

func TestManhattan3(t *testing.T) {
	testManhattanDistance(2, point{x:1, y:1}, point{x:2, y:2}, t)
}

func TestManhattan4(t *testing.T) {
	testManhattanDistance(3, point{x:3, y:1}, point{x:1, y:2}, t)
}

// Test parser
func testParseTranslation(expected translation, token string, t *testing.T) {
	actual, err := parseTranslation(token)
	if err != nil {
		t.Error(err)
	}
	if actual != expected {
		t.Errorf("actual '%v' does not match expected '%v'", actual, expected)
	}
}

func TestParseTranslation1(t *testing.T) {
	testParseTranslation(translation{axis:X, magnitude:1}, "R1", t)
}

func TestParseTranslation2(t *testing.T) {
	testParseTranslation(translation{axis:Y, magnitude:2}, "U2", t)
}

func TestParseTranslation3(t *testing.T) {
	testParseTranslation(translation{axis:X, magnitude:-11}, "L11", t)
}

func TestParseTranslation4(t *testing.T) {
	testParseTranslation(translation{axis:Y, magnitude:-1337}, "D1337", t)
}

func TestParseTranslationErr(t *testing.T) {
	_, err := parseTranslation("foobar")
	if err == nil {
		t.Error("Expected error when parsing invalid string")
	}
}

// Test create path
func testCreatePath(expected []point, point point, translation translation, t *testing.T) {
	actual := point.createPath(translation)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("actual '%v' does not match expected '%v'", actual, expected)
	}
}

func TestCreatePath1(t *testing.T) {
	testCreatePath([]point{{x:0, y:1}}, point{x:0, y:0}, translation{magnitude:1, axis:Y}, t)
}

func TestCreatePath2(t *testing.T) {
	testCreatePath([]point{{x:2, y:0}, {x:1, y:0}, {x:0, y:0}}, point{x:3, y:0}, translation{magnitude:-3, axis:X}, t)
}

// Test wire append
func testWireAppend(expected wire, wire wire, translation translation, t *testing.T) {
	wire.appendTranslation(translation)
	if expected.origin != wire.origin || expected.head != wire.head || !reflect.DeepEqual(wire.path, expected.path) {
		t.Errorf("actual '%v' does not match expected '%v'", wire, expected)
	}
}

func TestWireAppend1(t *testing.T) {
	expected := wire{
		path:   []point{{x:2, y:0}, {x:1, y:0}, {x:0, y:0}},
		head:   point{x:0, y:0},
		origin: point{x:3, y:0},
	}
	sut := wire{
		path:   []point{},
		head:   point{x:3, y:0},
		origin: point{x:3, y:0},
	}
	testWireAppend(expected, sut, translation{magnitude:-3, axis:X}, t)
}

func TestWireAppend2(t *testing.T) {
	expected := wire{
		path:   []point{{x:2, y:0}, {x:1, y:0}, {x:0, y:0}, {x:0, y:1}},
		head:   point{x:0, y:1},
		origin: point{x:3, y:0},
	}
	sut := wire{
		path:   []point{{x:2, y:0}, {x:1, y:0}, {x:0, y:0}},
		head:   point{x:0, y:0},
		origin: point{x:3, y:0},
	}
	testWireAppend(expected, sut, translation{magnitude:1, axis:Y}, t)
}