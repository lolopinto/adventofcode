package main

import "fmt"

func day2() {
	opp := map[string]string{"A": "R", "B": "P", "C": "S"}
	// me := map[string]string{"X": "R", "Y": "P", "Z": "S"}

	lines := readFile("day2input")
	sum := 0
	for _, line := range lines {
		parts := splitLength(line, " ", 2)

		// sum += score(opp[parts[0]], me[parts[1]])
		sum += score(opp[parts[0]], getMine(opp[parts[0]], parts[1]))
	}
	fmt.Println(sum)
}

func getMine(opp, me string) string {
	switch opp {
	case "R":
		if me == "X" {
			return "S"
		} else if me == "Y" {
			return "R"
		} else {
			return "P"
		}
	case "P":
		if me == "X" {
			return "R"
		} else if me == "Y" {
			return "P"
		} else {
			return "S"
		}
	default:
		if me == "X" {
			return "P"
		} else if me == "Y" {
			return "S"
		} else {
			return "R"
		}
	}
}

func score(opp, me string) int {
	s := 0
	switch me {
	case "R":
		s += 1
		if opp == "R" {
			s += 3
		}
		if opp == "S" {
			s += 6
		}

	case "P":
		s += 2
		if opp == "P" {
			s += 3
		}
		if opp == "R" {
			s += 6
		}
	default:
		s += 3
		if opp == "S" {
			s += 3
		}
		if opp == "P" {
			s += 6
		}
	}
	return s
}
