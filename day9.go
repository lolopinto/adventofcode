package main

import (
	"fmt"
	"sort"
)

type basintype map[int]map[int]map[int]bool

var seen [][]bool

func day9() {
	lines := readFile("day9input")
	sum := 0
	var basinCounts []int
	seen = make([][]bool, len(lines))
	for i := range seen {
		seen[i] = make([]bool, len(lines[0]))
	}
	for i, line := range lines {
		for j, v := range line {
			num := atoi(string(v))
			adj := getAdjacent(lines, i, j)
			mincan := make([]int, len(adj))
			for i := range adj {
				mincan[i] = adj[i].val
			}
			if num < min(mincan) {
				//				fmt.Println("lowpoint", num)
				sum += num + 1
				// rewritten approach that worked
				bc := findBasinCountWorking(lines, adj, pos{
					x:   i,
					y:   j,
					val: num,
				})
				// _, bc := findBasinCountFail(line, lines, i, j)
				basinCounts = append(basinCounts, bc)
			}
		}
	}
	sort.Ints(basinCounts)
	multiple := 1
	count := 0

	for i := len(basinCounts) - 1; i > 0; i-- {
		multiple *= basinCounts[i]
		count++
		if count == 3 {
			break
		}
	}

	fmt.Println(sum)
	fmt.Println(multiple)
}

func convToNum(c rune) int {
	return atoi(string(c))
}

type pos struct {
	x, y, val int
}

func newPos(lines []string, x, y int) pos {
	return pos{
		x:   x,
		y:   y,
		val: convToNum(rune(lines[x][y])),
	}
}

func getAdjacent(lines []string, linePos, strPos int) []pos {
	var ret []pos
	if strPos-1 >= 0 {
		ret = append(ret, newPos(lines, linePos, strPos-1))
	}
	if strPos+1 < len(lines[linePos]) {
		ret = append(ret, newPos(lines, linePos, strPos+1))
	}
	if linePos-1 >= 0 {
		ret = append(ret, newPos(lines, linePos-1, strPos))
	}
	if linePos+1 < len(lines) {
		ret = append(ret, newPos(lines, linePos+1, strPos))
	}
	return ret
}

func findBasinCountWorking(lines []string, adj []pos, origPos pos) int {
	if seen[origPos.x][origPos.y] {
		return 0
	}
	count := 1
	seen[origPos.x][origPos.y] = true

	for _, p := range adj {

		next := convToNum(rune(lines[p.x][p.y]))

		if next != 9 {
			adj2 := getAdjacent(lines, p.x, p.y)
			count += findBasinCountWorking(lines, adj2, p)
		}
	}
	return count
}

// everything below here is from failed initial attempt

func findBasinCountFail(line string, lines []string, linePos, strPos int) (basintype, int) {
	//	fmt.Println("findBasinCount", line, linePos, strPos)
	// count in both dirs
	lineMap := make(basintype)
	nextLine := countOnLine(line, strPos)
	lineMap[linePos] = nextLine

	setLineMap := func(i int, totalLine map[int]map[int]bool) {
		value, ok := lineMap[i]
		if !ok {
			value = make(map[int]map[int]bool)
		}
		for k, v := range totalLine {
			v2, ok := value[k]
			if !ok {
				v2 = make(map[int]bool)
			}
			for k2 := range v {
				v2[k2] = true
			}
			value[k] = v2
		}
		lineMap[i] = value
	}
	curNum := convToNum(rune(lines[linePos][strPos]))

	// go up
	prev := linePos
	for i := linePos - 1; i >= 0; i-- {
		totalLine := countNextLine(lines, i, lineMap[prev], curNum)
		prev = i
		setLineMap(i, totalLine)
	}

	// go down
	prev = linePos
	for i := linePos + 1; i < len(lines); i++ {
		totalLine := countNextLine(lines, i, lineMap[prev], curNum)
		prev = i
		setLineMap(i, totalLine)
	}

	count := 0

	for _, v := range lineMap {
		for _, v2 := range v {
			count += len(v2)
		}
	}

	return lineMap, count
}

func countNextLine(lines []string, linePos int, positions map[int]map[int]bool, curNum int) map[int]map[int]bool {
	totalLine := make(map[int]map[int]bool)

	for _, newPos := range positions {
		for pos := range newPos {
			next := convToNum(rune(lines[linePos][pos]))
			if next == 9 {
				continue
			}
			nextLine := countOnLine(lines[linePos], pos)
			for k, v := range nextLine {
				for k2, v2 := range v {
					m, ok := totalLine[k]
					if !ok {
						m = make(map[int]bool)
					}
					m[k2] = v2
					totalLine[k] = m
				}
			}
		}
	}
	return totalLine
}

func countOnLine(line string, pos int) map[int]map[int]bool {
	ret := make(map[int]map[int]bool)

	// go up
	prevVal := convToNum(rune(line[pos]))

	ret[prevVal] = map[int]bool{}

	for i := pos; i < len(line); i++ {
		val := convToNum(rune(line[i]))

		if val == 9 {
			break
		}
		//		fmt.Println("count going right", line, i, val)
		l, ok := ret[val]
		if !ok {
			l = make(map[int]bool)
		}
		l[i] = true
		ret[val] = l
	}

	// can't count yourself twice
	for i := pos - 1; i >= 0; i-- {
		val := convToNum(rune(line[i]))
		if val == 9 {
			break
		}
		l, ok := ret[val]
		if !ok {
			l = make(map[int]bool)
		}
		l[i] = true
		ret[val] = l
	}

	return ret
}
