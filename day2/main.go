package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

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

func getCount(str string, label string) int64 {
	regex := regexp.MustCompile(`(?P<num>\d+) ` + label)

	names := getRegexNames(str, *regex)

	n, _ := strconv.ParseInt(names["num"], 0, 64)

	return n
}

func isValidGame(game string, red int64, green int64, blue int64) bool {
	if getCount(game, "red") > red {
		return false
	}
	if getCount(game, "green") > green {
		return false
	}
	if getCount(game, "blue") > blue {
		return false
	}
	return true
}

func main() {
	captureIdRegex := regexp.MustCompile(`^Game (?P<id>\d+)\:`)

	r, err := os.Open("input.txt")

	if err != nil {
		os.Exit(1)
	}

	scanner := bufio.NewScanner(r)

	var count int64 = 0

	for scanner.Scan() {
		game := scanner.Text()

		data := getRegexNames(game, *captureIdRegex)

		fmt.Println(game)

		colonPlace := strings.Index(game, ":") + 2

		valid := true

		for _, grab := range strings.Split(game[colonPlace:], ";") {
			valid = valid && isValidGame(grab, 12, 13, 14)
		}
		if valid {
			idNumber, _ := strconv.ParseInt(data["id"], 0, 64)
			fmt.Println("Game", idNumber, "is valid.")
			count = count + idNumber
		}
	}
	fmt.Println(count)
}
