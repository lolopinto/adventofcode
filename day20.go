package main

import (
	"fmt"
	"strings"
)

type gridType map[[2]int]rune

func day20() {
	parts := readFileChunks("day20input", 2)
	input := parts[0][0]
	//	fmt.Println(parts[0])
	gridInput := parts[1]
	//	length := len(parts[1][0])
	//	mid := int(math.Floor(float64(length / 2)))

	m := make(gridType)
	// initialize grid
	for i := 0; i < len(gridInput); i++ {
		for j := 0; j < len(gridInput); j++ {
			m[[2]int{i, j}] = rune(gridInput[i][j])
		}
	}

	_, x1, y1 := gridStats(m)
	fmt.Println(x1, y1)

	cp := makeMapCopy(m)
	for x := x1[0]; x < x1[1]; x++ {
		for y := y1[0]; y < y1[1]; y++ {
			val := getNumberValue(m, x, y)
			cp[[2]int{x, y}] = rune(input[val])
		}
	}

	ct1, x2, y2 := gridStats(cp)
	fmt.Println(ct1, x2, y2)

	cp2 := makeMapCopy(cp)
	for x := x2[0]; x < x2[1]; x++ {
		for y := y2[0]; y < y2[1]; y++ {
			val := getNumberValue(cp, x, y)
			cp2[[2]int{x, y}] = rune(input[val])
		}
	}
	fmt.Println(gridStats(cp2))

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

func makeMapCopy(m gridType) gridType {
	ret := make(gridType)
	for k, v := range m {
		ret[k] = v
	}
	return ret
}

// count lit up, x range, y range
func gridStats(m gridType) (int, [2]int, [2]int) {
	ct := 0
	var xes []int
	var yes []int
	for k, v := range m {
		if v == '#' {
			ct++
		}
		xes = append(xes, k[0])
		yes = append(yes, k[1])
	}
	// fmt.Println(min(xes), max(xes))
	// fmt.Println(min(yes), max(yes))
	return ct, [2]int{min(xes) - 2, max(xes) + 2}, [2]int{min(yes) - 2, max(yes) + 2}
}

//5868 is wrong
// 6194 is wrong
// 5520 is wrong
// 6316 is wrong
// 5962 is wrong
