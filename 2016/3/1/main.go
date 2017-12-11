package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	numValid := 0
	tris := buildTriangleSlice()
	for _, t := range tris {
		if triangleStringIsValid(t) {
			numValid++
		}
	}
	fmt.Println(numValid)
}

func buildTriangleSlice() []string {
	pwd, err := os.Getwd()
	chk(err)
	datBytes, err := ioutil.ReadFile(pwd + "/input.txt")
	chk(err)
	dat := string(datBytes)
	triangles := strings.Split(dat, "\n")
	return triangles
}

func getTVals(t string) []int {
	dat := strings.Split(t, " ")
	tri := make([]int, 3, 3)
	triIndex := 0
	for _, str := range dat {
		if str != "" {
			val, err := strconv.Atoi(str)
			chk(err)
			tri[triIndex] = val
			triIndex++
		}
	}
	return tri
}

func validateTriangle(tri []int) bool {
	return (isValid(tri[0], tri[1], tri[2]) &&
		isValid(tri[0], tri[2], tri[3]) &&
		isValid(tri[1], tri[0], tri[2]) &&
		isValid(tri[1], tri[2], tri[0]) &&
		isValid(tri[2], tri[0], tri[1]) &&
		isValid(tri[2], tri[1], tri[0]))
}

func isValid(a int, b int, c int) bool {
	return (a + b) > c
}

func triangleStringIsValid(tString string) bool {
	t := getTVals(tString)
	return validateTriangle(t)
}

func chk(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
