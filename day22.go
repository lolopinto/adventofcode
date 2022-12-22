package main

import (
	"fmt"
	"unicode"

	"github.com/lolopinto/adventofcode2020/grid"
)

type move struct {
	dir   rune
	steps int
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
				dir: rune(c),
			})
			i++
		}
	}
	// fmt.Println(moves)

	// part 1 regular movement
	moveDelta := func(curr grid.Pos, facing rune) (grid.Pos, rune) {
		var delta grid.Pos
		switch facing {
		case 'R':
			delta = grid.NewPos(0, 1)
		case 'L':
			delta = grid.NewPos(0, -1)
		case 'U':
			delta = grid.NewPos(-1, 0)
		case 'D':
			delta = grid.NewPos(1, 0)
		}
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
			v := ' '

			val := g.At(r, c).Data()
			if val != nil {
				v = g.At(r, c).Rune()
			}

			if v == '.' || v == '#' {
				// valid spot
				return from, facing
			}
		}

	}

	moveDeltaCube := func(curr grid.Pos, facing rune) (grid.Pos, rune) {
		var delta grid.Pos
		switch facing {
		case 'R':
			delta = grid.NewPos(0, 1)
		case 'L':
			delta = grid.NewPos(0, -1)
		case 'U':
			delta = grid.NewPos(-1, 0)
		case 'D':
			delta = grid.NewPos(1, 0)
		}
		from := curr
		// for {

		r := (from.Row + delta.Row)
		c := (from.Column + delta.Column)
		jump := false
		if r < 0 || r > g.XLength {
			// TODO need to jump dimensions
			// r = r + g.XLength
			jump = true
		}
		if c < 0 || c > g.YLength {
			jump = true
		}

		if !jump {
			from = grid.NewPos(r, c)
			v := ' '

			val := g.At(r, c).Data()
			if val != nil {
				v = g.At(r, c).Rune()
			}
			if v == ' ' {
				jump = true
			}
		}

		if jump {
			fmt.Println(r, c, curr, string(facing))
			panic("TODO ola figure it out?")
		} else {
			return grid.NewPos(r, c), facing
		}
		// from = grid.NewPos(r, c)
		// v := ' '

		// val := g.At(r, c).Data()
		// if val != nil {
		// 	v = g.At(r, c).Rune()
		// }

		// if v == '.' || v == '#' {
		// 	// valid spot
		// 	return from, facing
		// }
		// // }
	}

	part1Ans := processMoves(moves, start, g, moveDelta)
	part2Ans := processMoves(moves, start, g, moveDeltaCube)

	// fmt.Println(curr, facing)
	fmt.Println("part 1", part1Ans)
	fmt.Println("part 2", part2Ans)
}

func processMoves(moves []move, start grid.Pos, g *grid.Grid, moveFn func(curr grid.Pos, facing rune) (grid.Pos, rune)) int {
	var facing = 'R'
	curr := start
	for _, move := range moves {
		if move.dir != 0 {
			switch move.dir {
			case 'R':
				if facing == 'R' {
					facing = 'D'
				} else if facing == 'D' {
					facing = 'L'
				} else if facing == 'L' {
					facing = 'U'
				} else if facing == 'U' {
					facing = 'R'
				}

			case 'L':
				if facing == 'R' {
					facing = 'U'
				} else if facing == 'D' {
					facing = 'R'
				} else if facing == 'L' {
					facing = 'D'
				} else if facing == 'U' {
					facing = 'L'
				}
			}
			continue
		}

		for i := 0; i < move.steps; i++ {
			newPos, newFacing := moveFn(curr, facing)
			if g.At(newPos.Row, newPos.Column).Rune() != '.' {
				break
			}
			// fmt.Println("changing pos", curr, newPos, string(facing))
			curr = newPos
			facing = newFacing
		}
	}

	finalRow := curr.Row + 1
	finalColumm := curr.Column + 1
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
