package main

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/lolopinto/adventofcode2020/grid"
)

func day15() {
	lines := readFile("day15input")

	parse := func(s string) int {
		s = strings.TrimRight(s, ":")
		s = strings.TrimRight(s, ",")
		parts := splitLength(s, "=", 2)
		return atoi(parts[1])
	}

	mandistance := func(p1, p2 grid.Pos) int {
		return abs(p1.Row, p2.Row) + abs(p1.Column, p2.Column)
	}

	where := make(map[grid.Pos]rune)

	sensors := []grid.Pos{}
	beacons := []grid.Pos{}

	for _, line := range lines {
		parts := strings.Split(line, " ")
		sensor := grid.NewPos(parse(parts[2]), parse(parts[3]))
		beacon := grid.NewPos(parse(parts[8]), parse(parts[9]))
		where[sensor] = 'S'
		where[beacon] = 'B'

		sensors = append(sensors, sensor)
		beacons = append(beacons, beacon)
	}

	type missingRange struct {
		// just rows
		start, end int
	}

	ranges := map[int][]missingRange{}

	miny := 0

	// example
	// searchy := 10
	// maxy := 20

	// question
	searchy := 2000000
	maxy := 4000000

	// does a contain b
	contains := func(a, b missingRange) bool {
		return a.start <= b.start &&
			b.end <= a.end
	}

	intersects := func(a, b missingRange) bool {
		return a.start <= b.start &&
			b.start <= a.end
	}

	maybeAddRange := func(currentCol int, r missingRange) {

		currentRanges := ranges[currentCol]
		if currentRanges == nil {
			currentRanges = []missingRange{}
		}
		for idx, existing := range currentRanges {

			if contains(existing, r) {
				// drop new one
				return
			}

			if contains(r, existing) {
				// update existing
				currentRanges[idx] = r
				ranges[currentCol] = currentRanges
				return
			}

			// intersects + r bigger
			if intersects(existing, r) {

				// extend the back
				currentRanges[idx] = missingRange{
					start: existing.start,
					end:   r.end,
				}
				ranges[currentCol] = currentRanges
				return
			}

			// intersects + existing bigger
			if intersects(r, existing) {

				// extend the back
				currentRanges[idx] = missingRange{
					start: r.start,
					end:   existing.end,
				}
				ranges[currentCol] = currentRanges
				return
			}
		}

		currentRanges = append(currentRanges, r)
		ranges[currentCol] = currentRanges
	}

	sum := 0

	for i := 0; i < len(sensors); i++ {
		sensor := sensors[i]
		beacon := beacons[i]

		dist := mandistance(sensor, beacon)

		for i := 0; i <= dist; i++ {
			delta := dist - i

			cols := []int{sensor.Column + i, sensor.Column, sensor.Column - i}

			for _, col := range cols {
				maybeAddRange(
					col,
					missingRange{
						start: sensor.Row - delta,
						end:   sensor.Row + delta,
					})
			}
		}
	}

	// fix any duplicate ranges so that we end up with something easily parseable
	// sort ranges and remove ranges already contained in sorted version

	for k, potentialranges := range ranges {

		sort.Slice(potentialranges, func(i, j int) bool {
			return potentialranges[i].start < potentialranges[j].start
		})

		var result []missingRange
		for _, v := range potentialranges {
			add := true
			for _, safe := range result {
				if contains(safe, v) {
					add = false
					break
				}
			}
			if add {
				result = append(result, v)
			}
		}
		ranges[k] = result
	}

	result := ranges[searchy]

	lastend := -1
	// sum up all missing beacons based on ranges
	for i, v := range result {
		sum += (v.end - v.start) + 1

		if i != 0 && lastend > v.start {
			sum -= (lastend + 1 - v.start)
		}
		lastend = v.end
	}

	// subtract locations that have beacons or sensors
	for v := range where {
		if v.Column == searchy {
			sum--
		}
	}

	fmt.Println("part 1:", sum)

	for col, cands := range ranges {
		if col > maxy || col < miny || len(cands) == 1 {
			continue
		}
		lastend := -1
		// search for the missing beacon
		for i, c := range cands {
			if i != 0 {
				if c.start-lastend > 1 {
					fmt.Println("part 2:", (lastend+1)*4000000+col)
					os.Exit(0)
				}
			}
			lastend = c.end
		}
	}
}
