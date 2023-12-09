package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func allEqualTo(numbers []int, exp int) bool {
	for _, n := range numbers {
		if n != exp {
			return false
		}
	}
	return true
}

func build(numbers []int, all *[][]int) {
	line := make([]int, len(numbers)-1)

	for i := 1; i < len(numbers); i++ {
		line[i-1] = numbers[i] - numbers[i-1]
	}
	*all = append(*all, line)

	if !allEqualTo(line, 0) {
		build(line, all)
	}
}

func extrapolate(all *[][]int) int {
	last := 0

	for i := len(*all) - 1; i >= 0; i-- {
		current := (*all)[i]
		last = current[0] - last

		(*all)[i] = append([]int{last}, current...)
	}
	return last
}

func main() {
	r, err := os.Open("input.txt")
	if err != nil {
		os.Exit(1)
	}
	scanner := bufio.NewScanner(r)

	total := 0
	for scanner.Scan() {
		line := scanner.Text()

		stringNumbers := strings.Split(line, " ")
		numbers := make([]int, len(stringNumbers))

		for i, s := range stringNumbers {
			n, err := strconv.Atoi(s)
			if err != nil {
				os.Exit(1)
			}
			numbers[i] = n
		}

		result := [][]int{numbers}

		build(numbers, &result)

		total = total + extrapolate(&result)

		fmt.Println(result)

		fmt.Println("\n")
	}
	fmt.Println(total)
}
