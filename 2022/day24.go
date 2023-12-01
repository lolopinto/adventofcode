package main

import (
	"fmt"

	"github.com/lolopinto/adventofcode2020/grid"
)

type valleyInfo struct {
	walls     map[grid.Pos]bool
	blizzards map[grid.Pos][]byte
	xLength   int
	yLength   int
}

func newValley(lines []string) *valleyInfo {
	g := grid.NewRectGrid(len(lines), len(lines[0]))

	walls := map[grid.Pos]bool{
		// fake walls near the beginning and end so we don't go farther than we should
		grid.NewPos(-1, 1):                  true,
		grid.NewPos(g.XLength, g.YLength-2): true,
	}
	blizarrds := map[grid.Pos][]byte{}

	for i := 0; i < g.XLength; i++ {
		for j := 0; j < g.YLength; j++ {
			val := lines[i][j]

			pos := grid.NewPos(i, j)
			if val == '#' {
				walls[pos] = true
				continue
			}

			_, ok := blizzardDeltas[val]
			if ok {
				l := blizarrds[pos]
				l = append(l, val)
				blizarrds[pos] = l
				continue
			}
		}
	}

	return &valleyInfo{
		walls:     walls,
		blizzards: blizarrds,
		xLength:   g.XLength,
		yLength:   g.YLength,
	}
}

var blizzardDeltas = map[byte]grid.Pos{
	'^': grid.NewPos(-1, 0),
	'v': grid.NewPos(1, 0),
	'>': grid.NewPos(0, 1),
	'<': grid.NewPos(0, -1),
}

func day24() {
	lines := readFile("day24input")
	v := newValley(lines)

	start := grid.NewPos(0, 1)
	end := grid.NewPos(v.xLength-1, v.yLength-2)

	part1, b2 := moveThroughBlizzards(v.walls, v.blizzards, start, end)
	part1b, b3 := moveThroughBlizzards(v.walls, b2, end, start)
	part1c, _ := moveThroughBlizzards(v.walls, b3, start, end)
	fmt.Println("part1", part1)
	fmt.Println("part2", part1+part1b+part1c)
}

// bfs till you find it.
// i need to use bfs more often for these things.
// simpler to grok than the dfs appraoch I was trying
func moveThroughBlizzards(walls map[grid.Pos]bool, blizzards map[grid.Pos][]byte, start, end grid.Pos) (int, map[grid.Pos][]byte) {
	locations := map[grid.Pos]bool{
		start: true,
	}
	minute := 0

	for {
		if locations[end] {
			break
		}
		minute++

		newBlizzards := map[grid.Pos][]byte{}

		for pos, list := range blizzards {
			for _, blizzard := range list {

				delta := blizzardDeltas[blizzard]
				newPos := pos.Add(delta)

				// instead of doing math to figure out where to go, just go all the way back until we hit a wall and then move
				// forward one step
				// math the correct thing to do but i wasn't getting it to work on real input only sample
				_, ok := walls[newPos]
				if ok {
					newPos = newPos.Sub(delta)
					for {
						_, ok := walls[newPos]
						if ok {
							break
						}
						newPos = newPos.Sub(delta)
					}
					newPos = newPos.Add(delta)
				}
				l := newBlizzards[newPos]
				l = append(l, blizzard)
				newBlizzards[newPos] = l
			}
		}

		newLocations := map[grid.Pos]bool{}
		for currLocation := range locations {
			for _, delta := range blizzardDeltas {
				newCurr := currLocation.Add(delta)

				_, ok := walls[newCurr]
				if ok {
					continue
				}
				if len(newBlizzards[newCurr]) != 0 {
					continue
				}

				newLocations[newCurr] = true
			}
			_, ok := walls[currLocation]
			if !ok && newBlizzards[currLocation] == nil {
				newLocations[currLocation] = true
			}
		}

		blizzards = newBlizzards
		locations = newLocations
	}
	return minute, blizzards
}
