package main

import (
	"fmt"
	"regexp"
	"strings"
)

type valve struct {
	name     string
	flowRate int
	tunnels  map[string]bool
	paths    map[string]int
}

// get the depth to each destination from the given tunnels
// if in given list of tunnels, would be 1
// otherwise, go through next set of tunnels till we find it
func getDepthToPath(valves map[string]*valve, tunnels map[string]bool, dest string) int {
	depth := 1
	for {
		next_tunnels := map[string]bool{}
		for item := range tunnels {
			if item == dest {
				return depth
			}
			for item2 := range valves[item].tunnels {
				next_tunnels[item2] = true
			}
		}
		tunnels = next_tunnels
		depth += 1
	}
}

func day16() {
	lines := readFile("day16input")

	valves := map[string]*valve{}
	r := regexp.MustCompile(`Valve (.+) has flow rate=(.+); (?:tunnels|tunnel) (?:lead|leads) to (?:valves|valve) (.+)`)

	openable := []string{}

	for _, line := range lines {
		match := r.FindStringSubmatch(line)
		if match == nil {
			panic("invalid regex")
		}
		v := &valve{
			name:     match[1],
			flowRate: atoi(match[2]),
			tunnels:  mapify(strings.Split(match[3], ", ")),
			paths:    map[string]int{},
		}
		if v.flowRate > 0 {
			openable = append(openable, v.name)
		}
		valves[v.name] = v
	}
	beginning := openable[:]
	beginning = append(beginning, "AA")
	for _, v := range beginning {
		for _, v2 := range openable {
			if v != v2 {
				valves[v].paths[v2] = getDepthToPath(valves, valves[v].tunnels, v2)
			}
		}
	}

	// best1 := 0
	// best2 := 0
	opened := map[string]bool{
		"AA": true,
	}
	// part 1
	best1 := traverse(valves["AA"], valves, opened, 29, 0)

	// // part 2
	best2 := traverseElephant(valves["AA"], valves, opened, 25, 0, false)

	fmt.Println("part 1", best1)
	fmt.Println("part 2", best2)
}

// traverse based on depths
// inspired by a solution on reddit since what i was doing wasn't working
func traverse(curr *valve, valves map[string]*valve, opened map[string]bool, depth, current int) int {
	best := current

	if depth <= 0 {
		return best
	}

	if !opened[curr.name] {
		opened2 := copyOpened(opened)
		opened2[curr.name] = true
		current += (depth * curr.flowRate)
		v := traverse(curr, valves, opened2, depth-1, current)
		if v > best {
			best = v
		}
	} else {
		for v, new_depth := range valves[curr.name].paths {
			if !opened[v] {
				v := traverse(valves[v], valves, opened, depth-new_depth, current)
				if v > best {
					best = v
				}
			}
		}
	}
	return best
}

func traverseElephant(curr *valve, valves map[string]*valve, opened map[string]bool, depth, current int, elephant bool) int {
	best := current

	if depth <= 0 {
		return best
	}

	if !opened[curr.name] {
		opened2 := copyOpened(opened)
		opened2[curr.name] = true
		current += (depth * curr.flowRate)
		v := traverseElephant(curr, valves, opened2, depth-1, current, elephant)
		if v > best {
			best = v
		}
		if !elephant {
			// elephant starts here
			v := traverseElephant(valves["AA"], valves, opened2, 25, current, true)
			if v > best {
				best = v
			}
		}
	} else {
		for v, new_depth := range valves[curr.name].paths {
			if !opened[v] {
				v := traverseElephant(valves[v], valves, opened, depth-new_depth, current, elephant)
				if v > best {
					best = v
				}
			}
		}
	}
	return best
}

func copyOpened(open map[string]bool) map[string]bool {
	ret := make(map[string]bool, len(open))
	for k, v := range open {
		ret[k] = v
	}
	return ret
}
