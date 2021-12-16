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
	fmt.Println(sumVersions(p))
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
	literal       bool
	literalvalue  int
	operator      bool
	packets       []*packet
	initialString string
	endIdx        int
}

func parsePacket(str string, m map[rune]string) *packet {
	//	fmt.Println("packet", str)
	//	fmt.Println(len(str))
	ret := &packet{
		version:       convertToBinary(str[0:3]),
		id:            convertToBinary(str[3:6]),
		initialString: str,
	}
	//	fmt.Println("start version", ret.version, "id", ret.id, "literal", ret.literalvalue)

	// literal
	if ret.id == 4 {
		i := 6
		ret.literal = true
		var sb strings.Builder
		for i < len(str) {
			first := rune(str[i])
			sb.WriteString(str[i+1 : i+5])
			if first == '0' {
				break
			}
			i += 5
		}
		ret.literalvalue = convertToBinary(sb.String())
		ret.endIdx = i + 5
	} else {
		ret.operator = true
		//		fmt.Println("id", rune(str[6]))
		// length type id
		if rune(str[6]) == '0' {
			//			fmt.Println("zero", str, len(str))
			// length of subpackets
			//			if end
			//			fmt.Println("0 len str", len(str), str)
			length := convertToBinary(str[7:22])
			//			fmt.Println("length", length, str, str[7:22])
			sub := str[22 : 22+length]
			ret.endIdx = 22 + length

			for {
				//				fmt.Println("sub", sub)
				if convertToBinary(sub) == 0 {
					break
				}
				if checkIfLiteral(sub) {
					p := parsePacket(sub[0:11], m)
					ret.packets = append(ret.packets, p)
					sub = sub[11:]
				} else {
					p := parsePacket(sub, m)
					//					fmt.Println(sub)
					if p == nil {
						break
					}
					ret.packets = append(ret.packets, p)
					//					fmt.Println("no literal", sub, p.endIdx)
					sub = sub[p.endIdx:]
				}
			}

			// 15
		} else {

			num := convertToBinary(str[7:18])

			//			fmt.Println("length to be subdivided", len(str[18:]), str[18:])
			start := str[18:]
			ret.endIdx = 18
			for i := 0; i < num; i++ {
				// need to return remnant from parsing...
				if checkIfLiteral(start) {
					p := parsePacket(start[0:11], m)
					ret.endIdx += p.endIdx
					ret.packets = append(ret.packets, p)
					start = start[11:]

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
