package grid

import (
	"log"
	"strconv"
)

func NewGrid(length int) *Grid {
	data := make([][]*Data, length)
	for i := 0; i < length; i++ {
		data[i] = make([]*Data, length)
		for j := 0; j < length; j++ {
			data[i][j] = &Data{}
		}
	}
	return &Grid{
		Length: length,
		data:   data,
	}
}

// copied from util for now
func atoi(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		log.Fatal(err)
	}
	return i
}

// int grid from a list of strings
func NewIntGrid(lines []string) *Grid {
	g := NewGrid(len(lines))
	for i := 0; i < g.Length; i++ {
		for j := 0; j < g.Length; j++ {
			g.data[i][j] = &Data{
				data: atoi(string(rune(lines[i][j]))),
			}
		}
	}
	return g
}

type Grid struct {
	Length int
	data   [][]*Data
}

func (g *Grid) At(r, c int) *Data {
	return g.data[r][c]
}

// returns up to 4 neighbors not including diagonals
func (g *Grid) Neighbors(r, c int) []Pos {
	var ret []Pos
	if r-1 >= 0 {
		ret = append(ret, Pos{Row: r - 1, Column: c})
	}
	if r+1 < g.Length {
		ret = append(ret, Pos{Row: r + 1, Column: c})
	}
	if c-1 >= 0 {
		ret = append(ret, Pos{Row: r, Column: c - 1})
	}
	if c+1 < g.Length {
		ret = append(ret, Pos{Row: r, Column: c + 1})
	}

	return ret
}

// returns up to 8 neighbors including diagonals
func (g *Grid) Neighbors8(r, c int) []Pos {
	var ret []Pos

	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			if i == 0 && j == 0 {
				continue
			}
			newR := r + i
			newC := c + j
			if newR >= 0 && newR < g.Length && newC >= 0 && newC < g.Length {
				ret = append(ret, Pos{Row: newR, Column: newC})
			}
		}
	}

	return ret
}

type Data struct {
	Visited bool
	data    interface{}
}

func (d *Data) Int() int {
	// returns 0 if not valid
	return d.data.(int)
}

func (d *Data) SetValue(val interface{}) {
	d.data = val
}

type Pos struct {
	Row, Column int
}
