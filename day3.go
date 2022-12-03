package main

import "fmt"

func day3() {
	lines := readFile("day3input")
	priority := 0

	for _, line := range lines {
		l := len(line)
		a := line[0 : l/2]
		b := line[(l+1)/2:]
		// fmt.Println(a, " .   ", b)

		m := map[rune]int{}
		for _, c := range a {
			m[c] += 1
		}
		for _, c := range b {
			_, ok := m[c]
			if ok {
				// fmt.Println(string(c))
				if c >= 'a' {
					priority += (int(c) - 'a') + 1
				} else if c >= 'A' {
					priority += (int(c) - 'A') + 27
				}
				break
			}
		}
	}
	fmt.Println(priority)
}
