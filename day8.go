package main

import (
	"fmt"
	"strings"
)

func day8() {
	// 7 3 segements
	// 4 4 segments
	// 1 2 segments
	// 8 7 segments
	lines := readFile("day8input")
	sum := 0
	for _, line := range lines {
		parts := strings.Split(line, " | ")
		//		input := strings.Split(parts[0], " ")
		output := strings.Split(parts[1], " ")
		for _, v := range output {
			switch len(v) {
			case 3, 2, 4, 7:
				sum += 1
			}
		}
		//		fmt.Println(len(input), len(output))
	}
	fmt.Println(sum)

}
