package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var curMax = 0

const (
	N  = "n"
	NE = "ne"
	SE = "se"
	S  = "s"
	SW = "sw"
	NW = "nw"
)

func main() {
	fmt.Println("hello")
	dirs := buildInstructionSlice()
	dt := countDirs(dirs)
	fmt.Printf("N: %d\nS: %d\nNE: %d\nNW: %d\nSE: %d\nSW: %d\n", dt.N, dt.S, dt.NE, dt.NW, dt.SE, dt.SW)
	fmt.Println(curMax)
}

type DirectionTracker struct {
	N  int
	NE int
	SE int
	S  int
	SW int
	NW int
}

func countDirs(directions []string) *DirectionTracker {
	dt := DirectionTracker{}
	for i, direction := range directions {
		if direction != "" {
			adjustDistances(direction, &dt)
			if i > 20 {
				if getDist(&dt) > curMax {
					curMax = getDist(&dt)
				}
			}
		}
	}
	return &dt
}

func getDist(dirs *DirectionTracker) int {
	dists := make([]int, 0)
	if dirs.N > 0 {
		dists = append(dists, dirs.N)
	}
	if dirs.NE > 0 {
		dists = append(dists, dirs.NE)
	}
	if dirs.SE > 0 {
		dists = append(dists, dirs.SE)
	}
	if dirs.S > 0 {
		dists = append(dists, dirs.S)
	}
	if dirs.SW > 0 {
		dists = append(dists, dirs.SW)
	}
	if dirs.NW > 0 {
		dists = append(dists, dirs.NW)
	}
	distance := 0
	min := dists[0]
	for _, val := range dists {
		distance += val
		if val < min {
			min = val
		}
	}
	return distance - min
}

func adjustDistances(direction string, dt *DirectionTracker) {
	switch direction {
	case (N):
		dt.N += 1
		dt.S -= 1
	case (NE):
		dt.NE += 1
		dt.SW -= 1
	case (SE):
		dt.SE += 1
		dt.NW -= 1
	case (S):
		dt.S += 1
		dt.N -= 1
	case (SW):
		dt.SW += 1
		dt.NE -= 1
	case (NW):
		dt.NW += 1
		dt.SE -= 1
	}
}

func buildInstructionSlice() []string {
	pwd, err := os.Getwd()
	chk(err)
	datBytes, err := ioutil.ReadFile(pwd + "/input.txt")
	chk(err)
	dat := string(datBytes)
	directions := strings.Split(dat, ",")
	return directions
}

func chk(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
