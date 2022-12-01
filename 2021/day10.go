package main

import (
	"fmt"
	"math"
	"sort"
)

func day10() {
	lines := readFile("day10input")
	var incomplete []int
	for _, line := range lines {
		inc := parseLine(line)
		if inc != 0 {
			incomplete = append(incomplete, inc)
		}
	}
	fmt.Println(incomplete)
	sort.Ints(incomplete)
	idx := int(math.Ceil(float64(len(incomplete) / 2)))
	fmt.Println(incomplete[idx])
}

var closing = map[rune]rune{
	'(': ')',
	'[': ']',
	'{': '}',
	'<': '>',
}

// used in part1
var invalidpoints = map[rune]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

var incompletepoints = map[rune]int{
	')': 1,
	']': 2,
	'}': 3,
	'>': 4,
}

func parseLine(s string) int {
	var list []rune
	for _, c := range s {
		switch c {
		case '(', '[', '{', '<':
			list = append(list, c)

		case ')', ']', '}', '>':
			last := list[len(list)-1]
			// valid
			if closing[last] != c {
				return 0
				// used in part1
				//				sum += invalidpoints[c]
			}
			list = list[0 : len(list)-1]
		}
	}

	// incomplete
	result := 0
	var l2 []rune
	for i := len(list) - 1; i >= 0; i-- {
		c := list[i]
		l2 = append(l2, closing[c])
	}
	for _, c := range l2 {
		result = result * 5
		result = result + incompletepoints[c]
	}
	return result
}
