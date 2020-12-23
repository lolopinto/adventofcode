package main

import (
	"fmt"
	"log"
	"strings"
)

func day23() {
	lines := readFile("day23input")

	numbers := make([]int, len(lines[0]))

	for i, c := range lines[0] {
		numbers[i] = atoi(string(c))
	}

	correctLen := len(numbers)
	// move 8 broken

	for i := 1; i <= 100; i++ {
		currentIdx := (i - 1) % len(numbers)
		currentCup := numbers[currentIdx]

		log.Println("move", i, numbers, "current idx", currentIdx, "current cup", currentCup)
		//		newIdx := currentIdx + 3
		//		pickedUp := [3]int{currentIdx + 1, currentIdx + 2, currentIdx + 3}
		//		log.Println("new slice", currentIdx+1, currentIdx+4)
		pickedUp := newSlice(numbers, currentIdx+1, currentIdx+4)
		//		numbers[currentIdx+1 : currentIdx+4]
		//		left := numbers[0 : currentIdx+1]
		//		var left []int
		//		if currentIdx+4 < len(numbers) {
		leftIdx := 0
		if currentIdx+4 > len(numbers) {
			leftIdx = (currentIdx + 4) % len(numbers)
		}
		//		log.Println("left idx", leftIdx)

		left := newSlice(numbers, leftIdx, currentIdx+1)
		// } else {
		// 	// wraps around
		// 	left = newSlice(numbers, currentIdx+4, currentIdx+1)
		// }

		right := newSlice(numbers, currentIdx+4, -1)
		//		right := numbers[currentIdx+4:]

		destination := currentCup - 1
		for {
			if destination <= 0 {
				destination = 9
			}

			inDestination := false
			for _, p := range pickedUp {
				if p == destination {
					//					log.Println("TODO", p, destination)
					inDestination = true
					break
				}
			}
			if !inDestination {
				break
			}
			destination--
		}
		log.Println("pickedup", pickedUp, "left", left, "righth", right, "destination", destination)
		foundIdx := -1
		for idx, v := range right {
			if v == destination {
				foundIdx = idx
				break
			}
		}
		// it's in theh left
		if foundIdx == -1 {
			for idx, v := range left {
				if v == destination {
					foundIdx = idx
					break
				}
			}
			if foundIdx == -1 {
				log.Fatalf("whaaa")
			}

			// nextCupIdx := currentIdx + 1
			// nextCup := numbers[nextCupIdx]
			numbers = make([]int, 0)
			//			numbers = append(numbers, left...)
			// it's found in the left
			numbers = append(numbers, newSlice(left, 0, foundIdx+1)...)
			numbers = append(numbers, pickedUp...)
			numbers = append(numbers, newSlice(left, foundIdx+1, -1)...)
			numbers = append(numbers, right...)

			// now change
			//			log.Println("lefttt", numbers, currentIdx, currentCup)

			for k, v := range numbers {
				if v == currentCup {
					//					log.Println("next", k, currentIdx)
					temp := numbers
					numbers = make([]int, 0)
					// rearrange it so that current index remains where it's at
					numbers = append(numbers, newSlice(temp, k-currentIdx, -1)...)
					numbers = append(numbers, newSlice(temp, 0, k-currentIdx)...)
					break
				}
			}

		} else {
			numbers = make([]int, 0)
			numbers = append(numbers, left...)
			numbers = append(numbers, newSlice(right, 0, foundIdx+1)...)
			numbers = append(numbers, pickedUp...)
			numbers = append(numbers, newSlice(right, foundIdx+1, -1)...)
			//		currentIdx = len(left) + foundIdx

		}

		if len(numbers) != correctLen {
			log.Fatal("numbers incorrect", numbers, "end of move ", i)
		}
	}

	idx := -1
	for k, v := range numbers {
		if v == 1 {
			idx = k
			break
		}
	}
	//	log.Println(idx)
	var sb strings.Builder
	for i := idx + 1; i < idx+len(numbers); i++ {
		j := i % len(numbers)
		sb.WriteString(fmt.Sprintf("%v", numbers[j]))
	}
	log.Println(sb.String())

	//	log.Println(numbers)
}

// func getIdx(numbers []int, idx int) int {
// 	log.Println(idx)
// 	idx = idx % len(numbers)
// 	log.Println(idx % len(numbers))
// 	return numbers[idx]

// 	// TODO
// 	return -1
// }

func newSlice(slice []int, left, right int) []int {
	if left == -1 {
		left = 0
		//		return slice[0:getIdx(slice, right)]
		// from beginning
	} else if right == -1 {
		right = len(slice)
		//		return slice[getIdx(slice, left):]
		// from idx to end
	}

	var ret []int
	for i := left; i < right; i++ {
		ret = append(ret, slice[i%len(slice)])
	}
	return ret
}
