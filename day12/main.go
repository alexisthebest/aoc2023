package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func doesMatch(line string, groups []int) (bool, int) {
	regex := regexp.MustCompile(`#+`)
	m := regex.FindAllString(line, -1)
	howMany := len(m)

	if howMany != len(groups) {
		return false, howMany
	}
	for i, count := range groups {
		if len(m[i]) != count {
			return false, howMany
		}
	}
	return true, howMany
}

func process(originalLine []rune, newLine string, groups []int, index int) int {
	total := 0
	matches, counted := doesMatch(newLine, groups)
	if len(newLine) == len(originalLine) && matches {
		fmt.Println(newLine)
		total += 1
	}
	if counted <= len(groups) && index < len(originalLine) {
		current := originalLine[index]

		if current == '#' || current == '.' {
			total += process(originalLine, string(append([]rune(newLine), current)), groups, index+1)
		} else {
			total += process(originalLine, string(append([]rune(newLine), '#')), groups, index+1)
			total += process(originalLine, string(append([]rune(newLine), '.')), groups, index+1)
		}
	}
	return total
}

func main() {
	r, err := os.Open("input.txt")
	if err != nil {
		os.Exit(1)
	}
	scanner := bufio.NewScanner(r)

	count := 0
	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Split(line, " ")

		spr := []rune(parts[0])

		counts := []int{}

		for _, c := range strings.Split(parts[1], ",") {
			n, err := strconv.Atoi(string(c))
			if err != nil {
				os.Exit(1)
			}
			counts = append(counts, n)
		}
		count += process(spr, "", counts, 0)
		fmt.Println()
	}
	fmt.Println(count)
}
