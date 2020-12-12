package main

import (
	"log"
	"math"

	"github.com/davecgh/go-spew/spew"
)

type pt struct {
	x      int
	y      int
	curDir byte
}

func direction(dir byte, num int, p *pt) {
	switch dir {
	case 'F':
		direction(p.curDir, num, p)
		break
	case 'N':
		p.y += num
		break
	case 'S':
		p.y -= num
		break
	case 'E':
		p.x += num
		break
	case 'W':
		p.x -= num
		break
	case 'L':
		switch p.curDir {
		case 'N':
			p.curDir = 'W'
			break
		case 'S':
			p.curDir = 'E'
			break
		case 'E':
			p.curDir = 'N'
			break
		case 'W':
			p.curDir = 'S'
			break
		}
		if num-90 != 0 {
			direction('L', num-90, p)
		}
		return
	case 'R':
		switch p.curDir {
		case 'N':
			p.curDir = 'E'
			break
		case 'S':
			p.curDir = 'W'
			break
		case 'E':
			p.curDir = 'S'
			break
		case 'W':
			p.curDir = 'N'
			break
		}
		if num-90 != 0 {
			direction('R', num-90, p)
		}
		return
	}
}

func day12() {
	lines := readFile("day12input")

	p := &pt{curDir: 'E'}
	for _, line := range lines {
		dir := line[0]
		num := atoi(line[1:])

		direction(dir, num, p)
		log.Println(dir, num, p)

	}
	spew.Dump(math.Abs(float64(p.x)) + math.Abs(float64(p.y)))
}
