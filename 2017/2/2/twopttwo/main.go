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
	for r, row := range nums {
		found := false
		for i, num := range row {
			if found {
				break
			}
			for y, num2 := range row {
				if found == true {
					break
				}
				if y != i {
					if num >= num2 {
						if num%num2 == 0 {
							fmt.Printf("row:%d\nnum:%d\nnum2:%d\n", r, num, num2)
							found = true
							sum = sum + (num / num2)
						}
					}
				}
			}
		}
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
