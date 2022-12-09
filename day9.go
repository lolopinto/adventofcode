package main

import "fmt"

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

				diff := point{follower.x - leader.x, follower.y - leader.y}

				if diff.x == 2 && diff.y == 2 {
					updateFollowerPos(point{x: leader.x + 1, y: leader.y + 1}, j)
				} else if diff.x == 2 && diff.y == -2 {
					updateFollowerPos(point{x: leader.x + 1, y: leader.y - 1}, j)
				} else if diff.x == -2 && diff.y == 2 {
					updateFollowerPos(point{x: leader.x - 1, y: leader.y + 1}, j)
				} else if diff.x == -2 && diff.y == -2 {
					updateFollowerPos(point{x: leader.x - 1, y: leader.y - 1}, j)
				} else if diff.x == 2 {
					updateFollowerPos(point{x: leader.x + 1, y: leader.y}, j)
				} else if diff.x == -2 {
					updateFollowerPos(point{x: leader.x - 1, y: leader.y}, j)
				} else if diff.y == 2 {
					updateFollowerPos(point{x: leader.x, y: leader.y + 1}, j)
				} else if diff.y == -2 {
					updateFollowerPos(point{x: leader.x, y: leader.y - 1}, j)
				}
			}
		}
	}

	fmt.Println(len(tail))
}
