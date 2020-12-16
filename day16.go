package main

import (
	"log"
	"strings"
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
	var myticket string
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
			myticket = line
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

	l := len(ranges)
	result := make(map[int]map[string]bool)
	for i := 0; i < l; i++ {

		//		all := true
		var data map[string]bool
		for idx, t := range validTickets {
			if idx == 0 {
				data = t.validPos[i]
			} else {
				current := t.validPos[i]
				for k := range data {
					if !current[k] {
						delete(data, k)
					}
				}
			}
		}
		result[i] = data
	}

	m := make(map[string]int)
	for len(result) > 0 {
		for k, v := range result {
			if len(v) == 1 {

				// set the key in the final map...
				var found string
				for key := range v {
					m[key] = k
					found = key
				}

				for k2, v2 := range result {
					if k2 == k {
						continue
					}
					// delete key from all other results
					delete(v2, found)
				}
				// delete what we just found
				delete(result, k)
			}
		}
	}

	numbers := ints(strings.Split(myticket, ","))

	mult := 1
	for k, v := range m {
		if strings.HasPrefix(k, "departure") {
			mult *= numbers[v]
		}
	}
	log.Println(mult)
}
