package main

import "log"

type point2 struct{ x, y int }

func (p *point2) clone() point2 {
	return point2{p.x, p.y}
}

func applyDir(pt *point2, dir string) {
	switch dir {
	case "e":
		pt.x++
		break
	case "w":
		pt.x--
		break
	case "ne":
		pt.x++
		pt.y++
		break
	case "se":
		pt.y--
		//					pt.x++
		break
	case "sw":
		pt.y--
		pt.x--
		break
	case "nw":
		pt.y++
		//					pt.x--
		break
	}
}

func flipAll(lines []string, coords map[point2]int) map[point2]int {
	//	coords := make(map[point2]int)
	for _, line := range lines {

		pt := point2{}
		for i := 0; i < len(line); i++ {
			c := rune(line[i])

			var dir string
			if c == 'e' || c == 'w' {
				dir = string(c)
			} else {
				dir = line[i : i+2]
				i++
			}
			applyDir(&pt, dir)
		}
		v, ok := coords[pt]
		//		log.Println(pt)
		if !ok {
			//			log.Println("black")
			coords[pt] = 1
		} else {
			coords[pt] = v + 1
		}
	}
	return coords
}

var directions = []string{"e", "w", "sw", "se", "ne", "nw"}

func eachDayChange(coords map[point2]int) map[point2]int {
	coords2 := make(map[point2]int)
	adjBlack := make(map[point2]int)

	// all done simultaneously
	for pt, v := range coords {
		// only do this for black...
		if v%2 == 0 {
			continue
		}
		for _, dir := range directions {
			pt2 := pt.clone()
			applyDir(&pt2, dir)
			v2, ok := adjBlack[pt2]
			if !ok {
				v2 = 0 //white
			}
			adjBlack[pt2] = v2 + 1
		}
	}

	for pt, v := range coords {
		if v%2 == 1 {
			ctBlack, ok := adjBlack[pt]
			if !ok {
				ctBlack = 0
			}
			if ctBlack == 0 || ctBlack > 2 {
			} else {
				coords2[pt] = 1
			}
		}
	}

	for pt, v := range adjBlack {
		if v == 2 {
			ctBlack, ok := coords[pt]
			if !ok {
				ctBlack = 0
			}
			if ctBlack%2 == 0 {
				coords2[pt] = 1
			}
		}
	}
	return coords2
}

//func cloneCoords()
func day24() {
	lines := readFile("day24input")

	base := make(map[point2]int)
	base = flipAll(lines, base)

	for i := 1; i <= 100; i++ {
		coords2 := eachDayChange(base)

		// 	v, ok := coords[pt2]
		// 	if !ok {
		// 		v = 0 //white
		// 	}
		// 	if v%2 == 1 {
		// 		ctBlack++
		// 	} else {
		// 		//					ctWhite++
		// 	}
		// }

		// black
		// 	if v%2 == 1 {
		// 		if ctBlack == 0 || ctBlack > 2 {
		// 			// flip to white
		// 			//coords2[pt] = v - 1
		// 		} else {
		// 			coords2[pt] = 1
		// 		}
		// 	} else {
		// 		if ctBlack == 2 {
		// 			coords2[pt] = v + 1
		// 		} else {
		// 			coords2[pt] = v
		// 		}
		// 	}

		// 	for k,v := range adjBlack {
		// 		if v %2 == 0 &&
		// 	}
		// }

		count := 0
		for _, v := range coords2 {
			if v%2 == 1 {
				count++
			}
		}
		base = coords2
		log.Println("day", i, "count", count)
	}

	// day 1

}
