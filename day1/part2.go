package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func convert(line string) string {
	numbers := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	newString := line

	for word, value := range numbers {
		newString = strings.Replace(newString, word, fmt.Sprintf("%s %d %s", word, value, word), -1)
	}

	return newString
}

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

		converted := convert(line)

		fmt.Println(converted)

		together := fmt.Sprintf("%d%d", getFirstDigit(converted), getLastDigit(converted))

		total, err := strconv.ParseInt(together, 0, 64)

		fmt.Println(total)

		if err != nil {
			os.Exit(1)
		}

		count = count + total

		fmt.Println("******")
	}

	fmt.Println(count)
}
