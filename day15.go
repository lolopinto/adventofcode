package main

import (
	"fmt"
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

	mins := grid.NewRectGrid(g.XLength, g.YLength)
	mins.At(0, 0).SetValue(0)

	makeKey := func(i, j int) string {
		return fmt.Sprintf("%d-%d", i, j)
	}
	// initialize queue
	q := make(map[string]bool)
	for i := 0; i < g.XLength; i++ {
		for j := 0; j < g.YLength; j++ {
			k := makeKey(i, j)
			q[k] = true
		}
	}
	currPos := grid.Pos{Row: 0, Column: 0}
	for len(q) > 0 {
		neighbors := mins.Neighbors(currPos.Row, currPos.Column)
		// there should be something if we're visiting it...
		currVal := mins.At(currPos.Row, currPos.Column).Int()

		for _, v := range neighbors {
			neighVal := g.At(v.Row, v.Column).Int()
			newMin := currVal + neighVal

			neigh := mins.At(v.Row, v.Column)

			neighMin, ok := neigh.Data().(int)
			if !ok || newMin < neighMin {
				neigh.SetValue(newMin)
			}
		}

		// mark visited
		delete(q, makeKey(currPos.Row, currPos.Column))

		// done
		if currPos.Row == g.XLength && currPos.Column == g.YLength {
			break
		}
		var unvisited []int
		for k := range q {
			parts := strings.Split(k, "-")
			r := atoi(parts[0])
			c := atoi(parts[1])
			v, ok := mins.At(r, c).Data().(int)
			if !ok {
				continue
			}
			unvisited = append(unvisited, v)
		}
		minunivisted := min(unvisited)
		for k := range q {
			parts := strings.Split(k, "-")
			r := atoi(parts[0])
			c := atoi(parts[1])
			v, ok := mins.At(r, c).Data().(int)
			if !ok {
				continue
			}
			if v == minunivisted {
				currPos = grid.Pos{Row: r, Column: c}
			}
		}

	}

	last := mins.At(g.XLength-1, g.YLength-1).Int()
	fmt.Println(last)
}

// part 2
func transformGrid(g *grid.Grid) *grid.Grid {
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

					xpos := i*10 + x
					ypos := j*10 + y

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
