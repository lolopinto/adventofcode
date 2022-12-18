package main

import (
	"fmt"
	"strings"
)

type cube struct {
	x, y, z int
}

func (c *cube) neighbors() []cube {
	return []cube{
		{c.x, c.y + 1, c.z},
		{c.x, c.y - 1, c.z},
		{c.x + 1, c.y, c.z},
		{c.x - 1, c.y, c.z},
		{c.x, c.y, c.z + 1},
		{c.x, c.y, c.z - 1},
	}
}
func day18() {
	lines := readFile("day18input")

	// cubes := make([]cube, len(lines))
	cubes := map[cube]bool{}

	for _, line := range lines {
		v := ints(strings.Split(line, ","))
		c := cube{
			x: v[0],
			y: v[1],
			z: v[2],
		}
		cubes[c] = true
	}
	sum := 0
	for c := range cubes {
		for _, neigh := range c.neighbors() {
			if !cubes[neigh] {
				sum++
			}
		}
	}
	fmt.Println(sum)
}
