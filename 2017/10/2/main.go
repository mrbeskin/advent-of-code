package main

import (
	"fmt"
)

func main() {
	steps := getSteps()
	array := generateDefaultArray()
	curIndex := 0
	skipLength := 0
	length := 256
	for i := 0; i < 64; i++ {
		for _, val := range steps {
			reverseSubArray(array, curIndex, val)
			curIndex = curIndex + skipLength + val
			for curIndex >= length {
				curIndex = curIndex - length
			}
			skipLength++
		}
	}
	dense := denseHash(array)
	hexString := ""
	for _, val := range dense {
		hexVal := fmt.Sprintf("%x", val)
		fmt.Println(hexVal)
		hexString += hexVal
	}
	fmt.Println(hexString)
}

func denseHash(in []int) []int {
	dense := make([]int, 16, 16)
	prev := 0
	denseIndex := 0
	for i, val := range in {
		if (i%16 == 0) && (i > 0) {
			dense[denseIndex] = prev
			fmt.Println(prev)
			denseIndex++
			prev = 0
		}
		prev = prev ^ val
	}
	dense[denseIndex] = prev
	return dense
}

func reverseSubArray(s []int, start int, offset int) {
	subArray := make([]int, offset, offset)
	pos := start
	for i := 0; i < offset; i++ {
		subArray[i] = s[pos]
		pos++
		if pos == len(s) {
			pos = 0
		}

	}
	reverse(subArray)
	pos = start
	for _, val := range subArray {
		s[pos] = val
		pos++
		if pos == len(s) {
			pos = 0
		}
	}
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func generateDefaultArray() []int {
	def := make([]int, 256, 256)
	for i, _ := range def {
		def[i] = i
	}
	return def
}

func getSteps() []int {
	steps := "157,222,1,2,177,254,0,228,159,140,249,187,255,51,76,30"
	stepsBytes := []byte(steps)
	s := make([]int, len(stepsBytes))
	for i, b := range stepsBytes {
		s[i] = int(b)
	}
	end := []int{17, 31, 73, 47, 23}
	s = append(s, end...)
	return s
}
