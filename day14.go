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

	leftDiagonal := func(x, y int) (int, int) {
		return x - 1, y + 1
	}
	rightDiagonal := func(x, y int) (int, int) {
		return x + 1, y + 1
	}
	moveDown := func(x, y int) (int, int) {
		return x, y + 1
	}

	unitct := 0

	startx := 500
	starty := 0
	startkey := makeKey(startx, starty)
	// main for loop
	done := false
	for {
		if done {
			break
		}
		// new sand
		x := startx
		y := starty
		for {

			if done || y > maxy || m[startkey] == 'S' {
				fmt.Println("done", x, y)
				done = true
				break
			}

			var x2, y2 int
			var key string
			x2, y2 = moveDown(x, y)
			key = makeKey(x2, y2)
			_, ok := m[key]
			// fmt.Println("trying", x2, y2, string(v), ok)

			if !ok {
				x = x2
				y = y2
				continue
			}

			x2, y2 = leftDiagonal(x, y)
			key = makeKey(x2, y2)
			_, ok = m[key]
			// fmt.Println("trying", x2, y2, string(v), ok)
			if !ok {
				x = x2
				y = y2
				continue
			}

			x2, y2 = rightDiagonal(x, y)
			key = makeKey(x2, y2)
			_, ok = m[key]
			// fmt.Println("trying", x2, y2, string(v), ok)
			if !ok {
				x = x2
				y = y2
				continue
			}

			m[makeKey(x, y)] = 'S'
			unitct++
			// fmt.Println("resting", x, y, unitct)
			break

		}

	}
	fmt.Println(unitct)
}
