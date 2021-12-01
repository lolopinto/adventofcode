package main

import (
	"log"
)

type gridData map[[4]int]byte

type grid struct {
	data gridData
}

func initGridFromData(data []string) *grid {
	m := make(gridData)

	for y, line := range data {
		for x, c := range line {
			m[makeMapKey(x, y, 0, 0)] = byte(c)
		}
	}
	return initGrid(m)
}

func initGrid(m gridData) *grid {
	return &grid{
		data: m,
	}
}

func getRanges(m gridData, add int) [4][2]int {
	var xs []int
	var ys []int
	var zs []int
	var ws []int

	for val := range m {
		xs = append(xs, val[0])
		ys = append(ys, val[1])
		zs = append(zs, val[2])
		ws = append(ws, val[3])
	}

	// when 0
	// we go from -1 to 1 [so need to do -1 to 1] we need +2 because we do <
	// -1 -> 2
	// and then do that the more dimensions
	return [4][2]int{
		{min(xs) - add, max(xs) + 1 + add},
		{min(ys) - add, max(ys) + 1 + add},
		{min(zs) - add, max(zs) + 1 + add},
		{min(ws) - add, max(ws) + 1 + add},
	}
}

func makeMapKey(x, y, z, w int) [4]int {
	return [4]int{x, y, z, w}
}

func (g *grid) get(x, y, z, w int) byte {
	key := makeMapKey(x, y, z, w)
	val, ok := g.data[key]
	if ok {
		return val
	}
	return '.'
}

func (g *grid) countActiveNeighbors(x, y, z, w int) int {
	count := 0
	for l := -1; l < 2; l++ {
		for k := -1; k < 2; k++ {
			for i := -1; i < 2; i++ {
				for j := -1; j < 2; j++ {
					if k == 0 && i == 0 && j == 0 && l == 0 {
						continue
					}
					x2 := x + i
					y2 := y + j
					z2 := z + k
					w2 := w + l
					if g.get(x2, y2, z2, w2) == '#' {
						count++
					}
				}
			}
		}
	}
	return count
}

func (g *grid) print() {
	ranges := getRanges(g.data, 0)
	xrange := ranges[0]
	yrange := ranges[1]
	zrange := ranges[2]
	wrange := ranges[3]

	//	log.Println(zrange[0], zrange[1], wrange[0], wrange[1])
	for w := wrange[0]; w < wrange[1]; w++ {
		for z := zrange[0]; z < zrange[1]; z++ {
			//			log.Println(z, w)
			for y := yrange[0]; y < yrange[1]; y++ {
				s := ""
				for x := xrange[0]; x < xrange[1]; x++ {
					v, ok := g.data[[4]int{x, y, z, w}]
					if ok {
						s += string(v)
					}
				}
				log.Println(s)
			}
			log.Println()
		}
	}
}

func (g *grid) clone() *grid {
	data := make(gridData, len(g.data))
	for k, v := range g.data {
		data[k] = v
	}
	return initGrid(data)
}

func (g *grid) countActive() int {
	count := 0
	for _, v := range g.data {
		if v == '#' {
			count++
		}
	}

	return count
}

func day17() {
	data := readFile("day17input")

	g := initGridFromData(data)
	//	g.print()

	for i := 0; i < 6; i++ {
		g2 := make(gridData)

		ranges := getRanges(g.data, 1)
		xrange := ranges[0]
		yrange := ranges[1]
		zrange := ranges[2]
		wrange := ranges[3]

		for w := wrange[0]; w < wrange[1]; w++ {
			for z := zrange[0]; z < zrange[1]; z++ {
				for y := yrange[0]; y < yrange[1]; y++ {
					for x := xrange[0]; x < xrange[1]; x++ {
						var c byte
						c = '.'

						count := g.countActiveNeighbors(x, y, z, w)
						if g.get(x, y, z, w) == '#' {
							if count == 2 || count == 3 {
								c = '#'
							}
						} else {
							if count == 3 {
								c = '#'
							}
						}
						g2[makeMapKey(x, y, z, w)] = c
					}
				}
			}
		}
		ng := initGrid(g2)
		//	ng.print()

		g = ng.clone()
	}

	log.Println(g.countActive())
}
