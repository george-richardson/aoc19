package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// Data structures
type point struct {
	x int
	y int
}

type axis = bool
const (
	X axis = true
	Y axis = false
)

type translation struct {
	axis axis
	magnitude int
}

type wire struct {
	path []point
	head point
	origin point
}

type intersection struct {
	point point
	stepsFromOrigin int
	distanceFromOrigin int
}

// Translation methods
func parseTranslation(token string) (translation, error) {
	result := translation{}
	if len(token) < 2 {
		return result, fmt.Errorf("'%v' is not long enough to be a translation token", token)
	}

	magnitude, err := strconv.Atoi(token[1:])
	if err != nil {
		return result, fmt.Errorf("'%v' is a not a valid integer in token '%v'", token[1:], token)
	}
	result.magnitude = magnitude
	directionChar := token[0]

	switch directionChar {
	case 'U':
		result.magnitude = magnitude
		result.axis = Y
	case 'D':
		result.magnitude = magnitude * -1
		result.axis = Y
	case 'R':
		result.magnitude = magnitude
		result.axis = X
	case 'L':
		result.magnitude = magnitude * -1
		result.axis = X
	default:
		return result, fmt.Errorf("unknown direction '%v' in translation token '%v'", directionChar, token)
	}

	return result, nil
}

// Point methods
func (p1 point) distanceTo(p2 point) int {
	return int(math.Abs(float64(p1.x - p2.x))) + int(math.Abs(float64(p1.y - p2.y)))
}

func (p point) createPath(t translation) []point {
	var path []point
	step := 1
	if t.magnitude < 0 {
		step = -1
	}

	length := 0
	for length != t.magnitude {
		length += step
		if t.axis == X {
			path = append(path, point{x: p.x + length, y: p.y})
		} else {
			path = append(path, point{x: p.x, y: p.y + length})
		}
	}

	return path
}

// Wire methods
func (wire *wire) appendTranslation(translation translation) {
	wire.path = append(wire.path, wire.head.createPath(translation)...)
	wire.head = wire.path[len(wire.path) -1]
}

func (w1 wire) getIntersections(w2 wire) ([]intersection, error) {
	var intersections []intersection
	for i, p1 := range w1.path {
		for j, p2 := range w2.path {
			if p1 == p2 {
				intersections = append(intersections, intersection{
					point:              p1,
					stepsFromOrigin:    i + j + 2, //Add two because the array is zero based
					distanceFromOrigin: w1.origin.distanceTo(p1),
				})
			}
		}
	}
	if len(intersections) == 0 {
		return intersections, fmt.Errorf("unable to find intersections")
	}
	return intersections, nil
}

func (w1 wire) findManhattanClosestIntersection(w2 wire) (intersection, error) {
	intersections, err := w1.getIntersections(w2)
	if err != nil {
		return intersection{}, err
	}

	closest := intersections[0]
	for _, intersection := range intersections {
		if intersection.distanceFromOrigin < closest.distanceFromOrigin {
			closest = intersection
		}
	}

	return closest, nil
}

func (w1 wire) findStepsClosestIntersection(w2 wire) (intersection, error) {
	intersections, err := w1.getIntersections(w2)
	if err != nil {
		return intersection{}, err
	}

	closest := intersections[0]
	for _, intersection := range intersections {
		if intersection.stepsFromOrigin < closest.stepsFromOrigin {
			closest = intersection
		}
	}

	return closest, nil
}

func parseWire(str string) wire {
	tokens := strings.Split(str, ",")
	var translations []translation
	for _, token := range tokens {
		translation, _ := parseTranslation(token)
		translations = append(translations, translation)
	}

	wire := wire{
		path:   []point{},
		head:   point{x: 0, y: 0},
		origin: point{x: 0, y: 0},
	}

	for _, translation := range translations {
		wire.appendTranslation(translation)
	}

	return wire
}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	wire1 := parseWire(scanner.Text())
	scanner.Scan()
	wire2 := parseWire(scanner.Text())

	intersection, _ := wire1.findStepsClosestIntersection(wire2)

	answer := intersection.stepsFromOrigin
	fmt.Println(answer)
}