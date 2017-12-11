package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"strconv"
	"strings"
)

func main() {
	seq := buildInstructionSlice()
	recordedSeqs := make([][]int, 0)
	cop := make([]int, len(seq))
	copy(cop, seq)
	recordedSeqs = append(recordedSeqs, cop)
	cycles := 0
	overflow := false
	for overflow == false {
		cycles++
		highestIndex := findHighestIndex(seq)
		newSeq := redistribute(seq, highestIndex)
		if recordHasMatch(recordedSeqs, seq) {
			cop = make([]int, len(seq))
			copy(cop, newSeq)
			recordedSeqs = append(recordedSeqs, newSeq)
			fmt.Println(recordedSeqs)
			fmt.Println(cycles)
			overflow = true
		}
		cop = make([]int, len(seq))
		copy(cop, newSeq)
		recordedSeqs = append(recordedSeqs, newSeq)
		seq = newSeq
	}
}

func chk(err error) {
	if err != nil {
		fmt.Printf("error: %v\n", err)
		os.Exit(1)
	}
}

func recordHasMatch(recordedSeqs [][]int, seq []int) bool {
	for _, val := range recordedSeqs {
		if reflect.DeepEqual(seq, val) {
			return true
		}
	}
	return false
}

func findHighestIndex(seq []int) int {
	cur := 0
	for i, val := range seq {
		if val > seq[0] {
			cur = i
		}
	}
	return cur
}

func redistribute(seq []int, index int) []int {
	bag := seq[index]
	seq[index] = 0
	distIndex := index
	for bag > 0 {
		distIndex++
		if distIndex > (len(seq) - 1) {
			distIndex = 0
		}
		seq[distIndex] += 1
		bag -= 1
	}
	return seq
}

func buildInstructionSlice() []int {
	pwd, err := os.Getwd()
	chk(err)
	datBytes, err := ioutil.ReadFile(pwd + "/input.txt")
	chk(err)
	dat := string(datBytes)
	lines := strings.Split(dat, "\n")
	lines = strings.Split(lines[0], "\t")
	nums := make([]int, 0)
	for _, line := range lines {
		if line != "" {
			num, err := strconv.Atoi(line)
			chk(err)
			nums = append(nums, num)
		}
	}
	return nums
}
