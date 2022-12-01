package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/lolopinto/adventofcode2020/grid"
)

// https://en.wikipedia.org/wiki/Dijkstra%27s_algorithm
func day15() {
	lines := readFile("day15input")
	g := grid.NewIntGrid(lines)

	// part2
	g = transformGrid(g)
	// not visited, and infinity

	makeKey := func(i, j int) string {
		return fmt.Sprintf("%d-%d", i, j)
	}
	// initialize queue
	q := make(map[string]bool)
	mins := make(map[string]int)
	unvisitedmins := make(map[string]bool)

	for i := 0; i < g.XLength; i++ {
		for j := 0; j < g.YLength; j++ {
			k := makeKey(i, j)
			q[k] = true
		}
	}
	// set initial
	mins[makeKey(0, 0)] = 0

	currPos := &grid.Pos{Row: 0, Column: 0}
	for len(q) > 0 {
		neighbors := g.Neighbors(currPos.Row, currPos.Column)
		// there should be something if we're visiting it...
		key := makeKey(currPos.Row, currPos.Column)
		currVal := mins[key]

		for _, v := range neighbors {
			neigh := g.At(v.Row, v.Column)
			if neigh.Visited {
				continue
			}
			neighVal := g.At(v.Row, v.Column).Int()
			newMin := currVal + neighVal

			neighKey := makeKey(v.Row, v.Column)
			neighMin := mins[neighKey]

			if neighMin == 0 || newMin < neighMin {
				mins[neighKey] = newMin
				unvisitedmins[neighKey] = true
			}
		}

		// mark visited
		delete(q, key)
		delete(unvisitedmins, key)
		g.At(currPos.Row, currPos.Column).Visited = true

		if currPos.Row == g.XLength && currPos.Column == g.YLength {
			break
		}

		var newCurrPos *grid.Pos
		min := math.MaxInt
		for k := range unvisitedmins {
			v := mins[k]
			if v == 0 {
				continue
			}
			parts := strings.Split(k, "-")
			r := atoi(parts[0])
			c := atoi(parts[1])
			if g.At(r, c).Visited {
				continue
			}
			if v < min {
				min = v
				newCurrPos = &grid.Pos{Row: r, Column: c}
			}
		}
		if newCurrPos != nil {
			currPos = newCurrPos
			continue
		}
	}

	lastKey := makeKey(g.XLength-1, g.YLength-1)
	fmt.Println(mins[lastKey])
}

// part 2
func transformGrid(g *grid.Grid) *grid.Grid {
	initialLength := g.XLength
	// get 1st row
	var r1 []*grid.Grid
	r1 = append(r1, g)
	for i := 0; i < 4; i++ {
		g2 := transform1Grid(g)

		r1 = append(r1, g2)
		g = g2
	}

	var r2 []*grid.Grid
	for i := 0; i < 5; i++ {
		g2 := transform1Grid(r1[i])
		r2 = append(r2, g2)
	}

	var r3 []*grid.Grid
	for i := 0; i < 5; i++ {
		g2 := transform1Grid(r2[i])
		r3 = append(r3, g2)
	}

	var r4 []*grid.Grid
	for i := 0; i < 5; i++ {
		g2 := transform1Grid(r3[i])
		r4 = append(r4, g2)
	}

	var r5 []*grid.Grid
	for i := 0; i < 5; i++ {
		g2 := transform1Grid(r4[i])
		r5 = append(r5, g2)
	}

	ret := grid.NewRectGrid(g.XLength*5, g.YLength*5)
	rows := [][]*grid.Grid{
		r1,
		r2,
		r3,
		r4,
		r5,
	}
	for i, row := range rows {
		for j, g := range row {

			for x := 0; x < g.XLength; x++ {
				for y := 0; y < g.YLength; y++ {
					val := g.At(x, y).Int()

					xpos := i*initialLength + x
					ypos := j*initialLength + y

					ret.At(xpos, ypos).SetValue(val)
				}
			}
		}
	}

	return ret
}

func transform1Grid(g *grid.Grid) *grid.Grid {
	g2 := grid.NewRectGrid(g.XLength, g.YLength)
	for r := 0; r < g.XLength; r++ {
		for c := 0; c < g.YLength; c++ {
			newVal := g.At(r, c).Int() + 1
			if newVal == 10 {
				newVal = 1
			}
			g2.At(r, c).SetValue(newVal)
		}
	}
	return g2
}
