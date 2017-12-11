package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	nums := readNumbersIntoTwoDimensionalSlice()
	sum := 0
	for _, row := range nums {
		min := row[0]
		max := row[0]
		for _, num := range row {
			if num > max {
				max = num
			}
			if num < min {
				min = num
			}
		}
		sum = sum + (max - min)
	}
	fmt.Println(sum)
}

func readNumbersIntoTwoDimensionalSlice() [][]int {
	pwd, err := os.Getwd()
	chk(err)
	datBytes, err := ioutil.ReadFile(pwd + "/input.txt")
	chk(err)
	dat := string(datBytes)
	lines := strings.Split(dat, "\n")
	nums := [][]int{}
	for _, line := range lines {
		if line != "" {
			row := []int{}
			numStrings := strings.Split(line, "\t")
			chk(err)
			for _, numString := range numStrings {
				if numString != "" {
					num, err := strconv.Atoi(numString)
					chk(err)
					row = append(row, num)
				}
			}
			nums = append(nums, row)
		}
	}
	return nums
}

func chk(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
