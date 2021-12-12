package main

import (
	"fmt"
	"strings"
)

func day12() {
	lines := readFile("day12input")
	neighbors := make(map[string][]string)

	addNeighbor := func(n1, n2 string) {
		l, ok := neighbors[n1]
		if !ok {
			l = []string{}
		}
		l = append(l, n2)
		neighbors[n1] = l
	}

	for _, line := range lines {
		parts := strings.Split(line, "-")
		addNeighbor(parts[0], parts[1])
		addNeighbor(parts[1], parts[0])
	}

	fmt.Println(checkPaths([]string{"start"}, neighbors, false))
}

func contains(haystack []string, needle string) bool {
	for _, s := range haystack {
		if s == needle {
			return true
		}
	}
	return false
}

func checkPaths(l []string, neighbors map[string][]string, visited bool) int {
	count := 0
	lastIdx := len(l) - 1
	last := l[lastIdx]

	if last == "end" {
		return 1
	}

	for _, neighbor := range neighbors[last] {
		// can check for upper or for something not yet visited
		if strings.ToUpper(neighbor) == neighbor || !contains(l, neighbor) {
			count += checkPaths(append(l, neighbor), neighbors, visited)
		}

		if !visited && neighbor != "start" && neighbor != "end" && strings.ToLower(neighbor) == neighbor && contains(l, neighbor) {
			count += checkPaths(append(l, neighbor), neighbors, true)
		}
	}

	return count
}
