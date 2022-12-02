package main

import "fmt"

func day2() {
	opp := map[string]string{"A": "rock", "B": "paper", "C": "scissors"}
	me := map[string]string{"X": "rock", "Y": "paper", "Z": "scissors"}
	lines := readFile("day2input")
	sum := 0
	for _, line := range lines {
		parts := splitLength(line, " ", 2)
		sum += score(opp[parts[0]], me[parts[1]])
	}
	fmt.Println(sum)

}

func score(opp, me string) int {
	s := 0
	switch me {
	case "rock":
		s += 1
		if opp == "rock" {
			s += 3
		}
		if opp == "scissors" {
			s += 6
		}

	case "paper":
		s += 2
		if opp == "paper" {
			s += 3
		}
		if opp == "rock" {
			s += 6
		}
	default:
		s += 3
		if opp == "scissors" {
			s += 3
		}
		if opp == "paper" {
			s += 6
		}
	}
	return s

}
