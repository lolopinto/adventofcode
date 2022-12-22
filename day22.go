package main

import (
	"fmt"
	"unicode"

	"github.com/lolopinto/adventofcode2020/grid"
)

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
	type move struct {
		dir   rune
		steps int
	}
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
			// fmt.Println("num coords", i, end, description[i:end])
			var num int
			if i == end {
				// fmt.Println("atoi...")
				num = atoi(description[i:])
				// never went through loop
				i++
				// break
			} else {
				num = atoi(description[i:end])
				i = end
			}
			// fmt.Println("num", num, i, end, len(description))
			moves = append(moves, move{
				steps: num,
			})
		} else {
			moves = append(moves, move{
				dir: rune(c),
			})
			// fmt.Println(string(c))
			i++
		}
	}
	// fmt.Println(moves)

	// canMove := func(p grid.Pos) bool {
	// 	return g.At(p.Row, p.Column).Rune() == '.'
	// }
	var facing = 'R'
	curr := start
	for _, move := range moves {
		if move.dir != 0 {
			switch move.dir {
			case 'R':
				// fmt.Println("90 degree clockwise")

				// fmt.Println("was facing", string(facing))
				if facing == 'R' {
					facing = 'D'
				} else if facing == 'D' {
					facing = 'L'
				} else if facing == 'L' {
					facing = 'U'
				} else if facing == 'U' {
					facing = 'R'
				}
				// fmt.Println("now facing", string(facing))

			case 'L':
				// fmt.Println("90 degree counter-clockwise")

				// fmt.Println("was facing", string(facing))
				if facing == 'R' {
					facing = 'U'
				} else if facing == 'D' {
					facing = 'R'
				} else if facing == 'L' {
					facing = 'D'
				} else if facing == 'U' {
					facing = 'L'
				}
				// fmt.Println("now facing", string(facing))
			}
			continue
		}

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

		moveDelta := func() grid.Pos {
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
				// fmt.Println("in move delta", r, c, string(v))

				if v == '.' || v == '#' {
					// valid spot
					return from
				}
				// fmt.Println(string(facing), curr, r, c, string(v))
			}
		}

		for i := 0; i < move.steps; i++ {
			newPos := moveDelta()
			if g.At(newPos.Row, newPos.Column).Rune() != '.' {
				break
			}
			// fmt.Println("changing pos", curr, newPos, string(facing))
			curr = newPos
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

	// fmt.Println(curr, facing)
	fmt.Println("part 1", ans)
}
