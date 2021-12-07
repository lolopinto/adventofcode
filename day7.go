package main

import (
	"fmt"
	"math"
	"strings"
)

func day7() {
	lines := readFile("day7input")
	input := ints(strings.Split(lines[0], ","))

	least := math.MaxInt
	for i := range input {
		sum := 0
		for j := range input {
			diff := int(math.Abs(float64(input[j] - i)))
			sum += ((diff * (diff + 1)) / 2)
		}
		if sum < least {
			least = sum
		}
	}
	fmt.Println(least)
}
