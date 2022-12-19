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
	b2.robots = make(map[string]Robot, len(b.robots))
	for k, v := range b.robots {
		b2.robots[k] = v
	}
	b2.have = make(map[string]int, len(b.have))
	for k, v := range b.have {
		b2.have[k] = v
	}

	b2.robotsOwned = make(map[string]int, len(b.robotsOwned))
	for k, v := range b.robotsOwned {
		b2.robotsOwned[k] = v
	}

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

func (b *Blueprint) update(newRobots map[string]int) {
	for material, many := range b.robotsOwned {
		if material == "" {
			fmt.Println(b.robotsOwned, b.have)
			panic("invalid have ")
		}
		ct := b.have[material]
		ct += many

		b.have[material] = ct
	}
	fmt.Println("have", b.have)

	for material, ct := range newRobots {
		existing := b.robotsOwned[material]
		existing += ct
		b.robotsOwned[material] = existing
	}
	fmt.Println("robots owned", b.robotsOwned)
}

func (b *Blueprint) key(min int) string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("%d-%d", b.number, min))
	for m, ct := range b.have {
		sb.WriteString(fmt.Sprintf("have-%s-%d", m, ct))
	}
	for m, ct := range b.robotsOwned {
		sb.WriteString(fmt.Sprintf("robots-owned-%s-%d", m, ct))
	}

	return sb.String()
}

func (b *Blueprint) twoAway(m string, log bool) bool {
	r := b.robots[m]

	fmt.Println("checking two away for ", m)
	for _, cost := range r.costs {
		targetDiff := cost.cost - b.have[cost.material]
		owned := b.robotsOwned[cost.material]

		r2 := b.robots[cost.material]

		if log {
			fmt.Println(cost, b.have)
		}

		// what to change here to make it work when ore is cheap enough?
		if cost.material != "ore" && !r2.canAfford(b.have) {
			if log {
				fmt.Println("cannot afford", cost.material, r2.canAfford(b.have))
			}
			return false
		}

		if targetDiff < 0 || owned <= 0 {
			if log {
				fmt.Println("incorrect numbers", cost.material, targetDiff, owned)
			}
			return false
		}

		if targetDiff/owned > 2 {
			if log {
				fmt.Println("math", targetDiff, owned)
			}
			return false
		}
	}
	return true
}

func (b *Blueprint) shouldSpend(log bool) (bool, string) {
	// duh!
	if b.robots["geode"].canAfford(b.have) {
		return true, "geode"
	}
	if b.twoAway("geode", log) {
		fmt.Println("two away from geode, don't spend")
		return false, ""
	}

	if b.robots["obsidian"].canAfford(b.have) {
		return true, "obsidian"
	}

	if b.twoAway("obsidian", log) {
		fmt.Println("two away from obsidian, don't spend")
		return false, ""
	}

	// TODO this may need speial logic at the end...
	if b.robots["clay"].canAfford(b.have) {
		return true, "clay"
	}

	if b.robots["ore"].canAfford(b.have) {
		return true, "ore"
	}

	return false, ""
}

type Robot struct {
	costs    []RobotCost
	material string
}

func (r Robot) canAfford(have map[string]int) bool {
	for _, cost := range r.costs {
		v := have[cost.material]
		// fmt.Println(v, cost.cost)
		if v < cost.cost {
			return false
		}
	}
	return true
}

// assumes canAfford == true
func (r Robot) spend(have map[string]int) {
	for _, cost := range r.costs {
		have[cost.material] -= cost.cost
	}
}

type RobotCost struct {
	cost     int
	material string
}

func day19() {
	lines := readFile("day19input")
	// r := regexp.MustCompile(`Blueprint (.+): Each ore robot costs (.+) ore. Each clay robot costs (.+) ore. Each obsidian robot costs (.+) ore and (.+) clay. Each geode robot costs (.+) ore and (.+) obsidian.`)
	r := regexp.MustCompile(`Each (.+) robot costs (.+)`)
	for _, line := range lines {
		// match := r.FindStringSubmatch(line)
		parts := splitLength(line, ": ", 2)
		number := atoi(splitLength(parts[0], " ", 2)[1])
		costs := strings.Split(parts[1], ".")

		b := Blueprint{
			number: number,
			robots: map[string]Robot{},
			have:   map[string]int{},
			robotsOwned: map[string]int{
				"ore": 1,
			},
		}
		// fmt.Println(costs, len(costs))
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
		// b.print()
		geode := runBluePrint(&b, 1, cache)
		// fmt.Println(len(cache))
		fmt.Println("answer", geode)
	}
}

// eventually need
func runBluePrint(b *Blueprint, start int, cache map[string]int) int {
	key := b.key(start)
	v, ok := cache[key]
	if ok {
		// fmt.Println("cache hit", v)
		return v
	}

	materials := []string{
		"geode",
		"obsidian",
		"clay",
		"ore",
	}

	/// TODO 24
	for i := start; i <= 24; i++ {

		fmt.Printf("\nminute %d\n", i)
		// b.print()
		newRobots := map[string]int{}

		// spend, material := b.shouldSpend(false)

		skip := make(map[string]bool)
		// for _, m := range materials {
		// 	valid[m] = true
		// }

		// go in order,
		// flag some as invalid for this round if need me

		// lol somehow got more with 13 which is invalid
		for _, material := range materials {
			r := b.robots[material]

			if skip[material] {
				continue
			}

			if (material == "obsidian" || material == "geode") && b.twoAway(material, false) {
				for _, cost := range r.costs {
					skip[cost.material] = true
				}
				continue
			}

			if r.canAfford(b.have) {
				fmt.Println("spend on", material)
				robot := b.robots[material]
				robot.spend(b.have)
				ct := newRobots[robot.material]
				// fmt.Println()
				newRobots[robot.material] = ct + 1
			}
		}

		// if spend {
		// 	if material == "" {
		// 		panic("invald spend")
		// 	}
		// 	// b2 = b.clone()
		// 	fmt.Println("spend on", material)
		// 	robot := b.robots[material]
		// 	robot.spend(b.have)
		// 	ct := newRobots[robot.material]
		// 	// fmt.Println()
		// 	newRobots[robot.material] = ct + 1
		// 	// spendCt = runBluePrint(b2, i+1, cache)

		// 	// doesn't
		// 	// b = b2
		// 	// } else {
		// 	// 	fmt.Println("else path")
		// 	// 	b3 = b.clone()
		// 	// 	b3.update(nil)
		// 	// 	noSpendCt = runBluePrint(b3, i+1, cache)
		// }
		b.update(newRobots)

	}
	// return b.have["geode"]
	cache[key] = b.have["geode"]
	return cache[key]
}
