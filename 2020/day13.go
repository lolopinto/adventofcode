package main

import (
	"fmt"
	"log"
	"math/big"
	"strings"
)

type datapoint struct {
	id  int
	mod int
}

var one = big.NewInt(1)

// from
// https://rosettacode.org/wiki/Chinese_remainder_theorem#Go
func crt(a, n []*big.Int) (*big.Int, error) {
	p := new(big.Int).Set(n[0])
	for _, n1 := range n[1:] {
		p.Mul(p, n1)
	}
	var x, q, s, z big.Int
	for i, n1 := range n {
		q.Div(p, n1)
		z.GCD(nil, &s, n1, &q)
		if z.Cmp(one) != 0 {
			return nil, fmt.Errorf("%d not coprime", n1)
		}
		x.Add(&x, s.Mul(a[i], s.Mul(&s, &q)))
	}
	return x.Mod(&x, p), nil
}

func day13() {
	lines := readFile("day13input")

	parts := strings.Split(lines[1], ",")
	var data []datapoint
	var a []*big.Int
	var n []*big.Int
	for i, part := range parts {
		if part == "x" {
			continue
		}

		data = append(data, datapoint{
			id:  atoi(part),
			mod: i,
		})
		num := atoi(part)
		a = append(a, big.NewInt(int64(num-i)))
		n = append(n, big.NewInt(int64(num)))
	}

	// previous solution didn't work so have to use math
	// apparently, there's a CRT solution
	// crt solution
	log.Println(crt(a, n))

	// solution:	825305207525452
	// and math below...

	minValue := 0
	runningProduct := 1
	for _, v := range data {
		for (minValue+v.mod)%v.id != 0 {
			minValue += runningProduct
		}
		runningProduct *= v.id
	}
	log.Println(minValue)

	// and another explanation from @astonm
	// https://github.com/astonm/advent-of-code-2020/blob/main/day13/code.py#L35-L50
}
