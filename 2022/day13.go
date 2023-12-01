package main

import (
	"fmt"
	"sort"
	"unicode"
)

type intchild int

func (i intchild) childMarker() bool {
	return true
}

type child interface {
	childMarker() bool
}

type list struct {
	children []child
}

func (i *list) childMarker() bool {
	return true
}

func (l *list) add(v intchild) {
	l.children = append(l.children, v)
}

func (l *list) comp(right *list) int {
	for i := 0; i < len(l.children); i++ {

		if i >= len(right.children) {
			return 1
		}
		item := l.children[i]
		rightitem := right.children[i]

		lc, ok := item.(intchild)
		rc, ok2 := rightitem.(intchild)
		if ok && ok2 {
			if lc == rc {
				continue
			}
			if lc < rc {
				return -1
			}
			return 1
		}

		ll, ok3 := item.(*list)
		rl, ok4 := rightitem.(*list)

		if ok3 && ok4 {
			cmp := ll.comp(rl)
			if cmp != 0 {
				return cmp
			}

			continue
		}

		// was individual, but not list
		// convert to list
		if ok {
			ll = (&list{})
			ll.add(lc)
		}
		if ok2 {
			rl = (&list{})
			rl.add(rc)
		}

		cmp := ll.comp(rl)
		if cmp != 0 {
			return cmp
		}
	}

	if len(l.children) == len(right.children) {
		return 0
	}
	// left side ran out of items, right order
	return -1
}

func day13() {
	chunks := readFileChunks("day13input", -1)

	parseInput := func(line string) *list {
		ret := &list{}

		stack := []*list{ret}
		i := 0
		curr := ret

		for i < len(line) {
			c := line[i]
			if c == '[' {
				if i != 0 {
					temp := &list{}

					curr.children = append(curr.children, temp)
					stack = append(stack, temp)
					curr = temp
				}
				i++
			} else if c == ']' {
				i++
				stack = stack[:len(stack)-1]
				if len(stack) > 0 {
					curr = stack[len(stack)-1]
				}
			} else if unicode.IsDigit(rune(c)) {
				end := i
				for j := i; j < len(line); j++ {
					if !unicode.IsDigit(rune(line[j])) {
						end = j
						break
					}
				}
				curr.add(intchild(atoi(line[i:end])))
				i = end
			} else if c != ',' {
				panic(fmt.Sprintf("invalid character %s", string(rune(c))))
			} else {
				i++
			}
		}
		return ret
	}

	// part 1
	sum := 0
	for i, chunk := range chunks {
		left := parseInput(chunk[0])
		right := parseInput(chunk[1])

		if left.comp(right) < 0 {
			sum += (i + 1)
		}
	}

	fmt.Println("part 1 answer:", sum)

	// part 2
	var lists []*list
	for _, chunk := range chunks {
		left := parseInput(chunk[0])
		right := parseInput(chunk[1])
		lists = append(lists, left, right)
	}

	div1 := parseInput("[[2]]")
	div2 := parseInput("[[6]]")
	lists = append(lists, div1, div2)

	sort.Slice(lists, func(i, j int) bool {
		cmp := lists[i].comp(lists[j])
		return cmp < 0
	})

	mult := 1
	for idx, v := range lists {
		if v == div1 || v == div2 {
			mult *= (idx + 1)
		}
	}

	fmt.Println("part 2 answer:", mult)
}
