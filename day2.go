package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type policy struct {
	minOccurence int
	maxOccurence int
	letter       rune
	password     string
}

func (p policy) validPassword() bool {
	count := 0
	for _, c := range p.password {
		if p.letter == c {
			count++
		}
	}
	return count >= p.minOccurence && count <= p.maxOccurence
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
	minOccurence, err := strconv.Atoi(numParts[0])
	maxOccurence, err2 := strconv.Atoi(numParts[1])

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
		minOccurence: minOccurence,
		maxOccurence: maxOccurence,
		letter:       letter,
		password:     parts[2],
	}

}

func day2() {
	b, err := ioutil.ReadFile("day2input")
	if err != nil {
		log.Fatal(err)
	}

	str := string(b)
	lines := strings.Split(str, "\n")

	validPassword := 0
	for _, line := range lines {
		p := parsePolicy(line)
		if p.validPassword() {
			validPassword++
		}
	}
	log.Println(validPassword)
}
