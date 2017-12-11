package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	sl := buildInstructionSlice()
	val := findOutOfBoundsInstructionSteps(sl)
	fmt.Println(val)
}

func chk(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func findOutOfBoundsInstructionSteps(inst []int) int {
	steps := 0
	found := false
	curI := 0
	nextI := 0
	for !found {
		steps += 1
		nextI = curI + inst[curI]
		if (nextI < 0) || (nextI >= len(inst)) {
			found = true
		}
		if inst[curI] > 2 {
			inst[curI] -= 1
		} else {
			inst[curI] += 1
		}
		curI = nextI
	}
	return steps
}

func buildInstructionSlice() []int {
	pwd, err := os.Getwd()
	chk(err)
	datBytes, err := ioutil.ReadFile(pwd + "/input.txt")
	chk(err)
	dat := string(datBytes)
	lines := strings.Split(dat, "\n")
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
