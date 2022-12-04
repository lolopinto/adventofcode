package main

import (
	"fmt"
	"strings"
)

func day4() {

	lines := readFile("day4input")

	parseRange := func(s string) [2]int {
		parts := strings.Split(s, "-")
		return [2]int{atoi(parts[0]), atoi(parts[1])}
	}

	contains := func(r1, r2 [2]int) bool {
		if r1[0] >= r2[0] && r1[1] <= r2[1] {
			return true
		}
		return false
	}

	ct := 0
	for _, line := range lines {
		parts := strings.Split(line, ",")
		r1 := parseRange(parts[0])
		r2 := parseRange(parts[1])

		if contains(r1, r2) || contains(r2, r1) {
			ct += 1
		}
	}
	fmt.Println(ct)
}
