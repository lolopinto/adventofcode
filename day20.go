package main

import "fmt"

func day20() {
	lines := readFile("day20input")
	vals := ints(lines)
	// brute force but it's not going to work later...
	dup := make([]int, len(vals))

	indices := make(map[int]int)
	copy(dup, vals)
	for i := range vals {
		indices[i] = i
	}

	getNewIdx := func(i, delta int) int {
		new_idx := i + delta
		fmt.Println("getNewIdx", i, delta, new_idx)
		if new_idx == 0 && delta < 0 {
			return len(vals) - 1
		}
		if new_idx >= len(vals) {
			return new_idx%len(vals) + 1
		}
		if new_idx < 0 {
			for new_idx < 0 {
				new_idx = new_idx + len(vals) - 1
				fmt.Println("for new_idx<0", new_idx)
			}
		}

		return new_idx
	}

	getIdx := func(i int) int {
		if i >= 0 {
			return i % len(vals)
		}
		return i + len(vals)
	}
	for _, v := range dup {
		if v == 0 {
			continue
		}
		for idx, v2 := range vals {
			if v == v2 {
				// vals

				new_idx := getNewIdx(idx, v)
				// fmt.Println("searching for", v)
				// new_idx := (idx + v) % len(vals)
				// if new_idx < 0 {
				// 	getIdx(new)
				// 	// fmt.Println("negative tODO", new_idx)
				// }
				// for new_idx < 0 {
				// 	new_idx += len(vals) - 1
				// }
				// fmt.Println(new_idx)
				// let's swap!
				// left := vals[:new_idx]

				// right := vals[new_idx:]
				// fmt.Println(v, "old", idx, "new", new_idx)

				if new_idx > idx {
					for i := idx; i <= new_idx-1; i++ {
						temp := vals[i]
						vals[i] = vals[i+1]
						vals[i+1] = temp
					}
				} else {
					// going downwards but really upwards
					// swap the two numbers and then go just before new_idx
					// fmt.Println("v downwards", v)
					if v > 0 {
						temp := vals[new_idx]
						vals[new_idx] = vals[idx]
						vals[idx] = temp
						// fmt.Println(idx, new_idx)
						// this feels tooo hacky
						new_idx += 3
						// fmt.Println("after swap", vals, idx, new_idx)
					}

					for i := idx; i >= new_idx-1; i-- {
						// fmt.Println("start swap", i)
						temp := vals[getIdx(i)]

						vals[getIdx(i)] = vals[getIdx(i-1)]
						vals[getIdx(i-1)] = temp
					}
				}

				break
			}
		}
		// fmt.Println(vals, v)
		// vals = newVals
		// break
	}

	zero_idx := -1
	for idx, v := range vals {
		if v == 0 {
			zero_idx = idx
			break
		}
	}
	// fmt.Println(vals)
	fmt.Println(zero_idx)
	l := len(vals)

	v1 := vals[getIdx(zero_idx+1000%l)]
	v2 := vals[getIdx(zero_idx+2000%l)]
	v3 := vals[getIdx(zero_idx+3000%l)]
	fmt.Println(v1, v2, v3)
	fmt.Println("answer", v1+v2+v3)

}
