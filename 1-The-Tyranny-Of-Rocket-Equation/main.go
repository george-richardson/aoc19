package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func fuelRequired(mass int) int {
	return (mass / 3) - 2
}

func fuelRequiredRecursive(mass int) int {
	sum := 0
	for tmp := fuelRequired(mass); tmp >= 0; tmp = fuelRequired(tmp) {
		sum += tmp
	}
	return sum
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sum := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		sum += fuelRequiredRecursive(i)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	println(sum)
}
