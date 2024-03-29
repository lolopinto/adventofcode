package main

import (
	"fmt"
	"regexp"
	"strings"
)

type Blueprint struct {
	number      int
	robots      map[string]Robot
	have        map[string]int
	robotsOwned map[string]int
}

func (b *Blueprint) clone() *Blueprint {
	b2 := Blueprint{
		number: b.number,
	}
	b2.robots = copyMap(b.robots)
	b2.have = copyMap(b.have)
	b2.robotsOwned = copyMap(b.robotsOwned)

	return &b2
}

func (b *Blueprint) print() {
	fmt.Printf("Blueprint %d\n", b.number)
	fmt.Println("have")
	for str, ct := range b.have {
		fmt.Println(str, ct)
	}
	fmt.Println("robotsOwned")
	for str, ct := range b.robotsOwned {
		fmt.Println(str, ct)
	}
	// always the same...
	// fmt.Println("robots")
	// for _, r := range b.robots {
	// 	fmt.Println(r.material, r.costs)
	// }
}

func (b *Blueprint) collect() int {
	for material, many := range b.robotsOwned {
		if material == "" {
			fmt.Println(b.robotsOwned, b.have)
			panic("invalid have ")
		}
		ct := b.have[material]
		ct += many

		b.have[material] = ct
	}

	return b.have["geode"]
}

func (b *Blueprint) update(building string) int {
	ret := b.collect()
	if building != "" {
		ct := b.robotsOwned[building]
		b.robotsOwned[building] = ct + 1
	}

	return ret
}

var materials = []string{
	"geode",
	"obsidian",
	"clay",
	"ore",
}

func (b *Blueprint) key(min int, building string) string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("%d-%d-%s", b.number, min, building))
	for m, ct := range b.robotsOwned {
		sb.WriteString(fmt.Sprintf("robots-owned-%s-%d", m, ct))
	}

	return sb.String()
}

// not being used and doesn't work for part 2 of blueprint...
func (b *Blueprint) twoAway(m string) bool {
	r := b.robots[m]

	for _, cost := range r.costs {
		targetDiff := cost.cost - b.have[cost.material]
		owned := b.robotsOwned[cost.material]

		r2 := b.robots[cost.material]

		if cost.material != "ore" && !r2.canAfford(b.have) {
			return false
		}

		if targetDiff < 0 || owned <= 0 {

			return false
		}

		if targetDiff/owned > 2 {
			return false
		}
	}
	return true
}

// assumes canAfford
func (b Blueprint) shouldBuild(material string) bool {
	// based on what we have, is there enough time to build to use this???
	// should never get here actually
	if material == "geode" {
		return true
	}

	have := b.have[material]

	for _, r := range b.robots {
		for _, cost := range r.costs {
			if cost.material == material {
				if have < cost.cost {
					return true
				}
			}
		}
	}

	//  TODO this should be false... to theoretically speed up?
	// NOT WORKING!
	return true
}

type Robot struct {
	costs    []RobotCost
	material string
}

func (r Robot) canAfford(have map[string]int) bool {
	for _, cost := range r.costs {
		v := have[cost.material]
		if v < cost.cost {
			return false
		}
	}
	return true
}

// assumes canAfford == true
func (r Robot) spend(have map[string]int) {
	for _, cost := range r.costs {
		curr := have[cost.material]
		curr = curr - cost.cost
		have[cost.material] = curr
	}
}

type RobotCost struct {
	cost     int
	material string
}

func day19() {
	lines := readFile("day19input")
	r := regexp.MustCompile(`Each (.+) robot costs (.+)`)

	blueprints := make([]*Blueprint, len(lines))
	sum := 0
	product := 1
	for i, line := range lines {
		parts := splitLength(line, ": ", 2)
		number := atoi(splitLength(parts[0], " ", 2)[1])
		costs := strings.Split(parts[1], ".")

		b := &Blueprint{
			number: number,
			robots: map[string]Robot{},
			have: map[string]int{
				"ore":      0,
				"obsidian": 0,
				"geode":    0,
				"clay":     0,
			},
			robotsOwned: map[string]int{
				"ore":      1,
				"obsidian": 0,
				"geode":    0,
				"clay":     0,
			},
		}
		for _, cost := range costs {
			if strings.TrimSpace(cost) == "" {
				continue
			}
			match := r.FindStringSubmatch(cost)
			if len(match) == 0 {
				panic(fmt.Errorf("couldn't parse %s", cost))
			}
			material := match[1]
			costs := strings.Split(match[2], " and ")
			robotCosts := make([]RobotCost, len(costs))
			for idx, cost := range costs {
				parts := splitLength(cost, " ", 2)
				robotCosts[idx] = RobotCost{
					cost:     atoi(parts[0]),
					material: parts[1],
				}
			}

			b.robots[material] = Robot{
				material: material,
				costs:    robotCosts,
			}
		}

		cache := map[string]int{}
		geode := runBluePrint(b, 1, 24, 0, "", cache)

		blueprints[i] = b
		sum += (b.number * geode)
		fmt.Println("answer", geode)
	}

	fmt.Println("part 1 sum", sum)

	for _, b := range blueprints[:3] {
		cache := map[string]int{}
		// b.print()
		geode := runBluePrint(b, 1, 32, 0, "", cache)
		product *= geode
	}
	fmt.Println("part 2 answer", product)

}

func runBluePrint(b *Blueprint, start, target, best int, building string, cache map[string]int) int {

	key := b.key(start, building)
	v, ok := cache[key]
	if ok {
		return v
	}

	if start > target {
		return best
	}

	skip := make(map[string]bool)

	build := func(r Robot) {
		b2 := b.clone()
		robot := b2.robots[r.material]
		robot.spend(b2.have)

		v := b2.update(robot.material)
		v = runBluePrint(b2, start+1, target, v, r.material, cache)
		if v > best {
			best = v
		}
	}

	// if can afford a geode. do it. and we're done
	if b.robots["geode"].canAfford(b.have) {
		build(b.robots["geode"])
		return best
	}

	for _, material := range materials {
		r := b.robots[material]

		// two away from this, don't do anything
		// seems wrong
		// this is back to a weird hack that we just need to figure out a generic way to avoid
		// don't end up using it
		if (material == "obsidian" || material == "geode") && b.twoAway(material) {
			// fmt.Println("two minute away", material)
			// break
			for _, cost := range r.costs {
				skip[cost.material] = true
			}
			// TODO SHOULD BE WORKING BUT IT'S NOT
			// continue
		}

		if skip[material] {
			// TODO THIS SHOULD BE WORKING
			// continue
		}

		// see if we can afford all and then do every combination

		afford := r.canAfford(b.have)
		should := b.shouldBuild(material)
		if afford && should {
			build(r)
		}
	}
	// check the don't spend robot route
	b2 := b.clone()
	v = b2.update("")
	v = runBluePrint(b2, start+1, target, v, "", cache)
	if v > best {
		best = v
	}

	cache[key] = best
	return best
}
