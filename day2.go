package main

import "fmt"

func day2() {
	horizontal := 0
	depth := 0
	lines := readFile("day2input")
	for _, line := range lines {
		parts := splitLength(line, " ", 2)
		dir := parts[0]
		num := atoi(parts[1])

		switch dir {
		case "forward":
			horizontal += num
		case "down":
			depth += num
		case "up":
			depth -= num
		}
	}
	fmt.Println(horizontal * depth)
}

func day2part1() {

}
