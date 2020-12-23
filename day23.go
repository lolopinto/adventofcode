package main

import (
	"log"
)

func day23() {
	lines := readFile("day23input")

	count := 1_000_000
	numbers := make([]int, count)

	for i, c := range lines[0] {
		numbers[i] = atoi(string(c))
	}

	//	no zero
	for j := len(lines[0]) + 1; j <= count; j++ {
		numbers[j-1] = j
	}

	m, currentNode := buildNodes(numbers)

	for i := 1; i <= 10_000_000; i++ {

		destination := currentNode.num - 1
		var destinationNode *crabNode

		// pickup...
		first := currentNode.next
		second := first.next
		third := second.next
		fourth := third.next

		for {
			if destination <= 0 {
				destination = count
			}

			if destination != first.num && destination != second.num && destination != third.num {
				break
			}
			destination--
		}

		destinationNode = m[destination]

		// first thing
		currentNode.next = destinationNode
		temp := destinationNode.next
		// next
		destinationNode.next = first
		third.next = temp
		currentNode.next = fourth

		currentNode = currentNode.next
	}

	//		log.Println(n.num)
	// part 1
	n := m[1]
	// var sb strings.Builder
	// n = n.next
	// for n.num != 1 {
	// 	sb.WriteString(fmt.Sprintf("%v", n.num))
	// 	n = n.next
	// }
	// log.Println(sb.String())

	// part 2
	log.Println(n.next.num * n.next.next.num)

}

type crabNode struct {
	num  int
	next *crabNode
}

func buildNodes(numbers []int) (map[int]*crabNode, *crabNode) {
	m := make(map[int]*crabNode)
	var first, last *crabNode
	for idx, num := range numbers {
		var next *crabNode
		if idx == 0 {
			first = &crabNode{num: num}
			last = first
			m[num] = first
		}
		if idx != len(numbers)-1 {
			num2 := numbers[idx+1]
			next = &crabNode{num: num2}
			m[num2] = next
			last.next = next
		} else {
			last.next = first
		}
		last = next

	}
	return m, first
}
