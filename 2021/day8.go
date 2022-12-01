package main

import (
	"fmt"
	"math"
	"strings"
)

// 7 3 segements
// 4 4 segments
// 1 2 segments
// 8 7 segments
var segmentMap = map[int]int{
	1: 2,
	7: 3,
	8: 7,
	4: 4,
}

func day8() {
	lines := readFile("day8input")
	sum := 0
	for _, line := range lines {
		parts := strings.Split(line, " | ")
		input := strings.Split(parts[0], " ")
		output := strings.Split(parts[1], " ")

		known := decode(input)
		flipped := make(map[string]int)
		for k, v := range known {
			flipped[v] = k
		}

		num := 0
		for i, v := range output {

			decoded := decodeOutput(flipped, v)
			exp := len(output) - i - 1
			num += int(math.Pow(float64(10), float64(exp))) * decoded
		}
		//		fmt.Println(num)
		sum += num
	}
	fmt.Println(sum)
}

func decodeOutput(flipped map[string]int, output string) int {
	for k, decoded := range flipped {
		if len(output) == len(k) && count(k, output) == len(output) {
			return decoded
		}
	}
	panic("couldn't decode " + output)
}

func decode(input []string) map[int]string {
	known := make(map[int]string)

	for _, v := range input {
		switch len(v) {
		case 3, 2, 4, 7:
			for k, v2 := range segmentMap {
				if v2 == len(v) {
					known[k] = v
				}
			}
		}
	}

	for _, v := range input {
		switch len(v) {
		case 6:
			// if every character in there is a subset of 4, it's 9
			if count(known[4], v) == 4 {
				known[9] = v
			} else if count(known[7], v) == 3 {
				known[0] = v
			} else {
				known[6] = v
			}
		}
	}

	for _, v := range input {
		switch len(v) {
		case 5:
			if count(known[4], v) == 2 {
				known[2] = v
			} else if count(known[1], v) == 2 {
				known[3] = v
			} else {
				known[5] = v
			}
		}
	}

	return known
}

func count(s1, s2 string) int {
	check := make(map[rune]bool)
	for _, v := range s1 {
		check[v] = true
	}
	ret := 0
	for _, v := range s2 {
		_, ok := check[v]
		if ok {
			ret++
		}
	}
	return ret
}
