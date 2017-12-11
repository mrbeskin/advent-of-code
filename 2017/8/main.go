package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

const (
	ifTok = " if "
	lt    = "<"
	lte   = "<="
	gt    = ">"
	gte   = ">="
	ne    = "!="
	e     = "=="
	inc   = "inc"
	dec   = "dec"
)

var hiVal = 0

func main() {
	input := getInstructionSlice()
	registers := make(map[string]int)
	for _, line := range input {
		halves := strings.Split(line, ifTok)
		if len(halves) > 1 {
			if evaluate(halves[1], registers) {
				doRegisterOp(halves[0], registers)
			}
		}
	}
	fmt.Println(registers)
	fmt.Println(hiVal)
}

func evaluate(exp string, m map[string]int) bool {
	toks := strings.Split(exp, " ")
	lval := m[toks[0]]
	rval, err := strconv.Atoi(toks[2])
	opTok := toks[1]
	chk(err)
	return doOp(lval, rval, opTok)
}

func doRegisterOp(input string, m map[string]int) {
	toks := strings.Split(input, " ")
	reg := toks[0]
	instr := toks[1]
	val, err := strconv.Atoi(toks[2])
	chk(err)
	switch instr {
	case dec:
		m[reg] = m[reg] - val
	case inc:
		m[reg] = m[reg] + val
	default:
		fmt.Printf("invalid register op %v", instr)
		os.Exit(1)
	}
	if m[reg] > hiVal {
		hiVal = m[reg]
	}
}

func doOp(lval int, rval int, opTok string) bool {
	switch opTok {
	case lt:
		return lval < rval
	case lte:
		return lval <= rval
	case gt:
		return lval > rval
	case gte:
		return lval >= rval
	case ne:
		return lval != rval
	case e:
		return lval == rval
	default:
		fmt.Println("invalid opTok")
		os.Exit(1)
	}
	return false
}

func chk(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func getInstructionSlice() []string {
	pwd, err := os.Getwd()
	chk(err)
	datBytes, err := ioutil.ReadFile(pwd + "/input.txt")
	chk(err)
	dat := string(datBytes)
	lines := strings.Split(dat, "\n")
	return lines
}
