package main

import (
	"fmt"
	"math"
	"regexp"
	"strings"
)

type valve struct {
	name     string
	flowRate int
	valves   []string
}

func (v *valve) canOpen() bool {
	return v.flowRate != 0
}

type valveOpenInfo struct {
	valve        *valve
	minuteOpened int
}

func day16() {
	lines := readFile("day16input")

	m := map[string]*valve{}
	r := regexp.MustCompile(`Valve (.+) has flow rate=(.+); (?:tunnels|tunnel) (?:lead|leads) to (?:valves|valve) (.+)`)

	// var curr *valve
	// open := []valve{}

	openableCt := 0
	for _, line := range lines {
		match := r.FindStringSubmatch(line)
		if match == nil {
			panic("invalid regex")
		}
		v := &valve{
			name:     match[1],
			flowRate: atoi(match[2]),
			valves:   strings.Split(match[3], ", "),
		}
		if v.flowRate > 0 {
			openableCt++
		}
		// fmt.Println("flow rate", v.name, v.flowRate, v.valves)
		m[v.name] = v
	}

	open := map[string]*valveOpenInfo{}

	// always start at A

	start := m["AA"]
	// move if we cannot open
	fmt.Println(traverseValves(m["AA"], m, openableCt, open, 1, !start.canOpen()))
}

// TODO update logic to account for opening early. seems like it'll matter in part 2

func traverseValves(curr *valve, valves map[string]*valve, openableCt int, open map[string]*valveOpenInfo, minute int, move bool) int {
	// fmt.Println(curr.name, minute, move)
	if minute == 31 || len(open) == openableCt {
		sum := 0
		names := []string{}
		for _, v := range open {
			// fmt.Println("open", v.valve.name, v.minuteOpened, v.valve.flowRate)
			sum += ((30 - v.minuteOpened) * v.valve.flowRate)
			names = append(names, v.valve.name)
		}
		if sum == 1642 {
			if len(names) > 5 {
				fmt.Println("open", names)
				for _, v := range open {
					fmt.Println("v", v.minuteOpened, v.valve.name, v.valve.flowRate)
				}
			}
		}
		return sum
	}

	if move {
		// todo move to all combinations, do math and see which is the max
		// this would do all possibilities here and return up
		max := math.MinInt
		for _, cand := range curr.valves {
			moveTo := valves[cand]
			// fmt.Printf("moveFrom %s moveTo %s\n", curr.name, moveTo.name)

			// if move.To. flowRate == 0, move instead of opening
			v := traverseValves(moveTo, valves, openableCt, open, minute+1, !moveTo.canOpen())
			if v > max {
				max = v
			}
		}
		return max
		// moveTo := valves[curr.valves[0]]
		// fmt.Printf("moveFrom %s moveTo %s\n", curr.name, moveTo.name)
		// return traverseValves(moveTo, valves, open, minute+1, false)
	} else {
		vInfo := open[curr.name]
		// not open and worth opening
		newOpen := open
		if vInfo == nil && curr.flowRate > 0 {
			newOpen = copyOpenValves(open)
			newOpen[curr.name] = &valveOpenInfo{
				minuteOpened: minute,
				valve:        curr,
			}
		}
		return traverseValves(curr, valves, openableCt, newOpen, minute+1, true)
	}
}

func copyOpenValves(open map[string]*valveOpenInfo) map[string]*valveOpenInfo {
	ret := map[string]*valveOpenInfo{}
	for k, v := range open {
		ret[k] = v
	}
	return ret
}
