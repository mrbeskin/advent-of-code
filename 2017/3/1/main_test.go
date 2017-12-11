package main

import (
	"testing"
)

func TestGetDistanceFromCenter(t *testing.T) {
	c := GetMinCubeDetailsThatContains(9)
	exLayers := 1
	if c.Layers != exLayers {
		t.Fatalf("input 9, for layers, expected %d, got %d\n", exLayers, c.Layers)
	}
	if c.EdgeLength != 3 {
		t.Fatalf("input 9, for EdgeLength expected %d, got %d\n", 3, c.EdgeLength)
	}
	if c.BottomRightValue != 9 {
		t.Fatalf("input 9, for bottom right value, expected %d, got %d\n", 9, c.BottomRightValue)
	}

	dist, _ := c.GetDistanceFromCenter(9)
	if dist != 2 {
		t.Fatalf("input 9, distance expected %d got %d\n", 2, dist)
	}

	c = GetMinCubeDetailsThatContains(10)
	if c.Layers != 2 {
		t.Fatalf("input 10, for layers, expected %d, got %d\n", 2, c.Layers)
	}
	if c.EdgeLength != 5 {
		t.Fatalf("input 10, for EdgeLength expected %d, got %d\n", 5, c.EdgeLength)
	}
	if c.BottomRightValue != 25 {
		t.Fatalf("input 10, for bottom right value, expected %d, got %d\n", 25, c.BottomRightValue)
	}

	dist, _ = c.GetDistanceFromCenter(10)
	if dist != 3 {
		t.Fatalf("input 10, distance expected %d got %d\n", 3, dist)
	}
}
