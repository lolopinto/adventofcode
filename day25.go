package main

import (
	"fmt"
	"math"
	"os"
	"strings"
	"unicode"
)

func day25() {
	// powers of 5
	// instead of 0-4
	//nstead of using digits four through zero,
	// the digits are 2, 1, 0, minus (written -), and double-minus (written =). Minus is worth -1, and double-minus is worth -2."

	lines := readFile("day25input")
	var sum int64
	for _, line := range lines {
		sum += snafuToDecimal(line)
	}

	list := []rune{
		'=',
		'-',
		'0',
		'1',
		'2',
	}

	i := 0
	start := ' '
	for {
		pow := int64(math.Pow(5, float64(i)))

		if sum >= pow && sum <= pow*2 {
			if sum >= pow {
				start = '2'
			} else {
				start = '1'
			}
			break
		}
		i++
	}

	valid := make([]rune, i+1)
	valid[0] = start

	toString := func(l []rune) string {
		var sb strings.Builder
		for _, v := range l {
			sb.WriteString(string(v))
		}
		return sb.String()
	}

	// set the rest of the fields
	for j := 1; j < i+1; j++ {
		valid[j] = '0'
	}

	// go through each character from 1-N and pick the one that's closest difference
	// and use it until we end at difference 0
	for i := 1; i < len(valid); i++ {
		copy := copyList(valid)

		diffs := make([]int64, len(list))
		diffMap := make(map[int64]rune)
		for j, c := range list {
			copy[i] = c
			val := snafuToDecimal(toString(copy))

			diff := abs64(val, sum)
			if diff == 0 {
				fmt.Println("answer", toString(copy))
				os.Exit(0)
			}

			diffs[j] = diff
			diffMap[diff] = c
		}
		sm := min64(diffs)
		v, ok := diffMap[sm]
		if !ok {
			panic(fmt.Errorf("couldn't find diff in diffMap %v", sm))
		}
		valid[i] = v
	}

	fmt.Println(toString(valid), i)
}

func snafuToDecimal(line string) int64 {
	start := len(line) - 1

	var sum int64
	for i, c := range line {
		pow := start - i
		var val int
		switch c {
		case '=':
			val = -2
		case '-':
			val = -1
		default:
			if !unicode.IsDigit(c) {
				panic(fmt.Errorf("invalid digit %d", c))
			}
			val = atoi(string(c))
		}
		curr := int64(math.Pow(5, float64(pow)) * float64(val))
		sum += curr
	}
	return sum
}
