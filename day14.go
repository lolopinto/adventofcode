package main

import (
	"log"
	"math"
	"strings"
)

func day14() {
	lines := readFile("day14input")

	m := make(map[int]int)
	var mask string
	for _, line := range lines {
		if strings.HasPrefix(line, "mask") {
			parts := splitLength(line, "=", 2)
			mask = strings.TrimSpace(parts[1])
			log.Println(mask)
		} else {
			parts := splitLength(line, "=", 2)
			decimal := atoi(strings.TrimSpace(parts[1]))
			idx := strings.Index(parts[0], "]")
			address := atoi(strings.TrimSpace(parts[0][4:idx]))
			//			log.Println("address", address)
			result := 0

			for i := 35; i >= 0; i-- {
				pow := int(math.Pow(2, float64(i)))
				switch mask[35-i] {
				case 'X':
					// leave bit unchanged
					//					log.Println("idx", i)
					if pow <= decimal {
						result |= pow
						decimal = decimal - pow
					}
					break
				case '1':
					// mask and change the bit
					result |= pow
					if pow <= decimal {
						decimal = decimal - pow
					}
					break
				case '0':
					if pow <= decimal {
						decimal = decimal - pow
					}
					break
				}
				// //				if mask[35-i] == ''
				// if mask[35-i] == '1' {
				// 	log.Println("mask ", i)
				// 	result |= pow
				// 	decimal = decimal - pow
				// 	continue
				// }
				// if pow > decimal {
				// 	continue
				// }
				// if mask[i] == '0' {
				// 	continue
				// }
				// result |= pow
				// log.Println("idx", i)
				// decimal = decimal - pow
				// log.Println("result", result)
			}
			if result != 0 {
				m[address] = result
			}
		}
	}

	sum := 0
	//	log.Println(m)
	for _, v := range m {
		sum += v
	}
	log.Println("sum", sum)
}
