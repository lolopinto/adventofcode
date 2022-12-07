package main

import (
	"fmt"
	"strings"
)

type dir struct {
	dirs   []*dir
	files  []*file
	name   string
	root   bool
	parent *dir
}

func (d *dir) size() int {
	sum := 0
	for _, d := range d.dirs {
		sum += d.size()
	}
	for _, f := range d.files {
		sum += f.size
	}
	return sum
}

type file struct {
	size int
	name string
}

func day7() {
	lines := readFile("day7input")
	root := &dir{
		name: "/",
		root: true,
	}
	i := 0
	curr := root
	for i < len(lines) {
		// line := lines[i]
		next, idx := parseCommand(lines, i, curr)
		// fmt.Println("return", next, idx)
		i = idx
		curr = next
	}

	valid := []*dir{}
	checkAllDirs(root, &valid)
	// fmt.Println(root)
	// fmt.Println(root.size())
	total := 0
	for _, d := range valid {
		total += d.size()
	}
	fmt.Println(total)
}

func checkAllDirs(d *dir, valid *[]*dir) {
	if d.size() <= 100000 {
		*valid = append(*valid, d)
	}
	for _, d2 := range d.dirs {
		checkAllDirs(d2, valid)
	}
}

// return current and next index
func parseCommand(lines []string, idx int, curr *dir) (*dir, int) {
	line := lines[idx]

	// fmt.Println("line", line)
	if !strings.HasPrefix(line, "$") {
		panic(fmt.Sprintf("invalid line %s", line))
	}
	cmd := line[2:4]
	switch cmd {
	case "cd":
		to_cd := line[5:]
		switch to_cd {
		case "/":
			// nothing to do (for now). go to next line
			return curr, idx + 1

		case "..":
			if curr.parent == nil {
				panic("cannot cd to parent")
			}
			return curr.parent, idx + 1

		default:
			for _, child := range curr.dirs {
				if child.name == to_cd {
					return child, idx + 1
				}
			}
			panic(fmt.Sprintf("couldn't find dir %s to cd", to_cd))
		}

	case "ls":
		for i := idx + 1; i < len(lines); i++ {
			line := lines[i]
			if strings.HasPrefix(line, "$") {
				return curr, i
			}
			parts := strings.Split(line, " ")

			if parts[0] == "dir" {
				curr.dirs = append(curr.dirs, &dir{
					name:   parts[1],
					parent: curr,
				})
			} else {
				curr.files = append(curr.files, &file{
					name: parts[1],
					size: atoi(parts[0]),
				})
			}

		}

	default:
		panic(fmt.Sprintf("invalid cmd %s", cmd))
	}

	// TODO
	return nil, len(lines)
}
