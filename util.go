package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func readFile(path string) []string {
	b, err := os.ReadFile(path)
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

func itoa(i int) string {
	return fmt.Sprintf("%v", i)
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
	return min
}

func max(slice []int) int {
	max := math.MinInt32
	for _, val := range slice {
		if val > max {
			max = val
		}
	}
	return max
}

func min64(slice []int64) int64 {
	var min int64
	min = math.MaxInt64
	for _, val := range slice {
		if val < min {
			min = val
		}
	}
	return min
}

func max64(slice []int64) int64 {
	var max int64
	max = math.MinInt64
	for _, val := range slice {
		if val > max {
			max = val
		}
	}
	return max
}

func die(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// permutations of a string
func combos(s string) map[string]bool {
	ret := make(map[string]bool)
	adder := func(s2 string) {
		ret[s2] = true
		//		ret = append(ret, s2)
	}
	combo_helper([]rune(s), adder, 0)
	return ret
}

func combo_helper(r []rune, adder func(s string), i int) {
	if i > len(r) {
		adder(string(r))
		return
	}
	combo_helper(r, adder, i+1)
	for j := i + i; j < len(r); j++ {
		r[i], r[j] = r[j], r[i]
		combo_helper(r, adder, i+1)
		r[i], r[j] = r[j], r[i]
	}
}

func abs(i, j int) int {
	return int(math.Abs(float64(i) - float64(j)))
}

func convertToBinary(line string) int {
	res := make([]int, len(line))
	for i, c := range line {
		if c == '1' {
			res[i] = 1
		} else if c == '0' {
			res[i] = 0
		} else {
			panic(fmt.Errorf("invalid value %v", c))
		}
	}
	return binary(res)
}

func binary(list []int) int {
	sum := 0
	for i, v := range list {
		pow := len(list) - i - 1
		if v == 1 {
			sum += int(math.Pow(2, float64(pow)))
		}
	}
	return sum
}
