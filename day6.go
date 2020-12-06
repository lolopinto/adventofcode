package main

import "fmt"

type group struct {
	m map[rune]int
}

func (g group) processLine(line string) {
	for _, c := range line {
		val, ok := g.m[c]
		if ok {
			g.m[c] = val + 1
		} else {
			g.m[c] = 1
		}
	}
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
		sum += len(g.m)
		groups = append(groups, g)
		//		fmt.Println(len(g.m))
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
