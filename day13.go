package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/lolopinto/adventofcode2020/grid"
)

func day13() {
	//	lines := readFile("day13input")
	chunks := readFileChunks("day13input", 2)

	// TODO get max and do this instead of hardcoding
	width := 1311
	g := grid.NewGrid(width)

	for _, line := range chunks[0] {
		parts := strings.Split(line, ",")
		c := atoi(parts[0])
		r := atoi(parts[1])
		// set dot
		g.At(r, c).SetValue('.')
	}

	// first fold
	first := chunks[1][0]
	g2 := fold(g, first)
	fmt.Println(countDots(g2))

	// second := chunks[1][1]
	// g3 := fold(g2, second)
	// fmt.Println(countDots(g3))

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

	// don't change length for now. optimization exists here
	g2 := grid.NewGrid(g.Length)
	if dir == "y" {
		// folding up

		last := g.Length - 1
		for i := 0; i < pos; i++ {
			for j := 0; j < g.Length; j++ {
				first := g.At(i, j).Data()
				second := g.At(last-i, j).Data()
				if first == '.' || second == '.' {
					g2.At(i, j).SetValue('.')
				}
			}
		}
	} else {
		last := g.Length - 1
		for i := 0; i < g.Length; i++ {
			for j := 0; j < pos; j++ {
				first := g.At(i, j).Data()
				second := g.At(i, last-j).Data()
				if first == '.' || second == '.' {
					g2.At(i, j).SetValue('.')
				}
			}
		}
	}

	return g2
}

func countDots(g *grid.Grid) int {
	ct := 0
	for i := 0; i < g.Length; i++ {
		for j := 0; j < g.Length; j++ {
			if g.At(i, j).Data() == '.' {
				ct++
			}
		}
	}
	return ct
}
