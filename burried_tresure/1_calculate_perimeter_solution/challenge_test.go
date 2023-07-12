package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTCR(t *testing.T) {
	assert.Equal(t, true, true)
}

func TestCalculatePerimeter1(t *testing.T) {
	points := [][2]int{{0, 0}, {0, 1}, {1, 1}, {1, 0}}
	p := CalculatePerimeter(points)
	assert.Equal(t, 4.0, p)
}
