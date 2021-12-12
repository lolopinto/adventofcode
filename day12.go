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
	for _, neighbor := range findNode(nodes, "start").neighbors {
		q = append(q, []string{"start", neighbor.node})
	}

	uniquepaths := make(map[string]bool)
	//	fmt.Println(q)
	for len(q) > 0 {
		var q2 [][]string
		for i := 0; i < len(q); i++ {
			poss := q[i]
			//			q = q[1:]

			lastIdx := len(poss) - 1
			last := poss[lastIdx]
			paths := getPaths(last, nodes)
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
	fmt.Println(uniquepaths)
	fmt.Println(len(uniquepaths))
}

func checkValid(path []string) bool {
	if path[0] != "start" {
		return false
	}
	visited := make(map[string]bool)
	for _, v := range path {
		// lower
		if strings.ToLower(v) == v {
			if visited[v] {
				return false
			}
			visited[v] = true
		}
	}
	return true
}

func getPaths(s string, nodes map[string]*node) [][]string {
	var ret [][]string
	//	resetvisited(nodes)

	//	fmt.Println("paths", s)
	n := findNode(nodes, s)
	app := n.possiblesPaths()
	for _, pp := range app {
		var ss []string
		for _, p := range pp {
			ss = append(ss, p.node)
			//			fmt.Println(p.node)
		}
		//		fmt.Println(strings.Join(ss, ","))
		ret = append(ret, ss)
	}
	return ret
}

// func resetvisited(nodes map[string]*node) {
// 	// for _, v := range nodes {
// 	// 	v.visited = false
// 	// }
// }

type node struct {
	node string
	// start     bool
	// end       bool
	neighbors []*node
	//	visited   bool
	big bool
}

// // func (n *node) canVisit() bool {
// // 	if n.big || n.node == "end" {
// // 		return true
// // 	}
// // 	return !n.visited
// // }

// func (n *node) possiblePath(neighbor *node) bool {
// 	// can't start from end or end with start
// 	if n.node == "end" || neighbor.node == "start" {
// 		return false
// 	}
// 	if n.node == "start" || neighbor.node == "end" {
// 		return true
// 	}
// 	// can't visit the same node twice if small
// 	if n.node == neighbor.node && !n.big && !neighbor.big {
// 		return false
// 	}
// 	return true

// }

func (n *node) possiblesPaths() [][]*node {
	var ret [][]*node
	for _, neighbor := range n.neighbors {
		// if !n.possiblePath(neighbor) {
		// 	//			fmt.Println("not possible path", n.node, neighbor.node)
		// 	//			continue
		// }
		// if !neighbor.canVisit() {
		// 	//			fmt.Println("cannot visit", n.node, neighbor.node)
		// 	// /			continue
		// }

		ret = append(ret, []*node{n, neighbor})
		//		paths := []*node{n, neighbor}
		//		path := []*node{n, neighbor}
		//		fmt.Println(n.node, neighbor.node)
		for _, n2 := range neighbor.neighbors {
			//			if n2.canVisit() && neighbor.possiblePath(n2) {
			ret = append(ret, []*node{n, neighbor, n2})
			for _, n3 := range n2.neighbors {
				ret = append(ret, []*node{n, neighbor, n2, n3})

			}
			//			}
			//			fmt.Println(n2.canVisit(), neighbor.possiblePath(n2), neighbor.node, n2.node)
		}
		//		fmt.Println(n.node, neighbor.node, ret)
		//		neighbor.visited = true

		// for _, paths := range neighbor.possiblesPaths() {
		// 	path := []*node{n, neighbor}
		// 	//			path2 := make()
		// 	path = append(path, paths...)
		// 	ret = append(ret, path)
		// }

	}
	return ret
}

func parseNode(part string, nodes map[string]*node) {
	if nodes[part] != nil {
		return
	}
	// allcaps := true
	// for _, c := range part {
	// 	if !unicode.IsUpper(c) {
	// 		allcaps = false
	// 		break
	// 	}
	// }
	n := &node{
		node: part,
		// start: part == "start",
		// end:   part == "end",
		big: strings.ToUpper(part) == part,
	}
	nodes[part] = n
}

func addNeighbors(part1, part2 string, nodes map[string]*node) {
	n1 := nodes[part1]
	n2 := nodes[part2]
	n1.neighbors = append(n1.neighbors, n2)
	n2.neighbors = append(n2.neighbors, n1)
}

func findNode(nodes map[string]*node, s string) *node {
	for _, v := range nodes {
		if v.node == s {
			return v
		}
	}
	panic("couldn't find node " + s)
}
