package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type box struct {
	lenses *[]string
}

func (b *box) remove(code string) int {
	found := -1
	for i, lens := range *b.lenses {
		if len(lens) > 0 {
			kv := strings.Split(lens, " ")
			if kv[0] == code {
				found = i
				break
			}
		}
	}
	if found != -1 {
		(*b.lenses)[found] = ""
	}
	return found
}

func (b *box) insert(key string) {
	newIndex := 0
	for i := len(*b.lenses) - 1; i >= 0; i-- {
		if len((*b.lenses)[i]) > 0 {
			newIndex = i + 1
			break
		}
	}
	if newIndex > len(*b.lenses)-1 {
		*b.lenses = append(*b.lenses, key)
	} else {
		(*b.lenses)[newIndex] = key
	}
}

func (b *box) replace(code string, key string) bool {
	index := b.remove(code)
	if index == -1 {
		b.insert(key)
		return false
	}
	(*b.lenses)[index] = key
	return true
}

func (b *box) shift() bool {
	lenses := *b.lenses

	for i := 1; i < len(lenses); i++ {
		if lenses[i] != "" && lenses[i-1] == "" {
			lenses[i], lenses[i-1] = lenses[i-1], lenses[i]
			return true
		}
	}
	return false
}

func (b *box) shiftAll() {
	for b.shift() {
		fmt.Println("shifting...")
	}
}

func (b *box) calculate(number int) int {
	total := 0

	for i, lens := range *b.lenses {
		if len(lens) > 0 {
			split := strings.Split(lens, " ")

			focalLength, err := strconv.Atoi(split[1])
			if err != nil {
				os.Exit(1)
			}
			sum := (1 + number) * (1 + i) * focalLength
			total += sum
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

	boxes := make([]box, 256)
	for i, _ := range boxes {
		arr := make([]string, 0)
		boxes[i] = box{&arr}
	}
	for scanner.Scan() {
		line := scanner.Text()

		for _, chars := range strings.Split(line, ",") {
			kv := strings.Split(chars, "=")
			lookup := strings.ReplaceAll(kv[0], "-", "")

			boxIndex := 0
			for _, step := range lookup {
				boxIndex += int(step)
				boxIndex *= 17
				boxIndex %= 256
			}
			box := boxes[boxIndex]

			if len(kv) == 2 {
				box.replace(kv[0], fmt.Sprintf("%s %s", lookup, kv[1]))
			} else {
				box.remove(lookup)
				box.shiftAll()
			}
		}
	}
	total := 0
	for i, box := range boxes {
		if len(*box.lenses) > 0 {
			fmt.Print(i, "->")
			for j, lens := range *box.lenses {
				fmt.Print(j, ": [", lens, "] ")
			}

			fmt.Println()
		}

		total += box.calculate(i)
	}
	fmt.Println(total)
}
