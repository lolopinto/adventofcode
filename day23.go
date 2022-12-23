package main

import (
	"fmt"

	"github.com/lolopinto/adventofcode2020/grid"
)

type direction string

const (
	North     direction = "N"
	South     direction = "S"
	East      direction = "E"
	West      direction = "W"
	NorthEast direction = "NE"
	NorthWest direction = "NW"
	SouthEast direction = "SE"
	SouthWest direction = "SW"
)

var deltas = map[direction]grid.Pos{
	North:     grid.NewPos(-1, 0),
	South:     grid.NewPos(1, 0),
	East:      grid.NewPos(0, 1),
	West:      grid.NewPos(0, -1),
	NorthEast: grid.NewPos(-1, 1),
	NorthWest: grid.NewPos(-1, -1),
	SouthEast: grid.NewPos(1, 1),
	SouthWest: grid.NewPos(1, -1),
}

var directionsConsidered = map[direction][]direction{
	North: {
		North,
		NorthEast,
		NorthWest,
	},
	South: {
		South,
		SouthEast,
		SouthWest,
	},
	West: {
		West,
		NorthWest,
		SouthWest,
	},
	East: {
		East,
		NorthEast,
		SouthEast,
	},
}

func day23() {
	lines := readFile("day23input")

	m := map[grid.Pos]rune{}

	priorities := []direction{
		North,
		South,
		West,
		East,
	}
	for r, line := range lines {
		for c, val := range line {
			if val != '#' {
				continue
			}
			pos := grid.NewPos(r, c)
			m[pos] = val
		}
	}

	getMinMax := func() (int, int, int, int) {
		rows := []int{}
		cols := []int{}
		for k := range m {
			rows = append(rows, k.Row)
			cols = append(cols, k.Column)
		}
		// fmt.Println(m)
		minr := min(rows)
		minc := min(cols)
		maxr := max(rows)
		maxc := max(cols)
		return minr, maxr, minc, maxc
	}
	print := func() {

		// for r := -2; r <= len(lines)+3; r++ {
		// 	for c := -3; c <= len(lines[0])+3; c++ {
		// 		v := m[grid.NewPos(r, c)]
		// 		if v == '#' {
		// 			fmt.Print(string(v))
		// 		} else {
		// 			fmt.Print(".")
		// 		}
		// 	}
		// 	fmt.Println()
		// }
	}

	print()
	for i := 0; i < 10; i++ {
		fmt.Println("round", i+1)
		priorities = doRound(m, priorities)
		print()
	}
	minr, maxr, minc, maxc := getMinMax()

	ct := 0
	for r := minr; r <= maxr; r++ {
		for c := minc; c <= maxc; c++ {
			_, ok := m[grid.NewPos(r, c)]
			if !ok {
				ct++
			}
		}
	}
	// fmt.Println(len(m), minr, maxr, minc, maxc)
	fmt.Println(ct)
}

func doRound(m map[grid.Pos]rune, priorities []direction) []direction {
	proposalsMapping := map[grid.Pos]grid.Pos{}
	proposalsTo := map[grid.Pos]int{}

	// first half
	for pos := range m {

		found := false
		for _, delta := range deltas {
			pos2 := pos.Add(delta)
			v, ok := m[pos2]
			if ok && v == '#' {
				found = true
				break
			}
		}

		// no other elves in this position, nothing to do here
		if !found {
			// fmt.Println("no elf, elf does nothing")
			continue
		}

		var proposedDir direction
		for _, dir := range priorities {
			elfFound := false
			for _, toCheck := range directionsConsidered[dir] {
				pos2 := pos.Add(deltas[toCheck])
				v, ok := m[pos2]
				if ok && v == '#' {
					elfFound = true
					// fmt.Println("elf found @ ", pos2, " for ", pos)
					break
				}
			}
			if !elfFound {
				proposedDir = dir
				break
			}
		}
		// no proposal
		if proposedDir == "" {
			continue
		}
		proposal := pos.Add(deltas[proposedDir])
		// fmt.Println(pos, "proposing", proposedDir, proposal)
		// set proposal for elf
		proposalsMapping[pos] = proposal

		ct := proposalsTo[proposal]
		ct++
		proposalsTo[proposal] = ct
	}

	for pos, newPos := range proposalsMapping {
		ct := proposalsTo[newPos]

		if ct != 1 {
			// fmt.Println("exiting count")
			continue
		}

		// fmt.Println("moving")
		delete(m, pos)
		m[newPos] = '#'
	}

	newLast := priorities[0]
	priorities = append([]direction{}, priorities[1:]...)
	priorities = append(priorities, newLast)
	return priorities
}
