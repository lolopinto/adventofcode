package main

import (
	"fmt"
	"math"
)

func day3() {
	lines := readFile("day3input")
	lines2 := readFile("day3input")
	first := lines[0]

	for i := 0; i < len(first); i++ {
		val := mostcommon(counts(lines, i))
		lines = filterLines(lines, val, i)
		if len(lines) == 1 {
			break
		}
	}
	for i := 0; i < len(first); i++ {
		val := leastcommon(counts(lines2, i))
		lines2 = filterLines(lines2, val, i)
		if len(lines2) == 1 {
			break
		}
	}
	fmt.Println(lines, lines2)
	fmt.Println(convertToBinary(lines[0]) * convertToBinary(lines2[0]))
}

func filterLines(lines []string, val, pos int) []string {
	var res []string
	for _, line := range lines {
		num := convRuneToInt(line[pos])
		if num == val {
			res = append(res, line)
		}
	}
	return res
}

func convRuneToInt(c byte) int {
	if c == '1' {
		return 1
	}
	return 0
}

func day3part1() {
	lines := readFile("day3input")

	first := lines[0]
	gammainput := make([]int, len(first))
	for i := 0; i < len(first); i++ {
		gammainput = append(gammainput, mostcommon(counts(lines, i)))
	}
	gamma := binary(gammainput)

	epilsoninput := make([]int, len(first))
	for i := 0; i < len(first); i++ {
		epilsoninput = append(epilsoninput, leastcommon(counts(lines, i)))
	}
	epilson := binary(epilsoninput)

	fmt.Println(gamma * epilson)
}

func binary(list []int) int {
	sum := 0
	for i, v := range list {
		pow := len(list) - i - 1
		if v == 1 {
			sum += int(math.Pow(2, float64(pow)))
		}
	}
	return sum
}

func counts(lines []string, pos int) map[int]int {
	res := make(map[int]int)
	for _, line := range lines {
		c := line[pos]
		num := 0
		if c == '1' {
			num = 1
		}
		ct := res[num]
		ct++
		res[num] = ct
	}
	return res
}

func mostcommon(m map[int]int) int {
	one := m[1]
	zero := m[0]
	if one >= zero {
		return 1
	}
	return 0
}

func leastcommon(m map[int]int) int {
	one := m[1]
	zero := m[0]
	if zero <= one {
		return 0
	}
	return 1
}

func convertToBinary(line string) int {
	res := make([]int, len(line))
	for i, c := range line {
		if c == '1' {
			res[i] = 1
		} else {
			res[i] = 0
		}
	}
	return binary(res)
}
