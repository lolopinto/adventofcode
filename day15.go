package main

import (
	"fmt"

	"github.com/lolopinto/adventofcode2020/grid"
)

func day15() {
	lines := readFile("day15input")
	g := grid.NewIntGrid(lines)
	// for the sums
	sums := grid.NewGrid(g.Length)

	for x := 0; x < g.XLength; x++ {
		for y := 0; y < g.YLength; y++ {
			val := g.At(x, y).Int()
			sumAt := sums.At(x, y)
			data, ok := sumAt.Data().([]int)
			if !ok {
				data = []int{val}
			}

			neigh := g.RightAndDownNeighbors(x, y)
			//			fmt.Println(x, y, len(neigh))
			for _, pos := range neigh {

				curr := g.At(pos.Row, pos.Column).Int()

				d := sums.At(pos.Row, pos.Column)
				data2, ok := d.Data().([]int)
				if !ok {
					data2 = []int{}
				}

				for _, v := range data {
					data2 = append(data2, v+curr)
				}
				d.SetValue(data2)
			}
		}
	}

	last := sums.At(9, 9).Data().([]int)
	//	fmt.Println(len(last))
	fmt.Println(min(last) - g.At(0, 0).Int())
}
