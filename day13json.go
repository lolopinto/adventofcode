package main

import (
	"encoding/json"
	"fmt"
	"sort"
)

// for some reason it's float64 at some point, why?? but never int64??
func interfaceToint(i interface{}) (int, bool) {
	v, ok := i.(int)
	if ok {
		return v, ok
	}
	v2, ok := i.(int64)
	if ok {
		return int(v2), ok
	}
	v3, ok := i.(float64)
	if ok {
		return int(v3), ok
	}
	return 0, false
}

func cmplist(l, r []interface{}) int {
	for i := 0; i < len(l); i++ {
		if i >= len(r) {
			return 1
		}

		li := l[i]
		ri := r[i]

		lint, ok := interfaceToint(li)
		rint, ok2 := interfaceToint(ri)
		if ok && ok2 {
			if lint == rint {
				continue
			}
			if lint < rint {
				return -1
			}
			return 1
		}

		ll, ok3 := li.([]interface{})
		rl, ok4 := ri.([]interface{})
		if ok3 && ok4 {
			cmp := cmplist(ll, rl)
			if cmp != 0 {
				return cmp
			}
			continue
		}

		if ok {
			ll = append(ll, lint)
		}
		if ok2 {
			rl = append(rl, rint)
		}

		cmp := cmplist(ll, rl)
		if cmp != 0 {
			return cmp
		}
	}

	if len(l) == len(r) {
		return 0
	}
	return -1
}

func day13json() {
	chunks := readFileChunks("day13input", -1)

	parseInput := func(line string) []interface{} {
		var ret []interface{}
		if err := json.Unmarshal([]byte(line), &ret); err != nil {
			panic(fmt.Errorf("error parsing json: %v", err))
		}
		return ret
	}

	sum := 0
	for i, c := range chunks {
		l := parseInput(c[0])
		r := parseInput(c[1])

		if cmplist(l, r) < 0 {
			sum += (i + 1)
		}
	}

	// part 2
	var lists [][]interface{}

	for _, chunk := range chunks {
		left := parseInput(chunk[0])
		right := parseInput(chunk[1])
		lists = append(lists, left, right)
	}

	div1 := parseInput("[[2]]")
	div2 := parseInput("[[6]]")

	lists = append(lists, div1, div2)

	sort.Slice(lists, func(i, j int) bool {
		cmp := cmplist(lists[i], lists[j])
		return cmp < 0
	})

	mult := 1
	for idx, v := range lists {
		if len(v) != 1 {
			continue
		}
		l2, ok := v[0].([]interface{})
		if !ok || len(l2) != 1 {
			continue
		}
		v2, _ := interfaceToint(l2[0])
		if v2 == 2 || v2 == 6 {
			mult *= (idx + 1)
		}
	}

	fmt.Println("part 1", sum)
	fmt.Println("part 2", mult)
}
