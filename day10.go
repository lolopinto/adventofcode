package main

import (
	"log"
	"sort"
)

func day10() {
	lines := readFile("day10input")

	numbers := make([]int, len(lines))

	for i, line := range lines {
		numbers[i] = atoi(line)
	}

	sort.Ints(numbers)

	diff1 := 1
	diff3 := 1

	init := numbers[0]
	last := init
	for idx, num := range numbers {
		if idx == 0 {
			continue
		}
		if num-last > 3 {
			continue
		}
		if num-last == 1 {
			diff1++
		}
		if num-last == 3 {
			diff3++
		}
		last = num
	}
	log.Println(diff1, diff3, diff1*diff3)
}
