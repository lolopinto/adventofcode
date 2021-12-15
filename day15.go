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
	// not visited, and infinity
	mins := grid.NewGrid(g.Length)
	mins.At(0, 0).SetValue(0)

	makeKey := func(i, j int) string {
		return fmt.Sprintf("%d-%d", i, j)
	}
	// initialize queue
	q := make(map[string]bool)
	for i := 0; i < g.Length; i++ {
		for j := 0; j < g.Length; j++ {
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
