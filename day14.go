package main

import (
	"log"
	"math"
	"strconv"
	"strings"
)

func day14() {
	lines := readFile("day14input")

	m := make(map[int]int)
	var mask string
	for _, line := range lines {
		parts := splitLength(line, " = ", 2)
		if parts[0] == "mask" {
			mask = parts[1]
		} else {
			decimal := atoi(parts[1])
			idx := strings.Index(parts[0], "]")
			address := atoi(parts[0][4:idx])

			result := make(map[int]byte, 36)
			var floating []int

			for i := 35; i >= 0; i-- {
				pow := int(math.Pow(2, float64(i)))
				newIdx := 35 - i

				switch mask[newIdx] {
				case 'X':
					// floating
					result[newIdx] = 'X'
					floating = append(floating, newIdx)

					break
				case '1':
					// mask and change the bit
					result[newIdx] = 1
					break
				case '0':
					// leave bit unchanged
					if pow <= address {
						// bit is set
						result[newIdx] = 1
					}
					break
				}
				if pow <= address {
					address = address - pow
				}
			}

			numFloating := int(math.Pow(2, float64(len(floating))))

			for i := 0; i < numFloating; i++ {
				// get from 000 to 100(and more) of floating bits
				str := strconv.FormatInt(int64(i), 2)
				str = leftPad(str, "0", len(floating))

				// make copy of map
				result2 := make(map[int]byte)
				for k, v := range result {
					result2[k] = v
				}

				// set whatever floating values that are 1
				for j := 0; j < len(floating); j++ {
					if str[j] == '1' {
						result2[floating[j]] = 1
					}
				}
				// convert to a number
				number := 0
				for k, v := range result2 {
					if v == 1 {
						number += int(math.Pow(2, float64(35-k)))
					}
				}
				//				log.Println("number", number)
				if number != 0 {
					m[number] = decimal
				}
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
