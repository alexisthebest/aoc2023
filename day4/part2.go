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

func getCount(card string) int {
	split := strings.Split(card, ":")
	winningAndPlayed := strings.Split(split[1], "|")
	won := tickWinningNumbers(strings.Split(winningAndPlayed[0], " "))

	count := 0
	for _, n := range strings.Split(winningAndPlayed[1], " ") {
		playedNumber, _ := strconv.ParseInt(strings.Trim(n, " "), 0, 64)
		if won[playedNumber] {
			count = count + 1
		}
	}
	return count
}

func playGame(card string, cardNumber int, max int) []int {
	adjustment := cardNumber + 1 + getCount(card)

	winnings := []int{}

	if adjustment > max {
		adjustment = max
	}
	for i := cardNumber + 1; i < adjustment; i++ {
		winnings = append(winnings, i)
	}
	return winnings
}

func playAll(games []string, scratchcards []int, total int) ([]int, int) {
	cardNumber := scratchcards[0]

	wins := playGame(games[cardNumber], cardNumber, len(games))

	result := append(scratchcards, wins...)

	return result[1:], total + len(wins)
}

func main() {
	r, err := os.Open("input.txt")

	if err != nil {
		os.Exit(1)
	}

	scanner := bufio.NewScanner(r)

	games := []string{}
	scratchcards := []int{}
	i := 0

	for scanner.Scan() {
		games = append(games, scanner.Text())
		scratchcards = append(scratchcards, i)
		i = i + 1
	}

	total := len(scratchcards)
	for len(scratchcards) > 0 {
		scratchcards, total = playAll(games, scratchcards, total)
		fmt.Println(total)
	}

	fmt.Println(total)
}
