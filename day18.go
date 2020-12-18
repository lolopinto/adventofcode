package main

import (
	"log"
	"strings"
	"unicode"
)

type part struct {
	line  string
	parts []*part
}

func (p *part) calculate(prev int) int {
	if p.line == ")" {
		return prev
	}
	if len(p.parts) == 0 && p.line == "" {
		return prev
	}
	if len(p.parts) == 0 {
		return calculateLine(p.line, prev)
	}
	// has parts use that
	last := 0
	result := 0
	for i := 0; i < len(p.parts); i++ {
		p2 := p.parts[i]
		// starts with a digit
		// do in isolation
		if unicode.IsDigit(rune(p2.line[0])) {
			//			last = calculateLine(p2.line, 0)
			result = p2.calculate(last)
		} else if p2.line == "*" {
			// look ahead...
			//			result = last * calculateLine(p.parts[i+1].line, 1)
			result = result * p.parts[i+1].calculate(1)
			i++
		} else if p2.line == "+" {
			//			log.Println("last", last)
			// look ahead...
			//			result = last + calculateLine(p.parts[i+1].line, 1)
			result = result + p.parts[i+1].calculate(1)
			i++
		} else { // starts with a + or something
			//			result = calculateLine(p2.line, result)
			result = p2.calculate(result)
		}
	}
	return result
}

func calculateLine(line string, result int) int {
	ops := strings.Split(strings.TrimSpace(line), " ")
	var lastOp string
	for _, op := range ops {

		switch op {
		case "*":
			lastOp = "*"
			break
		case "+":
			lastOp = "+"
			break
		default:
			if lastOp == "" {
				result = atoi(op)
			} else if lastOp == "*" {
				if result == 0 {
					result = 1 * atoi(op)
					lastOp = ""
				} else {
					result = result * atoi(op)
					lastOp = ""
				}
			} else if lastOp == "+" {
				//				log.Println(op, "..", line)
				result += atoi(op)
				lastOp = ""
			}
		}
	}
	return result
}

func breakParts(line string) []*part {
	var results []*part
	var left []int
	var right []int

	var found bool
	lastLeft := 0
	lastRight := 0
	for i, c := range line {
		if c == '(' {
			left = append(left, i)
			found = true
		}
		if c == ')' {
			right = append(right, i)
			found = true
		}
		if len(left) > 0 && len(right) == len(left) {

			lastLeft = left[0]
			// everything before left
			if lastLeft != lastRight {
				results = append(results, &part{line: line[lastRight:lastLeft]})
			}
			results = append(results, &part{line: line[lastLeft:left[0]]})

			//			log.Println("match", line, left, right)
			idx := len(left) - 1
			p := &part{
				line: line[left[0]+1 : right[idx]],
			}
			if right[idx] != len(line) {
				// break into parts
				p.parts = breakParts(line[left[0]+1 : right[idx]])
			}

			results = append(results, p)

			lastRight = right[idx]
			left = []int{}
			right = []int{}
		}
	}
	if found {
		if lastRight != 0 {
			results = append(results, &part{line: line[lastRight+1:]})
		}
	} else {
		results = append(results, &part{line: line})
	}

	// cleanup...
	var ret []*part
	for _, r := range results {
		//		log.Println(r.line)
		r.line = strings.TrimSpace(strings.TrimPrefix(r.line, ")"))
		//		log.Println(r.line)
		if r.line == "" {
			continue
		}
		last := len(r.line) - 1
		//		log.Println(r.line, len(r.line))

		if (r.line[last] == '+' || r.line[last] == '*') && len(r.line) != 1 {
			// if !unicode.IsDigit(rune(r.line[last])) {
			old := r.line
			r.line = r.line[0 : last-1]
			ret = append(ret, r)
			ret = append(ret, &part{line: old[last:]})
		} else {
			ret = append(ret, r)
		}
	}
	return ret
}

func parseArith(line string) int {

	parts := breakParts(line)
	//	spew.Dump(len(parts), parts)
	p := &part{parts: parts}

	return p.calculate(0)
}

func day18() {
	lines := readFile("day18input")

	sum := 0
	for _, line := range lines {

		// parts := strings.Split(line, " ")
		// spew.Dump(parts)
		//		log.Println(line)
		ans := parseArith(line)
		//		log.Println(line, ans)
		sum += ans
	}
	log.Println(sum)

}
