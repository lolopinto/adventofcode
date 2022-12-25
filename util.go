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
	if length != -1 {
		if len(parts) != length {
			log.Fatal("unexpected length")
		}
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

func atoi64(str string) int64 {
	i, err := strconv.Atoi(str)
	if err != nil {
		log.Fatal(err)
	}
	return int64(i)
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

// TODO comparable
func uniq(slice []int) []int {
	m := make(map[int]bool)
	ret := []int{}
	for _, v := range slice {
		if m[v] {
			continue
		}
		ret = append(ret, v)
		m[v] = true
	}
	return ret
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

func abs64(i, j int64) int64 {
	return int64(math.Abs(float64(i) - float64(j)))
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

func groupLines(lines []string, by int) [][]string {
	var ret [][]string
	for i := 0; i < len(lines); i += by {
		ret = append(ret, lines[i:i+by])
	}
	return ret
}

func windowed[T any](list []T, n int) [][]T {
	ret := [][]T{}
	for i := 0; i < len(list)-n+1; i++ {
		ret = append(ret, list[i:i+n])
	}
	return ret
}

func mapify[T comparable](list []T) map[T]bool {
	m := make(map[T]bool, len(list))
	for _, v := range list {
		m[v] = true
	}
	return m
}

// insert in a slice at pos
func insert[T any](slice []T, idx int, val T) []T {
	if len(slice) == idx {
		// empty or appending to end
		return append(slice, val)
	}

	slice = append(slice[:idx+1], slice[idx:]...)
	slice[idx] = val
	return slice
}

// note if using this while iterating through a slice, do idx-- afterwards
func remove[T any](slice []T, idx int) []T {
	if idx+1 == len(slice) {
		return slice[:idx]
	}
	return append(slice[:idx], slice[idx+1:]...)
}

func replaceInString(s string, idx int, v rune) string {
	return s[:idx] + string(v) + s[idx+1:]
}

func validate(v bool, s string, args ...any) {
	if !v {
		panic(fmt.Errorf("%s %v", s, args))
	}
}

func copyMap[K comparable, V any](m map[K]V) map[K]V {
	ret := make(map[K]V, len(m))
	for k, v := range m {
		ret[k] = v
	}
	return ret
}

func copyList[T any](l []T) []T {
	ret := make([]T, len(l))
	copy(ret, l)
	return ret
}

func keys[K comparable, V any](m map[K]V) []K {
	ret := make([]K, len(m))
	i := 0
	for k := range m {
		ret[i] = k
		i++
	}
	return ret
}

// permutaations with repetitions until we find one
// didn't end up using it as it's 20! so doesn't work for 2022 day 25 input lol
func permutationsWithRepeat[T any](list []T, n int, decide func(l []T) bool) {
	pn := make([]int, n)
	p := make([]T, n)
	k := len(list)
	for {
		// generate permutaton
		for i, x := range pn {
			p[i] = list[x]
		}
		// show progress
		// pass to deciding function
		if decide(p) {
			return // terminate early
		}
		// increment permutation number
		for i := 0; ; {
			pn[i]++
			if pn[i] < k {
				break
			}
			pn[i] = 0
			i++
			if i == n {
				return // all permutations generated
			}
		}
	}
}
