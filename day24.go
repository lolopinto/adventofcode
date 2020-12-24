package main

import (
	"log"
)

type point2 struct{ x, y int }

func day24() {
	lines := readFile("day24input")

	coords := make(map[point2]string)
	for _, line := range lines {

		pt := point2{}
		for i := 0; i < len(line); i++ {
			c := rune(line[i])

			if c == 'e' {
				pt.x++
				//				log.Println(string(c), pt)
			} else if c == 'w' {
				pt.x--
				//				log.Println(string(c), pt)
			} else {
				dir := line[i : i+2]
				switch dir {
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
				//				log.Println(dir, pt)
				i++
			}
		}
		v, ok := coords[pt]
		//		log.Println(pt)
		if !ok {
			//			log.Println("black")
			coords[pt] = "b"
		} else if v == "w" {
			//			log.Println("flipped back")
			coords[pt] = "b"
		} else {
			//			log.Println("flipped back again")
			coords[pt] = "w"
		}

		//		spew.Dump(pt)
		//		spew.Dump(directions)
	}

	//	spew.Dump(coords)
	count := 0
	for _, v := range coords {
		if v == "b" {
			count++
		}
	}
	log.Println(count)
}
