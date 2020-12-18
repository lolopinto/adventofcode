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
		// do itself
	}
	// has parts use that
	last := 0
	result := 0
	for i := 0; i < len(p.parts); i++ {
		p2 := p.parts[i]
		// if len(p2.parts) != 0 {
		// 	result = p2.calculate(result)
		// 	continue
		// }
		//		line := strings.TrimSpace(strings.TrimPrefix(p.line, ")"))
		//		log.Println(p2.line)
		// if p2.line == "*" 	|| p2.line == "+" {
		// 	// TODO use last one and next one
		// }
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

			//result = p2.calculate(prev)
		}
		// result
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
					// how??
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
	for {
		//		log.Println(line)
		//line = strings.TrimPrefix(line, "(")
		idx := strings.Index(line, "(")
		lastIdx := strings.LastIndex(line, "(")
		// if idx == 0 {
		// 	idx = strings.Index(line[1:], "(")
		// }
		//		log.Println(idx)
		if idx == -1 {
			break
		}
		//		log.Println(line[0:idx])
		p := &part{line: line[0:idx]}
		//		log.Println(p.line)
		//		if checkForParts(p) {
		if lastIdx != idx {
			parts := breakParts(line[lastIdx:])

			//		if !(len(parts) == 1 && parts[0].line == p.line) {
			p.parts = parts
		}

		results = append(results, p)

		//		results = append(results, &part{[0:idx])
		//		log.Println(line)
		idx2 := strings.LastIndex(line, ")")
		if idx2 == -1 {
			log.Fatalf("shouldn't happen")
		}
		line = line[idx : idx2+1]
		if idx2 == len(line)-1 {
			break
		}
	}
	p := &part{line: line}
	//	log.Println(p.line)
	idx := strings.Index(p.line, "(")
	if idx != 0 && idx != -1 {
		//if checkForParts(p) {
		p.parts = breakParts(p.line)
	}
	results = append(results, p)
	return results
}

func breakParts2(line string) []*part {
	//	log.Println("line", line)
	//	var pairs [][2]int
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
			//			if left[0] != 0 {
			if lastLeft != lastRight {
				results = append(results, &part{line: line[lastRight:lastLeft]})
			}
			results = append(results, &part{line: line[lastLeft:left[0]]})
			//			}

			//			log.Println("match", line, left, right)
			idx := len(left) - 1
			p := &part{
				line: line[left[0]+1 : right[idx]],
			}
			if right[idx] != len(line) {
				// break into parts
				p.parts = breakParts2(line[left[0]+1 : right[idx]])
			}

			results = append(results, p)
			// results = append(results, &part
			// 	[2]int{left[0], right[idx]})
			// breakParts2(line[left[0]:right[idx]])
			// log.Println(pairs)

			//log.Println(right)
			if right[idx] < len(line) {
				//				results = append(results, breakParts2(line[right[idx]+1:])...)
			}
			// top level, everything below is a sub

			// if len(left) != 1 {

			// }
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

	parts := breakParts2(line)
	//	spew.Dump(len(parts), parts)
	p := &part{parts: parts}

	return p.calculate(0)
	// prev := 0
	// for _, part := range parts {
	// 	prev = part.calculate(prev)
	// }

	// return prev
	// var allResult int
	// var carryOverOp string

	// for i, p := range parts {
	// 	if p.line == "" {
	// 		continue
	// 	}
	// 	ops := strings.Split(strings.TrimSpace(p.line), " ")
	// 	var lastOp string
	// 	var result int

	// 	var left int
	// 	//	result = 0
	// 	lastOp = ""
	// 	//		var carryOverOp string
	// 	//		spew.Dump(ops)
	// 	for idx, op := range ops {
	// 		switch op {
	// 		case "*":
	// 			lastOp = "*"
	// 			if result == 0 && idx == len(ops)-1 {
	// 				result = left
	// 			}
	// 			break
	// 		case "+":
	// 			lastOp = "+"
	// 			if result == 0 && idx == len(ops)-1 {
	// 				result = left
	// 			}
	// 			break

	// 		default:

	// 			if lastOp == "" {
	// 				left = atoi(op)
	// 				//					log.Println("left", left)
	// 			} else if lastOp == "*" {
	// 				if result == 0 {
	// 					result = left * atoi(op)
	// 					//						log.Println("result", result, lastOp, op)

	// 					lastOp = ""
	// 				} else {
	// 					result = result * atoi(op)
	// 					//						log.Println("result", result, lastOp, op)
	// 					lastOp = ""
	// 				}
	// 			} else if lastOp == "+" {
	// 				result += atoi(op)
	// 				//					log.Println("result", result, lastOp, op)

	// 				lastOp = ""
	// 			}
	// 		}
	// 		//			log.Println("result", result, lastOp, op)
	// 	}
	// 	for _, p2 := range p.parts {
	// 		r2 := parseArith(p2.line)
	// 		log.Println("sub part", r2)
	// 	}
	// 	if i == 0 {
	// 		allResult = result
	// 	} else {
	// 		if carryOverOp == "*" {
	// 			if allResult == 0 {
	// 				allResult = result * 1
	// 			} else {
	// 				allResult = result * allResult
	// 			}
	// 		} else if carryOverOp == "+" {
	// 			allResult = result + allResult
	// 		}
	// 	}
	// 	l := len(ops)
	// 	if ops[l-1] == "*" || ops[l-1] == "+" {
	// 		carryOverOp = lastOp
	// 	}
	// 	//		log.Println("all", allResult, carryOverOp)
	// 	//		log.Println(result, lastOp)
	// }
	// //	spew.Dump(parts)
	// //	log.Println(allResult)
	// return allResult
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
