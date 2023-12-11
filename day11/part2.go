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

func find(grid [][]rune) ([]bool, []bool) {
	galaxiesX := make([]bool, len(grid[0]))
	galaxiesY := make([]bool, len(grid))

	for y := len(grid) - 1; y >= 0; y-- {
		hadGalaxy := false
		for x, col := range grid[y] {
			if col == '#' {
				hadGalaxy = true
				galaxiesX[x] = true
			}
		}
		if hadGalaxy {
			galaxiesY[len(grid)-y-1] = true
		}
	}
	return galaxiesX, galaxiesY
}

func expand(index int, lookup []bool) int {
	newIndex := index
	for _, galaxy := range lookup[:index] {
		if !galaxy {
			newIndex += 999_999
			//newIndex += 99
		}
	}
	return newIndex
}

func findGalaxies(grid [][]rune, gx []bool, gy []bool) []point {
	i := 0
	points := []point{}

	for y, row := range grid {

		for x, col := range row {

			if col == '#' {
				i = i + 1

				newX := expand(x, gx)
				newY := expand(len(grid)-1-y, gy)

				points = append(points, point{i, newX, newY})
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

	gx, gy := find(grid)

	points := findGalaxies(grid, gx, gy)

	sum := 0
	count := 0
	for l, one := range points {
		for r, two := range points {
			if l != r && (one.x < two.x || (one.x == two.x && one.y > two.y)) {
				dist := distance(one, two)
				sum += dist
				count += 1
			}
		}
	}
	fmt.Println(count, sum)
}
