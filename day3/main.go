package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type info struct {
	periods []bool
	numbers []bool
}

func getInfo(line string) info {
	periods := make([]bool, len(line))
	numbers := make([]bool, len(line))

	for i, c := range line {
		periods[i] = false
		numbers[i] = false

		if c > 47 && c < 58 {
			numbers[i] = true
		} else if c == 46 {
			periods[i] = true
		}
	}
	return info{periods, numbers}
}

func symbolNextTo(index int, particular info, max int) bool {
	if !particular.numbers[index] && !particular.periods[index] {
		// Must be a symbol
		return true
	}
	if max > 0 {
		if index > 0 && symbolNextTo(index-1, particular, -1) {
			// Symbol to the left
			return true
		}
		if index < max && symbolNextTo(index+1, particular, -1) {
			// Symbol to the right
			return true
		}
	}
	return false
}

func adjacentToSymbol(index int, this *info, above *info, below *info, max int) bool {
	if symbolNextTo(index, *this, max) {
		return true
	}
	if above != nil && symbolNextTo(index, *above, max) {
		return true
	}
	if below != nil && symbolNextTo(index, *below, max) {
		return true
	}
	return false
}

func recalculateTotal(adjacent bool, currentNumber string, total int64) int64 {
	if adjacent && len(currentNumber) > 0 {
		n, _ := strconv.ParseInt(currentNumber, 0, 64)
		return total + n
	}
	return total
}

func getTotal(line string, this info, above *info, below *info) int64 {
	var total int64 = 0

	max := len(line) - 1

	currentNumber := ""
	adjacent := false

	for i, x := range line {
		if this.numbers[i] {
			// Working on a number
			currentNumber = currentNumber + string(x)
			adjacent = adjacent || adjacentToSymbol(i, &this, above, below, max)
		} else {
			// Number finished
			total = recalculateTotal(adjacent, currentNumber, total)
			currentNumber = ""
			adjacent = false
		}
	}
	return recalculateTotal(adjacent, currentNumber, total)
}

func main() {
	r, err := os.Open("input.txt")

	if err != nil {
		os.Exit(1)
	}

	scanner := bufio.NewScanner(r)

	lines := []string{}
	infos := []info{}

	for scanner.Scan() {
		line := scanner.Text()
		infos = append(infos, getInfo(line))
		lines = append(lines, line)
	}

	var total int64 = 0

	for i, line := range lines {
		fmt.Println(i, line)

		this := infos[i]

		var above *info = nil
		if i > 0 {
			above = &infos[i-1]
		}
		var below *info = nil
		if i < len(lines)-1 {
			below = &infos[i+1]
		}
		total = total + getTotal(line, this, above, below)

		fmt.Println("")
		fmt.Println("")

	}
	fmt.Println(total)
}
