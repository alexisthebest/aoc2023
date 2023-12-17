package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	r, err := os.Open("input.txt")
	if err != nil {
		os.Exit(1)
	}
	scanner := bufio.NewScanner(r)

	sum := 0
	for scanner.Scan() {
		line := scanner.Text()

		for _, chars := range strings.Split(line, ",") {
			currentValue := 0

			for _, step := range chars {
				currentValue += int(step)
				currentValue *= 17
				currentValue %= 256
			}
			fmt.Println(chars, "becomes", currentValue)
			sum += currentValue
		}
	}
	fmt.Println(sum)
}
