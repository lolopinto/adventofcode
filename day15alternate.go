package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/lolopinto/adventofcode2020/grid"
)

func day15alternate() {
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

	searchfor := map[grid.Pos]bool{}

	for _, line := range lines {
		parts := strings.Split(line, " ")
		sensor := grid.NewPos(parse(parts[2]), parse(parts[3]))
		beacon := grid.NewPos(parse(parts[8]), parse(parts[9]))
		where[sensor] = 'S'
		where[beacon] = 'B'

		sensors = append(sensors, sensor)
		beacons = append(beacons, beacon)
	}

	// TODO this should just be rows
	// and not columns, can fix if it's working
	type missingRange struct {
		start grid.Pos
		end   grid.Pos
	}
	ranges := map[int][]missingRange{}

	searchy := 10
	// searchy := 2000000

	// addMaybe := func(r, c int) {
	// 	p := grid.NewPos(r, c)
	// 	_, ok := where[p]
	// 	if ok {
	// 		return
	// 	}
	// 	if c == searchy {
	// 		searchfor[p] = true
	// 	}
	// }

	// does a contain b
	contains := func(a, b missingRange) bool {
		return a.start.Row <= b.start.Row &&
			b.end.Row <= a.end.Row &&
			a.start.Column <= b.start.Column &&
			b.end.Column <= a.end.Column
	}

	intersectsRow := func(a, b missingRange) bool {
		return a.start.Column == b.start.Column && a.end.Column == b.end.Column &&
			a.start.Row <= b.start.Row &&
			b.start.Row <= a.end.Row
	}

	// this and everything else assumes column the same
	// extendLeft := func(a, b missingRange) bool {
	// 	return a.end.Row-b.start.Row == 1
	// }
	// extendRight := func(a, b missingRange) bool {
	// 	return b.start.Row-a.end.Row == 1
	// }

	// intersectsColumn := func(a, b missingRange) bool {
	// 	return a.start.Row == b.start.Row && a.end.Row == b.end.Row &&
	// 		a.start.Column <= b.start.Column &&
	// 		b.start.Column <= a.end.Column
	// }

	maybeAddRange := func(arg missingRange) {

		candidates := []missingRange{}
		if arg.start.Column == arg.end.Column {
			candidates = append(candidates, arg)
			// fmt.Println("keep range per column", r)
		} else {
			// need to split range??
			for c := arg.start.Column; c <= arg.end.Column; c++ {
				candidates = append(candidates, missingRange{start: grid.NewPos(arg.start.Row, c), end: grid.NewPos(arg.end.Row, c)})
			}
			// fmt.Println("split range", r, candidates)
		}

		for _, r := range candidates {
			currentCol := r.start.Column
			if currentCol != searchy {
				continue
			}

			currentRanges := ranges[currentCol]
			if currentRanges == nil {
				currentRanges = []missingRange{}
			}
			toAdd := true
			for idx, existing := range currentRanges {
				log := false
				if r.end.Column == searchy {
					log = true
				}

				if contains(existing, r) {
					// drop new one
					if log {
						fmt.Println("dropping new one", existing, r)
					}
					toAdd = false
					break
				}

				if contains(r, existing) {
					// update existing
					if log {
						fmt.Println("upating existing", existing, r)
					}
					currentRanges[idx] = r
					ranges[currentCol] = currentRanges
					toAdd = false
					break
				}

				// intersects + r bigger
				if intersectsRow(existing, r) {
					if log {
						fmt.Println("intersects update end", existing, r)
					}
					// extend the back
					currentRanges[idx] = missingRange{
						start: existing.start,
						end:   r.end,
					}
					ranges[currentCol] = currentRanges
					if log {
						fmt.Println("new val", currentRanges[idx])
					}
					toAdd = false
					break
				}

				// intersects + existing bigger
				if intersectsRow(r, existing) {
					if log {
						fmt.Println("intersects upadte new", existing, r)
					}
					// extend the back
					currentRanges[idx] = missingRange{
						start: r.start,
						end:   existing.end,
					}
					ranges[currentCol] = currentRanges
					if log {
						fmt.Println("new val", currentRanges[idx])
					}
					toAdd = false
					break
				}

				// if extendLeft(existing, r) {
				// 	if log {
				// 		fmt.Println("extends left", existing, r)
				// 	}
				// 	// extend the left
				// 	currentRanges[idx] = missingRange{
				// 		start: r.start,
				// 		end:   existing.end,
				// 	}
				// 	ranges[currentCol] = currentRanges
				// 	if log {
				// 		fmt.Println("new val", currentRanges[idx])
				// 	}
				// 	toAdd = false
				// 	break
				// }
				// if extendRight(existing, r) {
				// 	if log {
				// 		fmt.Println("extends right", existing, r)
				// 	}
				// 	// extend the left
				// 	currentRanges[idx] = missingRange{
				// 		start: existing.start,
				// 		end:   r.end,
				// 	}
				// 	ranges[currentCol] = currentRanges
				// 	if log {
				// 		fmt.Println("new val", currentRanges[idx])
				// 	}
				// 	toAdd = false
				// 	break
				// }

			}

			if toAdd {
				if r.end.Column == searchy {
					// fmt.Println("new", ranges, r)
				}
				fmt.Println("adding", r)
				currentRanges = append(currentRanges, r)
				ranges[currentCol] = currentRanges
			}
			// fmt.Println(len(ranges))
		}

	}

	sum := 0

	for i := 0; i < len(sensors); i++ {
		sensor := sensors[i]
		beacon := beacons[i]

		dist := mandistance(sensor, beacon)

		for i := 0; i <= dist; i++ {
			delta := dist - i

			rows := uniq([]int{sensor.Row + i, sensor.Row, sensor.Row - i})
			cols := uniq([]int{sensor.Column + i, sensor.Column, sensor.Column - i})

			for _, col := range cols {
				maybeAddRange(
					missingRange{
						start: grid.NewPos(
							sensor.Row-delta,
							col,
						),
						end: grid.NewPos(
							sensor.Row+delta,
							col,
						),
					})

				// if col != searchy {
				// 	continue
				// }
				// for r := delta; r >= 0; r-- {
				// 	addMaybe(sensor.Row-r, col)
				// }
				// for r := delta; r >= 0; r-- {
				// 	addMaybe(sensor.Row+r, col)
				// }
			}

			for _, row := range rows {
				maybeAddRange(missingRange{
					start: grid.NewPos(
						row,
						sensor.Column-delta,
					),
					end: grid.NewPos(
						row,
						sensor.Column+delta,
					),
				})

				// if sensor.Column != searchy {
				// 	continue
				// }
				// for c := delta; c >= 0; c-- {
				// 	addMaybe(row, sensor.Column-c)
				// }
				// for c := delta; c >= 0; c-- {
				// 	addMaybe(row, sensor.Column+c)
				// }
			}
		}
	}

	// old answer
	fmt.Println("old answer", len(searchfor))

	// var potentialranges []missingRange

	// for _, v := range ranges {
	// 	if v.start.Column >= searchy && v.end.Column <= searchy {
	// 		potentialranges = append(potentialranges, v)
	// 	}
	// }
	// do we wanna change it so ranges are stored per column??
	potentialranges := ranges[searchy]

	sort.Slice(potentialranges, func(i, j int) bool {
		return potentialranges[i].start.Row < potentialranges[j].start.Row
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

	lastend := -1
	for i, v := range result {
		sum += (v.end.Row - v.start.Row) + 1

		if i != 0 && lastend > v.start.Row {
			sum -= (lastend + 1 - v.start.Row)
		}
		lastend = v.end.Row
	}

	for v := range where {
		if v.Column == searchy {
			sum--
		}
	}

	fmt.Println(result)
	fmt.Println(sum)
}
