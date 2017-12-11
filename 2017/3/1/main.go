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

func getDistanceFromMiddle(corner int, edgeSize int, target int) int {
	mid := float64(corner) - math.Ceil(float64(edgeSize/2))
	return int(math.Abs(mid - float64(target)))
}

func GetMinCubeDetailsThatContains(value int) *Cube {
	cur := 1
	layers := 0
	edgeLength := 1
	toAdd := 8
	for cur < value {
		cur += toAdd
		toAdd += 8
		edgeLength += 2
		layers += 1
	}
	return &Cube{
		Layers:           layers,
		EdgeLength:       edgeLength,
		BottomRightValue: cur,
	}
}
