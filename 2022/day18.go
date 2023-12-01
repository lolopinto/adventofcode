package main

import (
	"fmt"

	"github.com/lolopinto/adventofcode2020/cube"
)

// For part 2, I basically lifted this from zamore's implementation
// https://github.com/fzamore/advent-of-code/blob/main/2022/day18.py
// because cubes and grids break my mind and it took me a while to figure
// out what the question was even asking

// other relevant things: https://technology.cpm.org/general/3dgraph/?graph3ddata=____bHxBExBExBEJxmuxBExBEKxQOxBExBELxBExmuxBEMxBExQOxBEHxBExBExmuIxBExBExQOJxBExBEx1YKxBExBEyweLxmuxBEyg4MxQOxBEyg4HxBExmuyg4IxBExQOyg4

// the edges inside are the ones whose surface area aren't being used for part 2?
// https://www.reddit.com/r/adventofcode/comments/zow3f3/2022_day_18_part_2_trapped_air/
func day18() {
	lines := readFile("day18input")

	cubes, err := cube.NewCSVMapCube(lines)
	die(err)

	coords := []int{}
	for c := range cubes {
		coords = append(coords, c.X, c.Y, c.Z)
	}

	minVal := min(coords)
	maxVal := max(coords)

	// fmt.Println(minVal, maxVal)
	sum := 0
	for c := range cubes {
		for _, neigh := range c.Neighbors() {
			if !cubes[neigh] {
				visited := map[cube.Cube]bool{}

				if canEscape(neigh, cubes, visited, minVal, maxVal) {
					sum += 1
				}
			}
		}
	}
	day18part1()
	fmt.Println("part 2 answer", sum)
}

func canEscape(c cube.Cube, cubes, visited map[cube.Cube]bool, min, max int) bool {
	if c.X < min || c.X > max ||
		c.Y < min || c.Y > max ||
		c.Z < min || c.Z > max {
		return true
	}

	for _, neigh := range c.Neighbors() {
		if visited[neigh] {
			continue
		}
		visited[neigh] = true

		if !cubes[neigh] {
			if canEscape(neigh, cubes, visited, min, max) {
				return true
			}
		}
	}
	return false
}

func day18part1() {
	lines := readFile("day18input")

	cubes, err := cube.NewCSVMapCube(lines)
	die(err)

	sum := 0

	for c := range cubes {
		for _, neigh := range c.Neighbors() {
			if !cubes[neigh] {
				sum++
			}
		}
	}

	fmt.Println("part 1", sum)
}
