package main

import (
	"fmt"
	"regexp"
)

func day21() {
	lines := readFile("day21input")

	r := regexp.MustCompile(`Player (\d+) starting position: (\d+)`)
	p1 := atoi(r.FindStringSubmatch(lines[0])[2])
	p2 := atoi(r.FindStringSubmatch(lines[1])[2])

	p1wins, p2wins := playGame2(p1, p2, 0, 0, true)
	//	fmt.Println(p1wins, p2wins)
	if p1wins > p2wins {
		fmt.Println(p1wins)
	} else {
		fmt.Println(p2wins)
	}
}

func moveGame(space, score, sum int) (int, int, bool) {
	newspace := space + sum
	if newspace > 10 {
		newspace = newspace % 10
		if newspace == 0 {
			newspace = 10
		}
	}
	newscore := score + newspace
	return newspace, newscore, newscore >= 21
}

var cache map[[5]int][2]int

func playGame2(p1Space, p2Space, p1Score, p2Score int, currentP1 bool) (int, int) {
	var currentP int
	if currentP1 {
		currentP = 1
	}
	k := [5]int{p1Space, p2Space, p1Score, p2Score, currentP}
	if cache == nil {
		cache = make(map[[5]int][2]int)
	}
	v := cache[k]
	if v != [2]int{0, 0} {
		return v[0], v[1]
	}
	p1wins := 0
	p2wins := 0

	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			for k := 1; k <= 3; k++ {
				if currentP1 {
					newSpace, newScore, done := moveGame(p1Space, p1Score, i+j+k)
					if done {
						p1wins++
					} else {
						p1Clone1wins, p2Clonewins := playGame2(newSpace, p2Space, newScore, p2Score, false)

						p1wins += p1Clone1wins
						p2wins += p2Clonewins
					}
				} else {
					newSpace, newScore, done := moveGame(p2Space, p2Score, i+j+k)

					if done {
						p2wins++
					} else {
						p1Clone1wins, p2Clonewins := playGame2(p1Space, newSpace, p1Score, newScore, true)

						p1wins += p1Clone1wins
						p2wins += p2Clonewins
					}
				}
			}
		}
	}
	cache[k] = [2]int{p1wins, p2wins}
	return p1wins, p2wins
}
