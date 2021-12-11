package main

import (
	"fmt"

	"github.com/lolopinto/adventofcode2020/grid"
)

func day11() {
	lines := readFile("day11input")

	g := grid.NewIntGrid(lines)
	length := len(lines)

	flashct := 0
	iterations := 0
	for {
		for r := 0; r < length; r++ {
			for c := 0; c < length; c++ {
				currData := g.At(r, c)

				if currData.Visited {
					continue
				}

				val := currData.Int() + 1
				currData.SetValue(val)
				if val > 9 {
					ct := flash(g, r, c, length)
					if iterations < 100 {
						flashct += ct
					}
				}
			}
		}

		g2 := grid.NewGrid(length)
		allzero := 0

		// copy values. keep visited false
		for i := 0; i < length; i++ {
			for j := 0; j < length; j++ {
				val := g.At(i, j).Int()
				g2.At(i, j).SetValue(val)
				if val == 0 {
					allzero++
				}
			}
		}

		// part 2
		if allzero == length*length {
			fmt.Println(iterations + 1)
			break
		}
		iterations++
		g = g2
	}
	// part1
	fmt.Println(flashct)
}

func flash(g *grid.Grid, r, c, length int) int {

	data := g.At(r, c)
	if data.Visited {
		return 0
	}
	data.Visited = true
	data.SetValue(0)

	ct := 1

	for _, pos := range g.Neighbors8(r, c) {
		// already flashed. nothing to do here
		neighbor := g.At(pos.Row, pos.Column)
		if neighbor.Visited {
			continue
		}

		val := neighbor.Int() + 1
		neighbor.SetValue(val)

		if val > 9 {
			ct += flash(g, pos.Row, pos.Column, length)
		}
	}

	return ct
}
