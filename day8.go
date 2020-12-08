package main

import (
	"log"
	"strconv"
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
	parts := strings.Split(line, " ")
	if len(parts) != 2 {
		log.Fatalf("line %s not as expected", line)
	}
	var num int
	var err error
	if strings.HasPrefix(parts[1], "+") {
		num, err = strconv.Atoi(strings.TrimPrefix(parts[1], "+"))
		if err != nil {
			log.Fatalf(err.Error())
		}
	} else {
		num, err = strconv.Atoi(strings.TrimPrefix(parts[1], "-"))
		if err != nil {
			log.Fatalf(err.Error())
		}
		num = num * -1
	}
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

func (p *program) exec() {
	for {
		p.seenIns[p.index] = true
		p.commands[p.index].ins.exec(p)
		_, ok := p.seenIns[p.index]
		if ok {
			break
		}
	}
	log.Println(p.acc)
}

func day8() {
	lines := readFile("day8input")

	p := &program{}
	p.seenIns = make(map[int]bool)

	for _, line := range lines {
		p.parseInstructionLine(line)
	}
	//	spew.Dump(p)
	p.exec()
}
