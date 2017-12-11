package main

import (
	"fmt"
	"math"
)

var target = 265149

type Cube struct {
	Layers           int
	EdgeLength       int
	BottomRightValue int
}

func (c *Cube) GetDistanceFromCenter(target int) (int, error) {
	corner1 := c.BottomRightValue
	for i := 0; i < 4; i++ {
		corner2 := corner1 - (c.EdgeLength - 1)
		if (target <= corner1) && (target >= corner2) {
			return getDistanceFromMiddle(corner1, c.EdgeLength, target) + c.Layers, nil
		}
	}
	return 0, fmt.Errorf("value not on edge")
}

func main() {
	c := GetMinCubeDetailsThatContains(target)
	dist, err := c.GetDistanceFromCenter(target)
	if err != nil {
		fmt.Println("err :%v", err)
	} else {
		fmt.Println(dist)
	}
}

func getWrapAroundIndex(a []int, start int, move int) int {
	length := len(a)
	if start+move >= length {
		return start + move - length
	}
	if start+move < 0 {
		return (start + move) * -1
	}
	return start + move
}

func getValueAfter(n int) int {
	innerSlice := make([]int, 1)
	innerSlice[0] = 1
	outerSlice := make([]int, 9)
	for latestVal < n {
		for i := 0; i < len(outerSlice); i++ {

		}
	}
}
