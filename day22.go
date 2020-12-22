package main

import "github.com/davecgh/go-spew/spew"

func day22() {
	chunks := readFileChunks("day22input", 2)

	p1 := parsePlayer(chunks[0])
	p2 := parsePlayer(chunks[1])

	// spew.Dump(p1)
	// spew.Dump(p2)

	for {
		if len(p1.cards) == 0 || len(p2.cards) == 0 {
			break
		}

		c1 := p1.pop()
		c2 := p2.pop()

		if c1 > c2 {
			p1.append(c1, c2)
		} else {
			p2.append(c1, c2)
		}
	}

	spew.Dump(p1.calcScore())
	spew.Dump(p2.calcScore())
}

type player struct {
	cards []int
}

func (p *player) pop() int {
	v := p.cards[0]
	p.cards = p.cards[1:]
	return v
}

func (p *player) append(c1, c2 int) {
	if c1 > c2 {
		p.cards = append(p.cards, c1)
		p.cards = append(p.cards, c2)
	} else {
		p.cards = append(p.cards, c2)
		p.cards = append(p.cards, c1)
	}
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
