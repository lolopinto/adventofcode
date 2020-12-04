package main

import (
	"log"
	"strconv"
	"strings"
	"unicode"
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

type info interface {
	required() bool
	valid(str string) bool
}

type requiredInfo struct{}

func (i requiredInfo) required() bool {
	return true
}

func validRange(str string, min, max int) bool {
	i, err := strconv.Atoi(str)
	if err != nil {
		return false
	}
	return i >= min && i <= max
}

type byr struct {
	requiredInfo
}

func (b *byr) valid(str string) bool {
	return validRange(str, 1920, 2002)
}

type iyr struct {
	requiredInfo
}

func (b *iyr) valid(str string) bool {
	return validRange(str, 2010, 2020)
}

type eyr struct {
	requiredInfo
}

func (b *eyr) valid(str string) bool {
	return validRange(str, 2020, 2030)
}

type hgt struct {
	requiredInfo
}

func (b *hgt) valid(str string) bool {
	if strings.HasSuffix(str, "cm") {
		return validRange(strings.TrimSuffix(str, "cm"), 150, 193)
	} else if strings.HasSuffix(str, "in") {
		return validRange(strings.TrimSuffix(str, "in"), 59, 76)
	}
	return false
}

type hcl struct {
	requiredInfo
}

func (b *hcl) valid(str string) bool {
	if len(str) != 7 {
		return false
	}
	for i, c := range str {
		if i == 0 {
			if c != '#' {
				return false
			}
		} else {
			if !unicode.IsDigit(c) && !unicode.IsLower(c) {
				return false
			}
		}
	}
	return true
}

type ecl struct {
	requiredInfo
}

func (b *ecl) valid(str string) bool {
	switch str {
	case "amb", "blu", "brn", "gry", "grn", "hzl", "oth":
		return true
	default:
		return false
	}
}

type pid struct {
	requiredInfo
}

func (b *pid) valid(str string) bool {
	if len(str) != 9 {
		return false
	}
	for _, c := range str {
		if !unicode.IsDigit(c) {
			return false
		}
	}
	return true
}

type cid struct{}

func (b *cid) required() bool {
	return false
}

func (b *cid) valid(str string) bool {
	return true
}

var mKeys = map[string]info{
	"byr": &byr{},
	"iyr": &iyr{},
	"eyr": &eyr{},
	"hgt": &hgt{},
	"hcl": &hcl{},
	"ecl": &ecl{},
	"pid": &pid{},
	"cid": &cid{},
}

func day4() {
	lines := readFile("day4input")

	data := make(map[string]string)
	numValid := 0

	checkValid := func() {
		for k, infoInstance := range mKeys {
			val, ok := data[k]
			if infoInstance.required() && (!ok || !infoInstance.valid(val)) {
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
