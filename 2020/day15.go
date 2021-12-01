package main

import (
	"log"
	"strings"
)

func day15() {
	lines := readFile("day15input")

	line := lines[0]
	nos := strings.Split(line, ",")
	numbers := ints(nos)

	lastNumber := 0

	m := make(map[int][]int)
	takeTurn := func(turn int) int {
		var number int
		if turn-1 < len(numbers) {
			number = numbers[turn-1]
		} else {
			prevTurns, ok := m[lastNumber]
			if !ok {
				log.Fatal("shouldn't happen")
			} else {
				l := len(prevTurns)
				if l == 1 {
					number = 0
				} else {
					number = prevTurns[l-1] - prevTurns[l-2]
				}
			}
		}

		prevTurns, ok := m[number]
		if ok {
			m[number] = append(prevTurns, turn)
		} else {
			m[number] = []int{turn}
		}

		return number
	}

	for turn := 1; turn <= 30000000; turn++ {
		lastNumber = takeTurn(turn)
	}
	log.Println(lastNumber)
}
