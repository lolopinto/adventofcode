package main

import (
	"fmt"

	"github.com/lolopinto/adventofcode2020/grid"
)

func day25() {
	lines := readFile("day25input")

	g := grid.NewRectGrid(len(lines), len(lines[0]))
	for i := 0; i < g.XLength; i++ {
		for j := 0; j < g.YLength; j++ {
			c := lines[i][j]
			val := rune(c)
			if val == '>' || val == 'v' {
				g.At(i, j).SetValue(val)
			}
		}
	}
	//	printCucumber(g)

	step := 1
	for {
		moves := 0
		g2 := grid.NewRectGrid(g.XLength, g.YLength)

		// east pass
		for i := 0; i < g.XLength; i++ {
			for j := 0; j < g.YLength; j++ {
				val, ok := g.At(i, j).Data().(rune)
				if !ok || val != '>' {
					continue
				}
				r := i
				c := (j + 1) % g.YLength
				_, ok2 := g.At(r, c).Data().(rune)
				if !ok2 {
					if err := g2.At(r, c).SetValueOnce('>'); err != nil {
						fmt.Println(err, r, c)
					}
					moves++
				} else {
					// no moving
					if err := g2.At(i, j).SetValueOnce('>'); err != nil {
						fmt.Println(err, i, j)
					}
				}
			}
		}

		// south pass
		for i := g.XLength - 1; i >= 0; i-- {
			for j := 0; j < g.YLength; j++ {
				val, ok := g.At(i, j).Data().(rune)
				if !ok || val != 'v' {
					continue
				}
				r := (i + 1) % g.XLength
				c := j
				//				fmt.Println(i, j, r, c, val)
				val, ok2 := g.At(r, c).Data().(rune)
				_, ok3 := g2.At(r, c).Data().(rune)
				// if previous
				var canmove bool
				canmove = !ok2 && !ok3
				if !canmove {
					if ok2 && val == '>' {
						// if previous value was > and now empty, can move
						canmove = !ok3
					} else if ok2 && val == 'v' {
						canmove = false
					}
				}
				if canmove {
					// can move
					if err := g2.At(r, c).SetValueOnce('v'); err != nil {
						fmt.Println(err, r, c)
					}
					moves++
				} else {
					// no moving
					if err := g2.At(i, j).SetValueOnce('v'); err != nil {
						fmt.Println(err, i, j)

					}
				}
			}
		}
		if moves == 0 {
			break
		}
		//		if step == 58 {
		//		printCucumber(g2)

		//		}
		step++
		g = g2
	}
	fmt.Println(step)
}

func printCucumber(g *grid.Grid) {
	g.Print(func(val interface{}) string {
		if val == nil {
			return "."
		}
		return string(val.(rune))
	})
	fmt.Println()
}
