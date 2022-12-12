package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

type monkey struct {
	x            int
	starting     []int64
	op           string
	divsibleBy   int
	throwToTrue  int
	throwToFalse int
	ct           int
}

func (m *monkey) addCount(v int) {
	m.ct += v
}

func (m *monkey) evaluateOp(left int64) int64 {
	op := splitLength(m.op, " = ", 2)[1]
	parts := strings.Split(op, " ")
	if len(parts) != 3 {
		panic(fmt.Sprintf("rhs of op %s is not as expected", op))
	}

	var right int64
	if parts[2] == "old" {
		right = left
	} else {
		right = atoi64(parts[2])
	}
	switch parts[1] {

	case "*":
		return left * right

	case "+":
		return left + right
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

	product := 1
	for _, chunk := range chunks {
		if len(chunk) != 6 {
			panic(fmt.Sprintf("invalid chunk %s", chunk))
		}
		monkeyMatch := monkeyRegex.FindStringSubmatch(chunk[0])

		if monkeyMatch == nil {
			panic(fmt.Sprintf("invalid monkey parsing %v", chunk[0]))
		}

		x := atoi(monkeyMatch[1])

		var starting []int64
		startingStr := strings.Split(chunk[1], ":")[1]
		for _, v := range strings.Split(startingStr, ",") {
			starting = append(starting, atoi64(strings.TrimSpace(v)))
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
		product *= atoi(divisbleMatch[1])
	}

	for i := 1; i <= 10000; i++ {
		for i := 0; i < len(chunks); i++ {
			mon := m[i]
			for _, s := range mon.starting {
				mon.addCount(1)
				// part 1
				// v := mon.evaluateOp(s) / 3
				// part 2
				v := mon.evaluateOp(s) % int64(product)

				var mon2 *monkey
				if v%int64(mon.divsibleBy) == 0 {
					mon2 = m[mon.throwToTrue]
					if mon2 == nil {
						panic(fmt.Sprintf("could not find monkey %d", mon.throwToTrue))
					}
				} else {
					mon2 = m[mon.throwToFalse]
					if mon2 == nil {
						panic(fmt.Sprintf("could not find monkey %d", mon.throwToFalse))
					}
				}

				mon2.starting = append(mon2.starting, v)
			}
			mon.starting = []int64{}
		}
	}

	var cts []int
	for _, v := range m {
		cts = append(cts, v.ct)
	}
	sort.Ints(cts)
	l := len(cts)
	fmt.Println(cts[l-1] * cts[l-2])
}
