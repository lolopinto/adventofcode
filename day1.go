package main

// https://adventofcode.com/2020/day/1
import (
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func day1() {
	b, err := ioutil.ReadFile("day1input")
	if err != nil {
		log.Fatal(err)
	}

	str := string(b)
	lines := strings.Split(str, "\n")
	m := make(map[int]bool)
	for _, line := range lines {
		i, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		m[i] = true
	}

	for num := range m {
		target := 2020 - num
		for num2 := range m {
			target2 := target - num2

			if m[target2] {
				log.Println(num, num2, target2)
				log.Println(num * num2 * target2)
				os.Exit(1)
			}
		}
	}

}
