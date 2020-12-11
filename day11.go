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

func matrix(c, r int, matrix []string) byte {
	if floor(matrix[r][c]) {
		return matrix[r][c]
	}
	emp := empty(matrix[r][c])
	occ := occupied(matrix[r][c])

	emptyCount := 0
	occupiedCount := 0
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			c2 := c + i
			r2 := r + j

			if c2 == c && r2 == r {
				continue
			}
			// valid seats
			if c2 >= 0 && c2 < len(matrix[0]) && r2 >= 0 && r2 < len(matrix) && !floor(matrix[r2][c2]) {
				if empty(matrix[r2][c2]) {
					emptyCount++
				}
				if occupied(matrix[r2][c2]) {
					occupiedCount++
				}
			}
		}
	}

	if emp && occupiedCount == 0 {
		return '#'
	}
	if occ && occupiedCount >= 4 {
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
		//		spew.Dump(dup)
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
