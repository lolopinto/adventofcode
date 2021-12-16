package main

import (
	"fmt"
	"strings"
)

func day16() {
	m := map[rune]string{
		'0': "0000",
		'1': "0001",
		'2': "0010",
		'3': "0011",
		'4': "0100",
		'5': "0101",
		'6': "0110",
		'7': "0111",
		'8': "1000",
		'9': "1001",
		'A': "1010",
		'B': "1011",
		'C': "1100",
		'D': "1101",
		'E': "1110",
		'F': "1111",
	}

	lines := readFile("day16input")

	var sb strings.Builder
	for _, c := range lines[0] {
		sb.WriteString(m[c])
	}
	p := parsePacket(sb.String(), m)
	// part 1
	fmt.Println(sumVersions(p))
	// part 2
	fmt.Println(p.value())
}

func sumVersions(p *packet) int {
	if p == nil {
		return 0
	}
	sum := p.version
	for _, pp := range p.packets {
		sum += sumVersions(pp)
	}
	return sum
}

type packet struct {
	version       int
	id            int
	literal       int
	packets       []*packet
	initialString string
	endIdx        int
}

func (p *packet) value() int {
	switch p.id {
	case 0:
		sum := 0
		for _, pp := range p.packets {
			sum += pp.value()
		}
		return sum
	case 1:
		mult := 1
		for _, pp := range p.packets {
			mult *= pp.value()
		}
		return mult
	case 2:
		vals := make([]int, len(p.packets))
		for i, pp := range p.packets {
			vals[i] = pp.value()
		}
		return min(vals)
	case 3:
		vals := make([]int, len(p.packets))
		for i, pp := range p.packets {
			vals[i] = pp.value()
		}
		return max(vals)
	case 4:
		return p.literal
	case 5:
		if len(p.packets) != 2 {
			panic("invalid packet")
		}
		if p.packets[0].value() > p.packets[1].value() {
			return 1
		}
		return 0
	case 6:
		if len(p.packets) != 2 {
			panic("invalid packet")
		}
		if p.packets[0].value() < p.packets[1].value() {
			return 1
		}
		return 0
	case 7:
		if len(p.packets) != 2 {
			panic("invalid packet")
		}
		if p.packets[0].value() == p.packets[1].value() {
			return 1
		}
		return 0
	}
	panic("invalid packet id")
}

func parsePacket(str string, m map[rune]string) *packet {
	//	fmt.Println("packet", str)
	//	fmt.Println(len(str))
	ret := &packet{
		version:       convertToBinary(str[0:3]),
		id:            convertToBinary(str[3:6]),
		initialString: str,
	}

	// literal
	if ret.id == 4 {
		i := 6
		var sb strings.Builder
		for i < len(str) {
			first := rune(str[i])
			sb.WriteString(str[i+1 : i+5])
			if first == '0' {
				break
			}
			i += 5
		}
		ret.literal = convertToBinary(sb.String())
		ret.endIdx = i + 5
	} else {
		//		fmt.Println("id", rune(str[6]))
		// length type id
		if rune(str[6]) == '0' {
			length := convertToBinary(str[7:22])
			sub := str[22 : 22+length]
			ret.endIdx = 22 + length

			for {
				if convertToBinary(sub) == 0 {
					break
				}
				if checkIfLiteral(sub) {
					p := parsePacket(sub, m)
					ret.packets = append(ret.packets, p)
					sub = sub[p.endIdx:]
				} else {
					p := parsePacket(sub, m)
					if p == nil {
						break
					}
					ret.packets = append(ret.packets, p)
					sub = sub[p.endIdx:]
				}
			}

			// 15
		} else {

			num := convertToBinary(str[7:18])

			start := str[18:]
			ret.endIdx = 18
			for i := 0; i < num; i++ {

				if checkIfLiteral(start) {
					p := parsePacket(start, m)
					ret.endIdx += p.endIdx
					ret.packets = append(ret.packets, p)
					start = start[p.endIdx:]

				} else {
					p := parsePacket(start, m)
					ret.endIdx += p.endIdx

					ret.packets = append(ret.packets, p)
					if p.endIdx == 0 {
						panic("fail no ednIdx")
					}
					start = start[p.endIdx:]
				}
			}
		}
	}

	return ret
}

func checkIfLiteral(str string) bool {
	return convertToBinary(str[3:6]) == 4
}
