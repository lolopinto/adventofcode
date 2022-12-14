package main

import (
	"fmt"
	"math"
	"strings"
)

func day14() {

	lines := readFile("day14input")

	makeKey := func(i, j int) string {
		return fmt.Sprintf("%d-%d", i, j)
	}

	m := make(map[string]rune)

	makePos := func(s string) [2]int {
		pair := splitLength(s, ",", 2)
		return [2]int{atoi(pair[0]), atoi(pair[1])}
	}

	fillRock := func(start, end [2]int) int {
		startx, starty, endx, endy := 0, 0, 0, 0

		//
		if start[0] == end[0] {
			startx = start[0]
			endx = end[0]
			if start[1] < end[1] {
				starty = start[1]
				endy = end[1]
			} else {
				starty = end[1]
				endy = start[1]
			}
		} else {
			starty = start[1]
			endy = end[1]
			if start[0] < end[0] {
				startx = start[0]
				endx = end[0]
			} else {
				startx = end[0]
				endx = start[0]
			}
		}

		for x := startx; x <= endx; x++ {
			for y := starty; y <= endy; y++ {
				m[makeKey(x, y)] = 'R'
			}
		}
		return endy
	}

	maxy := math.MinInt
	for _, line := range lines {
		parts := strings.Split(line, " -> ")
		for _, pair := range windowed(parts, 2) {
			start := makePos(pair[0])
			end := makePos(pair[1])

			endy := fillRock(start, end)
			if endy > maxy {
				maxy = endy
			}
		}
	}

	canRest := func(x, y int) bool {
		y++
		_, ok := m[makeKey(x, y)]
		return ok
	}
	leftDiagonal := func(x, y int) (int, int) {
		return x - 1, y + 1
	}
	rightDiagonal := func(x, y int) (int, int) {
		return x + 1, y + 1
	}
	moveDown := func(x, y int) (int, int) {
		return x, y + 1
	}

	// canTryLeft := func(x, y int) (bool, int, int) {
	// 	x--
	// 	y++
	// 	_, ok := m[makeKey(x, y)]
	// 	if ok || !canRest(x, y) {
	// 		return false, x, y
	// 	}
	// 	m[makeKey(x, y)] = 'S'
	// 	fmt.Println("resting left", x, y)
	// 	return true, x, y
	// }
	// canTryRight := func(x, y int) (bool, int, int) {
	// 	x++
	// 	y++
	// 	_, ok := m[makeKey(x, y)]
	// 	if ok || !canRest(x, y) {
	// 		return false, x, y
	// 	}
	// 	m[makeKey(x, y)] = 'S'
	// 	fmt.Println("resting right", x, y)
	// 	return true, x, y
	// }

	unitct := 0

	startx := 500
	starty := 0
	startkey := makeKey(startx, starty)
	// main for loop
	for {
		done := false
		if m[startkey] == 'S' {
			// TODO BS
			// if unitct == 5 {

			fmt.Println("TODO as to when to stop")
			break
		}

		// for {
		x := startx
		y := starty

		fns := []func(int, int) (int, int){
			moveDown,
			leftDiagonal,
			rightDiagonal,
		}

		// for each direction, do...
		for idx, fn := range fns {

			// do this direction while we can

			for {

				x2, y2 := fn(x, y)
				key := makeKey(x2, y2)
				v, ok := m[key]
				fmt.Println("trying idx", idx, x2, y2, string(v), ok)

				// went this low, fail?
				if y2 > maxy {
					fmt.Println("breaking fail", x2, y2, x, y)
					done = true
					break
				}

				// nothing there, continue
				// only update this if it fails...
				if !ok {
					// if unitct >= 22 {
					// 	fmt.Println("trying for 23", x2, y2)
					// }
					x = x2
					y = y2
					// fmt.Println("assinging next dir", x, y)
					continue
				}
				fmt.Println("breakkking", x2, y2, x, y, string(v))
				break
			}
		}

		if canRest(x, y) {
			m[makeKey(x, y)] = 'S'
			unitct++
			fmt.Println("resting", x, y, unitct)
			continue
		} else {
			fmt.Println("cannot rest", x, y)
			// done = true
		}
		// done = true
		// break
		// }
		if done {
			break
		}
	}
	fmt.Println(m)
}
