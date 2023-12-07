package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func processLine(line string) int {
	result := []string{}
	for i, t := range strings.Split(line, " ") {
		if i > 0 {
			result = append(result, strings.Trim(t, " "))
		}
	}
	x := strings.Join(result, "")
	n, err := strconv.Atoi(x)
	if err != nil {
		fmt.Println(x)
		os.Exit(1)
	}
	return n
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

	lines := []int{}
	for scanner.Scan() {
		lines = append(lines, processLine(scanner.Text()))
	}
	fmt.Println(lines)

	time := lines[0]
	distance := lines[1]

	w := waysToWin(time, distance)

	fmt.Println(len(w))
}
