package main

import (
	"fmt"
	"math"
)

func day3() {
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

func day3part1() {

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
	ct := -1
	most := math.MinInt
	for k, v := range m {
		if v > ct {
			ct = v
			most = k
		}
	}
	return most
}

func leastcommon(m map[int]int) int {
	ct := math.MaxInt
	least := math.MaxInt
	for k, v := range m {
		if v < ct {
			ct = v
			least = k
		}
	}
	return least
}
