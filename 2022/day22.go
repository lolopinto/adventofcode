package main

import (
	"fmt"
	"log"
	"unicode"

	"github.com/lolopinto/adventofcode2020/grid"
)

type move struct {
	dir   cubeDirection
	steps int
}

type cubeDirection rune

const (
	Right cubeDirection = 'R'
	Left  cubeDirection = 'L'
	Up    cubeDirection = 'U'
	Down  cubeDirection = 'D'
)

var directionOrders = []cubeDirection{
	Right,
	Down,
	Left,
	Up,
}

var clockwise = map[cubeDirection]cubeDirection{
	Right: Down,
	Down:  Left,
	Left:  Up,
	Up:    Right,
}

var counterClockwise = map[cubeDirection]cubeDirection{
	Right: Up,
	Down:  Right,
	Left:  Down,
	Up:    Left,
}

var moveDeltas = map[cubeDirection]grid.Pos{
	Right: grid.NewPos(0, 1),
	Left:  grid.NewPos(0, -1),
	Up:    grid.NewPos(-1, 0),
	Down:  grid.NewPos(1, 0),
}

var flippedDirections = map[cubeDirection]cubeDirection{
	Right: Left,
	Down:  Up,
	Left:  Right,
	Up:    Down,
}

func day22() {
	chunks := readFileChunks("day22input", 2)

	max := 0
	lines := chunks[0]
	for _, line := range lines {
		if len(line) > max {
			max = len(line)
		}
	}

	g := grid.NewRectGrid(len(lines), max)

	var start grid.Pos
	foundStart := false
	for x, line := range lines {
		for y, v := range line {
			if x == 0 && !foundStart && v == '.' {
				start.Column = y
				foundStart = true
			}
			g.At(x, y).SetValue(v)
		}
	}

	// g.Print(func(val interface{}) string {
	// 	if val == nil {
	// 		return ""
	// 	}
	// 	return string(val.(rune))
	// })

	description := chunks[1][0]

	i := 0

	moves := []move{}
	// TODO refactor this into something that can parse ints and other known characters
	for i < len(description) {
		c := description[i]
		if unicode.IsDigit(rune(c)) {
			end := i
			for j := i; j < len(description); j++ {
				if !unicode.IsDigit(rune(description[j])) {
					end = j
					break
				}
			}
			var num int
			if i == end {
				num = atoi(description[i:])
				// never went through loop
				i++
				// break
			} else {
				num = atoi(description[i:end])
				i = end
			}
			moves = append(moves, move{
				steps: num,
			})
		} else {
			moves = append(moves, move{
				dir: cubeDirection(c),
			})
			i++
		}
	}

	// 4 in example, 50 in input
	cubeLength := abs(g.XLength, g.YLength)

	// part 2 inspired by this comment and solution
	// updated to be in line with how i was initially trying to solve it
	// https://www.reddit.com/r/adventofcode/comments/zsct8w/comment/j1bzuzm/?utm_source=reddit&utm_medium=web2x&context=3
	faces := []*cubeFace{}
	// add faces of cubes
	for r := 0; r < g.XLength; r += cubeLength {
		for c := 0; c < g.YLength; c += cubeLength {
			if g.At(r, c).RuneWithDefault(' ') != ' ' {
				c := &cubeFace{
					r:          [2]int{r, r + cubeLength - 1},
					c:          [2]int{c, c + cubeLength - 1},
					cubeLength: cubeLength,
					neighbors:  map[cubeDirection]*cubeFace{},
					crossCellNeighbors: map[cubeDirection]map[grid.Pos]grid.Pos{
						Right: {},
						Down:  {},
						Up:    {},
						Left:  {},
					},
				}
				faces = append(faces, c)
			}
		}
	}
	if len(faces) != 6 {
		log.Fatalf("expected 6 faces. got %d instead", len(faces))
	}

	for _, f1 := range faces {
		for _, f2 := range faces {
			if f1.eq(f2) {
				continue
			}
			// same row and distance is length we expected
			// connect right and left
			if f1.r == f2.r && f2.c[0]-f1.c[0] == cubeLength {
				connectSides(f1, f2, Right, Left)
			}
			// same column
			// connect up and down
			if f1.c == f2.c && f2.r[0]-f1.r[0] == cubeLength {
				connectSides(f1, f2, Down, Up)
			}
		}
	}

	// fold and set new neighbors by going through known faces and edges
	// and connecting
	// 5 done already ahead based on inputs
	for disconnected := 7; disconnected > 0; {
		for _, f := range faces {
			for _, dir := range directionOrders {
				if f.neighbors[dir] != nil {
					continue
				}

				next := clockwise[dir]
				prev := counterClockwise[dir]

				if f2 := f.neighbors[next]; f2 != nil {
					next = clockwise[f2.entrySide(f)]
					if fold := f2.neighbors[next]; fold != nil {
						next = clockwise[fold.entrySide(f2)]
						if fold.neighbors[next] == nil {
							connectSides(f, fold, dir, next)
							disconnected--
						}
					}
				} else if f2 := f.neighbors[prev]; f2 != nil {
					prev = counterClockwise[f2.entrySide(f)]
					if fold := f2.neighbors[prev]; fold != nil {
						prev = counterClockwise[fold.entrySide(f2)]
						if fold.neighbors[prev] == nil {
							connectSides(f, fold, dir, prev)
							disconnected--
						}
					}
				}
			}

		}
	}

	// part 1 regular movement
	moveDelta := func(curr grid.Pos, facing cubeDirection) (grid.Pos, cubeDirection) {
		delta := moveDeltas[facing]

		from := curr
		for {
			r := (from.Row + delta.Row) % g.XLength
			c := (from.Column + delta.Column) % g.YLength
			if r < 0 {
				r = r + g.XLength
			}
			if c < 0 {
				c = c + g.YLength
			}

			from = grid.NewPos(r, c)

			v := g.At(r, c).RuneWithDefault(' ')

			if v == '.' || v == '#' {
				// valid spot
				return from, facing
			}
		}
	}

	// part 2. cube
	moveDeltaCube := func(curr grid.Pos, facing cubeDirection) (grid.Pos, cubeDirection) {
		delta := moveDeltas[facing]
		newPos := curr.Add(delta)

		jump := false
		if !g.InGrid(newPos) || g.At(newPos.Row, newPos.Column).RuneWithDefault(' ') == ' ' {
			jump = true
		}

		if !jump {
			return newPos, facing
		}

		var currFace *cubeFace
		for _, f := range faces {
			if f.inFace(curr) {
				currFace = f
				break
			}
		}
		if currFace == nil {
			panic(fmt.Errorf("couldn't find the face that point %v is in", curr))
		}

		// what's the neighbor we're going to?
		newFace := currFace.neighbors[facing]

		// what's the direction currFace is going into new face
		newFacing := newFace.entrySide(currFace)
		// and direction flips from where we're facing
		newFacing = flippedDirections[newFacing]

		v, ok := currFace.crossCellNeighbors[facing][curr]
		if !ok {
			panic(fmt.Errorf("couldn't find cross-cell neighbor for %v facing %v", curr, facing))
		}
		return v, newFacing
	}

	part1Ans := processMoves(moves, start, g, moveDelta)
	part2Ans := processMoves(moves, start, g, moveDeltaCube)

	fmt.Println("part 1", part1Ans)
	fmt.Println("part 2", part2Ans)
}

type cubeFace struct {
	r, c               [2]int
	cubeLength         int
	neighbors          map[cubeDirection]*cubeFace
	crossCellNeighbors map[cubeDirection]map[grid.Pos]grid.Pos
}

func (cf *cubeFace) eq(cf2 *cubeFace) bool {
	return cf.r == cf2.r && cf.c == cf2.c
}

func (cf *cubeFace) entrySide(source *cubeFace) cubeDirection {
	for _, dir := range directionOrders {
		if cf.neighbors[dir] != nil && cf.neighbors[dir].eq(source) {
			return dir
		}
	}
	panic("couldn't find entry side for invalid value returned for entry side")
}

// TODO document
func (cf *cubeFace) sideClockwise(dir cubeDirection) []grid.Pos {
	cells := make([]grid.Pos, cf.cubeLength)
	n := len(cells) - 1
	var row, col, rowInc, colInc int
	switch dir {
	case Right:
		row, col, rowInc, colInc = 0, n, 1, 0
	case Down:
		row, col, rowInc, colInc = n, n, 0, -1
	case Left:
		row, col, rowInc, colInc = n, 0, -1, 0
	case Up:
		row, col, rowInc, colInc = 0, 0, 0, 1
	}
	start := grid.NewPos(cf.r[0], cf.c[0])
	for i := 0; i <= n; i++ {
		cells[i] = start.Add(grid.NewPos(row, col))
		row += rowInc
		col += colInc
	}
	return cells
}

func (cf *cubeFace) inFace(pos grid.Pos) bool {
	return pos.Row >= cf.r[0] && pos.Row <= cf.r[1] &&
		pos.Column >= cf.c[0] && pos.Column <= cf.c[1]
}

func connectSides(f1, f2 *cubeFace, facing, facing2 cubeDirection) {
	side1 := f1.sideClockwise(facing)
	side2 := f2.sideClockwise(facing2)
	n := len(side1) - 1

	for i := range side1 {
		cell1, cell2 := side1[i], side2[n-i]

		// keep track of cross cell neighbors for cells at the border
		// makes it easiest to go to them when crossing cubes
		f1.crossCellNeighbors[facing][cell1] = cell2
		f2.crossCellNeighbors[facing2][cell2] = cell1
	}

	f1.neighbors[facing] = f2
	f2.neighbors[facing2] = f1
}

func processMoves(moves []move, start grid.Pos, g *grid.Grid, moveFn func(curr grid.Pos, facing cubeDirection) (grid.Pos, cubeDirection)) int {
	var facing cubeDirection = 'R'
	curr := start
	for _, move := range moves {
		if move.dir != 0 {
			if move.dir == Right {
				facing = clockwise[facing]
			} else if move.dir == Left {
				facing = counterClockwise[facing]
			}
			continue
		}

		for i := 0; i < move.steps; i++ {
			newPos, newFacing := moveFn(curr, facing)
			val := g.At(newPos.Row, newPos.Column).Rune()
			// wall
			if val == '#' {
				break
			}
			if val == ' ' {
				panic(fmt.Errorf("sadness. hit air: %v", newPos))
			}

			facing = newFacing
			curr = newPos
		}
	}

	finalRow := curr.Row + 1
	finalColumm := curr.Column + 1
	// fmt.Println(finalRow, finalColumm)
	ans := (finalRow * 1000) + (finalColumm * 4)
	switch facing {
	// case 'R':
	case 'D':
		ans += 1
	case 'L':
		ans += 2
	case 'U':
		ans += 3

	}
	return ans
}
