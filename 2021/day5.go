package main

import (
	"fmt"
	"strings"
)

var length int = 1000

func day5() {
	lines := readFile("day5input")

	// cheat
	grid := make([][]int, length)
	for i, g := range grid {
		g = make([]int, length)
		grid[i] = g
	}

	for _, line := range lines {
		parts := strings.Split(line, " -> ")
		start := convertToCoords(strings.Split(parts[0], ","))
		end := convertToCoords(strings.Split(parts[1], ","))
		markGrid(grid, start, end)
	}
	count := 0
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			if grid[i][j] > 1 {
				count++
			}
		}
	}
	fmt.Println(count)
}

type coord struct {
	x, y int
}

func convertToCoords(p []string) coord {
	return coord{
		x: atoi(p[0]),
		y: atoi(p[1]),
	}
}

func markGrid(grid [][]int, start, end coord) {
	minx := min([]int{start.x, end.x})
	maxx := max([]int{start.x, end.x})
	miny := min([]int{start.y, end.y})
	maxy := max([]int{start.y, end.y})

	// keep y constant
	if miny == maxy && minx != maxx {
		for x := minx; x <= maxx; x++ {
			grid[x][start.y] += 1
		}
	} else if minx == maxx && miny != maxy {
		// keep x constant
		for y := miny; y <= maxy; y++ {
			grid[start.x][y] += 1
		}
	} else {
		// diagonal
		diff := maxx - minx
		// default going up to the right
		xdiff := 1
		ydiff := 1
		//going down
		if start.x > end.x {
			xdiff = -1
		}
		if start.y > end.y {
			ydiff = -1
		}

		for i := 0; i <= diff; i++ {
			x := start.x + (i * xdiff)
			y := start.y + (i * ydiff)
			if x >= minx && x <= maxx && y >= miny && y <= maxy {
				grid[x][y] += 1
			}
		}
	}
}
