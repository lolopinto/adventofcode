package main

import "fmt"

func day9() {
	type point struct {
		x, y int
	}
	lines := readFile("day9input")
	// head := map[point]bool{}
	start := point{0, 0}
	tail := map[point]bool{
		start: true,
	}

	headPos := start
	tailPos := start

	touching := func() bool {
		return abs(headPos.x, tailPos.x) <= 1 && abs(headPos.y, tailPos.y) <= 1
	}

	updateTailPos := func(pt point) {
		tail[pt] = true
		tailPos = pt
	}
	moveTailDiagonal := func() {
		// down left
		if headPos.x < tailPos.x && headPos.y > tailPos.y {
			updateTailPos(point{x: tailPos.x - 1, y: tailPos.y + 1})
		}
		if headPos.x < tailPos.x && headPos.y < tailPos.y {
			updateTailPos(point{x: tailPos.x - 1, y: tailPos.y - 1})
		}

		if headPos.x > tailPos.x && headPos.y < tailPos.y {
			updateTailPos(point{x: tailPos.x + 1, y: tailPos.y - 1})
		}

		if headPos.x > tailPos.x && headPos.y > tailPos.y {
			updateTailPos(point{x: tailPos.x + 1, y: tailPos.y + 1})
		}
	}

	for _, line := range lines {
		parts := splitLength(line, " ", 2)
		dir := parts[0]
		for i := 0; i < atoi(parts[1]); i++ {
			switch dir {
			case "R":
				headPos.x++
				if touching() {
					continue
				}
				if headPos.x-tailPos.x == 2 && headPos.y == tailPos.y {
					updateTailPos(point{x: tailPos.x + 1, y: tailPos.y})
				} else {
					moveTailDiagonal()
				}
			case "U":
				headPos.y--
				if touching() {
					continue
				}
				if headPos.y-tailPos.y == -2 && headPos.x == tailPos.x {
					updateTailPos(point{x: tailPos.x, y: tailPos.y - 1})
				} else {
					moveTailDiagonal()

				}

			case "D":
				headPos.y++
				if touching() {
					continue
				}
				if headPos.y-tailPos.y == 2 && headPos.x == tailPos.x {
					updateTailPos(point{x: tailPos.x, y: tailPos.y + 1})
				} else {
					moveTailDiagonal()

				}

			case "L":
				headPos.x--
				if touching() {
					continue
				}
				if headPos.x-tailPos.x == -2 && headPos.y == tailPos.y {
					updateTailPos(point{x: tailPos.x - 1, y: tailPos.y})
				} else {
					moveTailDiagonal()

				}

			}
		}
	}
	fmt.Println(len(tail))

}
