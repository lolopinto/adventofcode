package main

import (
	"log"
	"strings"
)

type command struct {
	line string
	ins  instruction
}

type program struct {
	commands []*command
	acc      int //also start with 0
	index    int // start with 0
	seenIns  map[int]bool
}

type instruction interface {
	exec(p *program)
}

type noOp struct{}

func (op *noOp) exec(p *program) {
	p.index++
}

type acc struct {
	num int
}

func (op *acc) exec(p *program) {
	p.acc += op.num
	p.index++
}

type jump struct {
	num int
}

func (op *jump) exec(p *program) {
	p.index += op.num
}

func (p *program) parseInstructionLine(line string) {
	parts := splitLength(line, " ", 2)
	num := atoi(parts[1])
	var i instruction
	switch parts[0] {
	case "nop":
		i = &noOp{}
		break
	case "acc":
		i = &acc{num: num}
		break
	case "jmp":
		i = &jump{num: num}
		break
	default:
		log.Fatalf("invalid instruction")
	}

	p.commands = append(p.commands, &command{
		line: line,
		ins:  i,
	})
}

func (p *program) execBrokenProg() bool {
	for {
		p.seenIns[p.index] = true
		p.commands[p.index].ins.exec(p)
		// executed correctly
		if p.index == len(p.commands) {
			return false
		}
		_, ok := p.seenIns[p.index]
		if ok {
			return true
		}
	}
}

func day8() {
	lines := readFile("day8input")

	for idx, line := range lines {
		if strings.HasPrefix(line, "acc") {
			continue
		}

		// make copy of program
		lines2 := make([]string, len(lines))
		copy(lines2, lines)
		line2 := strings.Replace(line, "nop", "jmp", 1)
		line2 = strings.Replace(line2, "jmp", "nop", 1)

		// replace one line
		lines2[idx] = line2

		p := &program{}
		p.seenIns = make(map[int]bool)

		for _, line3 := range lines2 {
			p.parseInstructionLine(line3)
		}

		if !p.execBrokenProg() {
			log.Println(p.acc)
			break
		}
	}
}
