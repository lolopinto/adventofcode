package main

import (
	"io/ioutil"
	"log"
	"math"
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

// read chunks separated by \n
func readFileChunks(path string, length int) [][]string {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	parts := strings.Split(string(b), "\n\n")
	if len(parts) != length {
		log.Fatal("unexpected length")
	}

	result := make([][]string, len(parts))
	for idx, part := range parts {
		result[idx] = strings.Split(part, "\n")
	}
	return result
}

func atoi(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		log.Fatal(err)
	}
	return i
}

func ints(lines []string) []int {
	numbers := make([]int, len(lines))

	for i, line := range lines {
		numbers[i] = atoi(line)
	}
	return numbers
}

func readInts(file string) []int {
	lines := readFile(file)
	return ints(lines)
}

func splitLength(str string, sep string, length int) []string {
	parts := strings.Split(str, sep)
	if len(parts) != length {
		log.Fatalf("length %s not as expected", str)
	}
	return parts
}

func leftPad(str string, c string, desiredLength int) string {
	if len(str) >= desiredLength {
		return str
	}
	if len(c) != 1 {
		log.Fatalf("can only left pad with 1 character")
	}
	var sb strings.Builder
	for j := 0; j < desiredLength-len(str); j++ {
		sb.WriteString(c)
	}
	sb.WriteString(str)
	return sb.String()
}

func rightPad(str string, c string, desiredLength int) string {
	if len(str) >= desiredLength {
		return str
	}
	if len(c) != 1 {
		log.Fatalf("can only right pad with 1 character")
	}
	var sb strings.Builder
	sb.WriteString(str)
	for j := 0; j < desiredLength-len(str); j++ {
		sb.WriteString(c)
	}
	return sb.String()
}

func min(slice []int) int {
	min := math.MaxInt32
	for _, val := range slice {
		if val < min {
			min = val
		}
	}
	//log.Println("min", min)
	return min
}

func max(slice []int) int {
	max := math.MinInt32
	for _, val := range slice {
		if val > max {
			max = val
		}
	}
	//	log.Println("max", max)

	return max
}

func die(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
