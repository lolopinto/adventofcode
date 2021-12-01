package main

import (
	"fmt"
	"log"
	"sort"
)

func countJoltPaths(list []int, m map[string]int, initial, max int) int {
	key := fmt.Sprintf("%d:%d", len(list), initial)
	v, ok := m[key]
	//	spew.Dump(m)
	if ok {
		return v
	}
	currCount := 0
	if max-initial <= 3 {
		currCount++
	}
	if len(list) == 0 {
		return currCount
	}
	if list[0]-initial <= 3 {
		currCount += countJoltPaths(list[1:], m, list[0], max)
	}
	currCount += countJoltPaths(list[1:], m, initial, max)
	m[key] = currCount
	//	log.Println(key, currCount)
	return currCount
}

func day10() {
	numbers := readInts("day10input")
	sort.Ints(numbers)

	// tried to do something iteratively initially and couldn't get it to work :(
	m := make(map[string]int)
	log.Println(countJoltPaths(numbers, m, 0, numbers[len(numbers)-1]+3))
}
