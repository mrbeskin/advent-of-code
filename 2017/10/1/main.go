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
	fmt.Println(array)
	for _, val := range steps {
		fmt.Println(curIndex)
		reverseSubArray(array, curIndex, val)
		curIndex = curIndex + skipLength + val
		if curIndex >= length {
			curIndex = curIndex - length
		}
		skipLength++
	}

	fmt.Println(array[0], array[1])
}

func reverseSubArray(s []int, start int, offset int) {
	fmt.Println(offset)
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
	fmt.Println(s)
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
	steps := []int{157, 222, 1, 2, 177, 254, 0, 228, 159, 140, 249, 187, 255, 51, 76, 30}
	return steps
}
