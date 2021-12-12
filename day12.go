package main

import (
	"fmt"
	"strings"
)

func day12() {
	lines := readFile("day12input")
	nodes := make(map[string]*node)
	for _, line := range lines {
		parts := strings.Split(line, "-")
		parseNode(parts[0], nodes)
		parseNode(parts[1], nodes)
		addNeighbors(parts[0], parts[1], nodes)
	}

	var q [][]string
	for _, neighbor := range nodes["start"].neighbors {
		q = append(q, []string{"start", neighbor.node})
	}

	uniquepaths := make(map[string]bool)
	for len(q) > 0 {
		var q2 [][]string
		for i := 0; i < len(q); i++ {
			poss := q[i]

			lastIdx := len(poss) - 1
			last := poss[lastIdx]
			paths := nodes[last].possiblesPaths()
			//			fmt.Println(len(paths))
			for _, p := range paths {
				newPoss := append(poss[:lastIdx], p...)
				if !checkValid(newPoss) {
					continue
				}
				l := len(newPoss)
				if newPoss[l-1] == "end" {
					// the end
					uniquepaths[strings.Join(newPoss, ",")] = true
				} else {
					q2 = append(q2, newPoss)
				}
			}
		}
		q = q2
	}
	//	fmt.Println(uniquepaths)
	fmt.Println(len(uniquepaths))
}

func checkValid(path []string) bool {
	visited := make(map[string]int)
	smallcttwice := 0
	l := len(path)
	for i, v := range path {
		if v == "end" && i != l-1 {
			return false
		}
		// lower
		if strings.ToLower(v) == v {
			visited[v] += 1
			ct := visited[v]

			if ct != 1 {
				if v == "start" || v == "end" {
					return false
				} else {
					if ct > 2 {
						return false
					}
					if ct == 2 {
						smallcttwice++
					}
				}
			}
		}
	}
	return smallcttwice < 2
}

type node struct {
	node      string
	neighbors []*node
	pp        [][]string
}

func (n *node) possiblesPaths() [][]string {
	if n.pp != nil {
		return n.pp
	}
	var ret [][]string
	for _, neighbor := range n.neighbors {

		tmp := []string{n.node, neighbor.node}
		if !checkValid(tmp) {
			continue
		}
		ret = append(ret, tmp)
		for _, n2 := range neighbor.neighbors {
			tmp := []string{n.node, neighbor.node, n2.node}
			if !checkValid(tmp) {
				continue
			}
			ret = append(ret, tmp)
			for _, n3 := range n2.neighbors {
				tmp := []string{n.node, neighbor.node, n2.node, n3.node}
				if !checkValid(tmp) {
					continue
				}
				ret = append(ret, tmp)
			}
		}
	}

	n.pp = ret
	return ret
}

func parseNode(part string, nodes map[string]*node) {
	if nodes[part] != nil {
		return
	}

	n := &node{
		node: part,
	}
	nodes[part] = n
}

func addNeighbors(part1, part2 string, nodes map[string]*node) {
	n1 := nodes[part1]
	n2 := nodes[part2]
	n1.neighbors = append(n1.neighbors, n2)
	n2.neighbors = append(n2.neighbors, n1)
}
