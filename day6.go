package main

import "fmt"

func day6() {

	lines := readFile("day6input")
	line := lines[0]
	for i := 4; i < len(line); i++ {
		seen := map[rune]bool{}
		for _, c := range line[i-4 : i] {
			seen[c] = true
		}
		if len(seen) == 4 {
			fmt.Println(i)
			break
		}
		// fmt.Println(i, line[i-4:i])
	}
}
