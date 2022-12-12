package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/lolopinto/adventofcode2020/grid"
)

// copied from 2021 day 15
// https://en.wikipedia.org/wiki/Dijkstra%27s_algorithm

func day12() {
	lines := readFile("day12input")
	g := grid.NewRuneGrid(lines)

	// g.Print(func(val interface{}) string {
	// 	return string(val.(rune))
	// 	// return fmt.Sprintf("%s", val)
	// })
	// return

	makeKey := func(i, j int) string {
		return fmt.Sprintf("%d-%d", i, j)
	}

	q := make(map[string]bool)
	mins := make(map[string]int)
	unvisitedmins := make(map[string]bool)

	var startPos *grid.Pos
	var endPos *grid.Pos
	for i := 0; i < g.XLength; i++ {
		for j := 0; j < g.YLength; j++ {
			k := makeKey(i, j)
			q[k] = true

			curr := g.At(i, j).Rune()
			if curr == 'S' {
				startPos = &grid.Pos{Row: i, Column: j}
			}
			if curr == 'E' {
				endPos = &grid.Pos{Row: i, Column: j}
			}
		}
	}

	// set initial
	mins[makeKey(startPos.Row, startPos.Column)] = 0

	currPos := startPos
	for len(q) > 0 {
		neighbors := g.Neighbors(currPos.Row, currPos.Column)
		currRune := g.At(currPos.Row, currPos.Column).Rune()

		// there should be something if we're visiting it...
		key := makeKey(currPos.Row, currPos.Column)
		currVal := mins[key]

		for _, v := range neighbors {
			neigh := g.At(v.Row, v.Column)
			if neigh.Visited {
				continue
			}
			neighRune := g.At(v.Row, v.Column).Rune()
			if neighRune == 'E' && currRune != 'z' {
				continue
			}
			if currRune != 'S' {
				if neighRune-currRune > 1 {
					continue
				}
			}

			newMin := currVal + 1

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

		if currRune == 'E' {
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

	lastKey := makeKey(endPos.Row, endPos.Column)
	fmt.Println(mins[lastKey])
}
