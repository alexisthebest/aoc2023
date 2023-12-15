package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func doesMatch(line string, groups []int, expectedTotalLength int) (bool, bool) {
	regex := regexp.MustCompile(`#+`)
	m := regex.FindAllString(line, -1)
	lastIndex := len(m) - 1

	if lastIndex >= 0 {
		j := 0
		totalLength := 0
		valid := false

		for i, str := range m {
			currentLength := len(str)
			valid = currentLength == groups[j]
			if i < lastIndex && !valid {
				return false, false
			}
			j++
			if j >= len(groups) {
				j = 0
			}
			totalLength += currentLength
		}
		return true, valid && totalLength == expectedTotalLength
	}
	return true, false
}

func process(originalLine []rune, newLine string, groups []int, index int, expected int) int {
	total := 0

	partialMatch, fullMatch := doesMatch(newLine, groups, expected)

	if partialMatch {
		if fullMatch && len(newLine) == len(originalLine) {
			return total + 1
		}
		if index < len(originalLine) {
			current := originalLine[index]

			if current == '#' || current == '.' {
				total += process(originalLine, string(append([]rune(newLine), current)), groups, index+1, expected)
			} else {
				total += process(originalLine, string(append([]rune(newLine), '#')), groups, index+1, expected)
				total += process(originalLine, string(append([]rune(newLine), '.')), groups, index+1, expected)
			}
		}
	}
	return total
}

func unfold(line string, join string) string {
	return strings.Join([]string{line, line, line, line, line}, join)
}

func parseCounts(line string) []int {
	counts := []int{}

	for _, c := range strings.Split(line, ",") {
		n, err := strconv.Atoi(string(c))
		if err != nil {
			os.Exit(1)
		}
		counts = append(counts, n)
	}
	return counts
}

func addUp(groups []int) int {
	n := 0
	for _, g := range groups {
		n += g
	}
	return n
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

		spr := []rune(unfold(parts[0], "?"))
		counts := parseCounts(parts[1])

		fmt.Println("Running", string(spr), counts, len(spr))
		n := process(spr, "", counts, 0, addUp(counts)*5)
		fmt.Println(n)
		count += n
		fmt.Println()
		//break
	}
	fmt.Println(count)
}
