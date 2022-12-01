package main

import "fmt"

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

	fmt.Println(max(sums))
}
