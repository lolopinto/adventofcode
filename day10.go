package main

import (
	"fmt"
	"strings"

	"github.com/lolopinto/adventofcode2020/grid"
)

func day10() {
	lines := readFile("day10input")
	x := 1
	cycle := 0
	sum := 0

	// grid dimensions? way grid class is structured is wrong
	g := grid.NewRectGrid(6, 40)

	increaseCycle := func() {
		cycle++
		if cycle == 20 || (cycle-20)%40 == 0 {
			sum += (cycle * x)
		}
		// part 2
		pos := cycle - 1
		// took too long to understand what they were asking
		// i hate the wording of these things
		sprite := []int{x - 1, x, x + 1}
		cyclewidth := (pos % 40)
		cycleheight := (pos / 40)

		for _, v := range sprite {
			if v == cyclewidth {
				g.At(cycleheight, cyclewidth).SetValue("#")
			}
		}
	}

	for _, line := range lines {
		if strings.HasPrefix(line, "addx") {
			increaseCycle()
			increaseCycle()
			add := atoi(strings.Split(line, " ")[1])
			x += add
		} else {
			increaseCycle()
		}
	}

	fmt.Println(sum)

	// EALGULPG
	g.Print(func(val interface{}) string {
		if val == nil {
			return "."
		}
		return "#"
	})
}
