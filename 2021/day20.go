package main

import (
	"fmt"
	"strings"
)

type gridType map[[2]int]rune

func day20() {
	// fmt.Println(convertToBinary("000000000")) => 1
	// fmt.Println(convertToBinary("111111111")) => 0
	// flipped in the example

	lines := readFile("day20input")
	input := lines[0]
	gridInput := lines[2:]

	m := make(gridType)
	// initialize grid
	for i := 0; i < len(gridInput); i++ {
		for j := 0; j < len(gridInput); j++ {
			m[[2]int{i, j}] = rune(gridInput[i][j])
		}
	}
	// rigged input
	flip := input[0] == '#'

	length := len(gridInput)
	//	fmt.Println(length)
	// it works all the way consistently for the small input
	for i := 0; i < 50; i++ {
		cp := make(gridType)
		length += 2
		//		fmt.Println("length", length)
		for x := 0; x < length; x++ {
			for y := 0; y < length; y++ {
				val := getNumberValue(m, x-1, y-1, length-2, flip && i%2 == 1)
				cp[[2]int{x, y}] = rune(input[val])
			}
		}
		m = cp
		//		ct := countLitUpGrid(m)
		//		fmt.Println("count", i, ct)
	}
	ct := countLitUpGrid(m)
	fmt.Println("result", ct)
}

func getValue(m gridType, i, j, length int, flip bool) rune {
	k := [2]int{i, j}
	v, ok := m[k]
	if ok {
		if v == '#' {
			return '1'
		}
		return '0'
	}
	// we pass previous length
	if flip && (i < 0 || j < 0 || i >= length || j >= length) {
		return '1'
	}
	return '0'
}

func getNumberValue(m gridType, x, y, length int, flip bool) int {
	var sb strings.Builder

	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			sb.WriteRune(getValue(m, x+i, y+j, length, flip))
		}
	}
	return convertToBinary(sb.String())
}

// count lit up
func countLitUpGrid(m gridType) int {
	ct := 0
	for _, v := range m {
		if v == '#' {
			ct++
		}
	}
	return ct
}
