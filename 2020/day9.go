package main

import (
	"fmt"
	"log"
	"sort"
	"strings"
)

func day9() {
	lines := readFile("day9input")

	preamble := 25
	numbers := make([]int, len(lines))

	findNumber := func() int {
		for idx, line := range lines {
			num := atoi(line)
			numbers[idx] = num
		}

		for i := preamble; i < len(numbers); i++ {
			num := numbers[i]

			list1 := numbers[i-preamble : i]
			list2 := numbers[i-preamble : i]

			var found bool
			for _, candidate1 := range list1 {
				for _, candidate2 := range list2 {
					if candidate1+candidate2 == num {
						found = true
						break
					}
				}
			}

			if !found {
				return num
			}
		}
		panic("should have returned")
	}

	num := findNumber()

	ranges := make(map[string]int)
	for idx, num2 := range numbers {
		lastSum := num2
		for idx2 := idx + 1; idx2 < len(numbers); idx2++ {
			num3 := numbers[idx2]
			lastSum = lastSum + num3
			ranges[fmt.Sprintf("%d:%d", idx, idx2)] = lastSum
		}
	}

	for key, value := range ranges {
		if value == num {
			var result []int
			parts := strings.Split(key, ":")
			num := atoi(parts[0])
			num2 := atoi(parts[1])
			for i := num; i < num2; i++ {
				result = append(result, numbers[i])
			}
			sort.Ints(result)
			log.Println(result[0] + result[len(result)-1])
		}
	}
}
