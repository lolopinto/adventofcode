package main

import (
	"log"
	"strings"
)

func empty(c byte) bool {
	return c == 'L'
}

func occupied(c byte) bool {
	return c == '#'
}

func floor(c byte) bool {
	return c == '.'
}

func validSeat(c2, r2 int, matrix []string) bool {
	return c2 >= 0 && c2 < len(matrix[0]) && r2 >= 0 && r2 < len(matrix)
}

func matrix(c, r int, matrix []string) byte {
	if floor(matrix[r][c]) {
		return matrix[r][c]
	}
	emp := empty(matrix[r][c])
	occ := occupied(matrix[r][c])

	occupiedCount := 0
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			if i == j && i == 0 {
				continue
			}
			c2 := c + i // column
			r2 := r + j

			for validSeat(c2, r2, matrix) {
				if !floor(matrix[r2][c2]) {
					break
				}
				// column not changing so row is changing
				// SO BAD AT GRIDS
				if i == 0 {
					// down
					if j > 0 {
						r2++
					} else {
						// up
						r2--
					}
					// left and right
				} else if j == 0 {
					if i > 0 {
						c2++
					} else {
						c2--
					}
				} else if i > 0 && j > 0 {
					// up to the right
					r2++
					c2++
				} else if i < 0 && j < 0 {
					// down to the left
					r2--
					c2--
				} else if j > 0 { // 1 -1
					r2++
					c2--
				} else if i > 0 {
					c2++
					r2--
				} else {
					log.Println(i, j)
					panic("unknown direction")
				}
			}

			if validSeat(c2, r2, matrix) && !floor(matrix[r2][c2]) {
				if occupied(matrix[r2][c2]) {
					occupiedCount++
				}
			}
		}
	}

	if emp && occupiedCount == 0 {
		return '#'
	}
	// part 2 change
	if occ && occupiedCount >= 5 {
		return 'L'
	}

	return matrix[r][c]
}

func convertMatrix(lines []string) []string {
	dup := make([]string, len(lines))
	for idx, line := range lines {
		var sb strings.Builder
		for pos := range line {
			sb.WriteByte(matrix(pos, idx, lines))
		}
		dup[idx] = sb.String()
	}
	return dup
}

func day11() {
	lines := readFile("day11input")

	count := 0
	for {
		dup := convertMatrix(lines)
		allEqual := true
		for idx := range lines {
			if dup[idx] != lines[idx] {
				allEqual = false
				break
			}
		}
		if allEqual {
			break
		}
		tmp := make([]string, len(dup))
		copy(tmp, dup)
		lines = tmp
		count++
	}
	log.Println(count)
	occupied := 0
	for _, line := range lines {
		for _, c := range line {
			if c == '#' {
				occupied++
			}
		}
	}
	log.Println(occupied)

}
