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
		return r1[0] >= r2[0] && r1[1] <= r2[1]
	}
	overlaps := func(r1, r2 [2]int) bool {
		return r1[0] >= r2[0] && r1[0] <= r2[1]
	}

	ct1 := 0
	ct2 := 0
	for _, line := range lines {
		parts := strings.Split(line, ",")
		r1 := parseRange(parts[0])
		r2 := parseRange(parts[1])

		if contains(r1, r2) || contains(r2, r1) {
			ct1 += 1
		}
		if overlaps(r1, r2) || overlaps(r2, r1) {
			ct2 += 1
		}
	}
	fmt.Println(ct1)
	fmt.Println(ct2)
}
