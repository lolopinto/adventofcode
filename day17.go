package main

import (
	"fmt"
	"os"
	"strings"
)

func day17() {
	emptyline := "......."

	moveLeft := func(s string) (string, bool) {

		start := strings.Index(s, "@")
		end := strings.LastIndex(s, "@")
		// fallenIdx := strings.Index(s, "#")

		// didn't find it. assume it doesn't interfere
		if start == -1 || end == -1 {
			return s, true
		}
		// left edge
		if start == 0 {
			return s, false
		}
		// interferring fallen rock, can't come through
		// if fallenIdx != -1 && fallenIdx > end {
		// 	return s, false
		// }
		if s[start-1] != '.' {
			return s, false
		}

		s = replaceInString(s, start-1, '@')
		s = replaceInString(s, end, '.')
		return s, true
	}

	moveRight := func(s string) (string, bool) {
		start := strings.Index(s, "@")
		end := strings.LastIndex(s, "@")
		// fallenIdx := strings.LastIndex(s, "#")

		// didn't find it. assume it doesn't interfere
		if start == -1 || end == -1 {
			return s, true
		}
		// right edge
		if end+1 == len(s) {
			return s, false
		}
		// interferring fallen rock, can't come through
		// if fallenIdx != -1 && fallenIdx < start {
		// 	// return s, false
		// }
		if s[end+1] != '.' {
			return s, false
		}

		s = replaceInString(s, end+1, '@')
		s = replaceInString(s, start, '.')
		return s, true
	}

	// fmt.Println(moveRight("..#.@.."))
	// return
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
	// emptyline := "......."

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

	ct := 0
	printTower := func(s string) {
		// if ct < 9 {
		// 	// return
		// }
		// if !strings.HasPrefix(s, "new rock") {
		// 	// return
		// }
		// fmt.Println("printing tower ", s)
		// for _, line := range tower {
		// 	fmt.Println(line)
		// }
	}

	//TODO no cache for now in case we don't need it
	// since there's now @ and #
	// cache := map[string]bool{}

	// makeCacheKey := func(towerslice, rock []string) string {
	// 	return fmt.Sprintf("%s|%s", strings.Join(towerslice, ""), strings.Join(rock, ""))
	// }

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

	for {
		idx := ct % 5
		// rock begins falling
		rock := make([]string, len(rockmap[idx]))
		copy(rock, rockmap[idx])

		for _, v := range rock {
			if v == floor {
				fmt.Printf("rock %d corrupted", ct)
				fmt.Println(rock)
				os.Exit(0)
			}
		}

		// new rock at beginning of tower
		tower = append(rock, tower...)
		printTower(fmt.Sprintf("new rock %d", ct))
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

					// can we fall one to the next line??
					towerslice := tower[rockStartIndex+1 : rockStartIndex+1+rockLength]
					rockslice := tower[rockStartIndex : rockStartIndex+rockLength]
					// for j := len(tower) - 1; j >= len(rock); j-- {
					// 	begSliceIdx = j - sliceidx
					// 	towerslice := tower[j-sliceidx : j]
					fmt.Println("trying fill", towerslice, rockslice)
					// this is broken too
					canFill := canFillSlice(towerslice, rockslice)

					// if !canfill
					fmt.Println("can we rest?", !canFill)

					if !canFill {
						rest = true
					} else {

						// TODO: this logic is broken, not replacing correctly...
						// duprock := rockslice[:]
						// this isn't working. needs to be done one at a time...
						// temp := map[int]string{}
						for j := rockLength - 1; j >= 0; j-- {
							// for i := 0; i < rockLength; i++ {
							// currRockIdx := rockStartIndex + i
							currLineIdx := rockStartIndex + j
							nextLineIdx := rockStartIndex + j + 1
							// rockLine := tower[currRockIdx]
							nextLine := tower[nextLineIdx]
							currLine := tower[currLineIdx]
							currTemp := currLine
							for i, v := range currLine {
								if v == '@' {
									nextLine = replaceInString(nextLine, i, '@')
									currTemp = replaceInString(currTemp, i, '.')
								}
							}
							tower[nextLineIdx] = nextLine
							// // for i,v := range
							// tower[nextLineIdx] = nextLine
							tower[currLineIdx] = currTemp
							// tower[currLineIdx]
						}
						// fmt.Println(temp)
						// os.Exit(0)
						// for k, v := range temp {
						// 	tower[k] = v
						// }
						// remove first line in tower
						// only remove first line in tower if no "#"

						top := tower[0][:]
						top = strings.ReplaceAll(top, "@", ".")
						if ct == 9 {
							fmt.Println("replace top", tower[0], top, top == emptyline)
							// os.Exit(1)
						}
						if top == emptyline {
							// if strings.Index(top, "#") == -1 {
							tower = remove(tower, 0)
							printTower("fancy falling. removed first line")
						} else {
							// replace top line
							tower[0] = top
							rockStartIndex++
							printTower("fancy falling, keep top line")
						}
						// os.Exit(0)

					}

				}

				if rest {
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
		// new rock 3 logic broken...
		// broken before new rock 4
		if ct == 2022 {
			break
		}
	}

	printTower("end")

	// no floor
	fmt.Println(len(tower) - 1)
}
