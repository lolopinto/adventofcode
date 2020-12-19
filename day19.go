package main

import (
	"log"
	"strconv"
	"strings"
)

type rule struct {
	idx   int
	rule  string
	rules [][]int //
}

func optionsForRule(r *rule, rules map[int]*rule, result map[int]map[string]bool) map[string]bool {
	m, ok := result[r.idx]
	if ok {
		return m
	}
	m = make(map[string]bool)

	if r.rule != "" {
		//		log.Println(r.idx, r.rule)
		m[r.rule] = true
		result[r.idx] = m
		return m
	}

	for _, opts := range r.rules {
		var strs []string
		// can be more than 2
		for i, idx := range opts {
			r2 := rules[idx]
			options := optionsForRule(r2, rules, result)
			//			spew.Dump("result for options", r2, options)
			if i == 0 {
				for k := range options {
					strs = append(strs, k)
				}
				//				log.Println("1st", r.idx, strs)
			} else {
				var strs2 []string
				for _, s := range strs {
					for k := range options {
						strs2 = append(strs2, s+k)
					}
				}
				//				spew.Dump("2nd", r.idx, strs2)
				strs = strs2
			}
		}
		for _, str := range strs {
			m[str] = true
		}
	}

	//	spew.Dump(r.idx, m)
	result[r.idx] = m
	return m
}

func day19() {
	lines := readFile("day19input")

	result := make(map[int]map[string]bool)
	rules := make(map[int]*rule)
	readRules := true
	options := make(map[string]bool)
	count := 0
	for _, line := range lines {
		if line == "" {
			readRules = false

			// time to check the rules

			r := rules[0]
			options = optionsForRule(r, rules, result)
			//			spew.Dump(options)
			continue
		}

		//		log.Println(line)
		if readRules {
			parts := splitLength(line, ": ", 2)
			//			log.Println(parts[0], len(parts[0]), "dis", parts[1])
			idx := atoi(parts[0])
			r := &rule{
				idx: idx,
				//				rule:
			}
			if parts[1][0] == '"' {
				rule, err := strconv.Unquote(parts[1])
				die(err)
				r.rule = rule
			} else {
				rs := strings.Split(parts[1], " | ")
				for _, indices := range rs {
					//					log.Println(strings.Split(indices, " "))
					r.rules = append(r.rules, ints(strings.Split(indices, " ")))
				}
			}

			//			rules = append(rules)
			//			spew.Dump(rs)
			rules[idx] = r
		} else {
			_, ok := options[line]
			if ok {
				count++
			}
		}
	}
	//	spew.Dump(rules)
	log.Println(count)
}
