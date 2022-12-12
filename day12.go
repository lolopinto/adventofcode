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
	// })

	makeKey := func(i, j int) string {
		return fmt.Sprintf("%d-%d", i, j)
	}

	elevations := map[rune]int{
		'S': 1,
	}
	for c := 'a'; c <= 'z'; c++ {
		elevations[c] = int(c-'a') + 1
	}

	var q map[string]bool
	var mins map[string]int
	var unvisitedmins map[string]bool

	initIsh := func() {

		q = make(map[string]bool)
		mins = make(map[string]int)
		unvisitedmins = make(map[string]bool)

		for i := 0; i < g.XLength; i++ {
			for j := 0; j < g.YLength; j++ {
				k := makeKey(i, j)
				q[k] = true
				g.At(i, j).Visited = false
			}
		}
	}

	var startPoses []*grid.Pos
	var endPos *grid.Pos
	for i := 0; i < g.XLength; i++ {
		for j := 0; j < g.YLength; j++ {
			curr := g.At(i, j).Rune()
			if curr == 'S' || curr == 'a' {
				startPoses = append(startPoses, &grid.Pos{Row: i, Column: j})
			}
			if curr == 'E' {
				endPos = &grid.Pos{Row: i, Column: j}
			}
		}
	}
	fmt.Println("total starts", len(startPoses))

	min := math.MaxInt

	for _, startPos := range startPoses {
		initIsh()
		// fmt.Println("startPos", startPos)

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
				neighRune := g.At(v.Row, v.Column).Rune()

				if neighRune == 'E' && currRune != 'z' {
					continue
				}
				if neigh.Visited {
					// fmt.Println("visited continue", currPos, v.Row, v.Column, string(currRune), string(neighRune))
					continue
				}

				currElevation := elevations[currRune]
				neighElevation := elevations[neighRune]
				if neighElevation-currElevation > 1 {
					// fmt.Println("elevation continue", currPos, v.Row, v.Column, string(currRune), string(neighRune))

					continue
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
			} else {
				// why is this needed?
				break
			}
		}
		lastKey := makeKey(endPos.Row, endPos.Column)
		// cheating?
		if mins[lastKey] < min && mins[lastKey] != 0 {
			min = mins[lastKey]
		}
	}

	fmt.Println("answer", min)
}
