package main

import (
	"fmt"
	"regexp"
	"strings"
)

func day5() {
	chunks := readFileChunks("day5input", 2)

	input := chunks[0]

	lastInput := input[len(input)-1]
	stacksMaybe := strings.Split(lastInput, " ")
	var stacks []int
	for _, v := range stacksMaybe {
		v = strings.TrimSpace(v)
		if v == "" {
			continue
		}
		stacks = append(stacks, atoi(v))
	}

	stacksLen := len(stacks)
	stacksMap := map[int][]string{}
	for _, line := range input[:len(input)-1] {
		// fmt.Println("line.....", line)
		for i := 0; i < stacksLen; i++ {
			start := i * 4
			candidate := line[start : start+3]
			if candidate[0] != '[' {
				continue
			}
			letter := string(candidate[1])

			l, ok := stacksMap[i]
			if !ok {
				l = []string{}
			}
			l = append(l, letter)
			stacksMap[i] = l
		}
	}
	// fmt.Println(stacksMap)
	r := regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)

	for _, move := range chunks[1] {
		// fmt.Println(move)
		match := r.FindStringSubmatch(move)
		count := atoi(match[1])
		from := atoi(match[2]) - 1
		to := atoi(match[3]) - 1

		fromList := stacksMap[from]
		toList := stacksMap[to]

		for i := 0; i < count; i++ {
			topFrom := fromList[0]
			fromList = fromList[1:]
			toList = append([]string{topFrom}, toList...)
		}

		stacksMap[from] = fromList
		stacksMap[to] = toList
	}

	var sb strings.Builder
	for i := 0; i < len(stacks); i++ {
		list := stacksMap[i]
		sb.WriteString(list[0])
	}
	fmt.Println(sb.String())
	// fmt.Println(input[len(input)-1])

}
