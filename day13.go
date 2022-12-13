package main

import (
	"fmt"
	"unicode"
)

type listItem interface {
	rightOrder(v listItem) bool
}

type listcontainer interface {
	add(v intchild)
}

// TODO: this is wrong. it should be list. will it matter?
type intList []intchild

func (l *intList) rightOrder(v listItem) bool {
	r, ok := v.(*intList)
	if !ok {
		return false
	}
	list := *l
	list2 := *r
	// fmt.Println(list, list2)
	for i := 0; i < len(list); i++ {
		if i >= len(list2) {
			return false
		}

		if list[i] < list2[i] {
			return true
		} else if list2[i] < list[i] {
			return false
		}
	}
	// should only happen with malformed input..
	// return len(list) == len(list2) {
	// 	return true
	// }
	return false
}

func (l *intList) add(v intchild) {
	*l = append(*l, v)
}

type intchild int

func (i intchild) rightOrder(v listItem) bool {
	val, ok := v.(*intchild)
	if ok {
		// fmt.Println(i, *val)
		if i < *val {
			return true
		}
	}
	return false
}

type list struct {
	children []listItem
}

func (l *list) add(v intchild) {
	l.children = append(l.children, v)
}

func (l *list) rightOrder(right *list) bool {
	for i := 0; i < len(l.children); i++ {

		if i >= len(right.children) {
			return false
		}
		item := l.children[i]
		rightitem := right.children[i]

		lc, ok := item.(intchild)
		rc, ok2 := rightitem.(intchild)
		// fmt.Println(ok, ok2)
		if ok && ok2 {
			fmt.Println("both item", lc, rc)
			// fmt.Println("intchild")
			if lc < rc {
				return true
			}
			// done
			if rc < lc {
				return false
			}
			continue
		}

		ll, ok3 := item.(*intList)
		rl, ok4 := rightitem.(*intList)

		if ok3 && ok4 {
			fmt.Println("both list", ll, rl)
			if ll.rightOrder(rl) {
				return true
			}
			// fmt.Println("compare", ll, rl, 0)
			if rl.rightOrder(ll) {
				return false
			}
			// r = rl.comp(ll)
			// if r != 0 {
			// 	return 0
			// }

			continue
		}

		// was individual, but not list
		// convert to list
		if ok {
			ll = (&intList{})
			ll.add(lc)
		}
		if ok2 {
			rl = (&intList{})
			rl.add(rc)
		}

		fmt.Println("double comp", ll, rl)
		if ll.rightOrder(rl) {
			return true
		}
		if rl.rightOrder(ll) {
			return false
		}
	}

	// cheating because of what  we're doing
	// left side ran out of items, right order
	return true
}

func day13() {
	chunks := readFileChunks("day13input", -1)

	parseInput := func(line string) *list {
		// leftct := 0
		// var lists [][]int
		ret := &list{}

		// var curr []int
		//		curr := ret
		stack := []listcontainer{ret}
		i := 0

		for i < len(line) {
			c := line[i]
			// for i, c := range line {
			if c == '[' {
				if i != 0 {
					curr := &intList{}

					ret.children = append(ret.children, curr)
					stack = append(stack, curr)
				}
				i++
			} else if c == ']' {
				// leftct--
				i++
				// lists = append(lists, curr)
				// fmt.Println(lists)
				// curr = []int{}
				stack = stack[:len(stack)-1]
			} else if unicode.IsDigit(rune(c)) {
				end := i
				for j := i; j < len(line); j++ {
					if !unicode.IsDigit(rune(line[j])) {
						end = j
						break
					}
				}
				curr := stack[len(stack)-1]
				// fmt.Println("curr", curr)
				curr.add(intchild(atoi(line[i:end])))
				// fmt.Println(line[i:end])
				// curr = append(curr, atoi(line[i:end]))
				// fmt.Println(curr)
				i = end
			} else if c != ',' {
				panic(fmt.Sprintf("invalid character %s", string(rune(c))))
			} else {
				i++
			}
		}
		return ret
	}

	sum := 0
	for i, chunk := range chunks {
		// if i != 7 {
		// 	continue
		// }
		left := parseInput(chunk[0])
		right := parseInput(chunk[1])

		// fmt.Println("left")
		// for _, c := range left.children {
		// 	fmt.Println(c)
		// }
		// fmt.Println("right")
		// for _, c := range right.children {
		// 	fmt.Println(c)
		// }
		// fmt.Println(left, right)
		if left.rightOrder(right) {
			// TODO change what's being returned. misunderstood
			// if r != 0 {
			fmt.Println(i + 1)
			sum += (i + 1)
		}

		// break
	}

	fmt.Println("answer:", sum)
}
