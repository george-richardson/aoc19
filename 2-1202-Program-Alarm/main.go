package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input1202.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	program := scanner.Text()

	var answer int

outer:
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			_, output, err := executeProgramWithParams(program, i, j)
			log.Printf("p1: %v p2: %v o: %v", i, j, output)
			if err != nil {
				log.Fatal(err)
			}
			if output == 19690720 {
				answer = 100*i + j
				break outer
			}
		}
	}

	println(answer)
}

func executeProgramStep(program *[]int, address int) (int, bool, error) {
	operator := (*program)[address]
	var l, r, o, result int
	switch operator {
	case 99:
		return 99, true, nil
	case 1:
		l, r, o = retrieveOperators(*program, address)
		result = l + r
	case 2:
		l, r, o = retrieveOperators(*program, address)
		result = l * r
	default:
		err := fmt.Errorf("unknown operator '%v' found at position '%v' of program '%v'", operator, address, *program)
		return -1, true, err
	}
	(*program)[o] = result
	return result, false, nil
}

func executeProgram(program string) (string, error) {
	progArr, err := strToIntArr(program)
	halt := false
	for address := 0; !halt && err == nil; address += 4 {
		_, halt, err = executeProgramStep(&progArr, address)
	}
	if err != nil {
		return "err", err
	}
	return intArrToStr(progArr), err
}

func executeProgramWithParams(program string, param1 int, param2 int) (string, int, error) {
	progArr, err := strToIntArr(program)
	progArr[1] = param1
	progArr[2] = param2
	halt := false
	for address := 0; !halt && err == nil; address += 4 {
		_, halt, err = executeProgramStep(&progArr, address)
	}
	if err != nil {
		return "err", -1, err
	}
	return intArrToStr(progArr), progArr[0], err
}

func strToIntArr(str string) ([]int, error) {
	var intArr = []int{}
	for _, i := range strings.Split(str, ",") {
		j, err := strconv.Atoi(i)
		if err != nil {
			return nil, err
		}
		intArr = append(intArr, j)
	}
	return intArr, nil
}

func intArrToStr(intArr []int) string {
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(intArr)), ","), "[]")
}

func retrieveOperators(program []int, address int) (int, int, int) {
	return program[program[address+1]], program[program[address+2]], program[address+3]
}
