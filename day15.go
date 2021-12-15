package main

import (
	"fmt"

	"github.com/lolopinto/adventofcode2020/grid"
)

func day15() {
	lines := readFile("day15input")
	g := grid.NewIntGrid(lines)
	mins := grid.NewGrid(g.Length)

	for x := 0; x < g.XLength; x++ {
		for y := 0; y < g.YLength; y++ {
			minAt := mins.At(x, y)
			data, ok := minAt.Data().(int)
			if !ok {
				// min at entry point is 0
				data = 0
			}

			neigh := g.RightAndDownNeighbors(x, y)
			for _, pos := range neigh {

				curr := g.At(pos.Row, pos.Column).Int()

				neighMin := mins.At(pos.Row, pos.Column)
				neighVal, ok := neighMin.Data().(int)
				if ok {
					if data+curr < neighVal {
						neighMin.SetValue(data + curr)
					}
				} else {
					// first time
					neighMin.SetValue(data + curr)
					//					fmt.Println(pos.Row, pos.Column, data+curr)
				}
			}
		}
	}

	last := mins.At(g.XLength-1, g.YLength-1).Int()
	fmt.Println(last)

	mins.Print(func(val interface{}) string {
		if val == nil {
			return leftPad("0 ", " ", 3)
		}
		return leftPad(fmt.Sprintf("%v ", val), " ", 4)
	})
}

//404 too high
