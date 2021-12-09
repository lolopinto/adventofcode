package main

import (
	"fmt"
	"sort"
)

func day9() {
	lines := readFile("day9input")
	sum := 0
	//	countSeen := 0
	var basinCounts []int
	for i, line := range lines {
		for j, v := range line {
			num := atoi(string(v))
			// this is wrong??
			adj := getAdjacent(line, lines, j, i)
			if num < min(adj) {
				sum += num + 1
				// if countSeen == 2 {
				bc := findBasinCount(line, lines, i, j)
				//				fmt.Println(bc)
				basinCounts = append(basinCounts, bc)
				// }
				// countSeen++

				//				return
			}
		}
	}
	sort.Ints(basinCounts)
	//	fmt.Println(basinCounts)
	multiple := 1
	count := 0

	for i := len(basinCounts) - 1; i > 0; i-- {
		multiple *= basinCounts[i]
		count++
		if count == 3 {
			break
		}
	}
	fmt.Println(multiple)

	//	fmt.Println(sum)
}

func convToNum(c rune) int {
	return atoi(string(c))
}

// this is wrong???
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

func findBasinCount(line string, lines []string, linePos, strPos int) int {
	//	fmt.Println("findBasinCount", line, linePos, strPos)
	// count in both dirs
	lineMap := make(map[int]map[int][]int)
	nextLine := countOnLine(line, strPos)
	lineMap[linePos] = nextLine

	// go up
	prev := linePos
	for i := linePos - 1; i >= 0; i-- {

		totalLine := make(map[int][]int)
		for _, newPos := range lineMap[prev] {
			for _, pos := range newPos {
				if convToNum(rune(lines[i][pos])) == 9 {
					continue
				}
				nextLine := countOnLine(lines[i], pos)
				for k, v := range nextLine {
					totalLine[k] = v
				}
			}
		}
		prev = i
		lineMap[i] = totalLine

	}

	// go down
	prev = linePos

	for i := linePos + 1; i < len(lines); i++ {

		totalLine := make(map[int][]int)
		for _, newPos := range lineMap[prev] {
			for _, pos := range newPos {
				if convToNum(rune(lines[i][pos])) == 9 {
					continue
				}
				nextLine := countOnLine(lines[i], pos)
				for k, v := range nextLine {
					totalLine[k] = v
				}
			}
		}
		prev = i
		lineMap[i] = totalLine

	}
	//	fmt.Println(lineMap)
	count := 0
	for _, v := range lineMap {
		for _, v2 := range v {
			count += len(v2)
		}
	}
	return count
}

func countOnLine(line string, pos int) map[int][]int {
	//	count := 0
	//	var ret []int
	ret := make(map[int][]int)
	// go up
	//	fmt.Println("count on line", line, pos, convToNum(rune(line[pos])))
	for i := pos; i < len(line); i++ {
		val := convToNum(rune(line[i]))
		if val == 9 {
			break
		}
		//		fmt.Println("count going right", line, i, val)
		l, ok := ret[val]
		if !ok {
			l = []int{}
		}
		l = append(l, i)
		ret[val] = l
		//		ret = append(ret, i)
		//		count++
	}

	// can't count yourself twice
	for i := pos - 1; i >= 0; i-- {
		val := convToNum(rune(line[i]))
		if val == 9 {
			break
		}
		//		fmt.Println("count going left", line, i, val)

		//		ret = append(ret, i)
		l, ok := ret[val]
		if !ok {
			l = []int{}
		}
		l = append(l, i)
		ret[val] = l

		//		count++
	}

	return ret
}

// 1699200 too high
// 1179675 wrong. doesn't tell me high or low
// 1157520
// 1168440
// my ish is not deterministic
