package main

import (
	"log"
)

type grid struct {
	//	data          []string
	// data          map[int][]string // other z-coordinates
	// data2         map[string]byte
	realdata map[[3]int]byte
	//	width, height int
	xrange, yrange, zrange [2]int
	//	yrange                 [2]int
}

func initGridFromData(data []string) *grid {
	m := make(map[[3]int]byte)

	for y, line := range data {
		for x, c := range line {
			m[makeMapKey(x, y, 0)] = byte(c)
		}
	}
	return initGrid(m)
	//	return &grid{realdata: m}
}

func initGrid(m map[[3]int]byte) *grid {
	// var xs []int
	// var ys []int
	// var zs []int

	// for val := range m {
	// 	xs = append(xs, val[0])
	// 	ys = append(ys, val[1])
	// 	zs = append(zs, val[2])
	// }

	// possible ranges...
	ranges := getRanges(m, 1)
	return &grid{
		realdata: m,
		xrange:   ranges[0],
		yrange:   ranges[1],
		zrange:   ranges[2],
	}
}

func getRanges(m map[[3]int]byte, add int) [3][2]int {
	var xs []int
	var ys []int
	var zs []int

	for val := range m {
		xs = append(xs, val[0])
		ys = append(ys, val[1])
		zs = append(zs, val[2])
	}

	return [3][2]int{
		{min(xs) - add, max(xs) + 1 + add},
		{min(ys) - add, max(ys) + 1 + add},
		{min(zs) - add, max(zs) + 1 + add},
	}
}

// func initGrid(data []string) *grid {
// 	width := len(data[0])
// 	height := len(data)
// 	// string -> byte map
// 	m := make(map[[3]int]byte)

// 	for x, line := range data {
// 		for y, c := range line {
// 			m[makeMapKey(x, y, 0)] = byte(c)
// 		}
// 	}
// 	// z -> strings
// 	//	m := make(map[int][]string)
// 	//	m[0] = data
// 	return &grid{
// 		//		data:      data,
// 		width:    width,
// 		height:   height,
// 		realdata: m,
// 		// data:   m,
// 		// data2:  m2,
// 	}
// }

func makeMapKey(x, y, z int) [3]int {
	return [3]int{x, y, z}
	//	return fmt.Sprintf("%d:%d:%d", x, y, z)
}

func (g *grid) get(x, y, z int) byte {
	key := makeMapKey(x, y, z)
	val, ok := g.realdata[key]
	if ok {
		return val
	}
	return '.'
	// data, ok := g.data[z]
	// if !ok {
	// 	// inactive
	// 	return '.'
	// 	//		data = g.data[0]
	// }

	// // // x = x + g.width
	// // // y = y + g.height

	// // //	return data[x%g.width][y%g.height]

	// if !g.validPos(x, y, data) {
	// 	// inactive
	// 	return '.'
	// }
	// return data[y][x]

	// }

	// data, ok = g.data[0]
	// if !ok {
	// 	panic("should have gotten data for 0")
	// }
	// return data[x%g.width][y%g.height]
}

func (g *grid) validPos(x, y int, grid []string) bool {
	return x >= 0 && x < len(grid[0]) && y >= 0 && y < len(grid)
}

func (g *grid) countActiveNeighbors(x, y, z int) int {
	count := 0
	//	for l :=
	// TODO active
	for k := -1; k < 2; k++ {
		for i := -1; i < 2; i++ {
			for j := -1; j < 2; j++ {

				if k == 0 && i == 0 && j == 0 {
					continue
				}
				x2 := x + i
				y2 := y + j
				z2 := z + k
				//				log.Println("neighbor", x2, y2, z2, g.get(x2, y2, z2))
				if g.get(x2, y2, z2) == '#' {
					count++
				}
			}
		}
	}
	return count
}

func (g *grid) print() {
	ranges := getRanges(g.realdata, 0)
	xrange := ranges[0]
	yrange := ranges[1]
	zrange := ranges[2]

	log.Println(zrange[0], zrange[1])
	for z := zrange[0]; z < zrange[1]; z++ {
		for y := yrange[0]; y < yrange[1]; y++ {
			s := ""
			for x := xrange[0]; x < xrange[1]; x++ {
				v, ok := g.realdata[[3]int{x, y, z}]
				if ok {
					s += string(v)
				}
			}
			log.Println(s)
		}
		log.Println()
	}
}

func (g *grid) clone() *grid {
	data := make(map[[3]int]byte, len(g.realdata))
	for k, v := range g.realdata {
		data[k] = v
	}
	return initGrid(data)
}

func (g *grid) countActive() int {
	count := 0
	for _, v := range g.realdata {
		if v == '#' {
			count++
		}
	}

	return count
}

func day17() {
	data := readFile("day17input")

	g := initGridFromData(data)
	g.print()

	// TODO change value of this...
	//	g2 := make(map[int][]string)
	//	g2:= make(map[3]intbyte)

	for i := 0; i < 6; i++ {
		g2 := make(map[[3]int]byte)

		for z := g.zrange[0]; z < g.zrange[1]; z++ {

			// TODO
			//		newGrid := make([]string, g.height+2)
			for y := g.yrange[0]; y < g.yrange[1]; y++ {
				//			var sb strings.Builder
				for x := g.xrange[0]; x < g.xrange[1]; x++ {
					var c byte
					c = '.'

					//				log.Println(x, y, z)
					count := g.countActiveNeighbors(x, y, z)
					//				log.Println("count", x, y, z, count)
					if g.get(x, y, z) == '#' {
						if count == 2 || count == 3 {
							c = '#'
							//						sb.WriteByte('#')
						} else {
							//						sb.WriteByte('.')
						}
					} else {
						if count == 3 {
							c = '#'
							//					sb.WriteByte('#')
						} else {
							//						sb.WriteByte('.')
						}
					}
					g2[makeMapKey(x, y, z)] = c

				}
				//			newGrid[y] = sb.String()
			}
			// g2
			// g2[z] = newGrid
			//		spew.Dump(z, newGrid)
		}
		//	spew.Dump(g2)

		ng := initGrid(g2)
		ng.print()

		g = ng.clone()
	}

	log.Println(g.countActive())
}
