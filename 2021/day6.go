package main

import (
	"fmt"
	"strings"
)

func day6() {
	lines := readFile("day6input")
	line := lines[0]
	input := ints(strings.Split(line, ","))
	counts := make(map[int]int)
	// keep track of count
	for _, v := range input {
		counts[v] += 1
	}

	for i := 1; i <= 256; i++ {
		newcounts := make(map[int]int)
		for k, ct := range counts {
			newkey := k - 1
			if newkey == -1 {
				newkey = 6
				newcounts[8] += ct
			}
			newcounts[newkey] += ct
		}

		sum := 0
		for _, v := range newcounts {
			sum += v
		}
		fmt.Println("day", i, sum)
		counts = newcounts
	}
}
