package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func tickWinningNumbers(numbers []string) []bool {
	slice := make([]bool, 100)

	for _, n := range numbers {
		d, _ := strconv.ParseInt(strings.Trim(n, " "), 0, 64)
		if d > 0 {
			slice[d] = true
		}
	}
	return slice
}

func calculateScore(score int, won bool) int {
	if won {
		if score == 0 {
			return 1
		} else {
			return score * 2
		}
	}
	return score
}

func main() {
	r, err := os.Open("input.txt")

	if err != nil {
		os.Exit(1)
	}

	scanner := bufio.NewScanner(r)

	totalScore := 0

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)

		split := strings.Split(line, ":")
		winningAndPlayed := strings.Split(split[1], "|")

		won := tickWinningNumbers(strings.Split(winningAndPlayed[0], " "))

		score := 0

		for _, n := range strings.Split(winningAndPlayed[1], " ") {
			playedNumber, _ := strconv.ParseInt(strings.Trim(n, " "), 0, 64)
			score = calculateScore(score, won[playedNumber])
		}
		totalScore = totalScore + score
	}
	fmt.Println(totalScore)
}
