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

	output, err := executeProgram(program)
	if err != nil {
		log.Fatal(err)
	}

	intArr, err := strToIntArr(output)
	if err != nil {
		log.Fatal(err)
	}

	println(intArr[0])
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
