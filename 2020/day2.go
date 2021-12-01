package main

// https://adventofcode.com/2020/day/2#part2
import (
	"log"
)

type policy struct {
	firstPos  int
	secondPos int
	letter    rune
	password  string
}

func (p policy) validPassword() bool {
	count := 0
	for i, c := range p.password {
		// there's a more efficient way of doing this but just easier to do it this way
		if i+1 != p.firstPos && i+1 != p.secondPos {
			continue
		}
		if p.letter == c {
			count++
		}
	}
	return count == 1
}

func parsePolicy(line string) policy {
	parts := splitLength(line, " ", 3)
	numParts := splitLength(parts[0], "-", 2)
	firstPos := atoi(numParts[0])
	secondPos := atoi(numParts[1])

	var letter rune
	for _, c := range parts[1] {
		letter = c
		break
	}

	return policy{
		firstPos:  firstPos,
		secondPos: secondPos,
		letter:    letter,
		password:  parts[2],
	}

}

func day2() {
	lines := readFile("day2input")

	validPassword := 0
	for _, line := range lines {
		p := parsePolicy(line)
		if p.validPassword() {
			validPassword++
		}
	}
	log.Println(validPassword)
}
