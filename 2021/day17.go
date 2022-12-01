package main

import (
	"fmt"
	"strings"
)

func day17() {
	input := readFile("day17input")[0]
	input = strings.TrimPrefix(input, "target area: ")
	parts := strings.Split(input, ", ")
	minx, maxx := parseParts(parts[0])
	miny, maxy := parseParts(parts[1])
	//	result := math.MinInt
	ct := 0
	for i := 0; i <= maxx; i++ {
		// what's the right range here lol?
		for j := miny; j < 10000; j++ {
			high := move(i, j, minx, maxx, miny, maxy)
			if high {
				ct++
			}
			// part 1
			// if high > result {
			// 	result = high
			// }
		}
	}
	//	fmt.Println(result)
	fmt.Println(ct)
}

func parseParts(s string) (int, int) {
	s = s[2:]
	parts := strings.Split(s, "..")
	return atoi(parts[0]), atoi(parts[1])
}

func move(startx, starty, minx, maxx, miny, maxy int) bool {
	x, y := 0, 0

	//	highesty := math.MinInt

	//
	for i := 0; i < 1000; i++ {
		x += startx
		y += starty

		// if y > highesty {
		// 	highesty = y
		// }
		//		fmt.Println(x, y)

		if x >= minx && x <= maxx && y >= miny && y <= maxy {
			return true
		}

		// what's the right logic here? does it matter?
		// if x > maxx || y < miny {
		// 	//			fmt.Println("here2")
		// 	//			break
		// }
		if startx > 0 {
			startx -= 1
		} else if startx < 0 {
			startx += 1
		}
		starty -= 1
	}

	// if hittarget {
	// 	return highesty
	// }
	// return 0
	return false
}
