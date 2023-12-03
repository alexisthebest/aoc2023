package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type info struct {
	numbers []bool
	gears   []bool
}

type gearmatch struct {
	num int64
	row int
	col int
}

func getInfo(line string) info {
	numbers := make([]bool, len(line))
	gears := make([]bool, len(line))

	for i, c := range line {
		numbers[i] = false
		gears[i] = false

		if c > 47 && c < 58 {
			numbers[i] = true
		} else if c == 42 {
			gears[i] = true
		}
	}
	return info{numbers, gears}
}

func symbolNextTo(index int, particular info, max int) (bool, int) {
	if particular.gears[index] {
		// Must be a symbol
		return true, 0
	}
	if max > 0 {
		if index > 0 {
			found, _ := symbolNextTo(index-1, particular, -1)
			if found {
				// Symbol to the left
				return true, -1
			}
		}
		if index < max {
			found, _ := symbolNextTo(index+1, particular, -1)
			if found {
				// Symbol to the right
				return true, 1
			}
		}
	}
	return false, 0
}

func adjacentToGear(num int64, row int, col int, this *info, above *info, below *info, max int) []gearmatch {
	all := []gearmatch{}

	found, adjustment := symbolNextTo(col, *this, max)
	if found {
		all = append(all, gearmatch{num, row, col + adjustment})
	}
	if above != nil {
		found, adjustment = symbolNextTo(col, *above, max)
		if found {
			all = append(all, gearmatch{num, row - 1, col + adjustment})
		}
	}
	if below != nil {
		found, adjustment = symbolNextTo(col, *below, max)
		if found {
			all = append(all, gearmatch{num, row + 1, col + adjustment})
		}
	}
	return all
}

func recalculateTotal(subTotal []gearmatch, currentNumber string, total []gearmatch) []gearmatch {
	if len(subTotal) > 0 && len(currentNumber) > 0 {
		n, _ := strconv.ParseInt(currentNumber, 0, 64)
		total = append(total, gearmatch{n, subTotal[0].row, subTotal[0].col})
	}
	return total
}

func getGearIndices(row int, line string, this info, above *info, below *info) []gearmatch {
	total := []gearmatch{}
	subTotal := []gearmatch{}

	max := len(line) - 1

	currentNumber := ""

	for col, x := range line {
		if this.numbers[col] {
			// Working on a number
			currentNumber = currentNumber + string(x)

			for _, found := range adjacentToGear(-1, row, col, &this, above, below, max) {
				subTotal = append(subTotal, found)
			}
		} else {
			// Number finished
			total = recalculateTotal(subTotal, currentNumber, total)
			currentNumber = ""
			subTotal = []gearmatch{}
		}
	}
	return recalculateTotal(subTotal, currentNumber, total)
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

	numbersFoundAdjacentToGears := []gearmatch{}
	numberAdjacentToGearAtIndex := make([][]int64, len(lines))

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
		for _, gi := range getGearIndices(i, line, this, above, below) {
			numbersFoundAdjacentToGears = append(numbersFoundAdjacentToGears, gi)
		}
		numberAdjacentToGearAtIndex[i] = make([]int64, len(line))
	}

	fmt.Println(numbersFoundAdjacentToGears)

	var sum int64 = 0
	for _, gearmatch := range numbersFoundAdjacentToGears {
		existing := numberAdjacentToGearAtIndex[gearmatch.row][gearmatch.col]
		if existing == 0 {
			numberAdjacentToGearAtIndex[gearmatch.row][gearmatch.col] = gearmatch.num
		} else {
			fmt.Println("sum of", gearmatch.num, existing)

			sum = sum + (gearmatch.num * existing)
		}
	}

	fmt.Println(sum)
}
