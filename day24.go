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
			// we only set directions
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
	fmt.Println(start, end)

	// fmt.Println(g.XLength, g.YLength)

	locations := map[grid.Pos]bool{
		start: true,
	}

	minute := 0
	walls := v.walls
	blizzards := v.blizzards

	for {
		// fmt.Println(minute, locations)
		if locations[end] {
			break
		}
		minute++

		newBlizzards := map[grid.Pos][]byte{}
		// g2 := grid.NewRectGrid(g.XLength, g.YLength)

		for pos, list := range blizzards {
			for _, blizzard := range list {

				delta := blizzardDeltas[blizzard]
				newPos := pos.Add(delta)
				// log := false
				// if pos == grid.NewPos(1, 5) && newPos == grid.NewPos(0, 5) {
				// 	log = true
				// 	fmt.Println("newpos")
				// }

				_, ok := walls[newPos]
				if ok {

					newPos = newPos.Sub(delta)
					// fmt.Println("looping")
					for {
						_, ok := walls[newPos]
						if ok {
							break
						}
						newPos = newPos.Sub(delta)
					}
					newPos = newPos.Add(delta)

					// if newPos.Row == 0 {
					// 	newPos.Row = v.xLength - 1
					// } else if newPos.Row == v.xLength-1 {
					// 	newPos.Row = 1
					// }
					// if newPos.Column == 0 {
					// 	newPos.Column = v.yLength - 2
					// } else if newPos.Column == v.yLength-1 {
					// 	newPos.Column = 1
					// }
				}
				// }
				// if log {
				// 	fmt.Println(newPos)
				// }
				// fmt.Println(pos, "moving:", string(blizzard), newPos)

				l := newBlizzards[newPos]
				l = append(l, blizzard)
				newBlizzards[newPos] = l
			}
		}

		newLocations := map[grid.Pos]bool{}
		for currLocation := range locations {

			for _, delta := range blizzardDeltas {
				newCurr := currLocation.Add(delta)

				// if newCurr.Row < 0 || newCurr.Column > v.yLength-1 || newCurr.Column < 0 {
				// 	fmt.Println(currLocation, newCurr, "sadness")
				// 	continue
				// }
				_, ok := walls[newCurr]
				if ok {
					// fmt.Println("wall", newCurr)
					continue
				}
				if len(newBlizzards[newCurr]) != 0 {
					// fmt.Println("blizzard", newCurr)
					continue
				}

				// fmt.Println("new location", newCurr)
				newLocations[newCurr] = true
			}
			// if newCurr.Row <= 0 || newCurr.Column <= 0 || newCurr.Column == g.YLength-1 {
			// 	// fmt.Println("skipping", newCurr)
			// 	continue
			// }
			_, ok := walls[currLocation]
			if !ok && newBlizzards[currLocation] == nil {
				// fmt.Println("add current location")
				newLocations[currLocation] = true
			}
		}

		blizzards = newBlizzards
		locations = newLocations

	}

	fmt.Println("part1", minute)
}

// func movePlusBlizzard(g *grid.Grid, curr grid.Pos, minute int) int {
// 	// if minute > 20 {
// 	// 	fmt.Println("failed", minute)
// 	// 	os.Exit(1)
// 	// }
// 	fmt.Println(minute, curr)
// 	// fmt.Println(curr.Column)
// 	if curr.Row == g.XLength-1 {
// 		fmt.Println("found")
// 		os.Exit(0)
// 		return minute
// 	}

// 	g2 := grid.NewRectGrid(g.XLength, g.YLength)
// 	for r := 0; r < g.XLength; r++ {
// 		for c := 0; c < g.YLength; c++ {
// 			v := g.At(r, c).Data()
// 			if v == nil {
// 				continue
// 			}

// 			runes := v.([]byte)
// 			// fmt.Println("runes at", r, c, runes)
// 			for _, dir := range runes {
// 				delta := blizzardDeltas[dir]
// 				pos := grid.NewPos(r, c).Add(delta)
// 				if pos.Row == 0 {
// 					pos.Row = g.XLength - 2
// 				}
// 				if pos.Row == g.XLength-1 {
// 					pos.Row = 1
// 				}
// 				if pos.Column == 0 {
// 					pos.Column = g.YLength - 2
// 				}
// 				if pos.Column == g.YLength-1 {
// 					pos.Column = 1
// 				}
// 				// fmt.Println(r, c, "moving->", string(dir), pos)

// 				vals := g.At(pos.Row, pos.Column).Data()
// 				var l []byte
// 				if vals != nil {
// 					l = vals.([]byte)
// 				}
// 				l = append(l, dir)

// 				// set new value
// 				g2.At(pos.Row, pos.Column).SetValue(l)
// 			}
// 		}
// 	}

// 	foundMins := []int{}
// 	for _, delta := range blizzardDeltas {
// 		newCurr := curr.Add(delta)

// 		if newCurr.Row <= 0 || newCurr.Column <= 0 || newCurr.Column == g.YLength-1 {
// 			// fmt.Println("skipping", newCurr)
// 			continue
// 		}
// 		// occupied by blizzard. bye
// 		if g2.At(newCurr.Row, newCurr.Column).Data() != nil {
// 			// fmt.Println("blizzard occupied", newCurr)
// 			continue
// 		}

// 		// if g.a
// 		min2 := movePlusBlizzard(g2, newCurr, minute+1)
// 		foundMins = append(foundMins, min2)
// 		// if found && best2 < minute {
// 		// 	minute = best2
// 		// }
// 		// if best2 < best {
// 		// 	best = best2
// 		// }
// 	}
// 	if len(foundMins) == 0 {
// 		fmt.Println("sadness", minute)
// 		os.Exit(0)
// 		return minute + 1
// 	}

// 	return min(foundMins)
// }
