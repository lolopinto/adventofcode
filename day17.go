package main

import (
	"fmt"
	"math"
	"strings"
)

func day17() {
	input := readFile("day17input")[0]
	input = strings.TrimPrefix(input, "target area: ")
	parts := strings.Split(input, ", ")
	// x,y flipped
	minx, maxx := parseParts(parts[0])
	miny, maxy := parseParts(parts[1])
	//	fmt.Println(minx, maxx, miny, maxy)
	// what's the range to check??
	result := math.MinInt
	for i := 0; i < minx; i++ {
		for j := maxy; j < 10000; j++ {
			high := move(i, j, minx, maxx, miny, maxy)
			if high > result {
				result = high
			}
		}
	}
	// 7,2 and 6,9 correct
	// high := move(17, -4, minx, maxx, miny, maxy)
	// fmt.Println(high)
	fmt.Println(result)
}

func parseParts(s string) (int, int) {
	s = s[2:]
	parts := strings.Split(s, "..")
	return atoi(parts[0]), atoi(parts[1])
}

func move(startx, starty, minx, maxx, miny, maxy int) int {
	x, y := 0, 0
	//	i := 0
	highesty := math.MinInt
	//	startx +=x
	hittarget := false
	for {
		x += startx
		y += starty

		if y > highesty {
			highesty = y
		}
		//		fmt.Println(x, y)

		if x >= minx && x <= maxx && y >= miny && y <= maxy {
			// target area or didn't find
			hittarget = true
			//			fmt.Println("here")
			break
		}

		// didn't hit
		if x > maxx || y < miny {
			//			fmt.Println("here2")
			break
		}
		if startx > 0 {
			startx -= 1
		} else if startx < 0 {
			startx += 1
		}
		starty -= 1
	}
	//	newx :=
	// 0,0 => 7,2 => 6,1 => 5,0 => 4,-1 => 3,-2, 2,-3 =>1,-4
	// trajectory...

	// 0,0 => 6,3 => 5,2 => 4,1 => 3,0 => 2,-1 => 1,-2 => 0, -3 => 0,4

	// 0,0 => 17, -4
	if hittarget {
		return highesty
	}
	return 0
}
