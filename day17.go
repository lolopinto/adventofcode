package main

import (
	"fmt"
	"strings"
)

func day17() {

	// rock massaged
	rocks := []string{
		"@@@@",
		".@.\n@@@\n.@.",
		"..@\n..@\n@@@",
		"@\n@\n@\n@",
		"@@\n@@",
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

	floor := "-------"
	tower := []string{floor}

	printTower := func(s string) {
		fmt.Println("printing tower ", s)
		for _, line := range tower {
			fmt.Println(line)
		}
	}

	//TODO no cache for now in case we don't need it
	// since there's now @ and #
	// cache := map[string]bool{}

	// makeCacheKey := func(towerslice, rock []string) string {
	// 	return fmt.Sprintf("%s|%s", strings.Join(towerslice, ""), strings.Join(rock, ""))
	// }

	// TODO

	// @...... -> .@.....
	// #.@..#. -> #..@.#.
	moveRight := func(s string) (string, bool) {
		prefix := ""
		suffix := ""
		main := ""
		idx := strings.Index(s, "#")
		if idx != -1 {
			prefix = s[:idx]
		} else {
			idx = 0
		}
		idx2 := strings.LastIndex(s, "#")
		if idx2 != -1 {
			suffix = s[idx:]
		} else {
			idx2 = len(s) - 1
		}
		main = s[idx:idx2]
		if len(main) == 0 || main[len(main)-1] != '.' {
			return "", false
		}
		return prefix + "." + main[:len(main)-1] + suffix, true

		// found_dot := true
		// found_stable_rock := false
		// move_idx := -1
		// for j := len(s) - 1; j > 0; j-- {
		// 	c := rune(s[j])
		// 	if c == '.' && !found_dot {
		// 		found_dot = true
		// 	}
		// 	if c == '#' && !found_stable_rock {
		// 		found_stable_rock = true
		// 		// don't touch anything to the right of this
		// 		move_idx = j
		// 	}
		// 	if c == '@' {
		// 		if found_stable_rock || !found_dot {
		// 			return "", false
		// 		}
		// 		break
		// 	}
		// }

		// if move_idx == -1 {
		// 	return "." + s[:6], true
		// }

		// return "", true
		// // first := s[:move_idx]
		// // untouched := s[move_idx:]
		// // return "." +
		// // // everything to first dot . is untouched
		// // return s[:move_idx] + "." + s[move_idx:6], true
	}

	// TODO change this also to account for # not moving
	// the moveLeft and moveRight functions need to change
	moveLeft := func(s string) (string, bool) {

		prefix := ""
		suffix := ""
		main := ""
		idx := strings.Index(s, "#")
		if idx != -1 {
			prefix = s[:idx]
		} else {
			idx = 0
		}
		idx2 := strings.LastIndex(s, "#")
		if idx2 != -1 {
			suffix = s[idx:]
		} else {
			idx2 = len(s) - 1
		}
		main = s[idx:idx2]
		if len(main) == 0 || main[0] != '.' {
			return "", false
		}
		return prefix + "." + main[1:] + suffix, true
	}

	canFillSlice := func(towerslice, rock []string) bool {
		// // key := makeCacheKey(towerslice, rock)
		// // v, ok := cache[key]
		// if ok {
		// 	return v
		// }
		topidx := len(towerslice) - 1
		canfill := true
		for j := len(rock) - 1; j >= 0; j-- {
			line := rock[j]
			if topidx < 0 {
				break
				// done
			}
			topline := towerslice[topidx]
			if topline == floor {
				canfill = false
				break
			}
			lineworks := true
			for idx, c := range line {
				// if '.' fine, if '@', assume current rock
				if c == '@' && topline[idx] == '#' {
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
		// cache[key] = canfill
		return canfill
	}

	ct := 0
	for {
		idx := ct % 5
		// rock begins falling
		rock := rockmap[idx][:]

		// new rock at beginning of tower
		tower = append(rockmap[idx][:], tower...)
		printTower("new rock")
		rockStartIndex := 0
		rockLength := len(rock)

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
					replacements := make([]string, rockLength)
					for i := 0; i < rockLength; i++ {
						line := tower[rockStartIndex+i]
						// ensure the only thing right of '@' is '.'

						line2, canmove2 := moveRight(line)
						if !canmove2 {
							canmove = false
							break
						}
						replacements[i] = line2
					}
					if canmove {
						for i := 0; i < rockLength; i++ {
							currIdx := rockStartIndex + i
							tower[currIdx] = replacements[i]
						}
						printTower("after moving right")
					} else {
						printTower("skipped moving right")
					}

				} else {
					canmove := true
					replacements := make([]string, rockLength)

					for i := 0; i < rockLength; i++ {
						line := tower[rockStartIndex+i]
						// // ensure the only thing left of '@' is '.'
						// found_dot := true
						// found_stable_rock := false
						// for _, c := range line {
						// 	if c == '.' {
						// 		found_dot = true
						// 	}
						// 	if c == '#' {
						// 		found_stable_rock = true
						// 	}
						// 	if c == '@' {
						// 		if found_stable_rock || !found_dot {
						// 			canmove = false
						// 			break
						// 		}
						// 	}
						// }
						line2, canmove2 := moveLeft(line)
						if !canmove2 {
							canmove = false
							break
						}
						replacements[i] = line2
					}
					// space to move left
					if canmove {
						for i := 0; i < rockLength; i++ {
							currIdx := rockStartIndex + i
							tower[currIdx] = replacements[i]
						}
						// for i := 0; i < rockLength; i++ {
						// 	currIdx := rockStartIndex + i
						// 	line := tower[currIdx]
						// 	// TODO THIS NEEDS TO CHANGE BECAUSE ROCK CAN MOVE WITHIN OTHER ROCKS
						// 	if line == emptyline {
						// 		break
						// 	}
						// 	if line[0] == '.' {
						// 		tower[currIdx] = line[1:] + "."
						// 	}
						// }
						printTower("after moving left")
					} else {
						printTower("skipped moving left")
					}
				}
				currjet++
				jettime = false
				continue
			}

			// jet next time
			jettime = true

			// can fall
			lastRockIdx := rockStartIndex + rockLength - 1
			if tower[lastRockIdx] == emptyline {
				// remove last
				tower = remove(tower, lastRockIdx)
				rockLength--
				printTower("rock falls")
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
					// sliceidx := len(tower)
					// if rockLength < sliceidx {
					// 	sliceidx = rockLength
					// }
					// top := tower[0:sliceidx]
					// fmt.Println("top", top, len(top))

					// canFill := false
					// var sliceToFill []string
					// var begSliceIdx int

					// THIS IS WRONG.
					// IT SHOULD FALL ONE UNIT AND NOT THIS
					// SO WE NEED TO KNOW WHERE THE ROCK IS IN WITHIN THE TOWER

					// can we fall one to the next line??
					towerslice := tower[rockStartIndex+1 : rockStartIndex+1+rockLength]
					rockslice := tower[rockStartIndex : rockStartIndex+rockLength]
					// for j := len(tower) - 1; j >= len(rock); j-- {
					// 	begSliceIdx = j - sliceidx
					// 	towerslice := tower[j-sliceidx : j]
					fmt.Println("trying fill", towerslice, rockslice)
					// this is broken too
					canFill := canFillSlice(towerslice, rockslice)
					// 	if slicefill {
					// 		// sliceToFill = towerslice
					// 		canFill = true
					// 		break
					// 	}

					// }
					// check to see if everything here can fill as needed
					// this logic is broken...

					// if !canfill
					fmt.Println("can we rest?", !canFill)

					if !canFill {
						rest = true
					} else {

						duprock := rockslice[:]
						for i := 0; i < rockLength; i++ {
							// currRockIdx := rockStartIndex + i
							nextLineIdx := rockStartIndex + i + 1
							// rockLine := tower[currRockIdx]
							nextLine := tower[nextLineIdx]
							for i, v := range duprock[i] {
								if v == '@' {
									nextLine = replaceInString(nextLine, i, '@')
								}
							}
							tower[nextLineIdx] = nextLine
						}
						// remove first line in tower
						tower = remove(tower, 0)

						// if len(rock) == len(sliceToFill) {
						// 	// TODO
						// 	// for i := 0; i < len(rock); i++ {
						// 	// 	towerIdx := begSliceIdx + i

						// 	// 	towerLine := tower[towerIdx]
						// 	// 	// replace every character we can
						// 	// 	for idx, c := range rock[i] {
						// 	// 		if c != '.' {
						// 	// 			towerLine = replaceInString(towerLine, '#', idx)
						// 	// 		}
						// 	// 	}
						// 	// 	// reassign back in tower
						// 	// 	tower[towerIdx] = towerLine
						// 	// }
						// 	// THIS IS WRONG. IT SHOULD F

						// 	printTower("filled rock")
						// 	os.Exit(0)

						// 	ct++
						// 	break
						// } else {
						// 	// len not equal. TODO handle
						// printTower("TODO handle")
						// fmt.Println(rock)
						// fmt.Println("fill slice", sliceToFill, begSliceIdx)
						// os.Exit(0)
						// }

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
					// tower = append(rock, tower...)
					for i := 0; i < rockLength; i++ {
						currIdx := rockStartIndex + i
						tower[currIdx] = strings.ReplaceAll(tower[currIdx], "@", "#")
					}
					ct++
					printTower("resting rock")
					break
				}
			}

		}

		// rocks stopped
		// toDO 2022
		if ct == 2 {
			break
		}
	}

	printTower("end")

	fmt.Println(len(tower))
	// fmt.Println(tower)
}
