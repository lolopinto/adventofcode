package main

import (
	"fmt"
)

// couldn't get this to work for real input for some reason with index math
// so implemented in javascript. my index math is broken apparently
// Python would have worked too
func day20() {
	lines := readFile("day20input")
	vals := ints(lines)
	dup_indices := make([]int, len(vals))

	indices := make([]int, len(vals))
	for i := range vals {
		indices[i] = i
	}
	copy(dup_indices, indices)

	getNewIdx := func(i, delta int) int {
		new_idx := (i + delta) % len(vals)
		if delta < 0 && new_idx < 0 {
			new_idx += len(vals) - 1
		} else if i+delta > len(vals) {
			new_idx += 1
		}

		return new_idx
	}

	for i := range indices {
		// fmt.Println(i)
		curr_idx := -1
		for curr, idx := range indices {
			// fmt.Println(curr, idx)
			if idx == i {
				curr_idx = curr
				break
			}
		}

		if curr_idx == -1 {
			panic("invalid idx")
		}

		new_idx := getNewIdx(curr_idx, vals[i])
		// fmt.Println("idx", i, curr_idx, vals[i], new_idx)

		// remove from slice
		// https://github.com/golang/go/wiki/SliceTricks#expand

		// delete from curr_idx
		indices = append(indices[:curr_idx], indices[curr_idx+1:]...)

		// add at new pos
		indices = append(indices[:new_idx], append([]int{i}, indices[new_idx:]...)...)

		// fmt.Println(indices)

		// OR
		// indices = append(indices, 0)
		// copy(indices[new_idx+1:], indices[new_idx:])
		// indices[new_idx] = i
	}

	zero_idx := -1
	zero_idx_in_orig := -1
	for idx, v := range vals {
		if v == 0 {
			zero_idx_in_orig = idx
			break
		}
	}

	for idx, v := range indices {
		if v == zero_idx_in_orig {
			zero_idx = idx
			break
		}
	}
	// fmt.Println(vals)
	fmt.Println(zero_idx)

	v1 := vals[indices[getNewIdx(zero_idx, 1000)]]
	v2 := vals[indices[getNewIdx(zero_idx, 2000)]]
	v3 := vals[indices[getNewIdx(zero_idx, 3000)]]
	fmt.Println(v1, v2, v3)
	fmt.Println("answer", v1+v2+v3)

}
