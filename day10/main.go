package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

const UP = 2
const DOWN = 4
const LEFT = 6
const RIGHT = 8

type movement struct {
	indicator int
	x         int
	y         int
}

func move(position rune) []movement {
	up := movement{UP, 0, -1}
	down := movement{DOWN, 0, 1}
	left := movement{LEFT, -1, 0}
	right := movement{RIGHT, 1, 0}

	switch position {
	case '|':
		return []movement{up, down}
	case '-':
		return []movement{left, right}
	case 'L':
		return []movement{up, right}
	case 'J':
		return []movement{up, left}
	case '7':
		return []movement{left, down}
	case 'F':
		return []movement{down, right}
	case '.':
		return []movement{}
	case 'S':
		return []movement{up, down, left, right}
	}
	return []movement{}
}

func permitted(moved movement, to rune) bool {
	if to == '.' || to == 'S' {
		return false
	}
	switch to {
	case '|':
		return moved.indicator == UP || moved.indicator == DOWN
	case '-':
		return moved.indicator == LEFT || moved.indicator == RIGHT
	case 'L':
		return moved.indicator == DOWN || moved.indicator == LEFT
	case 'J':
		return moved.indicator == DOWN || moved.indicator == RIGHT
	case '7':
		return moved.indicator == RIGHT || moved.indicator == UP
	case 'F':
		return moved.indicator == UP || moved.indicator == LEFT
	}
	return false
}

func walk(x int, y int, grid [][]rune, distances [][]int, count int) {
	position := grid[y][x]
	distances[y][x] = count

	fmt.Println(count, "at", string(position), "->", x, y)

	for _, movement := range move(position) {
		x2 := x + movement.x
		y2 := y + movement.y

		if y2 > -1 && y2 < len(grid) && x2 > -1 && x2 < len(grid[y2]) {
			currentDistance := distances[y2][x2]

			if permitted(movement, grid[y2][x2]) && (currentDistance == 0 || currentDistance > count) {
				walk(x2, y2, grid, distances, count+1)
			}
		}
	}
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

	sx, sy := 0, 0
	distances := make([][]int, len(grid))

	for y, row := range grid {
		distances[y] = make([]int, len(row))

		for x, col := range row {
			fmt.Print(string(col))
			if col == 'S' {
				sx, sy = x, y
			}
		}
		fmt.Println()
	}

	fmt.Println(sx, sy)

	walk(sx, sy, grid, distances, 0)

	highest := math.MinInt
	for _, row := range distances {
		fmt.Println(row)

		for _, col := range row {
			if col > highest {
				highest = col
			}
		}
	}
	fmt.Println(highest)
}
