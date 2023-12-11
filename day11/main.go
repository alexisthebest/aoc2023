package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type point struct {
	n int
	x int
	y int
}

func print(grid [][]rune) {
	for _, row := range grid {
		for _, col := range row {
			fmt.Print(string(col))
		}
		fmt.Println()
	}
}

func expand(grid [][]rune) [][]rune {
	galaxiesX := make([]bool, len(grid[0]))
	galaxiesY := make([]bool, len(grid))

	for y, row := range grid {
		hadGalaxy := false
		for x, col := range row {
			if col == '#' {
				hadGalaxy = true
				galaxiesX[x] = true
			}
		}
		if hadGalaxy {
			galaxiesY[y] = true
		}
	}
	for y := len(galaxiesY) - 1; y >= 0; y-- {
		if !galaxiesY[y] {
			grid = append(append(grid[:y], grid[y]), grid[y:]...)
		}
	}
	for y, _ := range grid {
		for x := len(galaxiesX) - 1; x >= 0; x-- {
			if !galaxiesX[x] {
				grid[y] = append(append(grid[y][:x], '.'), grid[y][x:]...)
			}
		}
	}
	return grid
}

func findGalaxies(grid [][]rune) []point {
	i := 0
	points := []point{}

	for y, row := range grid {
		for x, col := range row {
			if col == '#' {
				i = i + 1
				points = append(points, point{i, x, len(grid) - 1 - y})
			}
		}
	}
	return points
}

func distance(one point, two point) int {
	return int(math.Abs(float64(two.x-one.x) + math.Abs(float64(two.y-one.y))))
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

	print(grid)
	fmt.Println()

	grid = expand(grid)

	print(grid)

	points := findGalaxies(grid)

	fmt.Println(points)

	sum := 0
	count := 0
	for l, one := range points {
		for r, two := range points {
			if l != r && (one.x < two.x || (one.x == two.x && one.y > two.y)) {
				dist := distance(one, two)
				fmt.Println(one.n, "->", two.n, "=", dist, "(", one.x, one.y, "-", two.x, two.y, ")")
				sum += dist
				count += 1
			}
		}
	}
	fmt.Println(count, sum)
}
