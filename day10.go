package main

import (
	"fmt"
	"strings"
)

func day10() {
	lines := readFile("day10input")
	x := 1
	cycle := 0
	sum := 0
	increaseCycle := func() {
		cycle++
		if cycle == 20 || (cycle-20)%40 == 0 {
			// fmt.Println(cycle, x)
			sum += (cycle * x)
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
}
