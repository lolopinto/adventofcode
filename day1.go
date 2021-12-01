package main

import "fmt"

func day1part1() {
	lines := readFile("day1input")
	last := 0
	increased := 0
	for i, line := range lines {
		cur := atoi(line)
		if i != 0 {
			if cur > last {
				increased++
			}
		}
		last = cur
	}
	fmt.Println(increased)
}

func day1() {
	lines := readFile("day1input")

	last := 1
	windows := make([]int, len(lines))
	for i := range windows {
		windows[i] = last
		last++
	}

	sums := make(map[int]int)
	for i, line := range lines {
		num := atoi(line)
		start := i - 2
		if start < 0 {
			start = 0
		}
		nums := windows[start : i+1]
		for _, v := range nums {
			sum, ok := sums[v]
			if !ok {
				sum = 0
			}
			sum += num
			sums[v] = sum
		}
	}
	last = 0
	increased := 0
	for ii, key := range windows {
		if ii != 0 {
			cur := sums[key]
			if cur > last {
				increased++
			}
			last = cur
		}
	}
	fmt.Println(increased)
}
