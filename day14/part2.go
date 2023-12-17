package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

func rotateRight(grid [][]rune) [][]rune {
	result := make([][]rune, len(grid[0]))

	for y := len(grid) - 1; y >= 0; y-- {
		row := grid[y]

		for x, col := range row {
			result[x] = append(result[x], col)
		}
	}
	return result
}

func rotateLeft(grid [][]rune) [][]rune {
	result := make([][]rune, len(grid[0]))

	for y, row := range grid {
		for x := len(row) - 1; x >= 0; x-- {
			index := len(row) - 1 - x
			result[index] = append(result[index], grid[y][x])
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

func shiftRow(row []rune) bool {
	shifted := false
	for shift(row) {
		shifted = true
	}
	return shifted
}

func shiftAll(grid [][]rune) bool {
	shifted := false
	for _, row := range grid {
		shifted = shiftRow(row) || shifted
	}
	return shifted
}

func score(grid [][]rune) int {
	count := 0
	for y := len(grid) - 1; y >= 0; y-- {
		n := 0
		for _, col := range grid[y] {
			if col == 'O' {
				n++
			}
		}
		count += (len(grid) - y) * n
	}
	return count
}

func print(label string, grid [][]rune) {
	fmt.Println(label)
	for _, row := range grid {
		fmt.Println(string(row))
	}
	fmt.Println("\n")
}

func clear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func addToCounts(new int, counts []int) []int {
	for _, n := range counts {
		if n == new {
			return counts
		}
	}
	return append(counts, new)
}

func main() {
	r, err := os.Open("input.txt")
	if err != nil {
		os.Exit(1)
	}
	scanner := bufio.NewScanner(r)

	grid := [][]rune{}
	for scanner.Scan() {
		grid = append(grid, []rune(scanner.Text()))
	}

	print("raw", grid)

	for j := 0; j < 1000; j++ {
		for i := 0; i < 4; i++ {
			grid = rotateRight(grid)
			shiftAll(grid)
		}
	}

	fmt.Println(score(grid))
}
