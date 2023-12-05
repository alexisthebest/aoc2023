package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

const dir = "input"

const files = "seed-to-soil,soil-to-fertilizer,fertilizer-to-water,water-to-light,light-to-temperature,temperature-to-humidity,humidity-to-location"

type seed struct {
	start int
	end   int
}

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

func readSeeds() ([]seed, int) {
	scanner := readFile("seeds")
	var intNumbers []int

	for scanner.Scan() {
		intNumbers = append(intNumbers, readLine(scanner.Text())...)
	}
	result := []seed{}

	for i := 0; i < len(intNumbers); i += 2 {
		n := intNumbers[i]
		r := intNumbers[i+1]
		result = append(result, seed{n, n + r})
	}

	sort.SliceStable(result, func(i, j int) bool {
		return result[i].start < result[j].start
	})
	return result, result[len(result)-1].end
}

func readData() [][][]int {
	names := strings.Split(files, ",")
	store := make([][][]int, len(names))

	for i, key := range names {
		scanner := readFile(key)
		rows := [][]int{}
		for scanner.Scan() {
			rows = append(rows, readLine(scanner.Text()))
		}
		store[i] = rows
	}
	return store
}

func main() {
	lowestNumber := math.MaxInt

	data := readData()

	seeds, max := readSeeds()

	fmt.Println(seeds)

	last := 0
	for _, s := range seeds {
		fmt.Println("Running", s.start, "to", s.end)
		start := s.start
		if last > start {
			start = last
		}
		last = s.end

		for seed := start; seed <= last; seed++ {
			cursor := seed

			for _, file := range data {
				for _, line := range file {
					success, newCursor := processLine(line, cursor)
					if success {
						cursor = newCursor
						break
					}
				}
			}
			if cursor < lowestNumber {
				fmt.Println("Seed", seed, "of", max, "Lowest is now", cursor, "\n")
				lowestNumber = cursor
			}
		}
	}
	fmt.Println(lowestNumber)
}
