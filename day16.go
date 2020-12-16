package main

import (
	"log"
	"strings"
)

type ticketRange struct {
	low, hi, low2, hi2 int
}

func day16() {

	mytickets := false
	nearby := false
	lines := readFile("day16input")

	ranges := []ticketRange{}
	invalid := []int{}
	for _, line := range lines {
		//		log.Println("line", line)
		if line == "" {
			continue
		}

		if strings.HasPrefix(line, "your ticket:") {
			//			log.Println("your ticket")
			mytickets = true
			continue
		}

		if strings.HasPrefix(line, "nearby tickets:") {
			//			log.Println("neary tickets")
			nearby = true
			continue
		}

		if !mytickets && !nearby {
			//			log.Println("data")
			parts := splitLength(line, ": ", 2)
			data := splitLength(parts[1], " or ", 2)

			first := splitLength(data[0], "-", 2)
			second := splitLength(data[1], "-", 2)

			ranges = append(ranges, ticketRange{
				low:  atoi(first[0]),
				hi:   atoi(first[1]),
				low2: atoi(second[0]),
				hi2:  atoi(second[1]),
			})
			// strings.Split(line, ":")
			// strings.Split(line, "or")
		}

		if mytickets && !nearby {
			// ignore
			continue
		}

		log.Println(len(ranges), ranges)
		if nearby {
			//			log.Println("nearby")
			numbers := ints(strings.Split(line, ","))
			for _, num := range numbers {
				valid := false
				for _, r := range ranges {
					if (r.low <= num && r.hi >= num) || (r.low2 <= num && r.hi2 >= num) {
						log.Println("valid", num, r)
						valid = true
						break
					}
				}
				if !valid {
					invalid = append(invalid, num)
				}
			}
		}
		//		strings.Split(line, "or")
	}
	log.Println("invalid", invalid)

	sum := 0
	for _, inv := range invalid {
		sum += inv
	}
	log.Println(sum)
}
