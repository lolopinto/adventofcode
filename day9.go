package main

import "fmt"

func day9() {
	type point struct {
		x, y int
	}
	lines := readFile("day9input")
	start := point{0, 0}
	// tail 9
	type ropes map[int]map[point]bool
	type current map[int]point

	allData := make(ropes)
	currentPos := make(current)

	for i := 1; i < 10; i++ {
		allData[i] = map[point]bool{
			start: true,
		}
		currentPos[i] = start
	}

	headPos := start

	touching := func(leader, follower point) bool {
		return abs(leader.x, follower.x) <= 1 && abs(leader.y, follower.y) <= 1
	}

	updateFollowerPos := func(pt point, idx int) {
		// fmt.Println("update ", idx)

		allData[idx][pt] = true
		currentPos[idx] = pt
	}

	moveFollowerDiagonal := func(leader, follower point, idx int) {
		// down left
		if leader.x < follower.x && leader.y > follower.y {
			updateFollowerPos(point{x: follower.x - 1, y: follower.y + 1}, idx)
		}
		if leader.x < follower.x && leader.y < follower.y {
			updateFollowerPos(point{x: follower.x - 1, y: follower.y - 1}, idx)
		}

		if leader.x > follower.x && leader.y < follower.y {
			updateFollowerPos(point{x: follower.x + 1, y: follower.y - 1}, idx)
		}

		if leader.x > follower.x && leader.y > follower.y {
			updateFollowerPos(point{x: follower.x + 1, y: follower.y + 1}, idx)
		}
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

				if touching(leader, follower) {
					continue
				}

				switch dir {
				case "R":

					if leader.x-follower.x == 2 && leader.y == follower.y {
						updateFollowerPos(point{x: follower.x + 1, y: follower.y}, j)
					} else {
						moveFollowerDiagonal(leader, follower, j)
					}
				case "U":
					if leader.y-follower.y == -2 && leader.x == follower.x {
						updateFollowerPos(point{x: follower.x, y: follower.y - 1}, j)
					} else {
						moveFollowerDiagonal(leader, follower, j)
					}

				case "D":
					if leader.y-follower.y == 2 && leader.x == follower.x {
						updateFollowerPos(point{x: follower.x, y: follower.y + 1}, j)
					} else {
						moveFollowerDiagonal(leader, follower, j)
					}

				case "L":
					if leader.x-follower.x == -2 && leader.y == follower.y {
						updateFollowerPos(point{x: follower.x - 1, y: follower.y}, j)
					} else {
						moveFollowerDiagonal(leader, follower, j)
					}
				}
			}
		}
	}

	fmt.Println(len(allData[9]))
	// for k := range allData[9] {
	// 	fmt.Printf("%d:%d\n", k.x, k.y)
	// }
}
