package main

import (
	"bufio"
	"fmt"
	"os"
)

func transform(grid [][]rune) [][]rune {
	result := make([][]rune, len(grid[0]))

	for y := len(grid) - 1; y >= 0; y-- {
		row := grid[y]

		for x, col := range row {
			result[x] = append(result[x], col)
		}
	}
	return result
}

func shift(row []rune) bool {
	for i := len(row) - 2; i >= 0; i-- {
		if row[i] == 'O' && row[i+1] == '.' {
			row[i], row[i+1] = row[i+1], row[i]
			return true
		}
	}
	return false
}

func shiftAll(row []rune) {
	for shift(row) {
		// Nothing
	}
}

func main() {
	r, err := os.Open("input.txt")
	if err != nil {
		os.Exit(1)
	}
	scanner := bufio.NewScanner(r)
	count := 0

	grid := [][]rune{}
	for scanner.Scan() {
		grid = append(grid, []rune(scanner.Text()))
	}

	transformed := transform(grid)

	for _, row := range transformed {
		fmt.Println(string(row))
		shiftAll(row)
	}
	fmt.Println("\n")
	for _, row := range transformed {
		fmt.Println(string(row))
	}
	fmt.Println("\n")

	for x, _ := range transformed[0] {
		n := 0
		for _, row := range transformed {
			if row[x] == 'O' {
				n++
			}
		}
		count += (x + 1) * n
	}

	fmt.Println(count)
}
