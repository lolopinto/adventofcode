package main

import (
	"log"
	"math"
	"sort"
)

type boardingPass struct {
	row    int
	column int
}

func (bp boardingPass) seatID() int {
	return bp.row*8 + bp.column
}

func parseBoardingPass(line string) boardingPass {
	if len(line) != 10 {
		log.Fatalf("invalid boarding pass %s", line)
	}

	getVal := func(str string, pow int, hiR, lowR rune) int {
		hi := int(math.Pow(2, float64(pow)))
		list := make([]int, hi)
		for i := 0; i < hi; i++ {
			list[i] = i
		}
		for _, c := range str {
			switch c {
			case hiR:
				list = list[len(list)/2:]
				break
			case lowR:
				list = list[0 : len(list)/2]
				break
			}
		}
		//		log.Println(list)
		return list[0]
	}
	// F -> lower half
	// B -> higher half
	row := getVal(line[0:7], 7, 'B', 'F')
	// L -> lower half
	// R -> higher half
	col := getVal(line[7:], 3, 'R', 'L')

	return boardingPass{
		column: col,
		row:    row,
	}
}

func day5() {
	lines := readFile("day5input")

	sorted := make([]int, len(lines))
	for idx, line := range lines {
		bp := parseBoardingPass(line)
		id := bp.seatID()
		sorted[idx] = id
	}
	sort.Ints(sorted)

	last := sorted[0]
	for i := range sorted {
		last = sorted[i]
		if i == 0 || i+1 == len(sorted) {
			continue
		}
		if sorted[i+1] != last+1 {
			log.Println(last + 1)
		}
	}
	//	log.Println(sorted)
}
