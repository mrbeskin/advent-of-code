package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

const (
	arrowDelim = " -> "
	root       = "aapssr"
)

type Node struct {
	weight   int
	children []string
}

func main() {
	lines := getNodeLines()
	isChild := make(map[string]bool)
	childrens := make(map[string][]string)
	weights := make(map[string]int)
	for _, line := range lines {
		if line != "" {
			name, weight, children := getNameWeightChildrenFromString(line)
			childrens[name] = children
			weights[name] = weight
			for _, child := range children {
				isChild[child] = true
			}
			if isChild[name] != true {
				isChild[name] = false
			}
		}
	}
	fmt.Println(isChild)
	fmt.Println(getWeight(root, weights, childrens))
}

func getWeight(root string, weight map[string]int, children map[string][]string) int {
	childs := children[root]
	if len(childs) == 0 {
		return weight[root]
	}
	val := weight[root]
	for _, child := range childs {
		val += getWeight(child, weight, children)
	}
	return val

}

func getNameWeightChildrenFromString(line string) (string, int, []string) {
	pAndC := strings.Split(line, arrowDelim)
	children := make([]string, 0)
	if len(pAndC) > 1 {
		children = strings.Split(pAndC[1], ", ")
	}
	parentAndW := strings.Split(pAndC[0], " (")
	parent := parentAndW[0]
	weightString := strings.Split(parentAndW[1], ")")[0]
	weight, err := strconv.Atoi(weightString)
	chk(err)
	return parent, weight, children
}

func getNodeLines() []string {
	pwd, err := os.Getwd()
	chk(err)
	datBytes, err := ioutil.ReadFile(pwd + "/input.txt")
	chk(err)
	dat := string(datBytes)
	lines := strings.Split(dat, "\n")
	return lines
}

func chk(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
