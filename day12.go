package main

import (
	"fmt"
	"strings"
)

func day12() {
	lines := readFile("day12input")
	nodes := make(map[string]*node)
	//	neighbors := make
	for _, line := range lines {
		parts := strings.Split(line, "-")
		parseNode(parts[0], nodes)
		parseNode(parts[1], nodes)
		addNeighbors(parts[0], parts[1], nodes)
	}

	fmt.Println(checkPaths([]string{"start"}, nodes, false))
}

func contains(haystack []string, needle string) bool {
	for _, s := range haystack {
		if s == needle {
			return true
		}
	}
	return false
}

func checkPaths(l []string, nodes map[string]*node, visited bool) int {
	count := 0
	lastIdx := len(l) - 1
	last := l[lastIdx]

	if last == "end" {
		return 1
	}

	for _, neighbor := range nodes[last].neighbors {
		// can check for upper or for something not yet visited
		if strings.ToUpper(neighbor.node) == neighbor.node || !contains(l, neighbor.node) {
			count += checkPaths(append(l, neighbor.node), nodes, visited)
		}

		if !visited && neighbor.node != "start" && neighbor.node != "end" && strings.ToLower(neighbor.node) == neighbor.node && contains(l, neighbor.node) {
			count += checkPaths(append(l, neighbor.node), nodes, true)
		}
	}

	return count
}

type node struct {
	node      string
	neighbors []*node
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
