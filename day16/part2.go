package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type coordinates struct {
	x int
	y int
}

type dir int

const up dir = 1

const down dir = 2

const left dir = 4

const right dir = 8

func print(grid [][]rune) {
	for _, row := range grid {
		for _, col := range row {
			fmt.Print(string(col))
		}
		fmt.Println()
	}
	fmt.Println()
}

func count(grid [][]dir) int {
	count := 0
	for _, row := range grid {
		for _, col := range row {
			if col > 0 {
				count++
			}
		}
	}
	return count
}

func possible(grid [][]rune, position coordinates) bool {
	return position.y >= 0 && position.y < len(grid) && position.x >= 0 && position.x < len(grid[position.y])
}

func cpy(been [][]dir) [][]dir {
	new := make([][]dir, len(been))
	for y, row := range been {
		new[y] = make([]dir, len(row))
		copy(new[y], row)
	}
	return new
}

func walk(grid [][]rune, currentPath [][]dir, currentPosition coordinates, direction dir) {

	if possible(grid, currentPosition) {
		if currentPath[currentPosition.y][currentPosition.x] == direction {
			// Prevent infinite loop
			return
		}
		currentPath[currentPosition.y][currentPosition.x] = direction
	}
	var newPosition coordinates

	switch direction {
	case up:
		newPosition = coordinates{currentPosition.x, currentPosition.y - 1}
	case down:
		newPosition = coordinates{currentPosition.x, currentPosition.y + 1}
	case left:
		newPosition = coordinates{currentPosition.x - 1, currentPosition.y}
	case right:
		newPosition = coordinates{currentPosition.x + 1, currentPosition.y}
	}
	if possible(grid, newPosition) {
		switch grid[newPosition.y][newPosition.x] {
		case '|':
			if direction == left || direction == right {
				walk(grid, currentPath, newPosition, up)
				walk(grid, currentPath, newPosition, down)
			} else {
				walk(grid, currentPath, newPosition, direction)
			}
		case '\\':
			if direction == left {
				walk(grid, currentPath, newPosition, up)
			}
			if direction == right {
				walk(grid, currentPath, newPosition, down)
			}
			if direction == down {
				walk(grid, currentPath, newPosition, right)
			}
			if direction == up {
				walk(grid, currentPath, newPosition, left)
			}
		case '/':
			if direction == left {
				walk(grid, currentPath, newPosition, down)
			}
			if direction == right {
				walk(grid, currentPath, newPosition, up)
			}
			if direction == down {
				walk(grid, currentPath, newPosition, left)
			}
			if direction == up {
				walk(grid, currentPath, newPosition, right)
			}
		case '-':
			if direction == up || direction == down {
				walk(grid, currentPath, newPosition, left)
				walk(grid, currentPath, newPosition, right)
			} else {
				walk(grid, currentPath, newPosition, direction)
			}
		default:
			walk(grid, currentPath, newPosition, direction)
		}
	}
}

func calculate(channel chan int, start coordinates, grid [][]rune, currentPath [][]dir) {

	var direction dir
	if start.x == -1 {
		direction = right
	} else if start.y == -1 {
		direction = down
	} else {
		direction = left
	}

	walk(grid, currentPath, start, direction)
	channel <- count(currentPath)
}

func main() {
	r, err := os.Open("input.txt")
	if err != nil {
		os.Exit(1)
	}
	scanner := bufio.NewScanner(r)

	grid := [][]rune{}
	currentPath := [][]dir{}

	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []rune(line))
		currentPath = append(currentPath, make([]dir, len(line)))
	}

	print(grid)

	startingCoordinates := []coordinates{}

	// top and bottom row
	for x, _ := range grid[0] {
		startingCoordinates = append(startingCoordinates, coordinates{x, -1})
		startingCoordinates = append(startingCoordinates, coordinates{x, len(grid[0])})
	}
	// left column
	for y, _ := range grid {
		startingCoordinates = append(startingCoordinates, coordinates{-1, y})
	}

	channel := make(chan int)
	for _, start := range startingCoordinates {
		go calculate(channel, start, grid, cpy(currentPath))
	}

	highest := math.MinInt

	i := 0
	for total := range channel {
		i++
		fmt.Println(i, total, highest)

		if total > highest {
			highest = total
		}
		if i >= len(startingCoordinates) {
			break
		}
	}
	fmt.Println("highest", highest)
}
