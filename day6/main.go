package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func processLine(line string) []int {
	result := []int{}
	for _, t := range strings.Split(line, " ") {
		trimmed := strings.Trim(t, " ")
		n, err := strconv.Atoi(trimmed)
		if err == nil {
			result = append(result, n)
		}
	}
	return result
}

func waysToWin(time int, distance int) []int {
	winning := []int{}

	for i := 0; i < time; i++ {
		n := i * (time - i)
		if n > distance {
			winning = append(winning, i)
		}
	}

	return winning
}

func main() {
	r, err := os.Open("input.txt")
	if err != nil {
		os.Exit(1)
	}
	scanner := bufio.NewScanner(r)

	lines := [][]int{}
	for scanner.Scan() {
		lines = append(lines, processLine(scanner.Text()))
	}
	fmt.Println(lines)

	sum := 1
	for i := 0; i < len(lines[0]); i++ {
		time := lines[0][i]
		distance := lines[1][i]

		w := waysToWin(time, distance)
		sum = sum * len(w)
		fmt.Println(i, ":", w)
	}
	fmt.Println(sum)
}
