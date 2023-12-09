package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
)

type leftright struct {
	left  string
	right string
}

const rx = "^(?P<key>\\w{3}) \\= \\((?P<left>\\w{3})\\, (?P<right>\\w{3})\\)$"

func getRegexNames(line string, exp regexp.Regexp) map[string]string {
	result := make(map[string]string)
	match := exp.FindStringSubmatch(line)

	if len(match) > 0 {
		for i, name := range exp.SubexpNames() {
			if i != 0 && name != "" {
				result[name] = match[i]
			}
		}
	}
	return result
}

func create(str string) (string, leftright) {
	regex := regexp.MustCompile(rx)
	m := getRegexNames(str, *regex)
	return m["key"], leftright{m["left"], m["right"]}
}

func endsIn(key string, c rune) bool {
	return []rune(key)[len(key)-1] == c
}

func allEndIn(keys []string, c rune) bool {
	for _, k := range keys {
		if !endsIn(k, c) {
			return false
		}
	}
	return true
}

func countSteps(begin string, m map[string]leftright, lr []rune) int {
	cursor := begin
	steps := 0

	for true {
		for _, c := range lr {
			steps = steps + 1
			if c == 'L' {
				cursor = m[cursor].left
			} else {
				cursor = m[cursor].right
			}
			if endsIn(cursor, 'Z') {
				return steps
			}
		}
	}
	return -1
}

func allDivisible(counts []int, total int) bool {
	for _, n := range counts {
		if total%n != 0 {
			return false
		}
	}
	return true
}

func main() {
	r, err := os.Open("input.txt")
	if err != nil {
		os.Exit(1)
	}
	scanner := bufio.NewScanner(r)

	var lr []rune
	m := map[string]leftright{}

	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		if i == 0 {
			lr = []rune(line)
		} else if i > 1 {
			key, lookup := create(line)
			m[key] = lookup
		}
		i = i + 1
	}
	cursor := []string{}
	for k, _ := range m {
		if endsIn(k, 'A') {
			cursor = append(cursor, k)
		}
	}
	counts := make([]int, len(cursor))
	highest := math.MinInt
	for i, v := range cursor {
		n := countSteps(v, m, lr)
		if n > highest {
			highest = n
		}
		counts[i] = n
	}
	fmt.Println(counts)

	for total := highest; total < math.MaxInt; total += highest {
		if allDivisible(counts, total) {
			fmt.Println(total)
			break
		}
	}
}
