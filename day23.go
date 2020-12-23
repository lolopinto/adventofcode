package main

import (
	"fmt"
	"log"
	"strings"
)

func day23() {
	lines := readFile("day23input")

	count := 9
	numbers := make([]int, count)

	for i, c := range lines[0] {
		numbers[i] = atoi(string(c))
	}

	//	spew.Dump(numbers)
	//	correctLen := len(lines[0])

	// no zero
	// for j := len(lines[0]) + 1; j <= count; j++ {
	// 	numbers[j-1] = j
	// }

	//	spew.Dump(numbers)
	// return

	// num moves is different

	s := &inPlaceSlice{data: numbers}
	for i := 1; i <= 100; i++ {
		currentIdx := (i - 1) % len(s.data)
		currentCup := s.get(currentIdx)

		//		log.Println("move", i, s.data, "current idx", currentIdx, "current cup", currentCup)
		// get indices...
		pickedUpRange := [2]int{currentIdx + 1, currentIdx + 4}
		//		pickedUp := newSlice(numbers, currentIdx+1, currentIdx+4)
		leftIdx := 0
		if currentIdx+4 > len(numbers) {
			leftIdx = (currentIdx + 4) % len(numbers)
		}

		leftRange := [2]int{leftIdx, currentIdx + 1}
		//		left := newSlice(numbers, leftIdx, currentIdx+1)

		rightRange := [2]int{currentIdx + 4, len(s.data)}
		//		right := newSlice(numbers, currentIdx+4, -1)

		//		log.Println("ranges", "pickedUp", pickedUpRange, "left", leftRange, "righth", rightRange)
		destination := currentCup - 1
		for {
			if destination <= 0 {
				destination = count
			}

			inDestination := false
			for i := pickedUpRange[0]; i < pickedUpRange[1]; i++ {
				if s.get(i) == destination {
					inDestination = true
					break
				}
			}
			if !inDestination {
				break
			}
			destination--
		}
		//		log.Println("pickedup", pickedUp, "left", left, "righth", right, "destination", destination)
		foundIdx := -1
		for i := rightRange[0]; i < rightRange[1]; i++ {
			if s.get(i) == destination {
				foundIdx = i
				break
			}
		}
		// it's in the left
		if foundIdx == -1 {
			for i := leftRange[0]; i < leftRange[1]; i++ {
				if s.get(i) == destination {
					foundIdx = i
					break
				}
			}
			if foundIdx == -1 {
				//				log.Println("pickedup", pickedUp, "left", left, "righth", right, "destination", destination)
				log.Fatalf("whaaa, couldn't find destination")
			}

			//			log.Println("left")
			s = s.clone(
				[2]int{leftRange[0], foundIdx + 1},
				[2]int{pickedUpRange[0], pickedUpRange[1]},
				[2]int{foundIdx + 1, leftRange[1]},
				[2]int{rightRange[0], rightRange[1]},
			)

			// now change
			//			log.Println("lefttt", numbers, currentIdx, currentCup)

			// TODO can even simplify this more?
			// TODO hhmm too slow....
			for k, v := range s.data {
				if v == currentCup {
					//					log.Println("next", k, currentIdx)
					// rearrange it so that current index remains where it's at
					//					log.Println(k, currentIdx, temp, v)
					s = s.clone(
						[2]int{k - currentIdx, len(s.data)},
						[2]int{0, k - currentIdx},
					)
					break
				}
			}

		} else {
			//			log.Println("right", foundIdx)
			s = s.clone(
				[2]int{leftRange[0], leftRange[1]},
				[2]int{rightRange[0], foundIdx + 1},
				[2]int{pickedUpRange[0], pickedUpRange[1]},
				[2]int{foundIdx + 1, rightRange[1]},
			)
			//		currentIdx = len(left) + foundIdx

		}

		if len(s.data) != count {
			log.Fatal("numbers incorrect", s.data, "end of move ", i)
		}
	}

	idx := -1
	for k, v := range s.data {
		if v == 1 {
			idx = k
			break
		}
	}
	//	log.Println(idx)
	log.Println(s.get(idx+1) * s.get(idx+2))
	var sb strings.Builder
	for i := idx + 1; i < idx+len(s.data); i++ {
		j := i % len(s.data)
		sb.WriteString(fmt.Sprintf("%v", s.data[j]))
	}
	log.Println(sb.String())

	//	log.Println(numbers)
}

// it's this that's taking so long
// can we append in place?
func newSlice(slice []int, left, right int) []int {
	if left < 0 && left != -1 {
		panic(fmt.Sprintf("invalid number %d", left))
	}
	if left == -1 {
		left = 0
	} else if right == -1 {
		right = len(slice)
	}

	var ret []int
	for i := left; i < right; i++ {
		if i%len(slice) < 0 {
			log.Println(left, right, slice)
		}
		ret = append(ret, slice[i%len(slice)])
	}
	return ret
}

type inPlaceSlice struct {
	data []int
}

func (s *inPlaceSlice) getIndices(left, right int) []int {
	if left < 0 && left != -1 {
		panic(fmt.Sprintf("invalid number %d", left))
	}
	// TODO kill these two...
	if left == -1 {
		left = 0
	} else if right == -1 {
		right = len(s.data)
	}

	ret := make([]int, right-left)
	for i := left; i < right; i++ {
		idx := i - left
		ret[idx] = i % len(s.data)
	}
	return ret
}

func (s *inPlaceSlice) get(idx int) int {
	return s.data[idx%len(s.data)]
}

func (s *inPlaceSlice) clone(r ...[2]int) *inPlaceSlice {
	//	log.Println(r)
	ret := make([]int, len(s.data))
	idx := 0
	for _, rr := range r {
		for i := rr[0]; i < rr[1]; i++ {
			//			log.Println(s.get(i))
			ret[idx] = s.get(i)
			idx++
		}
	}
	if idx != len(s.data) {
		log.Fatalf("invalid clone. ended up with up to index %d", idx)
	}
	return &inPlaceSlice{data: ret}
}
