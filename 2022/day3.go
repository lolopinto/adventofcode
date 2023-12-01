package main

import "fmt"

func day3() {
	lines := readFile("day3input")

	getPriority := func(c rune) int {
		if c >= 'a' {
			return (int(c) - 'a') + 1
		} else if c >= 'A' {
			return (int(c) - 'A') + 27
		}
		return 0
	}

	priority := 0

	for _, line := range lines {
		l := len(line)
		a := line[0 : l/2]
		b := line[(l+1)/2:]

		m := map[rune]int{}
		for _, c := range a {
			m[c] += 1
		}
		for _, c := range b {
			_, ok := m[c]
			if ok {
				priority += getPriority(c)
				break
			}
		}
	}
	fmt.Println(priority)

	priority2 := 0
	for _, group := range groupLines(lines, 3) {
		m := map[rune]int{}
		for i, l := range group {
			seen := make(map[rune]bool)
			for _, c := range l {
				if !seen[c] {
					seen[c] = true
					m[c] += 1
				}

				if i == 2 && m[c] == 3 {
					priority2 += getPriority(c)
					break
				}
			}
		}
	}

	fmt.Println(priority2)
}
