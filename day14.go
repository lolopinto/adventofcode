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
		if strings.HasPrefix(line, "mask") {
			parts := splitLength(line, "=", 2)
			mask = strings.TrimSpace(parts[1])
			//			log.Println(mask)
		} else {
			parts := splitLength(line, "=", 2)
			decimal := atoi(strings.TrimSpace(parts[1]))
			idx := strings.Index(parts[0], "]")
			address := atoi(strings.TrimSpace(parts[0][4:idx]))

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

			// log.Println("floating", floating)
			// log.Println("num", numFloating)
			// log.Println("result", result)
			for i := 0; i < numFloating; i++ {
				str := strconv.FormatInt(int64(i), 2)
				if len(str) < len(floating) {
					var sb strings.Builder
					for j := 0; j < len(floating)-len(str); j++ {
						sb.WriteString("0")
					}
					sb.WriteString(str)
					str = sb.String()
				}
				result2 := make(map[int]byte)
				for k, v := range result {
					result2[k] = v
				}
				//				log.Println("str", str)
				//				log.Println("floating", floating, str)
				for j := 0; j < len(floating); j++ {
					if str[j] == '1' {
						result2[floating[j]] = 1
					}
				}
				//				log.Println("result2", result2)
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
