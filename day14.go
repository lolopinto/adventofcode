package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/lolopinto/adventofcode2020/grid"
)

func day14() {

	lines := readFile("day14input")

	m := make(map[grid.Pos]rune)

	makePos := func(s string) grid.Pos {
		pair := splitLength(s, ",", 2)
		return grid.Pos{Row: atoi(pair[0]), Column: atoi(pair[1])}
	}

	fillRock := func(start, end grid.Pos) int {
		endy := -1
		start.Line(&end, func(p *grid.Pos) {
			m[*p] = 'R'
			if p.Column > endy {
				endy = p.Column
			}
		})

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

	infinitey := maxy + 2

	leftDiagonal := func(x, y int) grid.Pos {
		return grid.NewPos(x-1, y+1)
	}
	rightDiagonal := func(x, y int) grid.Pos {
		return grid.NewPos(x+1, y+1)
	}
	moveDown := func(x, y int) grid.Pos {
		return grid.NewPos(x, y+1)
	}

	unitct := 0

	startx := 500
	starty := 0
	startPos := grid.NewPos(startx, starty)

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

			if m[startPos] == 'S' {
				done = true
				break
			}

			// part 2
			if y+1 == infinitey {
				m[grid.NewPos(x, y)] = 'S'
				unitct++
				break
			}

			if y > maxy {
				done = true
				break
			}

			toTry := []grid.Pos{
				moveDown(x, y),
				leftDiagonal(x, y),
				rightDiagonal(x, y),
			}
			cont := false
			for _, t := range toTry {
				_, ok := m[t]
				if !ok {
					x = t.Row
					y = t.Column
					cont = true
					break
				}
			}
			if cont {
				continue
			}

			m[grid.NewPos(x, y)] = 'S'
			unitct++
			break
		}
	}
	fmt.Println(unitct)
}
