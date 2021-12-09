package main

import (
	"fmt"
	"sort"
)

type basintype map[int]map[int]map[int]bool

func day9() {
	lines := readFile("day9input")
	sum := 0
	//	countSeen := 0
	var basinCounts []int
	var basins []basintype
	for i, line := range lines {
		for j, v := range line {
			num := atoi(string(v))
			// this is wrong??
			adj := getAdjacent(line, lines, j, i)
			if num < min(adj) {
				sum += num + 1
				// if countSeen == 2 {
				basin, bc := findBasinCount(line, lines, i, j)
				basins = append(basins, basin)
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
	//	fmt.Println(basinCounts)
	fmt.Println(multiple)

	for _, b := range basins {
		fmt.Println(b)
	}

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

func countNextLine(lines []string, linePos int, positions map[int]map[int]bool) map[int]map[int]bool {
	totalLine := make(map[int]map[int]bool)

	for _, newPos := range positions {
		for pos := range newPos {
			if convToNum(rune(lines[linePos][pos])) == 9 {
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
				//				totalLine[k] = v
			}
		}
	}
	return totalLine
}

func findBasinCount(line string, lines []string, linePos, strPos int) (basintype, int) {
	//	fmt.Println("findBasinCount", line, linePos, strPos)
	// count in both dirs
	lineMap := make(basintype)
	nextLine := countOnLine(line, strPos)
	lineMap[linePos] = nextLine

	// go up
	prev := linePos
	for i := linePos - 1; i >= 0; i-- {
		totalLine := countNextLine(lines, i, lineMap[prev])
		prev = i
		lineMap[i] = totalLine
	}

	// go down
	prev = linePos

	for i := linePos + 1; i < len(lines); i++ {
		totalLine := countNextLine(lines, i, lineMap[prev])
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
	// if count > 100 {
	// 	fmt.Println(lineMap)
	// }
	return lineMap, count
}

func countOnLine(line string, pos int) map[int]map[int]bool {
	ret := make(map[int]map[int]bool)
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

// 1699200 too high
// 1179675 wrong. doesn't tell me high or low
// 1157520
// 1168440
// my ish is not deterministic

//1168650 also wrong?

// [2 2 2 2 2 2 2 3 3 3 3 3 3 3 3 3 3 4 4 4 4 5 5 5 5 5 5 5 5 5 6 6 6 7 7 7 7 8 8 8 8 8 8 9 9 9 9 10 10 10 11 12 13 13 13 13 14 14 15 15 15 15 15 16 16 17 17 17 18 18 18 18 18 18 19 19 19 19 19 20 20 20 20 21 21 21 21 21 22 22 22 22 23 23 23 24 25 25 26 26 27 27 28 28 28 29 29 29 31 31 31 31 31 31 32 32 32 32 32 32 33 34 34 34 34 34 34 34 35 36 36 37 37 37 38 38 38 38 39 39 39 39 40 40 40 40 41 41 41 41 42 42 43 44 45 45 46 47 49 49 50 50 51 51 53 54 55 56 56 57 57 58 58 58 60 60 60 61 61 62 63 64 65 65 65 66 69 69 70 70 71 71 72 74 75 77 78 78 79 79 79 80 80 80 82 83 83 87 87 102 104 104 107]
// 1157312

// [2 2 2 2 2 2 2 3 3 3 3 3 3 3 3 3 3 4 4 4 4 5 5 5 5 5 5 5 5 5 6 6 6 7 7 7 7 8 8 8 8 8 8 9 9 9 9 10 10 10 11 12 13 13 13 13 14 14 15 15 15 15 15 16 16 17 17 17 18 18 18 18 18 18 19 19 19 19 19 20 20 20 20 21 21 21 21 21 22 22 22 22 23 23 23 24 25 25 26 26 27 27 28 28 28 29 29 29 30 31 31 31 31 31 32 32 32 32 32 32 33 33 34 34 34 34 34 34 36 36 37 37 37 37 38 38 38 39 39 39 39 39 39 40 40 40 41 41 41 41 42 42 44 45 45 45 46 48 49 49 50 51 51 51 53 54 55 56 56 57 57 58 58 58 60 60 61 61 61 63 64 64 65 65 65 66 68 69 70 70 71 71 72 74 75 76 78 78 78 79 79 80 80 80 82 83 84 87 87 102 105 105 107]
// 1179675

// [2 2 2 2 2 2 2 3 3 3 3 3 3 3 3 3 3 4 4 4 4 5 5 5 5 5 5 5 5 5 6 6 6 7 7 7 7 8 8 8 8 8 8 9 9 9 9 10 10 10 11 12 13 13 13 13 14 14 15 15 15 15 15 16 16 17 17 17 18 18 18 18 18 18 19 19 19 19 19 20 20 20 20 21 21 21 21 21 22 22 22 22 23 23 23 24 25 25 26 26 27 27 28 28 28 29 29 29 31 31 31 31 31 31 32 32 32 32 32 32 33 33 34 34 34 34 34 34 35 36 36 36 37 37 38 38 38 39 39 39 39 40 40 40 40 41 41 41 41 41 41 42 44 45 45 45 46 47 49 50 50 51 51 52 53 54 55 56 56 56 57 57 58 58 60 60 60 61 61 64 64 64 65 65 65 66 69 69 70 70 71 71 72 74 75 76 78 78 78 79 79 80 80 80 82 83 84 86 86 102 105 105 106]
// 1168650

// also wrong
// [2 2 2 2 2 2 2 3 3 3 3 3 3 3 3 3 3 4 4 4 4 5 5 5 5 5 5 5 5 5 6 6 6 7 7 7 7 8 8 8 8 8 9 9 9 9 9 10 10 10 11 12 13 13 13 14 14 15 15 15 15 15 15 16 16 17 17 18 18 18 18 18 18 18 19 19 19 19 20 20 20 20 20 21 21 21 21 22 22 22 23 23 23 23 24 25 25 27 27 27 27 28 28 29 29 29 30 31 31 31 31 31 31 32 32 32 32 32 33 33 34 34 34 35 35 35 35 37 37 37 37 38 38 38 38 39 39 40 40 40 40 40 41 41 41 42 42 42 42 43 43 44 45 45 47 47 48 50 52 52 55 55 57 57 58 58 58 59 59 59 60 60 60 61 62 62 65 65 65 65 66 66 67 68 70 70 71 72 74 74 75 75 75 76 77 79 80 81 81 82 83 84 84 86 86 86 86 88 90 104 108 109 111]
// 1306692
