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

	player1 := &player{space: p1}
	player2 := &player{space: p2}
	current := player1
	currentP1 := true
	d := &deterministicDie{}
	for {

		done := current.move(d)

		if currentP1 {
			player1 = current
			//			p1 = current
			current = player2
			currentP1 = false
		} else {
			player2 = current
			current = player1
			currentP1 = true
		}
		if done {
			fmt.Println("done")
			break
		}
	}
	if player1.score < player2.score {
		fmt.Println(player1.score * d.count)
	} else {
		fmt.Println(player2.score * d.count)
	}
}

type player struct {
	score int
	space int
}

func (p *player) move(d *deterministicDie) bool {
	sum := 0
	val := []int{}
	for i := 0; i < 3; i++ {
		v := d.roll()
		val = append(val, v)
		sum += v
	}
	p.space += sum
	if p.space > 10 {
		p.space = p.space % 10
		if p.space == 0 {
			p.space = 10
		}
	}
	p.score += p.space
	//	fmt.Println(val, sum, "space", p.space, "score", p.score)

	return p.score >= 1000
}

type deterministicDie struct {
	pos   int
	count int
}

func (d *deterministicDie) roll() int {
	d.count++
	ret := d.pos + 1
	if ret == 100 {
		d.pos = 0
	} else {
		d.pos = ret
	}
	return ret
}
