package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func getFirstDigit(line string) int64 {
	for i := 0; i < len(line); i++ {
		value, err := strconv.ParseInt(string(line[i]), 0, 64)
		if err != nil {
			continue
		}
		return value
	}
	return 0
}

func getLastDigit(line string) int64 {
	for i := len(line) - 1; i >= 0; i-- {
		value, err := strconv.ParseInt(string(line[i]), 0, 64)
		if err != nil {
			continue
		}
		return value
	}
	return 0
}

func main() {
	r, err := os.Open("input.txt")

	if err != nil {
		os.Exit(1)
	}

	scanner := bufio.NewScanner(r)

	var count int64 = 0

	for scanner.Scan() {
		line := scanner.Text()

		fmt.Println(line)

		together := fmt.Sprintf("%d%d", getFirstDigit(line), getLastDigit(line))

		total, err := strconv.ParseInt(together, 0, 64)

		fmt.Println(total)

		if err != nil {
			os.Exit(1)
		}

		count = count + total
	}

	fmt.Println(count)
}
