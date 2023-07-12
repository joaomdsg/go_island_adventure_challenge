package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculatePerimeter_Square_Unit(t *testing.T) {
	points := [][2]int{{0, 0}, {0, 1}, {1, 1}, {1, 0}}
	p, _ := CalculatePerimeter(points)
	assert.Equal(t, 4.0, p, "The perimeter of a unit square should be 4")
}

func TestCalculatePerimeter_Square_Large(t *testing.T) {
	points := [][2]int{{0, 0}, {10, 0}, {10, 10}, {0, 10}}
	p, _ := CalculatePerimeter(points)
	assert.Equal(t, 40.0, p, "The perimeter of a square with side length 10 should be 40")
}

func TestCalculatePerimeter_Complex_Shape(t *testing.T) {
	points := [][2]int{{-3, -2}, {-1, 4}, {6, 1}, {3, 10}, {-4, 9}}
	p, _ := CalculatePerimeter(points)
	assert.InDelta(t, 48.57564046313958, p, 0.0000001, "The perimeter calculation should be accurate to within a small delta")
}

func TestCalculatePerimeter_Insufficient_Points(t *testing.T) {
	points := [][2]int{{0, 0}, {0, 1}}
	_, err := CalculatePerimeter(points)
	assert.Error(t, err, "The function should return an error if fewer than 3 points are given")
}

func TestCalculatePerimeter_Too_Many_Points(t *testing.T) {
	points := make([][2]int, 10001)
	for i := range points {
		points[i] = [2]int{0, 0}
	}
	_, err := CalculatePerimeter(points)
	assert.Error(t, err, "The function should return an error if more than 10000 points are given")
}
