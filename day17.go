package main

import (
	"fmt"
	"os"
	"strings"
)

// don't need this.
// was one of the crazy things I tried to see if it'll spped things up
// in this case, we're storing the Tower with a reversed list to see if it'll help
// simpler to store with a normal list.
type Tower struct {
	// stored and accessed in reverse order
	items []string
}

func (t *Tower) addRock(rock []string) {
	t.items = append(t.items, rock...)
}

func (t *Tower) len() int {
	return len(t.items)
}

func (t *Tower) height() int {
	// no floor
	return len(t.items) - 1
}

func (t *Tower) get(idx int) string {
	idx = t.len() - 1 - idx
	return t.items[idx]
}

func (t *Tower) set(idx int, value string) {
	idx = t.len() - 1 - idx
	t.items[idx] = value
}

func (t *Tower) removeAt(idx int) {
	idx = t.len() - 1 - idx
	t.items = remove(t.items, idx)
}

func (t *Tower) getSlice(idx, length int) []string {
	idx = t.len() - 1 - idx

	ret := make([]string, length)
	for i := 0; i < length; i++ {
		new_idx := idx - i
		ret[i] = t.items[new_idx]
	}
	return ret
}

func (t *Tower) print(s string) {
	// fmt.Println("printing tower ", s)
	// for j := len(t.items) - 1; j >= 0; j-- {
	// 	line := t.items[j]
	// 	fmt.Println(line)
	// }
}

func day17() {
	emptyline := "......."

	moveLeft := func(s string) (string, bool) {

		start := strings.Index(s, "@")
		end := strings.LastIndex(s, "@")

		// didn't find it. assume it doesn't interfere
		if start == -1 || end == -1 {
			return s, true
		}
		// left edge
		if start == 0 {
			return s, false
		}

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

		// didn't find it. assume it doesn't interfere
		if start == -1 || end == -1 {
			return s, true
		}
		// right edge
		if end+1 == len(s) {
			return s, false
		}
		if s[end+1] != '.' {
			return s, false
		}

		s = replaceInString(s, end+1, '@')
		s = replaceInString(s, start, '.')
		return s, true
	}

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

	threelines := []string{
		emptyline,
		emptyline,
		emptyline,
	}
	for idx, rock := range rocks {
		rockLines := strings.Split(rock, "\n")

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
			// flipped with tower being reversed
			current = append([]string{line}, current...)
		}
		// flipped with tower being reversed now
		current = append(threelines, current...)

		rockmap[idx] = current
	}

	jets := readFile("day17input")[0]
	currjet := 0

	floor := "-------"
	tower := &Tower{}
	tower.addRock([]string{floor})

	rockCount := 0
	printTower := func(s string) {
		tower.print(s)
	}

	canFillSlice := func(towerslice, rock []string) bool {
		topidx := len(towerslice) - 1
		canfill := true
		for j := len(rock) - 1; j >= 0; j-- {
			line := rock[j]
			if topidx < 0 {
				break
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
		return canfill
	}

	seen := map[[2]int][2]int{}

	for {
		rock_idx := rockCount % 5

		// rock begins falling
		rock := make([]string, len(rockmap[rock_idx]))
		copy(rock, rockmap[rock_idx])

		// new rock at beginning of tower
		tower.addRock(rock)
		printTower(fmt.Sprintf("new rock %d", rockCount))
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
						line := tower.get(rockStartIndex + i)
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
							tower.set(currIdx, replacements[i])
						}
						printTower("after moving right")
					} else {
						printTower("skipped moving right")
					}

				} else {
					canmove := true
					replacements := make([]string, rockLength)

					for i := 0; i < rockLength; i++ {
						line := tower.get(rockStartIndex + i)

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
							tower.set(currIdx, replacements[i])
						}

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
			if tower.get(lastRockIdx) == emptyline {
				// remove last
				tower.removeAt(lastRockIdx)
				rockLength--
				printTower("rock falls")
			} else {

				rest := false
				// fell
				if tower.get(0) == floor {
					rest = true
				} else {

					// can we fall one to the next line??
					towerslice := tower.getSlice(rockStartIndex+1, rockLength)
					// [rockStartIndex+1 : rockStartIndex+1+rockLength]
					// rockslice := tower[rockStartIndex : rockStartIndex+rockLength]
					rockslice := tower.getSlice(rockStartIndex, rockLength)

					canFill := canFillSlice(towerslice, rockslice)

					// fmt.Println("can we rest?", !canFill)

					if !canFill {
						rest = true
					} else {

						for j := rockLength - 1; j >= 0; j-- {
							currLineIdx := rockStartIndex + j
							nextLineIdx := rockStartIndex + j + 1
							nextLine := tower.get(nextLineIdx)
							currLine := tower.get(currLineIdx)
							currTemp := currLine
							for i, v := range currLine {
								if v == '@' {
									nextLine = replaceInString(nextLine, i, '@')
									currTemp = replaceInString(currTemp, i, '.')
								}
							}
							tower.set(nextLineIdx, nextLine)
							// tower[nextLineIdx] = nextLine
							// tower[currLineIdx] = currTemp
							tower.set(currLineIdx, currTemp)
						}

						// only remove first line in tower if no "#"

						top := tower.get(0)
						top = strings.ReplaceAll(top, "@", ".")

						if top == emptyline {
							tower.removeAt(0)
							// tower = remove(tower, 0)
							printTower("fancy falling. removed first line")
						} else {
							// replace top line
							// tower[0] = top
							tower.set(0, top)
							rockStartIndex++
							printTower("fancy falling, keep top line")
						}
					}
				}

				if rest {
					for i := 0; i < rockLength; i++ {
						currIdx := rockStartIndex + i
						val := strings.ReplaceAll(tower.get(currIdx), "@", "#")
						tower.set(currIdx, val)
					}

					rockCount++
					printTower("resting rock")
					break
				}
			}
		}

		// rocks stopped
		// part 1 vs 2 here
		// part 2 1000000000000
		// part 1 2022
		if rockCount == 1000000000000 {
			break
		}

		if rockCount > 1000 {
			// key is rock | jet combos
			key := [2]int{rock_idx, currjet % len(jets)}
			v, ok := seen[key]
			if ok {

				prev_rock := v[0]
				height := v[1]

				period := rockCount - prev_rock

				// if we've seen this rock|height combo and it's the same modulo with the target, here we go
				if rockCount%period == 1000000000000%period {
					// fmt.Println(key, v, height, "period", period)
					fmt.Println("cycle detected")

					// find where cycle started
					cycleHeight := tower.height() - height
					rocksRemainining := 1000000000000 - rockCount
					cyclesRemaining := (rocksRemainining / period) + 1
					fmt.Println(height, cycleHeight, cyclesRemaining)
					fmt.Println("part 2 answer", height+(cycleHeight*cyclesRemaining))
					os.Exit(0)
				}
			} else {
				// keep track of how many rocks and height we've seen with this rock|height combo
				seen[key] = [2]int{rockCount, tower.height()}
			}
		}
	}

	printTower("end")

	fmt.Println(tower.height())
}
