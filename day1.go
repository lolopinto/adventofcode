package main

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

	for key := range m {
		val := 2020 - key

		if !m[val] {
			continue
		}

		log.Println(key * val)
		os.Exit(1)
	}

}
