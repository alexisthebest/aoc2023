package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func calculateMinimumRemaining(groups []int) int {
	minimumRemaining := -1 // Start at -1 to account for 1 added to the end
	for _, n := range groups {
		// Count group of hashes plus dot separator
		minimumRemaining += n + 1
	}
	return minimumRemaining
}

func process(index int, originalLine []rune, current rune, groups []int, seen int, totalSeen int, expectedTotal int, minimumRemaining []int) int {
	total := 0

	if index < len(originalLine) && totalSeen < expectedTotal {
		if seen > groups[0] {
			return 0
		}
		next := '?'
		if index < len(originalLine)-1 {
			next = originalLine[index+1]
		}
		switch current {
		case '#':
			{
				total += process(index+1, originalLine, next, groups, seen+1, totalSeen+1, expectedTotal, minimumRemaining)
			}
		case '.':
			{
				if seen > 0 {
					if seen != groups[0] {
						return 0
					}
					total += process(index+1, originalLine, next, groups[1:], 0, totalSeen, expectedTotal, minimumRemaining[1:])
				} else {
					remainingSlots := len(originalLine) - (index + 1)
					if minimumRemaining[0] > remainingSlots {
						return 0
					}
					total += process(index+1, originalLine, next, groups, 0, totalSeen, expectedTotal, minimumRemaining)
				}
			}
		default:
			{
				total += process(index, originalLine, '#', groups, seen, totalSeen, expectedTotal, minimumRemaining)
				total += process(index, originalLine, '.', groups, seen, totalSeen, expectedTotal, minimumRemaining)
			}
		}
	} else {
		if seen == groups[0] && totalSeen == expectedTotal {
			return total + 1
		}
	}
	return total
}

func buildMinimumRemaining(groups []int, minimumRemaining []int) []int {
	if len(groups) == 0 {
		return minimumRemaining
	}
	return buildMinimumRemaining(groups[1:], append(minimumRemaining, calculateMinimumRemaining(groups)))
}

func runProcess(channel chan int, line []rune, groups []int) {
	fmt.Println("Running", string(line), len(line))

	newGroups := make([]int, len(groups)*5)

	j := 0
	for j < len(newGroups) {
		for _, n := range groups {
			newGroups[j] = n
			j++
		}
	}

	minimumRemaining := buildMinimumRemaining(newGroups, []int{})

	total := process(0, line, line[0], newGroups, 0, 0, addUp(groups)*5, minimumRemaining)

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
		counts := parseCounts(parts[1])

		go runProcess(channel, spr, counts)
		channelCount += 1
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
