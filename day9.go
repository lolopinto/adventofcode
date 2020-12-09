package main

import (
	"log"
	"os"
	"strconv"
)

func day9() {
	lines := readFile("day9input")

	preamble := 25
	numbers := make([]int, len(lines))
	for idx, line := range lines {
		num, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		numbers[idx] = num
	}
	//	spew.Dump(numbers)

	for i := preamble; i < len(numbers); i++ {
		num := numbers[i]

		//		log.Println("initial", num)
		//		var found bool

		list1 := numbers[i-preamble : i]
		list2 := numbers[i-preamble : i]
		//		spew.Dump(list1, list2)

		var found bool
		for _, candidate1 := range list1 {
			for _, candidate2 := range list2 {
				if candidate1+candidate2 == num {
					found = true
					break
				}
			}
		}
		// for j := i - 1; j == preamble; j-- {
		// 	for k := i - 1; k == preamble; k-- {
		// 		log.Println(num, numbers[j], numbers[k])
		// 		if numbers[j]+numbers[k] == num {
		// 			found = true
		// 			break
		// 		}
		// 	}
		// }

		if !found {
			log.Println(num)
			os.Exit(1)
		}
	}
}
