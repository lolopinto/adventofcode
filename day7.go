package main

import (
	"fmt"
	"math"
	"strings"
)

func day7() {
	lines := readFile("day7input")
	input := ints(strings.Split(lines[0], ","))

	m := make(map[int]int)
	for i := range input {
		sum := 0
		for j := range input {
			if i != j {
				sum += int(math.Abs(float64(input[i] - input[j])))
			}
		}
		m[i] = sum
	}
	leasst := math.MaxInt
	for _, v := range m {
		if v < leasst {
			leasst = v
		}
	}
	fmt.Println(leasst)
}
