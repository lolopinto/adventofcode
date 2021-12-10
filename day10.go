package main

import (
	"fmt"
)

func day10() {
	lines := readFile("day10input")
	sum := 0
	for _, line := range lines {
		sum += parseLine(line)
	}
	fmt.Println(sum)
}

var closing = map[rune]rune{
	'(': ')',
	'[': ']',
	'{': '}',
	'<': '>',
}

var invalidpoints = map[rune]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

func parseLine(s string) int {
	sum := 0
	var list []rune
	for _, c := range s {
		switch c {
		case '(', '[', '{', '<':
			list = append(list, c)

		case ')', ']', '}', '>':
			last := list[len(list)-1]
			// valid
			if closing[last] != c {
				sum += invalidpoints[c]
				//				fmt.Println("invalid line", s)
			}
			list = list[0 : len(list)-1]
		}
	}
	return sum
}
