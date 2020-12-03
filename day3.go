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
	width      int
}

func (tm *topoMap) countTrees(pathRight, pathDown int) int {
	count := 0

	for tm.currentPos.y < len(tm.lines) {
		tm.currentPos.x += pathRight
		tm.currentPos.y += pathDown
		if tm.currentPos.y >= len(tm.lines) {
			break
		}

		// it repeats to the right so let's just use that to compare
		xComp := tm.currentPos.x % tm.width
		line := tm.lines[tm.currentPos.y]
		for i, c := range line {
			if i != xComp {
				continue
			}
			if c == tree {
				count++
			}
		}
	}
	return count
}

type iteration struct {
	pathRight int
	pathDown  int
}

func day3() {
	lines := readFile("day3input")

	var multiple int64
	multiple = 1
	iterations := []iteration{
		{
			pathRight: 1,
			pathDown:  1,
		},
		{
			pathRight: 3,
			pathDown:  1,
		},
		{
			pathRight: 5,
			pathDown:  1,
		},
		{
			pathRight: 7,
			pathDown:  1,
		},
		{
			pathRight: 1,
			pathDown:  2,
		},
	}

	for _, iteration := range iterations {
		tm := &topoMap{
			lines:      lines,
			currentPos: point{},
			width:      len(lines[0]),
		}
		count := tm.countTrees(iteration.pathRight, iteration.pathDown)
		log.Println(count)

		// 4953812864 wrong number
		// 5007658656 correct answer
		// 93, 164, 82, 91, 44 for each iteration
		// had 92 in previous way of doing this...
		multiple = multiple * int64(count)
	}

	log.Println(multiple)
}
