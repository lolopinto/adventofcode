package main

import (
	"log"
	"sort"
	"strings"
)

func day13() {
	lines := readFile("day13input")
	timestamp := atoi(lines[0])
	//	data := lines[1]

	parts := strings.Split(lines[1], ",")
	var ids []int
	for _, part := range parts {
		if part == "x" {
			continue
		}
		ids = append(ids, atoi(part))
	}
	earliest := [][]int{}
	for _, ts := range ids {
		log.Println(timestamp / ts)
		for i := 0; true; i++ {
			if ts*i > timestamp {
				earliest = append(earliest, []int{ts * i, ts})
				break
			}
		}
	}
	sort.Slice(earliest, func(i, j int) bool {
		return earliest[i][0] < earliest[j][0]
	})
	log.Println(ids)
	log.Println(earliest)
	log.Println((earliest[0][0] - timestamp) * earliest[0][1])

}
