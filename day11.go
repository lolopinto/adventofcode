package main

import (
	"fmt"
	"math/big"
	"regexp"
	"sort"
	"strings"
)

type monkey struct {
	x            int
	starting     []*big.Int
	op           string
	divsibleBy   int
	throwToTrue  int
	throwToFalse int
	ct           int
}

func (m *monkey) addCount(v int) {
	m.ct += v
}

func (m *monkey) evaluateOp(left *big.Int) *big.Int {
	op := splitLength(m.op, " = ", 2)[1]
	parts := strings.Split(op, " ")
	if len(parts) != 3 {
		panic(fmt.Sprintf("rhs of op %s is not as expected", op))
	}

	var right *big.Int
	if parts[2] == "old" {
		right = left
	} else {
		right = big.NewInt(atoi64(parts[2]))
	}
	switch parts[1] {

	case "*":
		return left.Mul(left, right)

	case "+":
		return left.Add(left, right)
	}

	panic(fmt.Sprintf("invalid op %s in %s", parts[1], m.op))
}

func day11() {
	chunks := readFileChunks("day11input", -1)

	m := map[int]*monkey{}

	monkeyRegex := regexp.MustCompile(`Monkey (\d+):`)
	testDivisible := regexp.MustCompile(`Test: divisible by (\d+)`)
	testTrue := regexp.MustCompile(`If true: throw to monkey (\d+)`)
	testFalse := regexp.MustCompile(`If false: throw to monkey (\d+)`)

	for _, chunk := range chunks {
		if len(chunk) != 6 {
			panic(fmt.Sprintf("invalid chunk %s", chunk))
		}
		monkeyMatch := monkeyRegex.FindStringSubmatch(chunk[0])

		if monkeyMatch == nil {
			panic(fmt.Sprintf("invalid monkey parsing %v", chunk[0]))
		}

		x := atoi(monkeyMatch[1])

		var starting []*big.Int
		startingStr := strings.Split(chunk[1], ":")[1]
		for _, v := range strings.Split(startingStr, ",") {
			starting = append(starting, big.NewInt(atoi64(strings.TrimSpace(v))))
		}

		opStr := strings.TrimSpace(strings.Split(chunk[2], ":")[1])

		divisbleMatch := testDivisible.FindStringSubmatch(strings.TrimSpace(chunk[3]))
		if divisbleMatch == nil {
			panic(fmt.Sprintf("invalid divisble by parsing %v", chunk[3]))
		}

		testTrueMatch := testTrue.FindStringSubmatch(strings.TrimSpace(chunk[4]))
		if testTrueMatch == nil {
			panic(fmt.Sprintf("invalid test true by parsing %v", chunk[4]))

		}

		testFalseMatch := testFalse.FindStringSubmatch(strings.TrimSpace(chunk[5]))
		if testFalseMatch == nil {
			panic(fmt.Sprintf("invalid test true by parsing %v", chunk[5]))
		}

		m[x] = &monkey{
			x:            x,
			starting:     starting,
			op:           opStr,
			divsibleBy:   atoi(divisbleMatch[1]),
			throwToTrue:  atoi(testTrueMatch[1]),
			throwToFalse: atoi(testFalseMatch[1]),
		}
	}

	for i := 1; i <= 10000; i++ {
		for i := 0; i < len(chunks); i++ {
			mon := m[i]
			for _, s := range mon.starting {
				mon.addCount(1)
				// how to evaluate the right number if we don't do the real math if this is "*"
				v := mon.evaluateOp(s)
				div := big.NewInt(0)
				div.Div(v, big.NewInt(3))
				// fmt.Println("v", v, div)
				// v = v / 3
				var mon2 *monkey
				mod := big.NewInt(0)
				// fmt.Println("large vvv", v, math.MaxInt64, v > math.MaxInt64)
				mod.Mod(div, big.NewInt(int64(mon.divsibleBy)))
				if mod.IsInt64() && mod.Int64() == 0 {
					// fmt.Println("mod 0")
					mon2 = m[mon.throwToTrue]
					if mon2 == nil {
						panic(fmt.Sprintf("could not find monkey %d", mon.throwToTrue))
					}
					// fmt.Printf("throwing %d to monkey %d. true path\n", v, mon.throwToTrue)
				} else {
					mon2 = m[mon.throwToFalse]
					if mon2 == nil {
						panic(fmt.Sprintf("could not find monkey %d", mon.throwToFalse))
					}
				}

				mon2.starting = append(mon2.starting, div)
			}
			mon.starting = []*big.Int{}
		}

	}

	var cts []int
	for _, v := range m {
		cts = append(cts, v.ct)
		// fmt.Println(v.x, v.starting)
	}
	fmt.Println(cts)
	sort.Ints(cts)
	l := len(cts)
	fmt.Println(cts[l-1] * cts[l-2])
}

// so now this is running and taking too much CPU time
