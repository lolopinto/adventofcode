package main

import (
	"strings"

	"github.com/davecgh/go-spew/spew"
)

type ticketRange struct {
	key                string
	low, hi, low2, hi2 int
	pos                int
}

type ticket struct {
	// line     string
	// numbers  []int
	validPos map[int]map[string]bool
}

func day16() {

	mytickets := false
	nearby := false
	lines := readFile("day16input")

	ranges := []ticketRange{}
	validTickets := []ticket{}
	var ticket string
	for _, line := range lines {
		if line == "" {
			continue
		}

		if strings.HasPrefix(line, "your ticket:") {
			mytickets = true
			continue
		}

		if strings.HasPrefix(line, "nearby tickets:") {
			nearby = true
			continue
		}

		if !mytickets && !nearby {
			parts := splitLength(line, ": ", 2)
			key := parts[0]
			data := splitLength(parts[1], " or ", 2)

			first := splitLength(data[0], "-", 2)
			second := splitLength(data[1], "-", 2)

			ranges = append(ranges, ticketRange{
				key:  key,
				low:  atoi(first[0]),
				hi:   atoi(first[1]),
				low2: atoi(second[0]),
				hi2:  atoi(second[1]),
			})
		}

		if mytickets && !nearby {
			ticket = line
			continue
		}

		if nearby {
			numbers := ints(strings.Split(line, ","))
			validPos := make(map[int]map[string]bool)
			validNum := 0
			for numCol, num := range numbers {
				valid := false
				for _, r := range ranges {
					if (r.low <= num && r.hi >= num) || (r.low2 <= num && r.hi2 >= num) {
						valid = true
						m, ok := validPos[numCol]
						if !ok {
							m = make(map[string]bool)

						}
						m[r.key] = true
						validPos[numCol] = m
					}
				}
				if valid {
					validNum++
				}

			}
			if validNum == len(numbers) {
				t := ticket{
					validPos: validPos,
				}
				validTickets = append(validTickets, t)
			}
		}
	}

	ignore := map[string]bool{
		"duration":           true,
		"arrival track":      true,
		"row":                true,
		"arrival location":   true,
		"seat":               true,
		"arrival station":    true,
		"departure platform": true,
		"departure location": true,
		"departure date":     true,
		"departure station":  true,
		"departure track":    true,
	}
	//	spew.Dump(validTickets)
	l := len(ranges)
	result := make(map[int][]string)
	for i := 0; i < l; i++ {

		//		all := true
		var last map[string]bool
		for idx, t := range validTickets {
			if idx == 0 {
				last = t.validPos[i]
			} else {
				current := t.validPos[i]
				for k := range last {
					if !current[k] || ignore[k] {
						last[k] = false
					}
				}
			}
		}
		l := []string{}
		for k, v := range last {
			if v {
				l = append(l, k)
			}
		}
		result[i] = l
	}

	spew.Dump(result)
}
