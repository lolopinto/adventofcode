package main

import (
	"fmt"
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

	// searchy := 10
	searchy := 2000000

	for i := 0; i < len(sensors); i++ {
		sensor := sensors[i]
		beacon := beacons[i]

		dist := mandistance(sensor, beacon)

		for i := 0; i <= dist; i++ {
			delta := dist - i

			rows := uniq([]int{sensor.Row + i, sensor.Row, sensor.Row - i})
			cols := uniq([]int{sensor.Column + i, sensor.Column, sensor.Column - i})

			addMaybe := func(r, c int) {
				p := grid.NewPos(r, c)
				_, ok := where[p]
				if ok {
					return
				}
				if c == searchy {
					searchfor[p] = true
				}
			}
			for _, col := range cols {
				if col != searchy {
					continue
				}
				for r := delta; r >= 0; r-- {
					addMaybe(sensor.Row-r, col)
				}
				for r := delta; r >= 0; r-- {
					addMaybe(sensor.Row+r, col)
				}
			}

			for _, row := range rows {
				if sensor.Column != searchy {
					continue
				}
				for c := delta; c >= 0; c-- {
					addMaybe(row, sensor.Column-c)
				}
				for c := delta; c >= 0; c-- {
					addMaybe(row, sensor.Column+c)
				}
			}
		}
	}

	fmt.Println(len(searchfor))
}
