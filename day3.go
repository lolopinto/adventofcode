package main

import (
	"log"
)

//https://adventofcode.com/2020/day/3

type point struct {
	x, y int
}

const tree = '#'

type topoMap struct {
	lines      []string
	currentPos point
	treesSeen  int
	width      int
}

func (tm *topoMap) advance() bool {
	newPos := point{
		x: tm.currentPos.x + 3,
		y: tm.currentPos.y + 1,
	}

	// it repeats to the right so let's just use that to compare
	xComp := newPos.x

	if newPos.x > tm.width {
		xComp = newPos.x % tm.width
	}

	if newPos.y >= len(tm.lines) {
		return false
	}

	line := tm.lines[newPos.y]
	for i, c := range line {
		if i != xComp {
			continue
		}
		if c == tree {
			tm.treesSeen++
		}
	}

	tm.currentPos = newPos

	if newPos.y >= len(tm.lines) && newPos.x >= len(tm.lines[0]) {
		return false
	}
	return true
}

func day3() {
	lines := readFile("day3input")

	tm := &topoMap{
		lines:      lines,
		currentPos: point{},
		width:      len(lines[0]),
	}

	for {
		if !tm.advance() {
			break
		}
	}
	log.Println(tm.treesSeen)
}
