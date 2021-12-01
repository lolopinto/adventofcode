package main

import (
	"fmt"
	"log"
	"math"
	"strings"
)

func findFirstPlus(ops []string) int {
	for i := 0; i < len(ops); i++ {
		op := ops[i]
		if op == "+" {
			return i
		}
	}
	return -1
}

func findInnerBrack(line string) *[2]int {
	left := []int{}

	for i, c := range line {
		if c == '(' {
			left = append(left, i)
		}
		if c == ')' {
			// found a matching pair!
			return &[2]int{left[len(left)-1], i}
		}
	}
	return nil
}

func calculateLine(line string) int {
	for {
		brack := findInnerBrack(line)
		if brack == nil {
			break
		}

		sub := line[brack[0]+1 : brack[1]]
		val := calculateLine(sub)
		line = strings.Replace(line, line[brack[0]:brack[1]+1], fmt.Sprintf("%d", val), 1)
	}
	ops := strings.Split(strings.TrimSpace(line), " ")

	for {
		idx := findFirstPlus(ops)
		if idx == -1 {
			break
		}
		sum := atoi(ops[idx-1]) + atoi(ops[idx+1])
		var ops2 []string
		var added bool
		for i, v := range ops {
			if math.Abs(float64(idx-i)) <= 1 {
				if !added {
					ops2 = append(ops2, fmt.Sprintf("%d", sum))
				}
				added = true
			} else {
				ops2 = append(ops2, v)
			}
		}
		ops = ops2
	}
	mult := 1
	for _, v := range ops {
		if v != "*" {
			mult *= atoi(v)
		}
	}
	return mult

}

func day18() {
	lines := readFile("day18input")

	sum := 0
	for _, line := range lines {
		ans := calculateLine(line)
		sum += ans
	}
	log.Println(sum)
}
