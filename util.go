package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func readFile(path string) []string {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	return strings.Split(string(b), "\n")
}

func atoi(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		log.Fatal(err)
	}
	return i
}

func splitLength(str string, sep string, length int) []string {
	parts := strings.Split(str, sep)
	if len(parts) != length {
		log.Fatalf("length %s not as expected", str)
	}
	return parts
}
