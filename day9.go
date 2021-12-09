package main

import "fmt"

func day9() {
	lines := readFile("day9input")
	sum := 0
	for i, line := range lines {
		for j, v := range line {
			num := atoi(string(v))
			adj := getAdjacent(line, lines, j, i)
			if num < min(adj) {
				sum += num + 1
			}
		}
	}

	fmt.Println(sum)
}

func convToNum(c rune) int {
	return atoi(string(c))
}

func getAdjacent(line string, lines []string, linePos, strPos int) []int {
	var ret []int
	if linePos-1 >= 0 {
		ret = append(ret, convToNum(rune(line[linePos-1])))
	}
	if linePos+1 < len(line) {
		ret = append(ret, convToNum(rune(line[linePos+1])))
	}
	if strPos-1 >= 0 {
		ret = append(ret, convToNum(rune(lines[strPos-1][linePos])))
	}
	if strPos+1 < len(lines) {
		ret = append(ret, convToNum(rune(lines[strPos+1][linePos])))
	}

	return ret
}
