package main

import (
	"log"
	"sort"
	"strings"
)

type food struct {
	ingredients map[string]bool
	allergens   map[string]bool
	done        bool
}

func day21() {
	lines := readFile("day21input")

	globalAllergens := make(map[string]map[int]bool)
	globalIngredients := make(map[string]map[int]bool)
	allToIngMap := make(map[string]map[string][]int)

	foods := make(map[int]*food, len(lines))
	for i, line := range lines {
		idx := strings.Index(line, "(contains")
		if idx == -1 {
			log.Fatal("wrong")
		}
		ingredients := strings.Split(line[0:idx-1], " ")
		ings := make(map[string]bool)
		for _, ing := range ingredients {
			ings[ing] = true
			m, ok := globalIngredients[ing]
			if !ok {
				m = make(map[int]bool)
			}
			// set the index
			m[i] = true
			// set the ingredients map
			globalIngredients[ing] = m
		}

		allergens := strings.Split(line[idx+10:len(line)-1], ", ")
		allgs := make(map[string]bool)
		for _, a := range allergens {
			allgs[a] = true
			m, ok := globalAllergens[a]
			if !ok {
				m = make(map[int]bool)
			}
			m[i] = true
			// set the ingredients map
			globalAllergens[a] = m

			for _, ing := range ingredients {
				m, ok := allToIngMap[a]
				if !ok {
					m = make(map[string][]int)
				}
				m2, ok2 := m[ing]
				if !ok2 {
					m2 = make([]int, 0)
				}
				m2 = append(m2, i)
				m[ing] = m2
				allToIngMap[a] = m
			}
		}

		foods[i] = &food{
			ingredients: ings,
			allergens:   allgs,
		}
	}

	cannotContain := make(map[string]bool)
	count := 0
	for k, v := range globalIngredients {
		allergen := false
		for _, v2 := range globalAllergens {
			allOk := true
			for k3 := range v2 {
				_, ok := v[k3]
				if !ok {
					allOk = false
					break
				}
			}
			if allOk {
				allergen = true
			}
		}

		if !allergen {
			cannotContain[k] = true
			count += len(v)
		}
	}
	log.Println(count)

	result := make(map[string]map[string]bool)

	for k, v := range allToIngMap {
		allergens := globalAllergens[k]
		for k2, v2 := range v {
			// ingredients + allergens match
			if len(v2) == len(allergens) {
				m, ok := result[k]
				if !ok {
					m = make(map[string]bool)
				}
				m[k2] = true
				result[k] = m

			}
		}
	}

	knownAllergens := make(map[string]string)
	var l []string
	// very similar logic to day 16
	for len(result) > 0 {
		for k, v := range result {
			if len(v) == 1 {
				var found string
				for key := range v {
					knownAllergens[k] = key
					l = append(l, k)
					found = key
				}

				for k2, v2 := range result {
					if k2 == k {
						continue
					}

					// delete key from all other results
					delete(v2, found)
				}
				// delete what we just found
				delete(result, k)
			}
		}
	}

	sort.Strings(l)
	realresult := make([]string, len(l))
	for k, v := range l {
		realresult[k] = knownAllergens[v]
	}
	log.Println(strings.Join(realresult, ","))
}
