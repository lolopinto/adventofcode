package main

import "fmt"

func day6() {

	// part 1 is 4
	ct := 14
	lines := readFile("day6input")
	line := lines[0]
	for i := ct; i < len(line); i++ {
		seen := map[rune]bool{}
		for _, c := range line[i-ct : i] {
			seen[c] = true
		}
		if len(seen) == ct {
			fmt.Println(i)
			break
		}
	}
}
