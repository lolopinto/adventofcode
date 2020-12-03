package main

// https://adventofcode.com/2020/day/2#part2
import (
	"log"
	"strconv"
	"strings"
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
	parts := strings.Split(line, " ")
	if len(parts) != 3 {
		log.Fatalf("invalid policy %s", line)
	}
	numParts := strings.Split(parts[0], "-")
	if len(numParts) != 2 {
		log.Fatalf("invalid range %s", parts[0])
	}
	firstPos, err := strconv.Atoi(numParts[0])
	secondPos, err2 := strconv.Atoi(numParts[1])

	if err != nil {
		log.Fatal(err)
	}
	if err2 != nil {
		log.Fatal(err2)
	}

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
