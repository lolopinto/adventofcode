package main

import (
	"fmt"
	"os"
	"unicode"
)

func day21binary() {
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

	// var low, high int64
	var lower int64 = 0
	// can't go too high because of integer overflow in go
	var highest int64 = 1000_000_000_000_000
	//
	// 3_715_799_488_132
	// "correct" answer is 3715799488132
	// using binary search gives 3715799488133

	// same issue from https://www.reddit.com/r/adventofcode/comments/zrbw7n/comment/j12r1ju/ apparently

	low := lower
	high := highest
	half := (low + high) / 2

	for {
		fmt.Println(low, half)
		newMonkeys := cloneMonkeys(dupMonkeys)
		// fmt.Println(half)

		newMonkeys["humn"].num = &half
		root := newMonkeys["root"]

		v := evalMonkeys(root, newMonkeys, true)
		if v == 0 {
			fmt.Println("part 2: ", half)
			os.Exit(0)
		}
		if v < 0 {
			low = half
		} else if v > 0 {
			high = half
		}
		// negative numbers in here so flip?
		if high-low == 1 {
			fmt.Println(high, low)
			fmt.Println("weird cycle??")
			// low = math.MaxInt64
			low = highest
			high = lower
		}
		half = (high + low) / 2
	}
}
