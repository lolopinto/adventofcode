package cube

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Cube struct {
	X, Y, Z int
}

func (c *Cube) Neighbors() []Cube {
	return []Cube{
		{c.X, c.Y + 1, c.Z},
		{c.X, c.Y - 1, c.Z},
		{c.X + 1, c.Y, c.Z},
		{c.X - 1, c.Y, c.Z},
		{c.X, c.Y, c.Z + 1},
		{c.X, c.Y, c.Z - 1},
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

func NewCSVListCube(l []string) ([]Cube, error) {
	ret := make([]Cube, len(l))
	for i, line := range l {
		parts := strings.Split(line, ",")
		if len(parts) != 3 {
			return nil, fmt.Errorf("invalid csv line %s", line)
		}

		ret[i] = Cube{
			X: atoi(parts[0]),
			Y: atoi(parts[1]),
			Z: atoi(parts[2]),
		}
	}
	return ret, nil
}

func NewCSVMapCube(l []string) (map[Cube]bool, error) {
	ret := make(map[Cube]bool, len(l))
	for _, line := range l {
		parts := strings.Split(line, ",")
		if len(parts) != 3 {
			return nil, fmt.Errorf("invalid csv line %s", line)
		}

		ret[Cube{
			X: atoi(parts[0]),
			Y: atoi(parts[1]),
			Z: atoi(parts[2]),
		}] = true
	}
	return ret, nil
}
