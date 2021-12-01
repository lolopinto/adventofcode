package main

// https://adventofcode.com/2020/day/1
import (
	"log"
	"os"
)

func day1() {
	lines := readFile("day1input")
	m := make(map[int]bool)
	for _, line := range lines {
		i := atoi(line)
		m[i] = true
	}

	for num := range m {
		target := 2020 - num
		for num2 := range m {
			target2 := target - num2

			if m[target2] {
				log.Println(num, num2, target2)
				log.Println(num * num2 * target2)
				os.Exit(1)
			}
		}
	}

}
