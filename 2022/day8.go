package main

import (
	"fmt"
	"math"

	"github.com/lolopinto/adventofcode2020/grid"
)

func day8() {
	lines := readFile("day8input")

	g := grid.NewIntGrid(lines)

	visibleCt := 0

	// part 1
	for r := 0; r < g.XLength; r++ {
		for c := 0; c < g.YLength; c++ {
			candidates := [][]grid.Pos{
				g.Top(r, c),
				g.Bottom(r, c),
				g.Right(r, c),
				g.Left(r, c),
			}
			curr := g.At(r, c).Int()
			for _, cand := range candidates {
				visible := true
				for _, pos := range cand {
					comp := g.At(pos.Row, pos.Column).Int()
					if comp >= curr {
						visible = false
						break
					}
				}
				if visible {
					visibleCt++
					break
				}
			}
		}
	}
	fmt.Println(visibleCt)

	// part2
	max := math.MinInt

	for r := 0; r < g.XLength; r++ {
		for c := 0; c < g.YLength; c++ {
			candidates := [][]grid.Pos{
				g.Top(r, c),
				g.Bottom(r, c),
				g.Right(r, c),
				g.Left(r, c),
			}
			curr := g.At(r, c).Int()
			mult := 1
			for _, cand := range candidates {

				ct := 0
				for _, pos := range cand {
					comp := g.At(pos.Row, pos.Column).Int()
					ct++
					if comp >= curr {
						break
					}
				}
				mult *= ct
			}
			if mult > max {
				max = mult
			}
		}
	}

	fmt.Println(max)
}
