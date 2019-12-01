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

func main() {
	file, err := os.Open("part1_input.txt")
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
		sum += fuelRequired(i)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	println(sum)
}