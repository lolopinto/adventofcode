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
		//		fmt.Println(grid, start, end)
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
	if start.x != end.x && start.y != end.y {
		return
	}
	// x equal
	if start.x == end.x {
		if start.y < end.y {
			// going up
			for i := start.y; i <= end.y; i++ {
				grid[start.x][i] += 1
			}
			//			fmt.Println("y going up", start.y)

		} else {
			// going down
			for i := start.y; i >= end.y; i-- {
				grid[start.x][i] += 1
				//				fmt.Println("y going down", start.y)

			}
		}
	} else {
		// y equal
		if start.x < end.x {
			for i := start.x; i <= end.x; i++ {
				grid[i][start.y] += 1

			}
			// fmt.Println("x going up", start.x, end.x)
			// fmt.Println(grid)

		} else {
			for i := start.x; i >= end.x; i-- {
				grid[i][start.y] += 1
			}
			//			fmt.Println("x going down", start.x)
		}
	}
}
