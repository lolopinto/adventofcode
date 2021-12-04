package main

import (
	"fmt"
	"os"
	"strings"
)

func day4() {
	lines := readFile("day4input")

	csv := ints(strings.Split(lines[0], ","))

	var grids [][][]int
	c := 0
	var last [][]int
	for _, line := range lines[2:] {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		var g []int

		line = strings.ReplaceAll(line, "  ", " ")
		g = ints(strings.Split(line, " "))

		if len(g) == 0 {
			continue
		}
		last = append(last, g)
		c++
		if c == 5 {
			grids = append(grids, last)
			last = [][]int{}
			c = 0
		}
	}
	//	fmt.Println(grids)

	drawn := make(map[int]bool)
	for _, num := range csv {
		//		drawn = append(drawn, num)
		drawn[num] = true
		for _, grid := range grids {
			b, sum := checkBoard(grid, drawn)
			if b {
				fmt.Println(sum * num)
				os.Exit(0)
			}
		}
	}
}

func checkBoard(grid [][]int, drawn map[int]bool) (bool, int) {

	marked := false
	for r := 0; r < 5; r++ {
		maybemarked := true
		for _, num := range grid[r] {
			if !drawn[num] {
				maybemarked = false
				break
			}
		}
		if maybemarked {
			marked = true
			break
		}
	}

	for c := 0; c < 5; c++ {
		maybemarked := true
		for i := 0; i < 5; i++ {
			num := grid[i][c]
			if !drawn[num] {
				maybemarked = false
				break
			}
		}

		if maybemarked {
			marked = true
			break
		}
	}

	sum := 0
	if marked {
		for r := 0; r < 5; r++ {
			for c := 0; c < 5; c++ {
				num := grid[r][c]
				if !drawn[num] {
					sum += num
				}
			}
		}
		return marked, sum
	}

	return false, 0

}

func day4part1() {

}
