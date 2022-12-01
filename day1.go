package main

import (
	"fmt"
	"sort"
)

func day1() {
	chunks := readFileChunks("day1input", -1)
	sums := make([]int, len(chunks))
	for i, chunk := range chunks {
		sum := 0
		for _, v := range chunk {
			sum += atoi(v)
		}
		sums[i] = sum
	}

	sort.Ints(sums)

	l := len(sums)

	// part 1
	fmt.Println(max(sums))
	fmt.Println(sums[l-1] + sums[l-2] + sums[l-3])
}
