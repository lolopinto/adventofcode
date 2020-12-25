package main

import "log"

func day25() {
	lines := readFile("day25input")

	pkey1 := atoi(lines[0])
	pkey2 := atoi(lines[1])

	//loopSize1 := calcLoopSize((pkey1))
	loopSize2 := calcLoopSize((pkey2))
	log.Println(calcEncryptionkey(pkey1, loopSize2))
}

func calcLoopSize(num int) int {
	val := 1
	for i := 1; true; i++ {
		val = val * 7
		val = val % 20201227

		//		log.Println(val)
		if val == num {
			//			log.Println(i)
			return i
		}
	}
	return -1
}

func calcEncryptionkey(num, loopSize int) int {
	val := 1
	for i := 1; i <= loopSize; i++ {
		val = val * num
		val = val % 20201227
	}
	return val
}
