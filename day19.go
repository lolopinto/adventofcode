package main

import (
	"fmt"
	"math"
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
	// fmt.Println("have", b.have)

	for material, ct := range newRobots {
		existing := b.robotsOwned[material]
		existing += ct
		b.robotsOwned[material] = existing
	}
	// fmt.Println("robots owned", b.robotsOwned)
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

func (b *Blueprint) shouldSpend() (bool, string) {
	// always go in this order for now
	// should always spend on geode if we need it
	materials := []string{
		"geode",
		"obsidian",
		"clay",
		// "ore",
	}

	for _, material := range materials {
		r := b.robots[material]
		// fmt.Println("spenddddd geode")
		if len(r.costs) == 0 {
			fmt.Println(material, b, r)
			panic("invalid robot")
		}
		if r.canAfford(b.have) {
			fmt.Println("afford", material)
			return true, material
		}

		// start with most expensive dependency
		// two_away := false

		for j := len(r.costs) - 1; j >= 0; j-- {
			cost := r.costs[j]
			r = b.robots[cost.material]
			if r.canAfford(b.have) {
				if material == "obsidian" || material == "geode" {
					targetDiff := cost.cost - b.have[cost.material]

					// fmt.Println(cost.material, b.robotsOwned[cost.material], targetDiff)
					// you can get  it in 1 or 2 minutes, just wait
					if targetDiff > 0 && b.robotsOwned[cost.material]/targetDiff < 2 {
						return false, ""
					}
				}
				// }
				// purchase this
				// fmt.Println("afford", material, cost.material, b.have[cost.material])
				return true, cost.material
			}
		}
		// if two_away {
		// 	return false, ""
		// }
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

	/// TODO 24
	for i := start; i <= 24; i++ {

		fmt.Printf("\nminute %d\n", i)
		// b.print()

		// spend, material := b.shouldSpend()

		// if !spend {
		// 	b.update(nil)
		// 	continue
		// }

		// var b2 *Blueprint
		// // var b3 *Blueprint
		// // spendCt := math.MinInt
		// // noSpendCt := math.MinInt

		// if spend {
		// 	if material == "" {
		// 		panic("invald spend")
		// 	}
		// 	b2 = b.clone()
		// 	fmt.Println("spend", material)
		// 	robot := b2.robots[material]
		// 	robot.spend(b2.have)
		// 	ct := newRobots[robot.material]
		// 	// fmt.Println()
		// 	newRobots[robot.material] = ct + 1
		// 	b2.update(newRobots)
		// 	// spendCt = runBluePrint(b2, i+1, cache)

		// 	// doesn't
		// 	b = b2
		// 	// } else {
		// 	// 	fmt.Println("else path")
		// 	// 	b3 = b.clone()
		// 	// 	b3.update(nil)
		// 	// 	noSpendCt = runBluePrint(b3, i+1, cache)
		// }

		// // fmt.Println(spendCt, noSpendCt)
		// if spendCt > noSpendCt {
		// 	// fmt.Println("assign b2")
		// 	b = b2
		// } else {
		// 	// not once???
		// 	// fmt.Println("assign b3")
		// 	b = b3
		// }

		max := math.MinInt
		var currMax *Blueprint

		for _, robot := range b.robots {
			newRobots := map[string]int{}

			if robot.canAfford(b.have) {
				b2 := b.clone()
				// fmt.Println("can afford", robot.material)
				// purchase
				robot.spend(b2.have)
				ct := newRobots[robot.material]
				newRobots[robot.material] = ct + 1
				b2.update(newRobots)

				geodeCt := runBluePrint(b2, i+1, cache)

				// fmt.Println(geodeCt)
				if geodeCt > max {
					currMax = b2
				}
			}
		}

		b2 := b.clone()

		// check the no new robot path
		b2.update(nil)
		geodeCt := runBluePrint(b2, i+1, cache)
		if geodeCt > max {
			currMax = b2
		}
		// fmt.Println(currMax)

		b = currMax
	}
	// return b.have["geode"]
	cache[key] = b.have["geode"]
	return cache[key]
}
