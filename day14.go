package main

import (
	"fmt"
	"strings"
)

func day14() {
	chunks := readFileChunks("day14input", 2)
	input := chunks[0][0]
	m := make(map[string]string)
	for _, line := range chunks[1] {
		parts := strings.Split(line, " -> ")
		m[parts[0]] = parts[1]
	}

	for i := 0; i < 10; i++ {
		var sb strings.Builder

		for j := 0; j < len(input); j++ {
			if j != len(input)-1 {
				k := input[j : j+2]
				sb.WriteString(string(rune(k[0])))
				sb.WriteString(m[k])
				//				sb.WriteString(string(rune(k[1])))
			} else {
				sb.WriteString(string(rune(input[j])))
			}
		}
		input = sb.String()
	}
	ct := make(map[rune]int)
	for _, c := range input {
		ct[c] += 1
	}
	var vals []int
	for _, v := range ct {
		vals = append(vals, v)
	}

	fmt.Println(max(vals) - min(vals))
	//fmt.Println(input)
	// fmt.Println(m)
}
