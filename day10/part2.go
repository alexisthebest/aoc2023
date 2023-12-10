package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

func walk(x int, y int, grid [][]rune, pipes [][]rune, count int) {
	position := grid[y][x]
	pipes[y][x] = position

	fmt.Println(count, "at", string(position), "->", x, y)

	for _, movement := range move(position) {
		x2 := x + movement.x
		y2 := y + movement.y

		if y2 > -1 && y2 < len(grid) && x2 > -1 && x2 < len(grid[y2]) {
			if permitted(movement, grid[y2][x2]) && pipes[y2][x2] == 0 {
				walk(x2, y2, grid, pipes, count+1)
			}
		}
	}
}

func cnt(arr []rune) int {
	new := []rune{}
	for _, c := range arr {
		if c > 0 {
			new = append(new, c)
		}
	}
	str := string(new)

	str = strings.Replace(str, " ", "", -1)
	str = strings.Replace(str, ".", "", -1)
	str = strings.Replace(str, "-", "", -1)
	str = strings.Replace(str, "L7", "|", -1)
	str = strings.Replace(str, "FJ", "|", -1)
	str = strings.Replace(str, "LJ", "", -1)
	str = strings.Replace(str, "F7", "", -1)

	return len(str)
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
	pipes := make([][]rune, len(grid))

	for y, row := range grid {
		pipes[y] = make([]rune, len(row))

		for x, col := range row {
			fmt.Print(string(col))
			if col == 'S' {
				sx, sy = x, y
			}
		}
		fmt.Println()
	}

	walk(sx, sy, grid, pipes, 0)

	count := 0
	for y, row := range pipes {
		for x, c := range row {
			inside := false
			if c == 0 {
				left := cnt(pipes[y][:x])
				right := cnt(pipes[y][x:])

				if left > 0 && right > 0 && left%2 != 0 && right%2 != 0 {
					//fmt.Println(left, right)
					inside = true
				}
			}
			if inside {
				fmt.Print("x")
				count = count + 1
			} else if c == 0 {
				fmt.Print(".")
			} else {
				fmt.Print("+")
			}
		}
		fmt.Println()
	}
	fmt.Println(count)
}
