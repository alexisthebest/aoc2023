package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

type leftright struct {
	left  string
	right string
}

const start = "AAA"

const end = "ZZZ"

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

	cursor := start
	steps := 0

	for cursor != end {
		for _, c := range lr {
			steps = steps + 1
			curr := m[cursor]
			if c == 'L' {
				cursor = curr.left
			} else {
				cursor = curr.right
			}
			if cursor == end {
				break
			}
		}
	}
	fmt.Println(steps)
}
