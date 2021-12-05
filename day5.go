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
	// x equal
	if start.x == end.x {
		if start.y < end.y {
			// going up
			for i := start.y; i <= end.y; i++ {
				grid[start.x][i] += 1
			}

		} else {
			// going down
			for i := start.y; i >= end.y; i-- {
				grid[start.x][i] += 1
			}
		}
	} else if start.y == end.y {
		// y equal
		if start.x < end.x {
			for i := start.x; i <= end.x; i++ {
				grid[i][start.y] += 1

			}

		} else {
			for i := start.x; i >= end.x; i-- {
				grid[i][start.y] += 1
			}
		}
	} else {
		// diagonal going up left
		if end.y > start.y && end.x > start.x {
			diff := end.x - start.x
			for i := 0; i <= diff; i++ {
				x := start.x + i
				y := start.y + i

				grid[x][y] += 1
			}
		} else if end.y < start.y && end.x < start.x {
			// inverse from up
			diff := start.x - end.x
			for i := 0; i <= diff; i++ {
				x := start.x - i
				y := start.y - i
				grid[x][y] += 1
			}
			// last 2 in opp directions

		} else if end.x > start.x {
			diff := end.x - start.x
			for i := 0; i <= diff; i++ {
				x := start.x + i
				y := start.y - i
				grid[x][y] += 1
			}
		} else {
			diff := end.y - start.y
			for i := 0; i <= diff; i++ {
				x := start.x - i
				y := start.y + i
				grid[x][y] += 1
			}
		}
	}
}
