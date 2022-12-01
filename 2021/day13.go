package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/lolopinto/adventofcode2020/grid"
)

func day13() {
	chunks := readFileChunks("day13input", 2)

	rows := make([]int, len(chunks[0]))
	cols := make([]int, len(chunks[0]))
	for i, line := range chunks[0] {
		parts := strings.Split(line, ",")
		cols[i] = atoi(parts[0])
		rows[i] = atoi(parts[1])
	}

	maxWidth := max(rows)
	maxHeight := max(cols)
	g := grid.NewRectGrid(maxWidth+1, maxHeight+1)

	for i := range rows {
		// set dot
		g.At(rows[i], cols[i]).SetValue('.')
	}

	for _, ins := range chunks[1] {
		g2 := fold(g, ins)
		fmt.Println(countDots(g2))
		g = g2
	}

	// print . when nothing there and # when a dot
	g.Print(func(val interface{}) string {
		if val == nil {
			return "."
		}
		return "#"
	})
}

var foldRegex = regexp.MustCompile(`(y|x)=(\d+)`)

// fold returns new grid
func fold(g *grid.Grid, fold string) *grid.Grid {
	match := foldRegex.FindStringSubmatch(fold)
	if len(match) != 3 {
		panic("invalid regex")
	}
	dir := match[1]
	pos := atoi(match[2])

	if dir == "y" {
		// folding up
		g2 := grid.NewRectGrid(g.XLength/2, g.YLength)

		last := g.XLength - 1
		for i := 0; i < pos; i++ {
			// keeping this axis
			for j := 0; j < g.YLength; j++ {
				first := g.At(i, j).Data()
				second := g.At(last-i, j).Data()
				if first == '.' || second == '.' {
					g2.At(i, j).SetValue('.')
				}
			}
		}
		return g2
	} else {
		// folding left
		g2 := grid.NewRectGrid(g.XLength, g.YLength/2)

		last := g.YLength - 1
		for i := 0; i < g.XLength; i++ {
			for j := 0; j < pos; j++ {
				first := g.At(i, j).Data()
				second := g.At(i, last-j).Data()
				if first == '.' || second == '.' {
					g2.At(i, j).SetValue('.')
				}
			}
		}
		return g2
	}
}

func countDots(g *grid.Grid) int {
	ct := 0
	for i := 0; i < g.XLength; i++ {
		for j := 0; j < g.YLength; j++ {
			if g.At(i, j).Data() == '.' {
				ct++
			}
		}
	}
	return ct
}
