package main

import (
	"fmt"
	"log"
	"strings"
)

func day22() {
	chunks := readFileChunks("day22input", 2)

	p1 := parsePlayer(chunks[0])
	p2 := parsePlayer(chunks[1])

	globalNum = 1
	g := &game{num: globalNum}
	winner := g.playGame(p1, p2)

	log.Println(winner.calcScore())
}

type game struct {
	player1Winner bool
	num           int
}

var globalNum int

func (g *game) playGame(p1, p2 *player) *player {
	round := 1
	for {
		//		log.Println("game", g.num, "round", round, p1.cards, p2.cards)

		if g.gameOver(p1, p2) {
			break
		}
		c1 := p1.pop()
		c2 := p2.pop()

		g.calcWinner(p1, p2, c1, c2)

		round++
	}
	if g.player1Winner {
		return p1
	}
	if len(p1.cards) == 0 {
		return p2
	}
	return p1
}

func (g *game) cloneAndPlay(p1, p2 *player, c1, c2 int) *player {
	p1cards := make([]int, c1)
	for i := 0; i < c1; i++ {
		p1cards[i] = p1.cards[i]
	}
	p2cards := make([]int, c2)
	for i := 0; i < c2; i++ {
		p2cards[i] = p2.cards[i]
	}

	globalNum++
	g2 := &game{num: globalNum}

	p1clone := &player{cards: p1cards}
	p2clone := &player{cards: p2cards}
	winner := g2.playGame(p1clone, p2clone)

	if winner == p1clone {
		return p1
	}
	return p2
}

func (g *game) calcWinner(p1, p2 *player, c1, c2 int) {
	if len(p1.cards) >= c1 && len(p2.cards) >= c2 {
		winner := g.cloneAndPlay(p1, p2, c1, c2)

		// win via clone, swap the lower number first
		if winner == p1 {
			winner.append(c1, c2)
		} else {
			winner.append(c2, c1)
		}
	} else if c1 > c2 {
		p1.append(c1, c2)
	} else {
		p2.append(c2, c1)
	}
}

func (g *game) gameOver(p1, p2 *player) bool {
	if p1.previouslySeen() || p2.previouslySeen() {
		g.player1Winner = true
		return true
	}
	return len(p1.cards) == 0 || len(p2.cards) == 0
}

type player struct {
	cards      []int
	prevRounds map[string]int
}

func (p *player) getKey() string {
	strs := make([]string, len(p.cards))
	for i, v := range p.cards {
		strs[i] = fmt.Sprintf("%v", v)
	}
	return strings.Join(strs, ",")
}

func (p *player) calcPreviousRound() {
	if p.prevRounds == nil {
		p.prevRounds = make(map[string]int)
	}
	key := p.getKey()
	val, ok := p.prevRounds[key]
	if !ok {
		p.prevRounds[key] = 1
	} else {
		p.prevRounds[key] = val + 1
	}
}

func (p *player) pop() int {
	v := p.cards[0]
	p.cards = p.cards[1:]
	return v
}

func (p *player) append(c1, c2 int) {
	p.cards = append(p.cards, c1, c2)
	p.calcPreviousRound()
}

func (p *player) previouslySeen() bool {
	key := p.getKey()
	return p.prevRounds[key] == 2
}

func (p *player) calcScore() int {
	i := len(p.cards)
	sum := 0
	for _, c := range p.cards {
		sum += (c * i)
		i--
	}
	return sum
}

func parsePlayer(lines []string) *player {
	var cards []int
	for i := 1; i < len(lines); i++ {
		line := lines[i]
		cards = append(cards, atoi(line))
	}
	return &player{cards: cards}
}
