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

func direction(dir byte, num int, p *pt, wp *pt) {
	switch dir {
	case 'F':
		p.x += wp.x * num
		p.y += wp.y * num
		break
	case 'N':
		wp.y += num
		break
	case 'S':
		wp.y -= num
		break
	case 'E':
		wp.x += num
		break
	case 'W':
		wp.x -= num
		break
	case 'L':
		var tmp int
		tmp = wp.x
		wp.x = -wp.y
		wp.y = tmp
		if num-90 != 0 {
			direction('L', num-90, p, wp)
		}
		return
	case 'R':
		// moving right...
		var tmp int
		//
		tmp = wp.x
		wp.x = wp.y
		wp.y = -tmp

		if num-90 != 0 {
			direction('R', num-90, p, wp)
		}
		return
	}
}

func day12() {
	lines := readFile("day12input")

	p := &pt{curDir: 'E'}
	wp := &pt{curDir: 'E', x: 10, y: 1}
	for _, line := range lines {
		dir := line[0]
		num := atoi(line[1:])

		direction(dir, num, p, wp)
		log.Println(dir, num, wp, p)

	}
	spew.Dump(math.Abs(float64(p.x)) + math.Abs(float64(p.y)))
}
