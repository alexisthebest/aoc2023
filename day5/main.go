package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const dir = "input"

const files = "seed-to-soil,soil-to-fertilizer,fertilizer-to-water,water-to-light,light-to-temperature,temperature-to-humidity,humidity-to-location"

func processLine(numbers []int, cursor int) (bool, int) {
	sourceRangeStart := numbers[1]
	rangeLength := numbers[2]

	if cursor < sourceRangeStart || cursor > (sourceRangeStart+rangeLength) {
		return false, cursor
	}
	destinationRangeStart := numbers[0]

	return true, destinationRangeStart + (cursor - sourceRangeStart)
}

func readLine(line string) []int {
	stringNumbers := strings.Split(line, " ")
	intNumbers := make([]int, len(stringNumbers))

	for i, s := range stringNumbers {
		n, err := strconv.Atoi(s)
		if err != nil {
			os.Exit(1)
		}
		intNumbers[i] = n
	}
	return intNumbers
}

func readFile(name string) *bufio.Scanner {
	path := fmt.Sprintf("%s/%s.txt", dir, name)

	r, err := os.Open(path)
	if err != nil {
		fmt.Println(path)
		os.Exit(1)
	}
	return bufio.NewScanner(r)
}

func readSeeds() []int {
	scanner := readFile("seeds")
	var intNumbers []int

	for scanner.Scan() {
		intNumbers = append(intNumbers, readLine(scanner.Text())...)
	}
	return intNumbers
}

func main() {
	lowestNumber := math.MaxInt

	for _, seed := range readSeeds() {
		cursor := seed

		for _, key := range strings.Split(files, ",") {
			fmt.Println("Seed", seed, "File", key)
			scanner := readFile(key)

			for scanner.Scan() {
				success, newCursor := processLine(readLine(scanner.Text()), cursor)
				if success {
					cursor = newCursor
					break
				}
			}
		}
		fmt.Println("Location", cursor, "\n")
		if cursor < lowestNumber {
			lowestNumber = cursor
		}
	}
	fmt.Println(lowestNumber)
}
