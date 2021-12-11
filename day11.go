package main

import "fmt"

type data struct {
	flashed bool
	value   int
}

func day11() {
	lines := readFile("day11input")
	// 10 * 10 grid
	//	visited := make([][]da, length)

	length := len(lines)

	// convert to int and set it
	input := make([][]data, length)
	for i := 0; i < length; i++ {
		input[i] = make([]data, length)
		for j := 0; j < length; j++ {
			//			fmt.Println(lines[i][j], convToNum(rune(lines[i][j])))
			input[i][j] = data{
				value: convToNum(rune(lines[i][j])),
			}
		}
	}

	flashct := 0
	for i := 0; i < 100; i++ {
		//		inputCopy := make([][]data, length)
		//		copy(inputCopy, input)
		//		fmt.Println(lines, input)
		for r := 0; r < length; r++ {
			for c := 0; c < length; c++ {
				if input[r][c].flashed {
					continue
				}

				input[r][c].value += 1
				val := input[r][c].value
				if val == 10 {
					//					fmt.Println("inital")
				}
				//				num := convRuneToInt(newLines[r][c]) + 1
				if val > 9 {
					//					input[r][c].value = 0

					//					fmt.Println(val > 9)
					flashct += flash(&input, r, c, length)
					//					newLines[r][c]
					//					flashct++
					// adj := getAllAdjacent(r, c, length)
					// for _, pos := range adj {
					// 	if inputCopy[pos.r][pos.c].flashed {
					// 		continue
					// 	}
					// }
				}
			}
		}

		inputcopy := make([][]data, length)
		for i := 0; i < length; i++ {
			inputcopy[i] = make([]data, length)

			for j := 0; j < length; j++ {
				inputcopy[i][j] = data{
					value: input[i][j].value,
				}
			}
			//			fmt.Println(inputcopy[i])
		}
		input = inputcopy
		//		fmt.Println(input)
	}
	fmt.Println(flashct)
}

func flash(d *[][]data, r, c, length int) int {
	d2 := *d

	if d2[r][c].flashed {
		return 0
	}
	d2[r][c].flashed = true
	d2[r][c].value = 0

	ct := 1
	adj := getAllAdjacent(r, c, length)

	//	var toflash []p
	for _, pos := range adj {
		// already flashed. nothing to do here
		if d2[pos.r][pos.c].flashed {
			continue
		}

		//		fmt.Println(d2[pos.r][pos.c].value)

		d2[pos.r][pos.c].value += 1
		val := d2[pos.r][pos.c].value
		//		fmt.Println(d2[pos.r][pos.c].value, val)

		if val == 10 {
			//			fmt.Println("loop")
		}
		if val == 1 {
			//			fmt.Println("whaa 1", pos)
		}

		if val > 9 {
			//			fmt.Println("subsequent flash")
			ct += flash(d, pos.r, pos.c, length)
			// d2[pos.r][pos.c].flashed = true
			// ct++
			// toflash = append(toflash, pos)
		}
	}
	// for _, v := range toflash {
	// 	ct += flash(d, v.r, v.c, length)
	// }
	return ct
}

type p struct {
	r, c int
}

func getAllAdjacent(r, c, length int) []p {
	var ret []p
	// if r-1 >= 0 {
	// 	ret = append(ret, p{r - 1, c})
	// }
	// if r+1 < length {
	// 	ret = append(ret, p{r + 1, c})
	// }
	// if c-1 >= 0 {
	// 	ret = append(ret, p{r, c - 1})
	// }
	// if c+1 < length {
	// 	ret = append(ret, p{r, c + 1})
	// }
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			if i == 0 && j == 0 {
				continue
			}
			newR := r + i
			newC := c + j
			if newR >= 0 && newR < length && newC >= 0 && newC < length {
				ret = append(ret, p{newR, newC})
			}
		}
	}

	return ret
}
