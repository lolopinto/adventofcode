package main

import (
	"log"
	"strings"
)

type passport struct {
	byr string
	iyr string
	eyr string
	hgt string
	hcl string
	ecl string
	pid string
	cid string
}

var mKeys = map[string]bool{
	"byr": true,
	"iyr": true,
	"eyr": true,
	"hgt": true,
	"hcl": true,
	"ecl": true,
	"pid": true,
	"cid": false,
}

func day4() {
	lines := readFile("day4input")

	data := make(map[string]string)
	numValid := 0

	checkValid := func() {
		for k, v := range mKeys {
			_, ok := data[k]
			if v && !ok {
				data = make(map[string]string)
				return
			}
		}
		numValid++
		data = make(map[string]string)
	}

	for _, line := range lines {
		if line == "" {
			checkValid()

		} else {
			parts := strings.Split(line, " ")
			for _, part := range parts {
				elem := strings.Split(part, ":")
				if len(elem) != 2 {
					log.Fatalf("the parts of a passport were not as expected %s", part)
				}
				key := elem[0]
				val := elem[1]
				data[key] = val
			}
		}
	}
	checkValid()

	log.Println(numValid)
}
