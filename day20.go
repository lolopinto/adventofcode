package main

import (
	"fmt"
	"strings"
)

type gridType map[[2]int]rune

// sparse grid approach not working. how do I write surrounding
func day20() {
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

	length := len(gridInput)
	fmt.Println(length)
	// it works all the way consistently for the small input
	for i := 0; i < 2; i++ {
		//		cp := makeMapCopy(m)
		//		length += 2
		cp := make(gridType)
		//		_, low, hi := gridStats(m)
		// low := length - length - 1
		// hi := length + 1
		//		fmt.Println("dim", low, hi)
		length += 2
		fmt.Println("length", length)
		for x := 0; x < length; x++ {
			for y := 0; y < length; y++ {
				val := getNumberValue(m, x-1, y-1)
				cp[[2]int{x, y}] = rune(input[val])
			}
		}
		m = cp
		ct := countLitUpGrid(m)
		fmt.Println("count", i, ct)
	}
	ct := countLitUpGrid(m)
	fmt.Println("end", ct)
}

func getValue(m gridType, i, j int) string {
	k := [2]int{i, j}
	v, ok := m[k]
	if ok && v == '#' {
		return "1"
	}
	return "0"
}

func getNumberValue(m gridType, x, y int) int {
	var sb strings.Builder

	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			sb.WriteString(getValue(m, x+i, y+j))
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

//5868 is wrong
// 6194 is wrong
// 5520 is wrong
// 6316 is wrong
// 5962 is wrong
