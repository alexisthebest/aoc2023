package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type sorter func(i, j int) bool

type hand struct {
	raw   string
	bid   int
	cards []rune
	score int
}

func count(l []rune) map[rune]int {
	m := map[rune]int{}
	for _, r := range l {
		m[r] = m[r] + 1
	}
	return m
}

func scoreCard(c rune) int {
	order := map[rune]int{
		'A': 13,
		'K': 12,
		'Q': 11,
		'J': 10,
		'T': 9,
		'9': 8,
		'8': 7,
		'7': 6,
		'6': 5,
		'5': 4,
		'4': 3,
		'3': 2,
		'2': 1,
	}
	return order[c]
}

func buildHand(line string) hand {
	r := strings.Split(line, " ")
	bid, err := strconv.Atoi(r[1])
	if err != nil {
		fmt.Println(r[1])
		os.Exit(1)
	}
	l := []rune(r[0])
	score := math.MinInt
	m := count(l)

	for _, v := range m {
		if v > score {
			score = v
		}
	}
	return hand{line, bid, l, score - len(m)}
}

func compare(a []rune, b []rune) bool {
	for i, c := range a {
		sa := scoreCard(c)
		sb := scoreCard(b[i])

		if sb == sa {
			continue
		}
		return sb > sa
	}
	return true
}

func sortfunc(hands []hand) sorter {
	return func(i, j int) bool {
		left := hands[i]
		right := hands[j]

		if left.score > right.score {
			return false
		}
		if right.score > left.score {
			return true
		}
		return compare(left.cards, right.cards)
	}
}

func main() {
	r, err := os.Open("input.txt")
	if err != nil {
		os.Exit(1)
	}
	scanner := bufio.NewScanner(r)

	hands := []hand{}
	for scanner.Scan() {
		hands = append(hands, buildHand(scanner.Text()))
	}
	fmt.Println(hands)

	sort.SliceStable(hands, sortfunc(hands))

	total := 0
	for i, hand := range hands {
		fmt.Println(i+1, hand.raw)
		total = total + ((i + 1) * hand.bid)
	}
	fmt.Println(total)
}
