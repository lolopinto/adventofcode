package main

import (
	"fmt"
	"unicode"
)

type monkey2 struct {
	name string
	num  *int
	deps []string
	op   string
}

func day21() {
	lines := readFile("day21input")

	monkeys := map[string]*monkey2{}
	for _, line := range lines {
		parts := splitLength(line, ": ", 2)
		name := parts[0]
		var num *int
		// num := 0
		deps := []string{}
		op := ""

		if unicode.IsDigit(rune(parts[1][0])) {
			v := atoi(parts[1])
			num = &v
		} else {
			// op = parts[1]
			parts := splitLength(parts[1], " ", 3)
			deps = []string{parts[0], parts[2]}
			op = parts[1]
		}
		monkeys[name] = &monkey2{
			name: name,
			num:  num,
			deps: deps,
			op:   op,
		}
		// fmt.Println(*monkeys[name])
	}
	fmt.Println(evalMonkeys(monkeys["root"], monkeys))
}

func evalMonkeys(curr *monkey2, monkeys map[string]*monkey2) int {
	// evaluated
	if curr.num != nil {
		fmt.Println(curr.name, *curr.num)
		return *curr.num
	}

	left := evalMonkeys(monkeys[curr.deps[0]], monkeys)
	right := evalMonkeys(monkeys[curr.deps[1]], monkeys)

	ans := 0
	switch curr.op {
	case "+":
		ans = left + right
	case "-":
		ans = left - right
	case "*":
		ans = left * right
	case "/":
		ans = left / right
	default:
		panic(fmt.Errorf("invalid op %s", curr.op))
	}

	curr.num = &ans
	return *curr.num
}
