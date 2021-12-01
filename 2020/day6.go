package main

import (
	"fmt"
)

type group struct {
	m     map[rune]int
	lines []string
}

func (g *group) processLine(line string) {
	g.lines = append(g.lines, line)
	for _, c := range line {
		val, ok := g.m[c]
		if ok {
			g.m[c] = val + 1
		} else {
			g.m[c] = 1
		}
	}
}

func (g *group) processCount() int {
	count := 0
	for _, v := range g.m {
		if v == len(g.lines) {
			count++
		}
	}
	return count
}

func initGroup() group {
	return group{
		m: make(map[rune]int),
	}
}

func day6() {
	lines := readFile("day6input")

	sum := 0
	var groups []group
	g := initGroup()

	endGroup := func() {
		sum += g.processCount()
		groups = append(groups, g)
		g = initGroup()
	}
	for _, line := range lines {
		if line == "" {
			endGroup()
		} else {
			g.processLine(line)
		}
	}
	endGroup()
	fmt.Println(sum)
}
