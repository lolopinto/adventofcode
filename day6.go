package main

import (
	"fmt"
	"strings"
)

func day6() {
	lines := readFile("day6input")
	line := lines[0]
	input := ints(strings.Split(line, ","))

	//	fmt.Println(input)
	// dst := make([]int, len(input))
	// copy(dst, input)
	for i := 1; i <= 80; i++ {
		var input2 []int
		newLanternCt := 0
		for _, num := range input {
			num2 := num - 1
			//			fmt.Println("num2", num2)

			// if num2 == -1 {
			// 	num2 = 6
			// 	input2[len(input2)-1] = 8
			// }
			switch num2 {
			case -1:
				num2 = 6
				//				newLantern = true
				newLanternCt++
			}
			input2 = append(input2, num2)

		}
		for i := 0; i < newLanternCt; i++ {
			input2 = append(input2, 8)
		}
		// sum := 0
		// for _, num := range input2 {
		// 	sum += num
		// }
		fmt.Println("day", i, len(input2))
		// if i == 2 {
		// 	fmt.Println(input2)
		// }

		input = make([]int, len(input2))
		copy(input, input2)
	}
}
