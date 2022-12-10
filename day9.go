package main

import "fmt"

func deltaChange(d int) int {
	if d == 0 {
		return 0
	}

	if d < 0 {
		return -1
	}
	return 1
}

func day9() {
	type point struct {
		x, y int
	}
	lines := readFile("day9input")
	start := point{0, 0}
	// tail 9

	currentPos := make(map[int]point)
	tail := make(map[point]bool)

	for i := 1; i < 10; i++ {
		currentPos[i] = start
	}
	tail[start] = true

	headPos := start

	updateFollowerPos := func(pt point, idx int) {
		if idx == 9 {
			tail[pt] = true
		}
		currentPos[idx] = pt
	}

	for _, line := range lines {
		parts := splitLength(line, " ", 2)
		dir := parts[0]
		for i := 0; i < atoi(parts[1]); i++ {

			// move head first
			switch dir {
			case "R":
				headPos.x++
			case "U":
				headPos.y--
			case "D":
				headPos.y++
			case "L":
				headPos.x--
			}

			for j := 1; j < 10; j++ {
				var leader point
				follower := currentPos[j]
				if j == 1 {
					leader = headPos
				} else {
					leader = currentPos[j-1]
				}

				dx := leader.x - follower.x
				dy := leader.y - follower.y

				if dx*dx+dy*dy > 2 {
					updateFollowerPos(point{x: follower.x + deltaChange(dx), y: follower.y + deltaChange(dy)}, j)
				}
			}
		}
	}

	day9part1()
	fmt.Println(len(tail))
}

func day9part1() {
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

	updateTailPos := func(pt point) {
		tail[pt] = true
		tailPos = pt
	}

	for _, line := range lines {
		parts := splitLength(line, " ", 2)
		dir := parts[0]
		for i := 0; i < atoi(parts[1]); i++ {
			switch dir {
			case "R":
				headPos.x++
			case "U":
				headPos.y--
			case "D":
				headPos.y++
			case "L":
				headPos.x--
			}

			dx := headPos.x - tailPos.x
			dy := headPos.y - tailPos.y

			if dx*dx+dy*dy > 2 {
				updateTailPos(point{x: tailPos.x + deltaChange(dx), y: tailPos.y + deltaChange(dy)})
			}
		}
	}
	fmt.Println(len(tail))
}
