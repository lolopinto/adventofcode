package main

import (
	"fmt"
	"math"
	"strings"
	"unicode"
)

func transformInput(s string) []string {
	input := make([]string, len(s))
	for i, c := range s {
		input[i] = string(c)
	}
	return input
}

func day18part1() {
	lines := readFile("day18input")

	var last []string
	for i, line := range lines {
		curr := transformInput(line)
		p := processSnailfish(curr)
		//		fmt.Println("processed", p)
		if i != 0 {
			last = processSnailfish(addSnaiflish(last, p))
		} else {
			last = p
		}
	}
	fmt.Println(strings.Join(last, ""))
	fmt.Println(calcMagnitude(last))
}

func day18() {
	lines := readFile("day18input")

	var sums []int
	for _, line := range lines {
		for _, line2 := range lines {
			p1 := processSnailfish(transformInput(line))
			p2 := processSnailfish(transformInput(line2))
			res := processSnailfish(addSnaiflish(p1, p2))
			sums = append(sums, calcMagnitude(res))
		}
	}
	fmt.Println(max(sums))
}

func processSnailfish(input []string) []string {
	explode := false
	needssplit := false

	leftct := 0

	var ret []string
	lastnum := math.MinInt
	lastnumpos := -1
	for i, s := range input {

		switch s {
		case "[":
			leftct++
		case "]":
			leftct--
		default:
			if isDigit(s) {
				lastnum = atoi(s)
				if lastnum >= 10 {
					needssplit = true
				}
				lastnumpos = i
			}
		}
		// explode
		// pair nested 4 in between
		if leftct == 5 {
			explode = true
			left := atoi(input[i+1])
			right := atoi(input[i+3])

			if lastnum != math.MinInt {
				ret[lastnumpos] = itoa(left + lastnum)
			}
			numright, pos := findNumToRight(input, i+5)

			if pos != -1 {
				numright += right
			}

			ret = append(ret, "0")

			// append the rest to be processed
			for j := i + 5; j < len(input); j++ {
				val := input[j]
				if pos == j {
					val = itoa(numright)
				}
				ret = append(ret, val)
			}
			break
		}

		ret = append(ret, s)
	}

	// didn't do anything and nothing needs to be done, we're done
	if !explode && !needssplit {
		return ret
	}

	// check if further processing needed
	if explode {
		return processSnailfish(ret)
	}
	if needssplit {
		// nothing should have changed
		var splitret []string
		for i, s := range input {
			if isDigit(s) {
				num := atoi(s)
				if num >= 10 {
					l := int(math.Floor(float64(num) / 2))
					r := int(math.Ceil(float64(num) / 2))
					splitret = append(splitret, "[", itoa(l), ",", itoa(r), "]")

					splitret = append(splitret, input[i+1:]...)
					return processSnailfish(splitret)
				}
			}
			splitret = append(splitret, s)
		}
	}
	return processSnailfish(ret)
}

func findNumToRight(input []string, pos int) (int, int) {
	for i := pos; i < len(input); i++ {
		s := input[i]
		if isDigit(s) {
			return atoi(s), i
		}
	}
	return 0, -1
}

func isDigit(s string) bool {
	for _, c := range s {
		if !unicode.IsDigit(c) {
			return false
		}
	}
	return true
}

func addSnaiflish(sf1, sf2 []string) []string {
	var ret []string
	ret = append(ret, "[")
	ret = append(ret, sf1...)
	ret = append(ret, ",")
	ret = append(ret, sf2...)
	ret = append(ret, "]")
	return ret
}

func calcMagnitude(sf []string) int {
	var next []string
	for i := 0; i < len(sf); i++ {
		s := sf[i]
		if isDigit(s) && validMagnitudeCheck(sf, i) {
			l := 3 * atoi(s)
			r := 2 * atoi(sf[i+2])
			next = sf[0 : i-1]
			next = append(next, itoa(l+r))
			next = append(next, sf[i+4:]...)
			break
		}
		next = append(next, s)
	}
	// we've reduced it to just one. done
	if len(next) == 1 {
		return atoi(next[0])
	}
	return calcMagnitude(next)
}

func validMagnitudeCheck(sf []string, pos int) bool {
	if sf[pos-1] != "[" {
		return false
	}
	if pos+2 > len(sf) {
		return false
	}
	return sf[pos+1] == "," && isDigit(sf[pos+2])
}
