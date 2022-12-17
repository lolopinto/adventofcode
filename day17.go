package main

import (
	"fmt"
	"os"
	"strings"
)

func day17() {

	//
	//	The tall, vertical chamber is exactly seven units wide. Each rock appears so that its left edge is two units away from the left wall and its bottom edge is three units above the highest rock in the room (or the floor, if there isn't one).

	// rock massaged
	rocks := []string{
		"####",
		".#.\n###\n.#.",
		"..#\n..#\n###",
		"#\n#\n#\n#",
		"##\n##",
	}

	// the caching is the combo the rocks can move within the 7 spaes

	// this is the building of the rockmap and what happens
	rockmap := map[int][]string{}
	emptyline := "......."

	threelines := []string{
		emptyline,
		emptyline,
		emptyline,
	}
	for idx, rock := range rocks {
		rockLines := strings.Split(rock, "\n")
		// fmt.Println(rockLines)

		l := 0
		for _, line := range rockLines {
			if len(line) > 0 {
				l = len(line)
			}
		}

		current := []string{}
		r := 7 - l - 2
		for _, line := range rockLines {
			line = ".." + line + strings.Repeat(".", r)
			if len(line) != 7 {
				panic(fmt.Errorf("invalid line %s of length %d", line, len(line)))
			}
			current = append(current, line)
		}
		current = append(current, threelines...)

		rockmap[idx] = current
	}

	jets := readFile("day17input")[0]
	currjet := 0

	// running this manually is going to be too slow
	// what's the caching here?
	// dfs?
	// let's do it manually first and then come back

	// rock stopp

	// not doing this yet and will change it if it becomes too slow
	// let's just do regular string lists for now
	// tower is represented in reverse order to make it easier to reason above
	// floor := "xxxxxxx"
	// tower := [][]string{{floor}}

	floor := "xxxxxxx"
	tower := []string{floor}

	printTower := func() {
		fmt.Println("printing tower")
		for _, line := range tower {
			fmt.Println(line)
		}
	}

	cache := map[string]bool{}

	makeCacheKey := func(towerslice, rock []string) string {
		return fmt.Sprintf("%s|%s", strings.Join(towerslice, ""), strings.Join(rock, ""))
	}

	canFillSlice := func(towerslice, rock []string) bool {
		key := makeCacheKey(towerslice, rock)
		v, ok := cache[key]
		if ok {
			return v
		}
		topidx := len(towerslice) - 1
		canfill := true
		for j := len(rock) - 1; j >= 0; j-- {
			line := rock[j]
			if topidx < 0 {
				break
				// done
			}
			topline := towerslice[topidx]
			lineworks := true
			for idx, c := range line {
				if c != '.' && topline[idx] != '.' {
					// fmt.Println("breaking for line", line, string(topline[idx]), idx)
					lineworks = false
					break
				}
			}
			if !lineworks {
				canfill = false
				break
			}
			topidx--
		}
		cache[key] = canfill
		return canfill
	}

	ct := 0
	for {
		idx := ct % 5
		// rock begins falling
		rock := rockmap[idx][:]
		fmt.Println("new rock", rock)
		// rock should be added to tower immediately
		// and then we manipulate rock within tower by keeping track of index and @@@

		jettime := true
		for {

			if jettime {
				// get direction
				dir := jets[currjet%len(jets)]

				// move right or left
				if dir == '>' {
					canmove := true
					for _, line := range rock {
						if line[6] != '.' {
							canmove = false
							break
						}
					}
					if canmove {
						for idx, line := range rock {
							if line == emptyline {
								break
							}
							if line[6] == '.' {
								rock[idx] = "." + line[:6]
							}
						}
						fmt.Println(rock, "after moving right")
					} else {
						fmt.Println("skipped moving right")
					}

				} else {
					canmove := true
					for _, line := range rock {
						if line[0] != '.' {
							canmove = false
							break
						}
					}
					// space to move left
					if canmove {
						for idx, line := range rock {
							if line == emptyline {
								break
							}
							if line[0] == '.' {
								rock[idx] = line[1:] + "."
							}
						}
						fmt.Println(rock, "after moving left")
					} else {
						fmt.Println("skipped moving left")
					}

				}
				currjet++
				jettime = false
				continue
			}

			// jet next time
			jettime = true

			// can fall
			if rock[len(rock)-1] == emptyline {
				// remove last
				rock = remove(rock, len(rock)-1)
				fmt.Println(rock, "rock falls")
			} else {

				// this is broken. what should be the representation?
				// if floor, nothing to do
				// if falling on top of something, need to do something

				rest := false
				// fell
				if tower[0] == floor {
					rest = true
				} else {
					// fmt.Println("len rock", len(rock), len(tower))
					sliceidx := len(tower)
					if len(rock) < sliceidx {
						sliceidx = len(rock)
					}
					// top := tower[0:sliceidx]
					// fmt.Println("top", top, len(top))

					canFill := false
					var sliceToFill []string
					var begSliceIdx int

					// THIS IS WRONG.
					// IT SHOULD FALL ONE UNIT AND NOT THIS
					// SO WE NEED TO KNOW WHERE THE ROCK IS IN WITHIN THE TOWER
					for j := len(tower) - 1; j >= len(rock); j-- {
						begSliceIdx = j - sliceidx
						towerslice := tower[j-sliceidx : j]
						fmt.Println("trying fill", begSliceIdx, j, towerslice, rock)
						slicefill := canFillSlice(towerslice, rock)
						if slicefill {
							sliceToFill = towerslice
							canFill = true
							break
						}

					}
					// check to see if everything here can fill as needed
					// this logic is broken...

					// if !canfill
					fmt.Println("can we rest?", !canFill)

					if !canFill {
						rest = true
					} else {

						if len(rock) == len(sliceToFill) {
							for i := 0; i < len(rock); i++ {
								towerIdx := begSliceIdx + i

								towerLine := tower[towerIdx]
								// replace every character we can
								for idx, c := range rock[i] {
									if c != '.' {
										towerLine = towerLine[:idx] + "#" + towerLine[idx+1:]
									}
								}
								// reassign back in tower
								tower[towerIdx] = towerLine
							}
							// THIS IS WRONG. IT SHOULD F

							fmt.Println("filled rock")
							printTower()
							os.Exit(0)

							ct++
							break
						} else {
							// len not equal. TODO handle
							printTower()
							fmt.Println(rock)
							fmt.Println("fill slice", sliceToFill, begSliceIdx)
							os.Exit(0)
						}

					}
					// check if it can rest next
					// fmt.Println("top", top)
					// fmt.Println("curr", rock)

					// for j := len(rock) - 1; j >= 0; j-- {
					// 	line := rock[j]
					// 	for idx, c := range line {
					// 		if c != '.' {

					// 		}
					// 	}
					// }

					// rest = true
				}

				if rest {
					tower = append(rock, tower...)
					ct++
					break
				}
			}

		}

		// rocks stopped
		// toDO 2022
		if ct == 10 {
			break
		}
	}

	printTower()

	fmt.Println(len(tower))
	// fmt.Println(tower)
}
