package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func toBinary(line string) uint64 {
	binary := make([]rune, len(line))

	for i, c := range line {
		newC := '0'
		if c == '#' {
			newC = '1'
		}
		binary[i] = newC
	}
	u, err := strconv.ParseUint(string(binary), 2, len(binary))
	if err != nil {
		os.Exit(1)
	}
	return u
}

func toBinaryList(lines []string) []uint64 {
	arr := make([]uint64, len(lines))
	for i, n := range lines {
		arr[i] = toBinary(n)
	}
	return arr
}

func findMidPoint(lines []uint64) []int {
	midPoints := []int{}
	for i := 0; i < len(lines)-1; i++ {
		if lines[i+1] == lines[i] {
			midPoints = append(midPoints, i)
		}
	}
	return midPoints
}

func compareSides(lines []uint64, midPointIndex int) (bool, int) {
	matches := 0
	i := midPointIndex
	j := i + 1

	for i >= 0 && j < len(lines) {
		if lines[i] == lines[j] {
			matches++
			i--
			j++
		} else {
			break
		}
	}
	return matches > 0 && (i == -1 || j == len(lines)), matches
}

func doesMirror(lines []uint64) (int, int) {
	mp := findMidPoint(lines)

	fmt.Println("there are ", len(mp), "midpoints")

	for _, midPointIndex := range mp {
		mirror, size := compareSides(lines, midPointIndex)
		if mirror {
			return midPointIndex + 1, size
		}
	}
	return -1, -1
}

func process(grid [][]rune) int {
	horizontalLines := make([]string, len(grid))
	verticalLines := make([]string, len(grid[0]))

	for y, row := range grid {
		horizontalLines[y] = string(row)

		for x, col := range row {
			verticalLines[x] = string(append([]rune(verticalLines[x]), col))
		}
	}

	h := toBinaryList(horizontalLines)
	fmt.Println(h)
	above, hSize := doesMirror(h)

	v := toBinaryList(verticalLines)
	fmt.Println(v)
	left, vSize := doesMirror(v)

	fmt.Println(left, "left=", vSize, above, "above=", hSize)

	if vSize < 0 && hSize < 0 {
		os.Exit(1)
	}
	if vSize > hSize {
		return left
	}
	return above * 100
}

func main() {
	r, err := os.Open("input.txt")
	if err != nil {
		os.Exit(1)
	}
	scanner := bufio.NewScanner(r)

	sum := 0
	grid := [][]rune{}
	for scanner.Scan() {
		line := []rune(scanner.Text())
		fmt.Println(string(line))
		if len(line) == 0 {
			sum += process(grid)
			fmt.Println("\n")
			grid = [][]rune{}
		} else {
			grid = append(grid, line)
		}
	}
	sum += process(grid)

	fmt.Println(sum)
}
