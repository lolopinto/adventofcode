package main

import (
	"fmt"
	"strings"
)

func translate(str, output string) []string {
	return []string{
		string(rune(str[0])) + output,
		output + string(rune(str[1])),
	}
}

func day14() {
	chunks := readFileChunks("day14input", 2)
	input := chunks[0][0]
	m := make(map[string][]string)

	// build map map from input to 2 letter combos we'll see
	for _, line := range chunks[1] {
		parts := strings.Split(line, " -> ")
		m[parts[0]] = translate(parts[0], parts[1])
	}

	// set up initial input
	data := make(map[string]int64)
	for i := 0; i < len(input)-1; i++ {
		two := input[i : i+2]
		data[two] += 1
	}
	//	fmt.Println(data)

	for i := 0; i < 40; i++ {
		data2 := make(map[string]int64)
		for k, v := range data {
			for _, v2 := range m[k] {
				data2[v2] += v
			}
		}
		data = data2
	}

	//	fmt.Println(data, last)
	fmt.Println(getCount(data, input))
}

func getCount(data map[string]int64, input string) int64 {
	ct := make(map[rune]int64)
	for k, v := range data {
		ct[rune(k[0])] += v
	}
	// add last character from last
	ct[rune(input[len(input)-1])] += 1

	var vals []int64
	for _, v := range ct {
		vals = append(vals, v)
	}
	return max64(vals) - min64(vals)
}
