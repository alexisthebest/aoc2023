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

func getCounts(game string) (int64, int64, int64) {
	return getCount(game, "red"), getCount(game, "green"), getCount(game, "blue")
}

func main() {
	r, err := os.Open("input.txt")

	if err != nil {
		os.Exit(1)
	}

	scanner := bufio.NewScanner(r)

	var sum int64 = 0

	for scanner.Scan() {
		game := scanner.Text()

		fmt.Println(game)

		colonPlace := strings.Index(game, ":") + 2

		var highestRed, highestGreen, highestBlue int64 = 0, 0, 0

		for _, grab := range strings.Split(game[colonPlace:], ";") {
			red, green, blue := getCounts(grab)

			if red > highestRed {
				highestRed = red
			}
			if green > highestGreen {
				highestGreen = green
			}
			if blue > highestBlue {
				highestBlue = blue
			}
		}
		fmt.Println("Game required", highestRed, "red", highestGreen, "green", highestBlue, "blue")

		sum = sum + (highestRed * highestGreen * highestBlue)
	}
	fmt.Println(sum)
}
