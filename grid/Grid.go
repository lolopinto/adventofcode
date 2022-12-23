package grid

import (
	"fmt"
	"log"
	"strconv"
	"strings"
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
		Length:  length,
		XLength: length,
		YLength: length,
		data:    data,
	}
}

func NewRectGrid(xLength, yLength int) *Grid {
	data := make([][]*Data, xLength)
	for i := 0; i < xLength; i++ {
		data[i] = make([]*Data, yLength)
		for j := 0; j < yLength; j++ {
			data[i][j] = &Data{}
		}
	}
	return &Grid{
		XLength: xLength,
		YLength: yLength,
		data:    data,
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

func NewRuneGrid(lines []string) *Grid {
	g := NewRectGrid(len(lines), len(lines[0]))
	for i := 0; i < g.XLength; i++ {
		for j := 0; j < g.YLength; j++ {
			g.data[i][j] = &Data{
				data: rune(lines[i][j]),
			}
		}
	}
	return g
}

type Grid struct {
	Length int

	XLength int
	YLength int
	data    [][]*Data
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
	if r+1 < g.XLength {
		ret = append(ret, Pos{Row: r + 1, Column: c})
	}
	if c-1 >= 0 {
		ret = append(ret, Pos{Row: r, Column: c - 1})
	}
	if c+1 < g.YLength {
		ret = append(ret, Pos{Row: r, Column: c + 1})
	}

	return ret
}

func (g *Grid) RightAndDownNeighbors(r, c int) []Pos {
	var ret []Pos
	if r+1 < g.XLength {
		ret = append(ret, Pos{Row: r + 1, Column: c})
	}

	if c+1 < g.YLength {
		ret = append(ret, Pos{Row: r, Column: c + 1})
	}

	return ret
}

func (g *Grid) Top(r, c int) []Pos {
	var ret []Pos
	for i := r - 1; i >= 0; i-- {
		ret = append(ret, Pos{Row: i, Column: c})
	}
	return ret
}

func (g *Grid) Bottom(r, c int) []Pos {
	var ret []Pos
	for i := r + 1; i < g.XLength; i++ {
		ret = append(ret, Pos{Row: i, Column: c})
	}
	return ret
}

func (g *Grid) Left(r, c int) []Pos {
	var ret []Pos
	for i := c - 1; i >= 0; i-- {
		ret = append(ret, Pos{Row: r, Column: i})
	}
	return ret
}

func (g *Grid) Right(r, c int) []Pos {
	var ret []Pos
	for i := c + 1; i < g.YLength; i++ {
		ret = append(ret, Pos{Row: r, Column: i})
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
			if newR >= 0 && newR < g.XLength && newC >= 0 && newC < g.YLength {
				ret = append(ret, Pos{Row: newR, Column: newC})
			}
		}
	}

	return ret
}

func (g *Grid) Print(fn func(val interface{}) string) {
	for i := 0; i < g.XLength; i++ {
		var sb strings.Builder
		for j := 0; j < g.YLength; j++ {
			val := g.data[i][j]
			sb.WriteString(fn(val.data))
		}
		fmt.Println(sb.String())
	}
}

type Data struct {
	Visited bool
	data    interface{}
}

func (d *Data) Int() int {
	// returns 0 if not valid
	return d.data.(int)
}

func (d *Data) Rune() rune {
	// returns 0 if not valid
	return d.data.(rune)
}

func (d *Data) Data() interface{} {
	return d.data
}

func (d *Data) SetValue(val interface{}) {
	d.data = val
}

func (d *Data) SetValueOnce(val interface{}) error {
	if d.data != nil {
		return fmt.Errorf("tried to set value %v when value already exists", val)
	}
	d.data = val
	return nil
}

type Pos struct {
	Row, Column int
}

func (p Pos) Add(p2 Pos) Pos {
	return NewPos(p.Row+p2.Row, p.Column+p2.Column)
}

func NewPos(r, c int) Pos {
	return Pos{Row: r, Column: c}
}

func (p *Pos) Line(p2 *Pos, fn func(pos *Pos)) {
	startr, startc, endr, endc := p.Row, p.Column, p2.Row, p2.Column

	dr := delta(p.Row, p2.Row)
	dc := delta(p.Column, p2.Column)

	if dr < 1 {
		startr = p2.Row
		endr = p.Row
	}
	if dc < 1 {
		startc = p2.Column
		endc = p.Column

	}

	for r := startr; r <= endr; r++ {
		for c := startc; c <= endc; c++ {
			fn(&Pos{r, c})
		}
	}
}

func delta(one, two int) int {
	if one < two {
		return 1
	}
	if one > two {
		return -1
	}
	return 0
}
