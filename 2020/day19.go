package main

import (
	"log"
	"strconv"
	"strings"
)

type rule struct {
	idx   int
	rule  string
	rules [][]int
}

// got inspiration from someone on reddit for this one
// go through and match recursively from the front to see if the letters match
// if we end up with "" after going through, we've matched
func applyRules(r *rule, rules map[int]*rule, str string) map[string]bool {
	log.Println(r, str)
	result := make(map[string]bool)

	if r.rule != "" {
		if str != "" && string(str[0]) == r.rule {
			result[str[1:]] = true
			return result
		}
	} else {
		for _, opts := range r.rules {
			m := make(map[string]bool)
			m[str] = true
			for _, idx := range opts {
				r2 := rules[idx]

				m2 := make(map[string]bool)
				for k := range m {
					vals := applyRules(r2, rules, k)
					for k := range vals {
						m2[k] = true
					}
				}
				m = m2
				if len(m) == 0 {
					break
				}
			}
			for k := range m {
				result[k] = true
			}
		}
	}

	return result
}

func day19() {
	chunks := readFileChunks("day19input", 2)
	rules := make(map[int]*rule)
	count := 0

	for _, r := range chunks[0] {
		parts := splitLength(r, ": ", 2)
		//			log.Println(parts[0], len(parts[0]), "dis", parts[1])
		idx := atoi(parts[0])
		r := &rule{
			idx: idx,
		}

		if parts[1][0] == '"' {
			rule, err := strconv.Unquote(parts[1])
			die(err)
			r.rule = rule
		} else {
			rs := strings.Split(parts[1], " | ")
			for _, indices := range rs {
				r.rules = append(r.rules, ints(strings.Split(indices, " ")))
			}
		}

		rules[idx] = r
	}

	// only thing needed for part2 crazy
	// rules[8] = &rule{idx: 8, rules: [][]int{[]int{42}, []int{42, 8}}}
	// rules[11] = &rule{idx: 42, rules: [][]int{[]int{42, 31}, []int{42, 11, 31}}}

	for _, msg := range chunks[1] {
		ret := applyRules(rules[0], rules, msg)
		_, ok := ret[""]
		if ok {
			count++
		}
	}

	log.Println(count)
}
