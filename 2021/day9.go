package main

import (
	"fmt"
	"sort"

	"github.com/lolopinto/adventofcode2020/grid"
)

func day9() {
	lines := readFile("day9input")
	sum := 0
	var basinCounts []int

	g := grid.NewIntGrid(lines)
	for r := 0; r < g.Length; r++ {
		for c := 0; c < g.Length; c++ {
			num := g.At(r, c).Int()
			adj := g.Neighbors(r, c)
			mincan := make([]int, len(adj))
			for i := range adj {
				mincan[i] = g.At(adj[i].Row, adj[i].Column).Int()
			}
			if num < min(mincan) {
				//				fmt.Println("lowpoint", num)
				sum += num + 1
				// rewritten approach that worked
				bc := findBasinCountWorking(g, adj, grid.Pos{
					Row:    r,
					Column: c,
				})
				basinCounts = append(basinCounts, bc)
			}
		}
	}
	sort.Ints(basinCounts)
	multiple := 1
	count := 0

	for i := len(basinCounts) - 1; i > 0; i-- {
		multiple *= basinCounts[i]
		count++
		if count == 3 {
			break
		}
	}

	fmt.Println(sum)
	fmt.Println(multiple)
}

func findBasinCountWorking(g *grid.Grid, neighbors []grid.Pos, origPos grid.Pos) int {
	at := g.At(origPos.Row, origPos.Column)
	if at.Visited {
		return 0
	}
	count := 1
	at.Visited = true

	for _, p := range neighbors {

		next := g.At(p.Row, p.Column).Int()

		if next != 9 {
			adj2 := g.Neighbors(p.Row, p.Column)
			count += findBasinCountWorking(g, adj2, p)
		}
	}
	return count
}
