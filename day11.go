package main

import "fmt"

type data struct {
	flashed bool
	value   int
}

func day11() {
	lines := readFile("day11input")

	length := len(lines)

	// convert to int and set it
	input := make([][]data, length)
	for i := 0; i < length; i++ {
		input[i] = make([]data, length)
		for j := 0; j < length; j++ {
			input[i][j] = data{
				value: convToNum(rune(lines[i][j])),
			}
		}
	}

	flashct := 0
	iterations := 0
	for {
		for r := 0; r < length; r++ {
			for c := 0; c < length; c++ {
				if input[r][c].flashed {
					continue
				}

				input[r][c].value += 1
				val := input[r][c].value
				if val > 9 {
					flashct += flash(&input, r, c, length)
				}
			}
		}

		inputcopy := make([][]data, length)
		allzero := 0
		for i := 0; i < length; i++ {
			inputcopy[i] = make([]data, length)

			for j := 0; j < length; j++ {
				inputcopy[i][j] = data{
					value: input[i][j].value,
				}
				if input[i][j].value == 0 {
					allzero++
				}
			}
		}

		// part 2
		if allzero == length*length {
			fmt.Println(iterations + 1)
			break
		}
		iterations++
		input = inputcopy
	}
	// part1
	//	fmt.Println(flashct)
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

	for _, pos := range adj {
		// already flashed. nothing to do here
		if d2[pos.r][pos.c].flashed {
			continue
		}

		d2[pos.r][pos.c].value += 1
		val := d2[pos.r][pos.c].value

		if val > 9 {
			ct += flash(d, pos.r, pos.c, length)
		}
	}

	return ct
}

type p struct {
	r, c int
}

func getAllAdjacent(r, c, length int) []p {
	var ret []p

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
