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

func cloneMonkeys(m map[string]*monkey2) map[string]*monkey2 {
	ret := copyMap(m)

	for k, v := range ret {
		ret[k] = &monkey2{
			name: v.name,
			num:  v.num,
			deps: v.deps,
			op:   v.op,
		}
	}
	return ret
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
	// part 1

	root := monkeys["root"]
	dupMonkeys := cloneMonkeys(monkeys)

	ans := evalMonkeys(root, monkeys, false)
	fmt.Println("part 1: ", ans)

	leftSide := monkeys[root.deps[0]]
	rightSide := monkeys[root.deps[1]]

	left := evalMonkeys(leftSide, monkeys, false)
	right := evalMonkeys(rightSide, monkeys, false)
	leftDeps := findAllDeps(leftSide, monkeys)
	rightDeps := findAllDeps(rightSide, monkeys)

	fmt.Println(left, right)
	fmt.Println(leftDeps["humn"], rightDeps["humn"])
	currHumn := *monkeys["humn"].num
	fmt.Println("currHumn", currHumn)
	// it's in left in sample and mine and less both times so assume that for now and don't make it generic

	less := true
	for {
		next := 0
		if less {
			next = currHumn * 2
		} else {
			next = currHumn - 1
		}
		// fmt.Println("trying", next)

		newMonkeys := cloneMonkeys(dupMonkeys)
		newMonkeys["humn"].num = &next
		// new root and left
		root := newMonkeys["root"]
		// leftSide := newMonkeys[root.deps[0]]

		v := evalMonkeys(root, newMonkeys, true)
		if v == 0 {
			fmt.Println("part 2", next)
			break
		}
		// fmt.Println("new left", left)
		// break

		if v > 0 {
			less = false
		}
		currHumn = next
	}

	// start :=
	// monkeys["humn"]
	// // binary search from 4 -> 301 and see if two halfs of root are the same
}

func findAllDeps(curr *monkey2, monkeys map[string]*monkey2) map[string]bool {
	m := make(map[string]bool)
	for _, dep := range curr.deps {
		m[dep] = true
		child := findAllDeps(monkeys[dep], monkeys)
		for k, v := range child {
			m[k] = v
		}
	}
	return m
}

func evalMonkeys(curr *monkey2, monkeys map[string]*monkey2, evalRootAsEq bool) int {
	// evaluated
	if curr.num != nil {
		// fmt.Println(curr.name, *curr.num)
		return *curr.num
	}

	left := evalMonkeys(monkeys[curr.deps[0]], monkeys, evalRootAsEq)
	right := evalMonkeys(monkeys[curr.deps[1]], monkeys, evalRootAsEq)

	// eval root and if we get back 0. it's correct
	if curr.name == "root" && evalRootAsEq {
		if left == right {
			return 0
		}
		if left < right {
			return -1
		}
		return 1
	}

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
