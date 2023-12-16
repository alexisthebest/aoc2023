package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func counts(line []bool) []int {
	arr := []int{}

	n := 0
	for _, c := range line {
		if c {
			n++
		} else {
			if n > 0 {
				arr = append(arr, n)
			}
			n = 0
		}
	}
	if n > 0 {
		return append(arr, n)
	}
	return arr
}

func doesMatch(line []bool, groups []int, expectedTotalLength int) (bool, bool) {
	m := counts(line)
	lastIndex := len(m) - 1

	if lastIndex >= 0 {
		j := 0
		totalLength := 0
		valid := false

		for i, currentLength := range m {
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

func nextGroupIndex(index int, groups []int, increment int) int {
	newIndex := index + increment

	if newIndex < 0 {
		newIndex = len(groups) - 1
	} else if newIndex > len(groups)-1 {
		newIndex = 0
	}
	return newIndex
}

func process(index int, originalLine []rune, current rune, groups []int, groupIndex int, seen int, totalSeen int, expectedTotal int) int {
	total := 0

	if index < len(originalLine) && totalSeen < expectedTotal {
		if seen > groups[groupIndex] {
			return 0
		}

		next := '?'
		if index < len(originalLine)-1 {
			next = originalLine[index+1]
		}
		switch current {
		case '#':
			{
				total += process(index+1, originalLine, next, groups, groupIndex, seen+1, totalSeen+1, expectedTotal)
			}
		case '.':
			{
				if seen > 0 {
					if seen != groups[groupIndex] {
						return 0
					}
					groupIndex = nextGroupIndex(groupIndex, groups, 1)
				}
				total += process(index+1, originalLine, next, groups, groupIndex, 0, totalSeen, expectedTotal)
			}
		default:
			{
				total += process(index, originalLine, '#', groups, groupIndex, seen, totalSeen, expectedTotal)
				total += process(index, originalLine, '.', groups, groupIndex, seen, totalSeen, expectedTotal)
			}
		}
	} else {
		if seen == groups[groupIndex] && totalSeen == expectedTotal {
			return total + 1
		}
	}

	return total
}

func runProcess(channel chan int, line []rune, groups []int) {
	fmt.Println("Running", string(line), counts, len(line))

	total := process(0, line, line[0], groups, 0, 0, 0, addUp(groups)*5)

	fmt.Println(string(line), ":", total)

	channel <- total
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

	channel := make(chan int)

	count := 0

	channelCount := 0
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")

		spr := []rune(unfold(parts[0], "?"))
		//spr := []rune(parts[0])
		counts := parseCounts(parts[1])

		go runProcess(channel, spr, counts)
		channelCount += 1
		//break
	}
	n := 0
	for total := range channel {
		count += total
		n++
		fmt.Println(n, count)

		if n >= channelCount {
			break
		}
	}
	close(channel)

	fmt.Println(count)
}
