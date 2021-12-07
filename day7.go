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
			diff := int(math.Abs(float64(input[j] - i)))
			sum += ((diff * (diff + 1)) / 2)
		}
		m[i] = sum
	}
	least := math.MaxInt
	for _, v := range m {
		if v < least {
			least = v
		}
	}
	fmt.Println(least)
}
