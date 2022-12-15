package main

import (
	"fmt"
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

	// missing
	where := make(map[grid.Pos]rune)
	dist := make(map[grid.Pos]int)
	missing := make(map[grid.Pos]rune)

	rows := []int{}
	sensors := []grid.Pos{}

	for _, line := range lines {
		parts := strings.Split(line, " ")
		sensor := grid.NewPos(parse(parts[2]), parse(parts[3]))
		beacon := grid.NewPos(parse(parts[8]), parse(parts[9]))
		where[sensor] = 'S'
		where[beacon] = 'B'

		dist[sensor] = mandistance(beacon, sensor)

		rows = append(rows, sensor.Row, beacon.Row)

		sensors = append(sensors, sensor)
	}

	c := 10
	for r := min(rows); r <= max(rows); r++ {
		// for c := min(cols); c <= max(cols); c++ {

		log := false
		if r == -2 || r == 14 || r == 24 {
			log = true
		}
		pos := grid.NewPos(r, c)
		_, ok := where[pos]
		// sensor or beacon, nothing to do here
		if ok {
			// if log {
			// 	fmt.Printf("something %v at %v \n", v, pos)
			// }
			// fmt.Println("existing", pos)
			continue
		}
		for _, sensor := range sensors {
			newdist := mandistance(pos, sensor)
			// tie logic means something else is broken
			// There is never a tie where two beacons are the same distance to a sensor.
			if newdist < dist[sensor] {
				// fmt.Println(pos)
				missing[pos] = '#'
			} else if log {
				fmt.Printf("conflict pos %v sensor %v newdist %d existing dist %d \n", pos, sensor, newdist, dist[sensor])
			}
			// else if log {
			// 	// fmt.Printf("distance fail %d pos %v mandist %d \n", newdist, pos, dist[sensor])
			// }
		}
		// }
	}

	// // fmt.Println(len(missing))
	colct := map[int]int{}
	for k := range missing {
		ct := colct[k.Column]
		ct++
		colct[k.Column] = ct
	}
	// // fmt.Println(len(missing))
	// fmt.Println(colct)
	// // fmt.Println(where)
	// fmt.Println(len(colct))
	fmt.Println(colct[c])
	// fmt.Println(ct)
}
