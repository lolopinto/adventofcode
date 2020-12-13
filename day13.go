package main

import (
	"log"
	"strings"
)

type datapoint struct {
	id int
	t  int
}

func day13() {
	lines := readFile("day13input")
	timestamp := atoi(lines[0])

	parts := strings.Split(lines[1], ",")
	var data []datapoint
	for i, part := range parts {
		if part == "x" {
			continue
		}

		data = append(data, datapoint{
			id: atoi(part),
			t:  i,
		})
	}

	first := data[0]
	ts := first.id

	// find minimum value
	var earliest int

	idx := 0
	for i := idx; true; i++ {
		if ts*i > timestamp {
			earliest = ts * i
			break
		}
	}

	for {
		correct := true
		for _, datum := range data {
			if (earliest+datum.t)%datum.id != 0 {
				correct = false
				break
			}
		}
		if correct {
			log.Println(earliest)
			break
		}
		idx++
		earliest = ts * idx
	}

	//		idx++
	//	}
	//		}
	// sort.Slice(earliest, func(i, j int) bool {
	// 	return earliest[i][0] < earliest[j][0]
	// })
	// //	log.Println(ids)
	// //	log.Println(earliest)
	// log.Println((earliest[0][0] - timestamp) * earliest[0][1])

}
