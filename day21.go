package main

import (
	"fmt"
	"unicode"
)

type monkey2 struct {
	name string
	num  *int64
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
		var num *int64
		deps := []string{}
		op := ""

		if unicode.IsDigit(rune(parts[1][0])) {
			v := atoi64(parts[1])
			num = &v
		} else {
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
	diff := left - right
	multiply := true
	fmt.Println("initial", left, right, diff)

	// less
	// less := left > right
	// which is greater
	// multiply := left < right

	// divide := left > right
	// addAfter := divide
	// subAfter := multiply
	// fmt.Println(multiply)
	// return

	// sample
	// left so multiply and add after

	// is it in left or right
	// it's in left in both inputs
	// but greater or not flipped
	fmt.Println(leftDeps["humn"], rightDeps["humn"])
	// return
	currHumn := *monkeys["humn"].num
	fmt.Println("currHumn", currHumn)
	// it's in left in sample and mine and less both times so assume that for now and don't make it generic

	// less := true
	// fake trying to get a sense
	//	btw 4277777777777 and 5277777777777
	// 3677777777777 1 => right bigger
	// 3777777777777 -1 => left bigger

	// 3715777777777 1
	// 3716777777777 -1

	// 371666666666 1
	// currHumn = 3715777777777

	// currHumn = 3715789999999
	// 2^9
	//	3715799999999

	currHumn = 3715791111110
	// 3715801111110
	// 2^8

	//	2^7
	// currHumn = 3715791111110
	//3715801111110

	// currHumn = 3715792222221
	// // 2^7 above

	// btw 3715792222221 and 3715802222221
	currHumn = 3715792222221

	// btw 3715793333332 and 3715803333332
	currHumn = 3715793333332

	// 3715794444443 and 3715804444443
	currHumn = 3715794444443

	// 3715795555554 and 3715805555554
	currHumn = 3715795555554

	// 3715796666665 and 3715806666665
	currHumn = 3715796666665

	// 3715797777776 and 3715807777776
	currHumn = 3715797777776

	// 3715798888887 and 3715808888887
	currHumn = 3715798888887

	// 3715798999998 and 3715799999998
	currHumn = 3715798999998

	// 3715799111109 and 3715800111109
	currHumn = 3715799111109

	// 3715799222220 and 3715800222220
	currHumn = 3715799222220

	// 3715799333331 and 3715800333331
	currHumn = 3715799333331

	// 3715799444442 and 3715800444442
	currHumn = 3715799444442

	// 3715799455553 and 3715799555553
	currHumn = 3715799455553

	// 3715799466664 and 3715799566664

	// eventually gave up and did some binary search by hand
	// and then eventually just ran it down between these numbers

	// could have coded it but had been making mistakes so just gave up lol
	for i := 3715799466664; i <= 3715799566664; i++ {
		// pw := math.Pow10(i)
		// currHumn = int64(pw) + (currHumn)
		currHumn = int64(i)

		if multiply {
			// currHumn = currHumn * 1_000_000_000
			// } else if divide {
			// 	currHumn = currHumn / 2
		} else {
			// if subAfter {
			// 	currHumn = currHumn - 1
			// }
			// if addAfter {
			// 	currHumn = currHumn + 1
			// }
		}

		// 	// currHumn = currHumn * 2
		// 	currHumn = currHumn / 2
		// } else {
		// 	currHumn = currHumn + 1
		// }
		// fmt.Println("trying", currHumn)

		newMonkeys := cloneMonkeys(dupMonkeys)

		newMonkeys["humn"].num = &currHumn
		// new root and left
		root := newMonkeys["root"]
		// leftSide := newMonkeys[root.deps[0]]

		v := evalMonkeys(root, newMonkeys, true)

		newLeft := *newMonkeys[root.deps[0]].num
		newRight := *newMonkeys[root.deps[1]].num
		// newDiff := newLeft - newRight
		fmt.Println(v, currHumn, newLeft, newRight, newLeft-newRight)

		if v == 0 {
			fmt.Println("part 2", currHumn)
			break
		}

		// it's flipped
		if multiply && v > 0 {
			// multiply = false
		}
		// it's flipped
		// if divide && v < 0 {
		// 	divide = false
		// }
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

func evalMonkeys(curr *monkey2, monkeys map[string]*monkey2, evalRootAsEq bool) int64 {
	// evaluated
	if curr.num != nil {
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

	var ans int64
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
